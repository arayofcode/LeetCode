class Solution:
    def maxProfit1(self, prices: List[int]) -> int:
        maximum = 0
        for i in range(len(prices)):
            for j in range(i + 1, len(prices)):
                if prices[j] - prices[i] > maximum:
                    maximum = prices[j] - prices[i]
        return maximum
    
    def maxProfit(self, prices: List[int]) -> int:
        minimum_so_far = prices[0]
        max_diff = 0
        
        for i in range(len(prices)):
            if minimum_so_far > prices[i]:
                minimum_so_far = prices[i]
            if prices[i] - minimum_so_far > max_diff:
                max_diff = prices[i] - minimum_so_far
        return max_diff