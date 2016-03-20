#!/bin/python

import sys
from itertools import cycle

t = int(raw_input().strip())
for a0 in xrange(t):
    n = int(raw_input().strip())

    f = cycle([lambda x: 2*x, lambda y:y+1])
    height = 1
    for i in range(n):
        height = f.next()(height)
    print height