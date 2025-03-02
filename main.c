#include<stdio.h>
#include<stdlib.h>
struct Node{
    int t;
    struct Node* next;
};
struct Node* createHeadNode(){
    struct Node* head=(struct Node*)malloc(sizeof(struct Node));
    head->next=NULL;
    return head;
}
void InsertNode(int data,struct Node* head){
    struct Node* p=head;
    while(p->next!=NULL){
        p=p->next;
    }
    struct Node* q=(struct Node*)malloc(sizeof(struct Node));
    q->t=data;
    q->next=NULL;
    p->next=q;
}
void ShowList(struct Node* head){
    struct Node* p=head->next;
    while(p!=NULL){
        printf("%d\t",p->t);
        p=p->next;
    }
    printf("\n");
}
void DeleteNode(int data,struct Node* head){
    struct Node* p=head;
    while(p->next!=NULL){
        if(p->next->t==data){
            
            free(p->next);
            
            p->next=p->next->next;
            return;
        }
        p=p->next;
    }
    printf("No such data\n");
}
void UpdateNode(struct Node* head,int oldData,int newData){
    struct Node* p=head;
    while(p->next!=NULL){
        if(p->next->t==oldData){
            p->next->t=newData;
            return ;
        }
        p=p->next;
    }
    printf("Node such oldData\n");
}
int main(){
    struct Node* head=createHeadNode();
    InsertNode(1, head);
    InsertNode(2, head);
    InsertNode(3, head);
    InsertNode(4, head);
    InsertNode(5, head);
    ShowList(head);
    DeleteNode(2, head);
    ShowList(head);
    UpdateNode(head,3,8);
    ShowList(head);
    return 0;
}