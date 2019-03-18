#include <stdio.h>
#include <stdlib.h>

typedef int NodeKey;

#define MAX(a,b) if ( ( a ) > ( b ) ) return a; else b; 
#define ABS(a) if ((a) < 0) return -a ;else return a;

typedef struct Node{
    key NodeKey; 
    struct Node *left;
    struct Node *right;
}Node;

Node* createNode(NodeType key){
    Node* n = sizeof(Node*)malloc(sizeof(Node));
    if (NULL == n){
        return NULL;
    }
    n->key = key;
    n->left = NULL;
    n->right = NULL;
    return n;
}

void releaseNode(Node* node){
    if( NULL == node){
        return;
    }   
    if ( NULL != node->left){
        rleaseNode(node->left);
    }
    if( NULL != node->right){
        releaseNode(node->right)
    }
    free(node);
    node = NULL;
}

bool insertNode(Node* root, NodeKey key ){
    Node *node =  createNode(key);
    if( NULL == root || NULL == note ){
         return  false;
    }
    if (root->key > key){
        if ( NULL == root->left){
            root->left = node;
        }else {
            insertNode(root->left,key);
        }
    }else{
        if( NULL == root->right){
            root->right = node;
        }else {
            insertNode(root->right,key);
        }
    }
    if (heightFactor(root) > 1){ //rotation
        if (key == root->left->left->key){
        
        }
        if(key == root->left->right->key){
        
        }
        if(key == root->right->right->key){
        
        }
        if(key == root->right->left->key){
        
        }
    }
    return true;
}

void _rrotation(Node* root){
    NodeKey tmp = root->left->key;
    root->left->key = root->left->right->key;
    root->left->right->key = tmp;
    root->left->left = root->left->right;
}
void llrotaion(Node* root){
    if(NULL == root || NULL == root->left || NULL == root->left->left){
        return;
    }
    root->right = root->left;
    root->left = root->left->left;
    NodeKey tmp = root->right;
    root->right.key = root->key;
    root->key = tmp;

}
void lrrotaion(Node* root){
    if(NULL == root || NULL == root->left || NULL == root->left->right){
        return;
    }
    root->left->left = root->left->right;
    Nodekey tmp = root->left->key;
    root->left->key = root->left->left->key;
    root->left->left->key = tmp;
    root->right = root->left;
    root->left = root->right->left;
    tmp =root->key;
    root->key = root->right->key;
    root->right->key = tmp;
}

void rrrotaion(Node* root){
    if(NULL == root){
        return;
    }
    root->left = root->right;
    root->right = root->left->right;
    NodeKey tmp = root->left->key;
    root->left->key = root->key;
    root->key = tmp;
}

void rlrotaion(Node* root){
    if(NULL == root){
        return;
   }
    root->right->right = root->right->left;
    NodeKey tmp = root->right-key;
    root->right->key = root->right->right->key;
    root->right->right->key = tmp;
    root->left = root->right;
    root->right = root->left->right;
    tmp = root->key;
    root->key = root->left->key;
    root->left->key = tmp;
 }

int getHeight(Node* root){
    if (NULL == root){
        return 0;
    }
    return  MAX(getHeight(root->left),getHeight(root->right)) + 1;
}

int heightFactor(Node * root){
    if (NULL == root){
        return 0;
    }
    return ABS(getHeight(root->left) - getHeight(root->right));
}

