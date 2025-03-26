#include<stdio.h>
#include<stdlib.h>
struct Node{
    int data;
    int height;//高度从下往上算
    struct Node* p[2];
};
typedef struct Node Node;
int max(int a,int b){
    if(a>b)
    return a;
    return b;
}
int Height(Node* p){
    if(p==NULL)
    return 0;
    return p->height;
}
Node* SingleSpin(Node* root,int id){   //id指出root下的不平衡子节点
    Node* position=root->p[id];
    root->p[id]=position->p[1-id];
    position->p[1-id]=root;
    root->height=max(Height(root->p[0]),Height(root->p[1]))+1;
    position->height=max(Height(position->p[id]),root->height)+1;//position原来的子节点高度未变
    return position;
}
Node* DoubleSpin(Node* root,int id){   //id指出root下的不平衡子节点
    root->p[id]=SingleSpin(root->p[id],1-id);
    return SingleSpin(root,id);
}
Node* Insert(Node* root, int data) {
    if (root == NULL) { // 如果当前节点为空，创建新节点
        Node* p = (Node*)malloc(sizeof(Node));
        if (p == NULL) {
            printf("error\n");
            return NULL;
        }
        p->data = data;
        p->height = 0; // 初始化高度为0
        p->p[0] = p->p[1] = NULL; 
        return p;
    }
    if (data < root->data) {
        root->p[0] = Insert(root->p[0], data);
    } else if (data > root->data) {
        root->p[1] = Insert(root->p[1], data);
    } else {
        // 数据已存在
        return root;
    }
    root->height = max(Height(root->p[0]), Height(root->p[1])) + 1;
    int balance = Height(root->p[1]) - Height(root->p[0]);
    if (balance > 1) { 
        if (data > root->p[1]->data) {
            root = SingleSpin(root, 1); 
        } else {
            root = DoubleSpin(root, 1);
        }
    } else if (balance < -1) { 
        if (data < root->p[0]->data) {
            root = SingleSpin(root, 0); 
        } else {
            root = DoubleSpin(root, 0); 
        }
    }
    return root;
}
Node* Find(Node* root,int data){
    if(root==NULL){
        printf("error\n");
        return NULL;
    }
    if(data==root->data){
        return root;
    }else if(data<root->data){
        return Find(root->p[0],data);
    }else{
        return Find(root->p[1],data);
    }
}
Node* Delete(Node* root, int data) {
    if (root == NULL) {
        printf("error\n");
        return NULL;
    }
    if (data < root->data) {
        root->p[0] = Delete(root->p[0], data);
    } else if (data > root->data) {
        root->p[1] = Delete(root->p[1], data);
    } else {
        if (root->p[0] == NULL || root->p[1] == NULL) {
            Node* temp = root->p[0] ? root->p[0] : root->p[1];
            if (temp == NULL) {
                temp = root;
                root = NULL;
            } else {
                *root = *temp; // 替换节点内容
            }
            free(temp);
        } else {
            Node* successor = root->p[1];
            while (successor->p[0] != NULL) {
                successor = successor->p[0];
            }
            root->data = successor->data;
            root->p[1] = Delete(root->p[1], successor->data);
        }
    }
    if (root == NULL) {
        return root;
    }
    root->height = max(Height(root->p[0]), Height(root->p[1])) + 1;
    int balance = Height(root->p[1]) - Height(root->p[0]);
    if (balance > 1) { 
        if (Height(root->p[1]->p[0]) > Height(root->p[1]->p[1])) {
            root = DoubleSpin(root, 1);
        } else {
            root = SingleSpin(root, 1);
        }
    } else if (balance < -1) { 
        if (Height(root->p[0]->p[1]) > Height(root->p[0]->p[0])) {
            root = DoubleSpin(root, 0);
        } else {
            root = SingleSpin(root, 0);
        }
    }
    return root;
}