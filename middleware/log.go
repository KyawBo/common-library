package middleware

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/KyawBo/common-library/logger"
	"github.com/blendle/zapdriver"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func LoggingWithDumbBody(projectId string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()

			httpPayLoad := &zapdriver.HTTPPayload{
				RequestMethod: req.Method,
				UserAgent:     req.UserAgent(),
				RemoteIP:      req.RemoteAddr,
				Referer:       req.Referer(),
				Protocol:      req.Proto,
			}
			if req.URL != nil {
				httpPayLoad.RequestURL = req.URL.String()
			}

			// reqBody
			reqBody := []byte{}
			if c.Request().Body != nil { // Read
				reqBody, _ = ioutil.ReadAll(c.Request().Body)
			}
			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // Reset

			// resBody
			resBody := new(bytes.Buffer)
			mw := io.MultiWriter(res.Writer, resBody)
			writer := &MsgResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
			c.Response().Writer = writer

			msgReq := fmt.Sprintf("http request: %s %s - body: %s", req.Method, c.Path(), string(reqBody))

			traceContext := req.Header.Get("X-Cloud-Trace-Context")

			logger.InfoWithTraceId(req.Context(), projectId,
				msgReq,
				zap.Any("request_trace_context", traceContext),
				zap.String("body", string(reqBody)),
				zapdriver.HTTP(httpPayLoad),
			)

			start := time.Now()

			// next
			if err := next(c); err != nil {
				c.Error(err)
			}

			stop := time.Now()

			httpPayLoad.Status = res.Status
			httpPayLoad.ResponseSize = strconv.FormatInt(res.Size, 10)

			l := stop.Sub(start)
			httpPayLoad.Latency = l.String()

			msgResp := fmt.Sprintf("http response:  %s %d %s - body: %s", req.Method, res.Status, c.Path(), resBody)

			logger.InfoWithTraceId(req.Context(), projectId,
				msgResp,
				zap.Any("request_trace_context", traceContext),
				zap.String("body", resBody.String()),
				zapdriver.HTTP(httpPayLoad),
			)

			return nil
		}
	}

}

type MsgResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *MsgResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *MsgResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
