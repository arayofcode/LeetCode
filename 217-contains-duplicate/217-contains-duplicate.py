class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        memo = {}
        for i in nums:
            if i in memo:
                return True
            else: 
                memo[i] = 0
        return False