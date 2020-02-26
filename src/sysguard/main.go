package main

import (
	"net"
	"net/http"
	"net/url"
	"os"
	"sysguard/utils/print"
	"sysguard/handler"
	"flag"
)

func main() {
	print.Welcome()
	print.Info("Starting SysGuard proxy...")

	var sysproxy *url.URL
	var sysgateway *url.URL
	var key string

	if parseArgs(&sysproxy, &sysgateway) == false {
		sysproxy = parseEnvUrl("SYSPROXY")
		sysgateway = parseEnvUrl("SYSGATEWAY")
		key = os.Getenv("SYSKEY")
	}

	handler.Host = sysproxy.Hostname()
	handler.Port = sysproxy.Port()
	handler.RHost = sysgateway.Hostname()
	handler.RPort = sysgateway.Port()
	handler.Key = key
	
	if ln, err := net.Listen("tcp", sysproxy.Hostname() + ":" + sysproxy.Port()); err == nil {
		print.Info("OK : ["+sysproxy.String()+"] -> ["+sysgateway.String()+"]")
		http.Serve(ln, handler.SysHttpHandler())
	} else {
		print.Critical("Can't init SysGuard proxy server")
		print.Critical(err.Error())
	}

	print.Info("Exiting...")
}

func parseEnvUrl(key string) *url.URL {
	if value := os.Getenv(key); len(value) > 0 {
		return strToUrl(value)
	} else {
		print.Critical("Can't Read "+key+" - ")
		os.Exit(2)
	}
	panic("Error")
}

func parseArgs(sysproxy **url.URL, sysgateway **url.URL) bool {
	var origin string
	var listen string
	flag.StringVar(&origin, "origin", "", "Original Application URL")
	flag.StringVar(&listen, "listen", "", "Where to listen")
	flag.Parse()
	if len(origin) > 0 && len(listen) > 0 {
		if rurl, err := url.Parse(origin); err == nil {
			*sysgateway = rurl
		} else {
			print.Critical("Can't read origin URL.")
			os.Exit(1)
		}
		if rurl, err := url.Parse(listen); err == nil {
			*sysproxy = rurl
		} else {
			print.Critical("Can't read listen URL")
			os.Exit(1)
		}
		return true
	}
	return false
}

func strToUrl(str string) *url.URL {
	if rurl, err := url.Parse(str); err == nil {
		return rurl
	} else {
		print.Critical("Can't parse Url")
		os.Exit(2)
	}
	panic("Error")
}