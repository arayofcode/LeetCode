

int getSum(int a, int b){
    /*
    * XOR for add
    * AND for carry
    * Left shift carry
    * Algorithm:
        1. Find the carries (AND)
        2. Do the addition (XOR)
        3. Do addition with result of 2 and 1 until result of 1 is 0
    */
    while(b){
        unsigned int carry = (a & b);
        a = a ^ b;
        b = carry << 1;
    }
    return a;
}