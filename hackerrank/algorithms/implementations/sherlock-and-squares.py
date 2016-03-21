#!/bin/python

from math import sqrt, floor, ceil
import sys

t = int(raw_input().strip())
for a0 in xrange(t):
    a, b = map(sqrt, (map(int, raw_input().strip().split(' '))))
    print int(floor(b) - ceil(a) + 1)