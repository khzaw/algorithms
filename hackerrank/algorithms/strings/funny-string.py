#!/bin/python

import sys
from itertools import izip

t = int(raw_input().strip())
for a0 in xrange(t):
    s = raw_input().strip()
    r = s[::-1]
    sfunny = map(lambda (x,y): abs(ord(x) - ord(y)), izip(s[1:], s))
    rfunny = map(lambda (x,y): abs(ord(x) - ord(y)), izip(r[1:], r))
    funny = 'Funny' if sfunny == rfunny else 'Not Funny'
    print funny