package main

/*
  Reverse a linked list and return pointer to the head
  The input list will have at least one element
  Node is defined as
  struct Node
  {
     int data;
     struct Node *next;
  }
*/
/*
Node* Reverse(Node *head)
{
  // Complete this method
  if (head == NULL) {
      return head;
  }

  Node* pre = head;
  Node* cur = head->next;
  pre->next = NULL;

  while(cur) {
      nxt = cur->next;
      cur->next = pre;
      pre = cur;
      cur = nxt;
  }

  return pre;


}
*/
