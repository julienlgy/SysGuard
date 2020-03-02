package main

import (
	"net"
	"net/http"
	"net/url"
	"os"
	"sysguard/utils/print"
	"sysguard/handler"
	"sysguard/utils/config"
	"flag"
	"fmt"
	"time"
	"crypto/tls"
)

type sargs struct {
	origin *url.URL
	listen *url.URL
	ssl bool
	key string
	certs string
}

func main() {
	var args *sargs = new(sargs)

	if parseArgs(args) == false {
		print.Critical("Missing parameters.")
		os.Exit(2)
	}

	print.Welcome()
	print.Info("Starting SysGuard proxy...")

	handler.RHost = args.origin
	handler.Key = args.key
	
	if args.ssl {
		tlsConfig := config.Parse(args.certs)
		tlsConfig.InsecureSkipVerify = true
		handler.Prepare()
		server := &http.Server{
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
			TLSConfig:      tlsConfig,
			Handler: handler.SysHttpHandler(),
		}

		listener, err := tls.Listen("tcp", args.listen.Hostname() + ":" + args.listen.Port(), tlsConfig)
		if err != nil {
			fmt.Print(err)
		}
		
		print.Info("OK : ["+args.listen.String()+"] -> ["+args.origin.String()+"]")
		server.Serve(listener)
	} else {
		if ln, err := net.Listen("tcp", args.listen.Hostname() + ":" + args.listen.Port()); err == nil {
			print.Info("OK : ["+args.listen.String()+"] -> ["+args.origin.String()+"]")
			handler.Prepare()
			http.Serve(ln, handler.SysHttpHandler())
		} else {
			print.Critical("Can't init SysGuard proxy server")
			print.Critical(err.Error())
		}
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

func parseArgs(args *sargs) bool {
	var origin, listen string
	var key, cert string
	var tssl * bool

	flag.StringVar(&origin, "origin", "", "Original Application URL")
	flag.StringVar(&listen, "listen", "", "Where to listen")
	flag.StringVar(&cert, "certs", "./certs.json", "Certifcat JSON file configuration.")
	flag.StringVar(&key, "key", "default", "SysGuardian key for locking your app.")
	tssl = flag.Bool("ssl", true, "Enabling SSL or not.")
	flag.Parse()

	args.ssl = *tssl
	args.key = key
	args.certs = cert
	if len(origin) > 0 && len(listen) > 0 {
		if rurl, err := url.Parse(origin); err == nil {
			args.origin = rurl
		} else {
			print.Critical("Can't read origin URL.")
			os.Exit(1)
		}
		if rurl, err := url.Parse(listen); err == nil {
			args.listen = rurl
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