package handler

import (
	"net/http"
	"net/url"
	"net/http/httputil"
	"sysguard/utils/print"
	"errors"
)

var RHost *url.URL
var Key string = "00"

func SysHandler(w http.ResponseWriter, r *http.Request) {
	print.Critical("internal Misconfiguration - Request couldn't be sent to origin.")
}

func SysHttpHandler() http.Handler {
	return HttpHandlerFunc(SysHandler)
}

type HttpHandlerFunc func(http.ResponseWriter, *http.Request)

var hostReverseProxy *httputil.ReverseProxy

func Prepare() error {
	hostReverseProxy = httputil.NewSingleHostReverseProxy(RHost)
	if (hostReverseProxy == nil) {
		print.Critical("an error occured creating hostReverseProxy..")
		return errors.New("Can't init HostReverseProxy")
	} else {
		return nil
	}
}
// ServeHTTP calls f(w, r).
func (f HttpHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostReverseProxy.ServeHTTP(w, r)
	print.Info("Request ["+r.RequestURI+"]")
}