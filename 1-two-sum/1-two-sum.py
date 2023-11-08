class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        nums_dict = {}
        for index, n in enumerate(nums):
            if target - n in nums_dict:
                return [nums_dict[target - n], index]
            nums_dict[n] = index