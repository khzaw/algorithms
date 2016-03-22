#!/bin/python

n, m = map(int, raw_input().strip().split(' '))
denom = map(int, raw_input().strip().split(' '))

def coin_change(given_denoms, p):
    given_denoms = sorted(given_denoms)
    c = [0]*(p+1)
    for i in xrange(1, p+1):
        denoms = min(c[i-d] for d in given_denoms if d <= i)
        c[i] = denoms + 1
    return c[p]

print coin_change(denom, n)