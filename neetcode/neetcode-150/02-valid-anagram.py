"""
Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.
"""

class Solution:
    def is_anagram(self, s: str, t: str) -> bool:
        m = {}
        for c in s:
            if c in m:
                m[c] += 1
            else:
                m[c] = 1

        for c in t:
            if c in m:
                if m[c] == 1:
                    m.pop(c, None)
                else:
                    m[c] -= 1
            else:
                return False
        return len(m) == 0



s = Solution()
print(s.is_anagram("abc", "cab"))
