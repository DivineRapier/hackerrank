package main

// https://www.hackerrank.com/challenges/compare-two-linked-lists

/*
  Compare two linked lists A and B
  Return 1 if they are identical and 0 if they are not.
  Node is defined as
  struct Node
  {
     int data;
     struct Node *next;
  }
*/
/*
int CompareLists(Node *headA, Node* headB)
{
  // This is a "method-only" submission.
  // You only need to complete this method

  int lenA = 0, lenB = 0;
  Node* tmp = headA;
  while(tmp) {
      ++lenA;
      tmp = tmp->next;
  }
  tmp = headB;
  while(tmp) {
      ++lenB;
      tmp = tmp->next;
  }

  if (lenA != lenB) {
      return 0;
  }

  while(headA) {
      if (headA->data != headB->data) {
          return 0;
      }
      headA = headA->next;
      headB = headB->next;
  }

  return 1;
}
*/
