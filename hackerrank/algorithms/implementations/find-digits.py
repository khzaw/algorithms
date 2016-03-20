#!/bin/python

import sys

digits = lambda x: filter(lambda y: y != 0, map(int, str(x)))

t = int(raw_input().strip())
for a0 in xrange(t):
    n = int(raw_input().strip())
    print len(filter(lambda x: n % x == 0, digits(n)))