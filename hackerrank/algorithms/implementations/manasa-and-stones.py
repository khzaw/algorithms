#!/bin/python
from itertools import product

t = int(raw_input().strip())
for a0 in xrange(t):
  n = int(raw_input().strip())
  a = int(raw_input().strip())
  b = int(raw_input().strip())

  end = set()
  a, b = sorted([a, b])
  if a == b:
    print (n-1)*a
  else:
    for i in xrange((n-1)*a, ((n-1)*b) + 1, (b-a)):
      end.add(i)
    print ' '.join(str(x) for x in sorted(end))
