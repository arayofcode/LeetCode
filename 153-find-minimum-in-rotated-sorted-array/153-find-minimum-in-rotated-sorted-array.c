

int findMin(int* nums, int numsSize){
    // Binary search
    int left = 0, right = numsSize - 1;
    while(left < right){
        int mid = (left + right) / 2;
        /*
          If nums[mid] > nums[right], it means this part of array is sorted already, and you won't find the solution here. 
        */
        if(nums[mid] > nums[right])
            left = mid + 1;
        else
            right = mid;
    }
    return nums[left];
}