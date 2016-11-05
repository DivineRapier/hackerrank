package linked_list

/*
   Reverse a doubly linked list, input list may also be empty
   Node is defined as
   struct Node
   {
     int data;
     Node *next;
     Node *prev;
   }
*/
// Node* Reverse(Node* head)
// {
//     // Complete this function
//     // Do not write the main method.
//     if (head == NULL || head->next == NULL) {
//         return head;
//     }

//      Node* temp = head;
//      Node* newHead = head;
//      while (temp) {
//         Node* prev = temp->prev;
//         temp->prev = temp->next;
//         temp->next = prev;
//         newHead = temp;
//         temp = temp->prev;
//     }
//     return newHead;
// }
