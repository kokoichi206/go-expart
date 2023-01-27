package main

import (
	"net/http"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	rules := []tracer.SamplingRule{tracer.RateRule(1)}
	tracer.Start(
		tracer.WithSamplingRules(rules),
		tracer.WithService("mux.test"),
		tracer.WithEnv("env.test"),
	)
	defer tracer.Stop()

	// Create a traced mux router.
	mux := muxtrace.NewRouter()
	// Continue using the router as you normally would.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":8084", mux)
}
