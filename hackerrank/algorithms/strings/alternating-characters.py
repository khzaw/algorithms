#!/bin/python

import sys

t = int(raw_input().strip())
for _ in range(t):
    st = raw_input().strip()
    count = 0
    for c in range(len(st) - 1):
        if st[c] == st[c+1]:
            count += 1
    print count
