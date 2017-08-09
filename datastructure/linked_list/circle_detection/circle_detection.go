package circle_detection

/*
Detect a cycle in a linked list. Note that the head pointer may be 'NULL' if the list is empty.

A Node is defined as:
    struct Node {
        int data;
        struct Node* next;
    }
*/
/*
bool has_cycle(Node* head) {
    // Complete this function
    // Do not write the main method
    if (head == NULL || head->next == NULL) {
        return false;
    }

    Node* pre = head->next->next;
    Node* cur = head;
    while(pre) {
        if (!pre->next) {
            return false;
        }
        if (pre == cur) {
            return true;
        }
        pre = pre->next->next;
        cur = cur->next;
    }
    return false;
}
*/
