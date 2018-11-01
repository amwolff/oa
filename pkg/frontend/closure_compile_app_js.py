#!/usr/bin/python2.4

import httplib
import urllib

f = open("unoptimized/app.js", "r")

params = urllib.urlencode([
    ('js_code', f.read()),
    ('compilation_level', 'SIMPLE_OPTIMIZATIONS'),
    ('output_format', 'text'),
    ('output_info', 'compiled_code'),
])

headers = {"Content-type": "application/x-www-form-urlencoded"}
conn = httplib.HTTPSConnection('closure-compiler.appspot.com')
conn.request('POST', '/compile', params, headers)
response = conn.getresponse()
c = open("optimized/app.js", "w")
c.write(response.read())
conn.close()
