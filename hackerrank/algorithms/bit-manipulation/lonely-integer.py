#!/usr/bin/py
def lonelyinteger(a):
    return reduce(int.__xor__, a)
if __name__ == '__main__':
    a = input()
    b = map(int, raw_input().strip().split(" "))
    print lonelyinteger(b)
