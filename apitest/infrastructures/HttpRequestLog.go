package infrastructures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"text/template"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

const (
	// DefaultDateTimeFormat Y-m-d H:s:i
	DefaultDateTimeFormat = "2006-01-02 15:04:05"
)

// LoggerDefaultDateFormat is the
// format used for date by the
// default Logger instance.
var LoggerDefaultDateFormat = time.RFC3339

// ALogger interface
type ALogger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

// LoggerMiddleware is a middleware handler that logs the request as it goes in and the response as it goes out.
type HttpLogger struct {
	// ALogger implements just enough log.Logger interface to be compatible with other implementations
	ALogger
	dateFormat string
	template   *template.Template
}

// NewLoggerMiddleware returns a new Logger instance
func NewHttpLogger() *HttpLogger {
	//logger := &HttpLogger{ALogger: log.New(os.Stdout, "[wallet router] ", 0), dateFormat: LoggerDefaultDateFormat}
	logger := &HttpLogger{ALogger: log.New(os.Stdout, "", 0), dateFormat: LoggerDefaultDateFormat}
	//logger.SetFormat(LoggerDefaultFormat)
	logger.SetDateFormat(DefaultDateTimeFormat)
	return logger
}

// SetDateFormat log time
func (l *HttpLogger) SetDateFormat(format string) {
	l.dateFormat = format
}

// ServeHTTP the http serve
func (l *HttpLogger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		r := recover()
		if r != nil {
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(http.StatusInternalServerError)

			rsp := struct {
				Code    int64       `json:"code"`
				Message interface{} `json:"message"`
			}{
				Code:    2000,
				Message: "internal server error",
			}

			json.NewEncoder(rw).Encode(&rsp)

			buf := make([]byte, 256)
			buf = buf[:runtime.Stack(buf, false)]
			logrus.Errorf("panic %v \n %s\n", r, buf)

			return
		}
	}()

	start := time.Now()

	next(rw, r)

	res := rw.(negroni.ResponseWriter)

	log := map[string]interface{}{
		"level":     logrus.InfoLevel.String(),
		"component": "http",
		"timestamp": start.Format(l.dateFormat),
		"duration":  time.Since(start),
		"status":    res.Status(),
		"hostname":  r.Host,
		"path":      r.URL.Path,
		"method":    r.Method,
		"remote_ip": r.RemoteAddr,
	}

	buff := &bytes.Buffer{}
	json.NewEncoder(buff).Encode(log)
	fmt.Println(buff.String())

}

func (l *HttpLogger) drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}
