from typing import List


class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        m = dict()
        for n in nums:
            if m in m:
                return True
            m[n] = True
        return False
