class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """

        """
        O(n^2)
        """
        # for i in range(len(nums)):
        #     for j in range(i+1, len(nums)):
        #         if nums[i] + nums[j] == target:
        #             return [i, j]

        """
        O(n) time and O(n) space
        """
        mydict = {}
        for i, v in enumerate(nums):
            compliment = target - v
            if compliment not in mydict:
                mydict[v] = i
            else:
                return [mydict[complement], i]
