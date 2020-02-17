#!/usr/bin/env python
"""
    SYSGUARD 2020
    EPITECH PROJECT
"""

from core import config, gateway

print config.get("proxy_host")
gateway.listen()

