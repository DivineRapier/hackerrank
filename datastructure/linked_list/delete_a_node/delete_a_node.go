package main

//https://www.hackerrank.com/challenges/delete-a-node-from-a-linked-list

/*
  Delete Node at a given position in a linked list
  Node is defined as
  struct Node
  {
     int data;
     struct Node *next;
  }
*/
/*
Node* Delete(Node *head, int position)
{
  // Complete this method

  Node* tmp = head;

  if (position == 0) {
      tmp = head->next;
      free(head);
      return tmp;
  }


  while(tmp && position-1 > 0) {
      tmp = tmp->next;
      --position;
  }

  Node* t = tmp->next;
  tmp->next = t->next;
  free(t);
  return head;
}
*/
