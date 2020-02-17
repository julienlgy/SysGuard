"""

    SYSGUARD 2020
    EPITECH PROJECT

"""
import socket
import select
import time
import sys

class Forwarder:
    def __init__(self):
        self.forward = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    def start(self, host, port):
        try:
            self.forward.connect((host, port))
            return self.forward
        except Exception, e:
            print e
            return False