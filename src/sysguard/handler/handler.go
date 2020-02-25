package handler

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"sysguard/utils/print"
	"time"
)

var Host string
var Port string
var RHost string
var RPort string
var Key string = "00"

func SysHandler(w http.ResponseWriter, r *http.Request) {}

func SysHttpHandler() http.Handler {
	return HttpHandlerFunc(SysHandler)
}

type HttpHandlerFunc func(http.ResponseWriter, *http.Request)

// ServeHTTP calls f(w, r).
func (f HttpHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	r.Header.Add("X-Sys", Key)
	ru := r.RequestURI
	if gu := localUrl(r.RequestURI); gu != nil {
		r.RequestURI = ""
		r.URL = gu
		client := &http.Client{}
		streq := time.Now()
		resp, err := client.Do(r)
		stres := time.Now()
		if err == nil {
			for k, vs := range resp.Header {
				for _, v := range vs {
					w.Header().Set(k, v)
				}
			}
			w.WriteHeader(resp.StatusCode)
			if responseData,err := ioutil.ReadAll(resp.Body); err == nil {
				w.Write(responseData)
			}
			t := time.Now()
			elapsed := t.Sub(start)
			elapsedN := stres.Sub(streq)
			elapsedSys := elapsed - elapsedN
			print.Info("Completed [" + ru + "] 100% in " + elapsed.String() + " ( waf took " + elapsedSys.String() + ")")
		} else {
			print.Warning("An error occured : " + err.Error())
		}
	} else {
		print.Warning("A Request has been unsuccessful")
	}
}

func localUrl(opt string) *url.URL {
	if rurl, err := url.Parse("http://"+RHost + ":" + RPort + opt); err == nil {
		return rurl
	} else {
		return nil
	}
}