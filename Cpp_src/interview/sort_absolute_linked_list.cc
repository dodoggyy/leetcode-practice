#include <iostream>
using namespace std;

calss Node
{
    Node *next;
    int data;
};

Node *sortList(Node *head)
{
    Node *pre = (head);
    Node *cur = (head)->next;

    while (cur != NULL)
    {
        if (cur->data < pre->data)
        {
            pre->next = cur->next;
            cur->next = head;
            head = cur;
            cur = pre;
        }
        else
        {
            pre = cur;
        }
        cur = cur->next;
    }

    return head;
}

// Utility function to print a linked list
void printList(Node *head)
{
    while (head != NULL)
    {
        cout << head->data;
        if (head->next != NULL)
            cout << " -> ";
        head = head->next;
    }
    cout << endl;
}

// Utility function to insert a node at the
// beginning
void push(Node **head, int data)
{
    Node *newNode = new Node;
    newNode->next = (*head);
    newNode->data = data;
    (*head) = newNode;
}

// Driver code
int main()
{
    Node *head = NULL;
    push(&head, -5);
    push(&head, 5);
    push(&head, 4);
    push(&head, 3);
    push(&head, -2);
    push(&head, 1);
    push(&head, 0);

    cout << "Original list :\n";
    printList(head);

    head = sortList(head);

    cout << "\nSorted list :\n";
    printList(head);

    return 0;
}