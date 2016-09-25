package linked_list

/*
  Print elements of a linked list in reverse order as standard output
  head pointer could be NULL as well for empty list
  Node is defined as
  struct Node
  {
     int data;
     struct Node *next;
  }
*/
/*
void ReversePrint(Node *head)
{
  // This is a "method-only" submission.
  // You only need to complete this method.

    if (head == NULL) {
        return ;
    }

    Node* pre = head;
    Node* cur = pre->next;
    pre->next = NULL;
    Node* nxt = NULL;
    while(cur) {
        nxt = cur->next;
        cur->next = pre;
        pre = cur;
        cur = nxt;
    }

    while(pre) {
        printf("%d\n", pre->data);
        pre = pre->next;
    }

}
*/
