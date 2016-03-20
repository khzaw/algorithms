#!/bin/python

import sys
import re
from string import lowercase

s = raw_input().strip().lower()
s = re.sub('[\s+]', '', s)
pangram = 'pangram' if set(lowercase) == set(s) else 'not pangram'
print pangram