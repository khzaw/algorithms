#!/bin/python

from math import sqrt, floor
import sys

t = int(raw_input().strip())
for a0 in xrange(t):
    a, b = map(int, raw_input().strip().split(' '))
    print filter(lambda x: floor(sqrt(x) + 0.5) ** 2 == x, range(a, b+1))