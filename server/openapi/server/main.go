package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"openapi-generator-server/gen/component"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/gin-gonic/gin"
)

type ave struct {
	num   int
	total int
}

var av ave

// https://pkg.go.dev/runtime/pprof
func startCPUProf(fileNameSuffix string) (closer func(), err error) {
	fc, err := os.Create(fmt.Sprintf("cpu-%s.prof", fileNameSuffix))
	if err != nil {
		return nil, fmt.Errorf("failed to create cpu profile: %w", err)
	}

	closer = func() {
		pprof.StopCPUProfile()

		if err := fc.Close(); err != nil {
			//nolint:forbidigo
			fmt.Printf("failed to close: %v\n", err)
		}
	}

	if err := pprof.StartCPUProfile(fc); err != nil {
		closer()

		return nil, fmt.Errorf("failed to start cpu profile: %w", err)
	}

	return
}

func main() {
	runtime.GOMAXPROCS(1)
	r := gin.Default()
	r.Handle("GET", "/comments", func(c *gin.Context) {
		bytes, err := os.ReadFile("res.json")
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		var res component.ListCommentsResponse
		json.Unmarshal(bytes, &res)

		start := time.Now()

		cc, err := startCPUProf(start.Format("2006-01-02T15-04-05"))
		if err != nil {
			log.Fatal(err)
		}
		defer cc()

		rawBytes, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}

		cost := time.Since(start).Milliseconds()
		fmt.Printf("time.Since(start).Milliseconds() %v\n", cost)

		av.num++
		av.total += int(cost)
		fmt.Printf("%v\n", av.total/av.num)

		c.Data(http.StatusOK, "application/json", rawBytes)
	})

	r.Run()
}
