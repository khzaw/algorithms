#!/bin/python

n, m = map(int, raw_input().strip().split(' '))
denom = map(int, raw_input().strip().split(' '))

def coin_change(coins, amount):
    ways = [0] * (amount + 1)
    ways[0] = 1
    for coin in coins:
        for j in xrange(coin, amount + 1):
            ways[j] += ways[j - coin]
    return ways[amount]

print coin_change(denom, n)