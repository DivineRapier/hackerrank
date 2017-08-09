package merge_two_linked_lists

// https://www.hackerrank.com/challenges/merge-two-sorted-linked-lists

/*
  Merge two sorted lists A and B as one linked list
  Node is defined as
  struct Node
  {
     int data;
     struct Node *next;
  }
*/
/*
Node* MergeLists(Node *headA, Node* headB)
{
  // This is a "method-only" submission.
  // You only need to complete this method

    if (!headA) {
        return headB;
    }
    if (!headB) {
        return headA;
    }

    Node* start, *tmp ;
    if (headA->data < headB->data) {
        start = headA;
        headA = headA->next;
    } else {
        start = headB;
        headB = headB->next;
    }
    tmp = start;
    while(headA || headB) {
        if (headA && !headB) {
            tmp->next = headA;
            break;
        } else if (!headA && headB) {
            tmp->next = headB;
            break;
        }else {
            if (headA->data < headB->data) {
                tmp->next = headA;
                headA = headA->next;
            } else {
                tmp->next = headB;
                headB = headB->next;
            }
            tmp = tmp->next;
        }
    }
    return start;
}
*/
