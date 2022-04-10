Finding maximum difference in an array where the larger number comes after smaller number. 

## Approach 1: Brute Force- O(n^2) time, O(1) space

Use nested loops to find difference

```python3
maxDiff = 0

for(i = 0; i < len - 1; i++):
  for(j = i+1; j < len; j++):
    if(price[j] - price[i] > maxDiff):
      maxDiff = price[j] - price[i]

return maxDiff
```

## Approach 2: Linear time, constant space

**Note:** The maximum difference should involve minimum element found in array until that point

```python3
min = prices[0]
maxDiff = 0

for(int i = 0; i < len; i++):
  # Difference from minimum found until now
  diff = prices[i] - min
  
  # Update maxDiff and minimum price
  if(diff > maxDiff):
    maxDiff = diff
   
  if(prices[i] < min):
    min = prices[i]

return maxDiff
```
