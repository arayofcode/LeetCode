

int missingNumber(int* nums, int numsSize){
    /*
      * Methods:
        1. Sort array and find number that isn't present: O(n log n)
        2. int sum = n(n+1) / 2 then subtract all elements from sum. Final result is the required number: O(n)
        3. XOR from 0 to n, and all elements of nums: O(n)
    */
    
    // Method 2
    // int sum = numsSize * (numsSize + 1) / 2;
    // for(int i = 0; i < numsSize; i++){
    //     sum -= nums[i];
    // }
    // return sum;
    
    // Method 3
    int result = numsSize;
    for(int i = 0; i < numsSize; i++){
        result ^= i ^ nums[i];
    }
    return result;
}