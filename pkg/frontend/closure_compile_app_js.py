#!/usr/bin/python2.7

import httplib
import urllib


def getReqParams(f):
    return urllib.urlencode([
        ('js_code', f.read()),
        ('compilation_level', 'SIMPLE_OPTIMIZATIONS'),
        ('output_format', 'text'),
        ('output_info', 'compiled_code'),
    ])


def compile(fname):
    f = open("raw/" + fname, "r")
    headers = {"Content-type": "application/x-www-form-urlencoded"}
    conn = httplib.HTTPSConnection('closure-compiler.appspot.com')
    conn.request('POST', '/compile', getReqParams(f), headers)
    resp = conn.getresponse()

    c = open("dist/" + fname, "w")
    c.write(resp.read())

    conn.close()


compile("app.js")
compile("app_dev.js")
