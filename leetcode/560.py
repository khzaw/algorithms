class Solution:
    def subarraySum(self, nums: list[int], k: int) -> int:
        count = 0
        m = {0: 1}
        curr_sum = 0

        for num in nums:
            curr_sum += num

            if curr_sum - k in m:
                count += m[curr_sum - k]

            m[curr_sum] = m.get(curr_sum, 0) + 1

        return count
