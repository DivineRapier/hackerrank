package insert_a_node_at_the_head_of_a_linked_list

//https://www.hackerrank.com/challenges/insert-a-node-at-the-head-of-a-linked-list

/*
  Insert Node at the begining of a linked list
  Initially head pointer argument could be NULL for empty list
  Node is defined as
  struct Node
  {
     int data;
     struct Node *next;
  }
return back the pointer to the head of the linked list in the below method.
*/
/*
Node* Insert(Node *head,int data)
{
  // Complete this method
  if (!head) {
      head = (Node*)malloc(sizeof(Node));
      head->data = data;
      head->next = NULL;
      return head;
  }

  Node* n = (Node*)malloc(sizeof(Node));
  n->data = data;
  n->next = head;
  return n;

}
*/
