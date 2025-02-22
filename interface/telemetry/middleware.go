package telemetry

import (
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"
)

func WrapHandler(handler http.Handler) http.Handler {
	return otelhttp.NewHandler(handler, "http-server")
}
