#!/bin/python
from itertools import chain

n,k = map(int, raw_input().strip().split(' '))
chapters = map(int, raw_input().strip().split(' '))
#  chapters = chain.from_iterable([c if c <= k else filter(lambda x: x, [k]*(c/k) + [c%k]) for c in chapters])


total = 0
for c in chapters:


#  is_special = lambda x, chapter, total: chapter.index(x) + 1 == total
special = 0
total = 0
for c in chapters:
    chapter = range(1, c+2)
    pages = [c if c <= k else filter(lambda x: x, [k]*(c/k)+[c%k])]
    for page in pages:
        total += 1
        if is_special(i, chapter, total):
            special += 1
print special