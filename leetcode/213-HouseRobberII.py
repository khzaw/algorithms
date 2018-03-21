class Solution:
    def rob(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        def rob(nums):
            prev, now = 0, 0
            for n in nums:
                prev, now = now, max(prev + n, now)
            return now

        if len(nums) == 1:
            return nums[0]

        return max(rob(nums[1:]), rob(nums[:-1]))
