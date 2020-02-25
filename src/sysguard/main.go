package main

import (
	"net"
	"net/http"
	"net/url"
	"os"
	"sysguard/utils/print"
	"sysguard/handler"
)

func main() {
	print.Welcome()
	print.Info("Starting SysGuard proxy...")

	sysproxy := parseEnvUrl("SYSPROXY")
	sysgateway := parseEnvUrl("SYSGATEWAY")
	key := os.Getenv("SYSKEY")

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
		if rurl, err := url.Parse(value); err == nil {
			return rurl
		} else {
			print.Critical("Can't Read "+key+" Environment variable")
			os.Exit(2)
		}
	} else {
		print.Critical("Can't Read "+key+" Environment variable")
		os.Exit(2)
	}
	panic("Error")
}