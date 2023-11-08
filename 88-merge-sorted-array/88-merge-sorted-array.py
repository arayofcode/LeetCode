class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        i = 0
        
        while i <= m and n != 0:
            if nums1[i] > nums2[0]:
                # Insert the first element of num2
                nums1.insert(i, nums2.pop(0))
                # Remove the last 0 from nums1
                nums1.pop()
                m += 1
                n -= 1
            
            i += 1
        
        i = 0
        
        # Adding remaining elements
        while i < n:
            nums1[i+m] = nums2.pop(0)
            i += 1