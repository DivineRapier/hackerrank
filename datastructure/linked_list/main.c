#include <stdio.h>

typedef struct {
    Node* next;
    int data;
}Node;

int length(Node* head) {
    int len = 0;
    while(head) {
        ++len;
        head = head->next;
    }
    return len;
}

Node* findNextNode(Node* head, int offset) {
    if (offset <1) {
        return head;
    }

    while (offset > 0) {
        head = head->next;
    }
    return head;
}


int FindMergeNode(Node *headA, Node *headB)
{
    // Complete this function
    // Do not write the main method.

    if (!headA || !headB) {
        return -1;
    }

    int lenA = length(headA);
    int lenB = length(headB);

    Node* tmpA = findNextNode(headA, (lenA - lenB));
    Node* tmpB = findNextNode(headB, (lenB - lenA));

    while(tmpA != tmpB) {
        tmpA = tmpA->next;
        tmpB = tmpB->next;
    }
    return tmpA->data;

}

int main() {
    printf("hello\n")
}