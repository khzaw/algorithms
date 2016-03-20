#!/bin/python

import sys
import string


n = int(raw_input().strip())
s = raw_input().strip()
k = int(raw_input().strip())

print string.translate(s, string.maketrans(string.uppercase + string.lowercase,
                                     string.uppercase[k:] + string.uppercase[:k] +
                                     string.lowercase[k:] + string.lowercase[:k]))