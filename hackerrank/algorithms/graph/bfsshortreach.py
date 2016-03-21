#!/bin/python

t = int(raw_input().strip())
for g in xrange(t):
    graph = map(int, raw_input().strip().split(' '))
    n, m = graph[0], graph[1]
