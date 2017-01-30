package cors

import (
	"net/http"
	"net/textproto"
	"strings"
)

const (
	AccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	AccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	AccessControlExposeHeaders    = "Access-Control-Expose-Headers"
)

var originKey = textproto.CanonicalMIMEHeaderKey("Origin")

type Response struct {
	AllowOrigin string
	// AllowMethods     []string
	// AllowHeaders     []string
	ExposeHeaders []string
	// MaxAge           string
	AllowCredentials bool
}

func Middleware(f func(*http.Request) Response) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// if origin := r.Header[originKey]; len(origin) != 0 {
			cr := f(r)
			if cr.AllowCredentials {
				w.Header().Set(AccessControlAllowCredentials, "true")
			}
			if cr.AllowOrigin != "" {
				w.Header().Set(AccessControlAllowOrigin, cr.AllowOrigin)
			}
			if len(cr.ExposeHeaders) > 0 {
				w.Header().Set(AccessControlExposeHeaders, strings.Join(cr.ExposeHeaders, ", "))
			}
			// }
			h.ServeHTTP(w, r)
		})
	}
}
