/*
 * Mqtts API
 *
 * Mqtts service
 *
 * API version: 1.0.0
 */
package swagger

import (
	"net/http"
	"time"
)

//Logger Logs each request
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		lg.Trace(

			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
