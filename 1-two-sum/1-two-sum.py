class Solution:
    # Brute-forced solution
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i, val in enumerate(nums):
            for j in range(i + 1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i, j]

    # Efficient solution
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # Create a dictionary, store index with the value as key
        # If target - val exists in dictionary, return indexes
        # else add val to dictionary
        memo = {}
        for i, val in enumerate(nums):
            if target - val in memo:
                return [i, memo[target-val]]
            memo[val] = i