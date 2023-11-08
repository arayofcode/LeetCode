int maxAreaBruteForce(int* height, int heightSize){
    // Brute Force solution: Find all possible pair of lines and calculate area between them. Return the one with largest size
    int maxCalculatedArea = 0;
    for(int i = 0; i < heightSize - 1; i++) {
        for(int j = i + 1; j < heightSize; j++) {
            int currentArea = 0;
            if(height[i] < height[j]){
                currentArea = (j - i) * height[i];
            }
            else {
                currentArea = (j - i) * height[j];
            }
            if(currentArea > maxCalculatedArea){
                maxCalculatedArea = currentArea;
            }
        }
    }
    return maxCalculatedArea;
}

int maxArea(int* height, int heightSize){
    // Max area is max length and max height
    // We use two points, one at the leftmost side and other rightmost
    int left = 0;
    int right = heightSize - 1;
    int maxArea = (right - left) * ((height[left] < height[right]) ? height[left] : height[right]);
    while(left <= right) {
        int currentArea = (right - left) * ((height[left] < height[right]) ? height[left] : height[right]);
        if(currentArea > maxArea) {
            maxArea = currentArea;
        }
        if(height[left] < height[right]) {
            left++;
        }
        else {
            right--;
        }
    }
    // printf("%d %d\n%d %d", left, height[left], right, height[right]);
    return maxArea;
}