#!/bin/python
from itertools import product

t = int(raw_input().strip())
for a0 in xrange(t):
  n = int(raw_input().strip())
  a = int(raw_input().strip())
  b = int(raw_input().strip())

  end = set()
  for i in product([a, b], repeat=(n-1)):
    end.add(reduce(int.__add__, i))
  print ' '.join(str(x) for x in sorted(end))
