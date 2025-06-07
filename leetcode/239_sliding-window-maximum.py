class Solution:
    def maxSlidingWindow(self, nums: list[int], k: int) -> list[int]:
        m = []

        for right in range(k - 1, len(nums)):
            m.append(max(nums[right - k + 1 : right + 1]))

        return m
