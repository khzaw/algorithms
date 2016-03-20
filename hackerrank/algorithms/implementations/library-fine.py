#!/bin/python

import sys
from datetime import date, timedelta
from calendar import monthrange


d1,m1,y1 = raw_input().strip().split(' ')
d1,m1,y1 = [int(d1),int(m1),int(y1)]
actual = date(y1, m1, d1)

d2,m2,y2 = raw_input().strip().split(' ')
d2,m2,y2 = [int(d2),int(m2),int(y2)]
expected = date(y2, m2, d2)

def monthdelta(d1, d2):
    delta = 0
    while True:
        mdays = monthrange(d1.year, d1.month)[1]
        d1 += timedelta(days=mdays)
        if d1 <= d2 and d1.month != d2.month:
            delta += 1
        else:
            break
    return delta

if actual <= expected:
    print 0
elif actual.year != expected.year:
    print 10000
elif actual.month == expected.month:
    print 15 * (actual-expected).days
elif actual.year == expected.year:
    print 500 * monthdelta(expected, actual)