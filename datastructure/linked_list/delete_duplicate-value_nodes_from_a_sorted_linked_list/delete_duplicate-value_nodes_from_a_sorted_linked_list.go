package main

// https://www.hackerrank.com/challenges/delete-duplicate-value-nodes-from-a-sorted-linked-list

/*
  Remove all duplicate elements from a sorted linked list
  Node is defined as
  struct Node
  {
     int data;
     struct Node *next;
  }
*/
/*
Node* RemoveDuplicates(Node *head)
{
  // This is a "method-only" submission.
  // You only need to complete this method.

  if (!head) {
      return head;
  }

  Node* pre = head;
  Node* cur = head->next;

  while(cur) {
      if (pre->data == cur->data){
          pre->next = cur->next;
          free(cur);
          cur = pre->next;
      } else {
          cur = cur->next;
          pre = pre->next;
      }
  }
  return head;
}
*/
