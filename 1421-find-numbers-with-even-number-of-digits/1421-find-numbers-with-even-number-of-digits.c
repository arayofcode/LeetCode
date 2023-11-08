

int findNumbers(int* nums, int numsSize){
    int count = 0;
    for(int i = 0; i < numsSize; i++){
        int copy = nums[i];
        int digits = 0;
        while(copy != 0){
            digits += 1;
            copy /= 10;
        }
        if(digits % 2 == 0)
            count++;
    }
    return count;
}