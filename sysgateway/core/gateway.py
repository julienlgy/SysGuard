
import config
import sys
from .proxy import server

def listen():
    pserv = server.Server(config.get('proxy_host'), int(config.get('proxy_port')))
    pserv.web_server["host"] = config.get('local_host')
    pserv.web_server["port"] = int(config.get('local_port'))
    try:
        pserv.main_loop()
    except KeyboardInterrupt:
        print "Ctrl C - Stopping server"
        sys.exit(1)
