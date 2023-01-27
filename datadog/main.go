package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type LogMessage struct {
	Host    string      `json:"hostname"`
	AppName string      `json:"appname"`
	Msg     string      `json:"message"`
	Session string      `json:"session"`
	Span    interface{} `json:"span"`
}

func main() {
	rules := []tracer.SamplingRule{tracer.RateRule(1)}
	tracer.Start(
		tracer.WithSamplingRules(rules),
		tracer.WithService("mux.test"),
		tracer.WithEnv("env.test"),
	)
	defer tracer.Stop()

	logfile, err := os.OpenFile("/usr/log/api/test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open test.log: " + err.Error())
	}
	defer logfile.Close()

	// Create a traced mux router.
	mux := muxtrace.NewRouter()
	// Continue using the router as you normally would.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		span := tracer.StartSpan("web.request", tracer.ResourceName("/posts"))
		defer span.Finish()

		w.Write([]byte("Hello World!"))
		m := LogMessage{
			Host:    "ubuntu",
			AppName: "go-test",
			Msg:     "HandleFunc called",
			Session: "13e934ps",
			Span:    span,
		}
		s, err := json.Marshal(m)
		if err != nil {
			logfile.WriteString("{\"Error\": \"Failed to Marshal Struct to Json\"}")
		}
		// FIXME: もっといいやり方がありそう
		s = append(s, []byte("\n")...)
		logfile.Write(s)

		// TODO: どうやってファイルに span の情報を入れるのか
		log.Printf("HandlerFunc called!! %v", span)
	})
	http.ListenAndServe(":8084", mux)
}
