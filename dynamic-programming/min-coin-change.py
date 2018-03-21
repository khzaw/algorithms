#!/bin/python

#
# Given coins of denominations m, find minimum number of coins
# required to make an amount n.

# n = amount
# m = number of denominations, followed by the denoms list


n, m = list(map(int, input().strip().split(' ')))
denom = list(map(int, input().strip().split(' ')))

# def coin_change(given_denoms, p):
    # given_denoms = sorted(given_denoms)
    # c = [0]*(p+1)
    # for i in xrange(1, p+1):
        # denoms = min(c[i-d] for d in given_denoms if d <= i)
        # c[i] = denoms + 1
    # return c[p]

def build_matrix(n, c):
    m = [[0 for _ in xrange(n + 1)] for _ in xrange(len(c) + 1)]
    for i in xrange(n + 1):
        m[0][i] = i
    return m

# def getWays(n, c):
#    m = build_matrix(n, c)
#    for i in xrange(1, len(c) + 1):
#        for j in xrange(1, n + 1):
#            # i = n, j = c
#            if j >= coin[i]:

def getWays(n, c):
    ways = [1] + [0]*n
    for coin in c:
        for i in range(coin, n+1):
            if i >= coin:
                ways[i] += ways[i-coin]
    return ways[n]
