class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        # Simplest solution is to merge the two arrays and find the middle element
        merged_array = []
        i = 0
        j = 0
        while i < len(nums1) or j < len(nums2):
            if i == len(nums1):
                merged_array.append(nums2[j])
                j += 1
                continue
            elif j == len(nums2):
                merged_array.append(nums1[i])
                i += 1
                continue

            if nums1[i] < nums2[j]:
                merged_array.append(nums1[i])
                i += 1
            else:
                merged_array.append(nums2[j])
                j += 1
        length = len(merged_array)
        if length % 2 == 1:
            return merged_array[(length + 1) // 2 - 1]
        else:
            return (merged_array[length // 2 - 1] + merged_array[length // 2]) / 2