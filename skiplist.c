#include <stdbool.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <time.h>

#define MAX_SL_LEVEL 10

typedef struct Node{
    int key;
    struct Node** forward;
}Node;

typedef struct SkipList{
    Node* head;
    int level;
}SkipList;

Node* newNode(int key,int level){
    Node* node = (Node*)malloc(sizeof(Node));
    node->key =  key;
    node->forward = (Node**)malloc(sizeof(Node*) * (level+1));
    memset(node->forward,0,sizeof(Node*) * level);
    return node;
}

void releaseNode(Node* node){
    if (NULL != node){
        free(node->forward);
        free(node);
        node = NULL;
    }
}

struct SkipList* createList(){
    SkipList* list = (SkipList*)malloc(sizeof(SkipList));
    if (NULL ==list) {
        return list;
    }
    list->head = newNode(-1,MAX_SL_LEVEL-1);
    list->level = 0;
    return list;
}


void releaseList(SkipList *list){
    if( NULL == list)
        return;
    Node *node = list->head;
    Node *fd;
    while( NULL != node && NULL != node->forward[0]){
        fd = node->forward[0];    
        releaseNode(node);
        node = fd;
    }
    free(list);
}

int randomLevel(){
    int level = 0;
    int rd = rand();
    while (rd%2&& level < MAX_SL_LEVEL){
        rd = rand();
        level ++;
    }
    return level;
}

void displayList(SkipList *skl) 
{ 
    int i = 0;
    printf("\n*****Skip List*****\n");
    for (i =skl->level;i>=0;i--) { 
        Node *node = skl->head->forward[i]; 
        printf("Level %d:\n",i);
        while (node != NULL) { 
            printf("%d ",node->key);
            node = node->forward[0]; 
        } 
        printf("\n");
    } 
};

void insertKey(SkipList* list,int key){
    int i  = 0;
    Node *updates[MAX_SL_LEVEL];
    Node *cur = list->head;
    for (i = list->level; i >=0; i--){
        while(NULL != cur->forward[i] && cur->forward[i]->key < key){
            cur = cur->forward[i];
        }
        updates[i] = cur;
    }
    int rlevel = randomLevel();
    Node* nd = newNode(key,rlevel); 
    if (rlevel > list->level) {
        for (i = list->level + 1;i <= rlevel;i++){
            updates[i] = list->head;
        } 
        list->level = rlevel;
    }
    for (i = 0;i <= rlevel;i++){
        nd->forward[i] = updates[i]->forward[i];
        updates[i]->forward[i] = nd;
    }
}

bool deleteKey(SkipList* list,int key){
    int i  = 0;
    Node *updates[MAX_SL_LEVEL];
    Node *cur = list->head;
    int rlevel = -1;
    Node* delN = NULL;
    for (i = list->level; i >=0; i--){
        while(NULL != cur->forward[i] && cur->forward[i]->key < key){
            cur = cur->forward[i];
        }
        if(NULL != cur->forward[i]&& cur->forward[i]->key == key){
            if(rlevel == -1){
                rlevel = i;
                delN = cur->forward[i];
            }
            updates[i] = cur;
        }
    }
    if(rlevel == -1){
        return false;
    }
    for (i = 0;i<=rlevel;i++){
        updates[i]->forward[i]= delN->forward[i];    
    }
    releaseNode(delN);
    return true;
}

bool searchKey(SkipList* list,int key){
    int i  = 0;
    Node *updates[MAX_SL_LEVEL];
    Node *cur = list->head;
    Node* keyN = NULL;
    for (i = list->level; i >=0; i--){
        while(NULL != cur->forward[i] && cur->forward[i]->key < key){
            cur = cur->forward[i];
        }
        if(NULL != cur->forward[i]&& cur->forward[i]->key == key){
            return true;
        }
    }
    return false;
}

int main(){
    srand(time(NULL));
    SkipList *skl =  createList();
    insertKey(skl,3);
    insertKey(skl,1);
    insertKey(skl,33);
    insertKey(skl,23);
    insertKey(skl,54);
    insertKey(skl,273);
    insertKey(skl,37);
    insertKey(skl,62);
    insertKey(skl,18);
    insertKey(skl,237);
    insertKey(skl,326);
    insertKey(skl,3264);
    insertKey(skl,46202164);
    insertKey(skl,46372);
    insertKey(skl,5736);
    insertKey(skl,126294);
    insertKey(skl,476);
    displayList(skl);
    if(searchKey(skl,126294)){
        printf("find key\n");
    }
    deleteKey(skl,126294);
    if(searchKey(skl,126294) == false){
        printf("can't find key\n");
    }
    displayList(skl);
    releaseList(skl);
}

