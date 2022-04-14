class Solution:
    # Using TwoSum solution
    def twoSum(self, nums: List[int], target) -> List[int]:
        memo = {}
        out = []
        for i, val in enumerate(nums):
            if target - val in memo:
                out.append([val, target-val])
            memo[val] = i
        # If no such pair found
        return out
    
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        out = []
        for i, val in enumerate(nums):
            remTwo = self.twoSum(nums[i+1:], 0 - val)
            if remTwo:
                triple = [[val] + two for two in remTwo]
                out.extend(triple)
        return set(tuple(sorted(i)) for i in out)