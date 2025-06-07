"""
Given an array of strings strs, group all anagrams together into sublists. You may return the output in any order.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.
"""

from collections import defaultdict


class Solution:
    def histogram(self, s: str) -> tuple[int, ...]:
        h = [0] * 26
        for c in s:
            h[ord(c) - ord("a")] += 1
        return tuple(h)

    def groupAnagrams(self, strs: list[str]) -> list[list[str]]:
        groups = defaultdict(list)
        for s in strs:
            groups[self.histogram(s)].append(s)
        return list(groups.values())
