#!/bin/python

a, b, n = map(int, raw_input().strip().split(' '))
for i in xrange(2, n):
    c = b*b + a
    a,b = b,c
print c