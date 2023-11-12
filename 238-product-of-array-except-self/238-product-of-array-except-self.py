class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        prefix = [1] * len(nums)
        suffix = [1] * len(nums)
        
        l = len(nums)
        
        for i in range(1, l):
            prefix[i] = prefix[i-1] * nums[i-1]
        
        for i in range(l-2, -1, -1):
            suffix[i] = suffix[i+1] * nums[i+1]
        
        out = [1] * l
        
        for i in range(l):
            out[i] = prefix[i] * suffix[i]
        
        return out