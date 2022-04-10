class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        """
        Two things to remember here:
          1. Max sum should include the max element found so far
          2. Two choices to make: Should we add num[i] to previous best                    solution or start a new solution from index i itself?
        The choices is an easy part. Which choice gives a better sum? 
        """
        max_ending_at_i = nums[0]
        maxSum = nums[0]
        
        for i in range(1, len(nums)):
            max_ending_at_i = max(max_ending_at_i + nums[i], nums[i])
            maxSum = max_ending_at_i if max_ending_at_i > maxSum else maxSum
        return maxSum