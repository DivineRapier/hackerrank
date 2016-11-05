package linked_list

/*
    Insert Node in a doubly sorted linked list
    After each insertion, the list should be sorted
   Node is defined as
   struct Node
   {
     int data;
     Node *next;
     Node *prev;
   }
*/

typedef struct Node
{
	int data;
	struct Node *next;
	struct Node *prev;
}Node;

void frontInsert(Node *node, int data)
{
	Node *prev = node->prev;
	Node *insert = (Node *)malloc(sizeof(Node));
	insert->data = data;
	if (prev != NULL)
	{
		prev->next = insert;
	}
	insert->prev = prev;
	insert->next = node;
	node->prev = insert;
}

void behindInsert(Node *node, int data)
{
	Node *insert = (Node *)malloc(sizeof(Node));
	insert->data = data;
	insert->next = node->next;
	insert->prev = node;
	if (node->next)
	{
		node->next->prev = insert;
	}
	node->next = insert;
}

Node *
SortedInsert(Node *head, int data)
{
	// Complete this function
	// Do not write the main method.
	struct Node *tmpHead = head;
	if (head == NULL)
	{
		head = (Node *)malloc(sizeof(Node));
		head->prev = NULL;
		head->next = NULL;
		head->data = data;
		return head;
	}


	while (tmpHead->data < data && tmpHead->next)
	{
		tmpHead = tmpHead->next;
	}

	behindInsert(tmpHead, data);

	if (tmpHead->data > data)
	{
		int t = tmpHead->next->data;
		tmpHead->next->data = tmpHead->data;
		tmpHead->data = t;
	}

	return head;
}



int main()
{
	Node *head = NULL;
	head = SortedInsert(head, 1);
	head = SortedInsert(head, 3);
	head = SortedInsert(head, 5);
	head = SortedInsert(head, 7);
	head = SortedInsert(head, 9);

	head = SortedInsert(head, 0);
	head = SortedInsert(head, 2);
	head = SortedInsert(head, 4);
	head = SortedInsert(head, 6);
	head = SortedInsert(head, 8);

	while (head) {
		printf("%d, prev = %d, next = %d\n", head->data,
        head->prev?head->prev->data:"-1",
        head->next?head->next->data:"-1");
		head = head->next;
	}
	system("pause");
	return 0;
}