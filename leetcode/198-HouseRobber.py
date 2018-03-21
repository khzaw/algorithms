class Solution:
    def rob(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        prev, now = 0, 0
        for n in nums:
            prev, now = now, max(prev + n, now)
            print(prev, now)
        return now
