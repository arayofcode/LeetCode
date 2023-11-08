/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2){
    struct ListNode* curr1 = l1;
    struct ListNode* curr2 = l2;
    struct ListNode* head = (struct ListNode*) malloc(sizeof(struct ListNode));
    head -> val = 0;
    head -> next = NULL;
    struct ListNode* curr = head;
    int carry = 0;
    while(curr1 != NULL || curr2 != NULL || carry == 1) {
        int sum = ((curr1 != 0) ? curr1 -> val  : 0) + ((curr2 != 0) ? curr2 -> val : 0) + carry;
        carry = sum / 10;
        curr -> val = sum % 10;
        if(curr1) {
            curr1 = curr1 -> next;
        }
        if(curr2) {
            curr2 = curr2 -> next;
        }
        // Create new node only if needed
        if(curr1 || curr2 || carry) {
            curr -> next = malloc(sizeof(struct ListNode));
            curr -> next -> val = 0;
            curr -> next -> next = NULL;
        }
        curr = curr -> next;
    }
    return head;
}