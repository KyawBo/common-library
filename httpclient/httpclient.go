package httpclient

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

func SetDefaultHeader(ctx context.Context, jwtTokenRaw string, isBody bool) http.Header {
	header := make(http.Header)
	if jwtTokenRaw != "" {
		header.Set("Authorization", "Bearer "+jwtTokenRaw)
	}
	if isBody {
		header.Set("Content-Type", "application/json")
	}
	header.Set("cache-control", "no-cache")
	header.Set("Connection", "keep-alive")

	spanContext := trace.SpanContextFromContext(ctx)
	if spanContext.IsValid() {
		traceId := spanContext.TraceID().String()
		spanId := spanContext.SpanID().String()
		traceTrue := spanContext.HasTraceID()

		var sampledFlag string
		if traceTrue {
			sampledFlag = "01"
		} else {
			sampledFlag = "00"
		}

		traceContext := fmt.Sprintf("00-%s-%s-%s", traceId, spanId, sampledFlag)
		header.Set("traceparent", traceContext)
	}

	return header
}
