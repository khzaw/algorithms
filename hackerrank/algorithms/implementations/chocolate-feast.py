#!/bin/python

import sys

t = int(raw_input().strip())
for a0 in xrange(t):
    n,c,m = raw_input().strip().split(' ')
    n,c,m = [int(n),int(c),int(m)]

    total = n/c
    wrappers = total
    while(wrappers >= m):
        total = total + (wrappers/m)
        leftover = wrappers % m
        wrappers = (wrappers/m) + leftover
    print total