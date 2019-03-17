// Package routes is router http request
package routes

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

const (
	panicMessage = "Oops, something went wrong!"
)

// StaticTextPanicFormatter implements interface negroni.PanicFormatter.
type StaticTextPanicFormatter struct{}

// FormatPanicError implement the same function from interface negroni.PanicFormatter.
// This function will return a static message rather than error message.
func (t *StaticTextPanicFormatter) FormatPanicError(rw http.ResponseWriter, r *http.Request, infos *negroni.PanicInformation) {
	if rw.Header().Get("Content-Type") == "" {
		rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	}
	fmt.Fprintf(rw, panicMessage)
}
