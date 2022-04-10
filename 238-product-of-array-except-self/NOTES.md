For some `nums[i]`, 

![image](https://user-images.githubusercontent.com/49439106/162609258-84197377-d4a7-4c64-9f05-bdf7f91fdc5e.png)


## Approach 1: Using two arrays for having prefix and suffix

Make two arrays `prefix` and `suffix`. 

```python
prefix[0] = 1
prefix[i] = prefix[i-1] * nums[i-1]

suffix[len - 1] = 1
suffix[i] = suffix[i+1] * nums[i+1]

out[i] = prefix[i] * suffix[i]

return out
```

Time = O(n)
Space = O(n)

## Approach 2: Improving previous approach to use O(1) extra space

**Challenge:** Finding Prefix and Suffix in O(1) space

Iterate through nums, keep storing product of all elements until that point. 

**Order:**
1. **Prefix:** Index `0` to `nums.length - 1`
2. **Suffix:** Index `nums.length - 1` to `0`



```python3
product = 1
output = []

for(int i = 0; i < num.length; i++):
  output.append(product)
  product *= nums[i]
  
product = 1
for(int i = num.length - 1; i >= 0; i--):
  output[i] *= product
  product *= nums[i] 
```
