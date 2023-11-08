class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        memo = {}
        for i, val in enumerate(nums):
            if target - val in memo:
                return [i, memo[target-val]]
            memo[val] = i