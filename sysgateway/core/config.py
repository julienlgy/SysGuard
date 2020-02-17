"""
    SYSGUARD 2020
    EPITECH PROJECT
"""
import configparser

repo = "DEFAULT"
conf = configparser.ConfigParser()
conf.read('sysguard.conf')

def get(key):
    return conf[repo][key]

def set(key, value):
    conf[repo][key] = value