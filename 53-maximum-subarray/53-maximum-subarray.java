class Solution {
    public int maxSubArray(int[] nums) {
        /* 
        Two things to remember here:
          1. Max sum should include the max element found so far
          2. Two choices to make: Should we add num[i] to previous best solution or start a new solution from index i itself?
        The choices is an easy part. Which choice gives a better sum? 
        */
        int max_ending_until_i = nums[0];
        int maxSum = nums[0];
        
        for(int i = 1; i < nums.length; i++){
            max_ending_until_i = Math.max(max_ending_until_i + nums[i], nums[i]);
            
            if(max_ending_until_i > maxSum)
                maxSum = max_ending_until_i;
        }
        return maxSum;
    }
}