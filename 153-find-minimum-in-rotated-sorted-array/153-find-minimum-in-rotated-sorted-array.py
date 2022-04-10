class Solution:
    def findMin(self, nums: List[int]) -> int:
        n = len(nums)
        
        low = 0
        high = n - 1
        
        minimum = nums[low]

        while low <= high:
            mid = (low + high) // 2
            if nums[mid] > nums[(mid + 1) % n]:
                minimum = nums[mid + 1]
                break
            elif nums[mid - 1] > nums[mid]:
                minimum = nums[mid]
                break
            elif nums[mid] < nums[high]:
                high = mid - 1
            else:
                low = mid + 1

        return minimum