#!/bin/python

import sys


n = int(raw_input().strip())
sticks = map(int,raw_input().strip().split(' '))

while len(sticks) >= 1:
    print len(sticks)
    minimum = min(sticks)
    sticks = filter(lambda y: y > 0, map(lambda x: x - minimum, sticks))