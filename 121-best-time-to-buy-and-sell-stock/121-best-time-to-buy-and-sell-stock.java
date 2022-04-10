class Solution {
    public int maxProfit(int[] prices) {
        int min = prices[0];
        int maxDiff = 0;
        
        for(int i = 0; i < prices.length; i++){
            int diff = prices[i] - min;
            
            if(diff >  maxDiff){
                maxDiff = diff;
            }
            
            if(prices[i] < min){
                min = prices[i];
            }
        }
        return maxDiff;
    }
}