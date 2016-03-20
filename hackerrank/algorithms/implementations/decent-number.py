#!/bin/python

import sys

digits = lambda x: map(int, str(x))
threefive = lambda x: set(digits(x)) in [set([3,5]), set([3]), set([5])]
divisible = lambda x, y: len(filter(lambda z: z == 3, digits(x))) % y == 0
is_decent = lambda x: threefive(x) and divisible(x, 3) and divisible(x, 5)

#  def largest_decent_number(n):
    #  numbers = [x for x in xrange(1, int('5'*n)+1)]
    #  return filter(is_decent, numbers)

#  print largest_decent_number(11)


def largest_decent_number(n):
    y = n
    z = y
    while (z % 3 != 0):
        z = z - 5
    if z < 0:
        return -1
    else:
        return z*'5' + (y-z) * '3'


t = int(raw_input().strip())
for a0 in xrange(t):
    n = int(raw_input().strip())
    print largest_decent_number(n)