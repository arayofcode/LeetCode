class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        product = 1
        output = []
        for i in range(len(nums)):
            output.append(product)
            product *= nums[i]
        
        product = 1
        for i in range(len(nums) - 1, -1, -1):
            output[i] *= product
            product *= nums[i]
        
        return output