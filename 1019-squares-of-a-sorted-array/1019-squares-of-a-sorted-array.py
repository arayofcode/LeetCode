class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        sorted_array = sorted(nums, key= lambda a: abs(a))
        squares = []
        [squares.append(i*i) for i in sorted_array]
        return squares