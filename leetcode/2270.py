# pyright: basic

# from itertools import accumulate


# class Solution:
#     def waysToSplitArray(self, nums: list[int]) -> int:
#         prefix = list(accumulate(nums))
#         return sum(left >= (prefix[-1] - left) for left in prefix[:-1])


class Solution:
    def waysToSplitArray(self, nums: list[int]) -> int:
        total = sum(nums)
        ways = left = 0

        for num in nums[:-1]:
            left += num
            if left >= total - left:
                ways += 1
        return ways
