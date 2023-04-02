#include <stdio.h>
#include <stdlib.h>

struct Node {
  int data;
  struct Node* next;
};

struct LinkedList {
  struct Node* head;
};

void insert(struct LinkedList* list, struct Node* newNode)
{
  if(list->head == NULL || list->head->data == 0){
    list->head = newNode;
  } else {
    struct Node* pointer = list->head;
    while(pointer->next != NULL){
      pointer = pointer->next;
    }
    pointer->next = newNode;
  }
}

void show(struct LinkedList* list){
  struct Node* pointer = list->head;
  while(pointer != NULL){
    printf("%d -> ", pointer->data);
    pointer = pointer->next;
  }
  printf("NULL\n");
};

int main(){
  printf("LinkedLists in C\n");
  struct LinkedList list;
  list.head = (struct Node*)malloc(sizeof(struct Node));
  struct Node* head = (struct Node*)malloc(sizeof(struct Node));
  struct Node* second = (struct Node*)malloc(sizeof(struct Node));
  struct Node* third = (struct Node*)malloc(sizeof(struct Node));
  head->data = 22;
  second->data = 30;
  third->data = 57;
  insert(&list, head);
  insert(&list, second);
  insert(&list, third);
  show(&list);
  
  return 0;
}