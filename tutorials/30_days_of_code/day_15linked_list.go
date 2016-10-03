package tutorials

/*


Node* insert(Node *head,int data);
Node* getlast(Node* head);

Node* insert(Node *head,int data)
{
    Node* last = NULL, * inc = NULL;
    if (head == NULL) {
        head = (Node*)malloc(sizeof(struct Node));
        head->data = data;
        head->next = NULL;
        goto end;
    }

    last = getlast(head);
    inc = (Node*)malloc(sizeof(struct Node));
    last->next = inc;
    inc->data = data;
    inc->next = NULL;
    
end:
    return head;
}

Node* getlast(Node* head) {
    while(head->next) {
        head = head->next;
    }
    return head;
}


*/