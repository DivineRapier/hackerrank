package linked_list

//https://www.hackerrank.com/challenges/insert-a-node-at-the-tail-of-a-linked-list

/*


//   Insert Node at the end of a linked list
//   head pointer input could be NULL as well for empty list
//   Node is defined as
  struct Node
  {
     int data;
     struct Node *next;
  }

Node* Insert(Node *head,int data)
{
  // Complete this method
    if (!head) {
        head = (Node*)malloc(sizeof(Node));
        head->data = data;
        head->next = NULL;
        return head;
    }
    Node* tmp = head;
    while(tmp->next) {
        tmp = tmp->next;
    }
    tmp->next = (Node*)malloc(sizeof(Node));
    tmp->next->data = data;
    tmp->next->next = NULL;
    return head;
}



*/
