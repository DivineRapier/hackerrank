package linked_list

//https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list

/*
  Insert Node at a given position in a linked list
  head can be NULL
  First element in the linked list is at position 0
  Node is defined as
  struct Node
  {
     int data;
     struct Node *next;
  }
*/
/*
Node* InsertNth(Node *head, int data, int position)
{
    Node *newNode = (Node*)malloc(sizeof(Node));
    newNode->data = data;

    if (head == NULL) {
        return newNode;
    }

    if (position == 0) {
       newNode->next = head;
       return newNode;
    }

    Node *currentNode = head;
    while (position - 1 > 0) {
        currentNode = currentNode->next;
        position--;
    }

    newNode->next = currentNode->next;
    currentNode->next = newNode;

    return head;
 }
*/
