#include <stdio.h>
#include <stdlib.h>
struct TreeNode {
	int data;
	struct TreeNode *left;
	struct TreeNode *right;
};

typedef struct queue{
    struct tree* numQ[50];
    int front;
    int rear;
}Queue;
void visit(char n) {
    printf("%c ", n);
}

//创建树
int initTree (struct TreeNode** bt) {
    char ch;
    ch = getchar();
    if (ch == '.') return 0;
    if (ch == '@') *bt = NULL;
    else 
    {
        *bt = (struct TreeNode*)malloc(sizeof(struct TreeNode));
        (*bt) -> data = ch;
        initTree(&((*bt) -> left));
        initTree(&((*bt) -> right));
    }
    return 0;
} 

//先序遍历
void PreOrder(struct TreeNode* root) {
    if (root != NULL) {
        visit(root-> data);
        PreOrder(root -> left);
        PreOrder(root -> right);
    }
}
//中序遍历
void InOrderTree(struct TreeNode* root) {
	if (root == NULL) {
		return;
	}
	InOrderTree(root->left);
	visit(root-> data);
	InOrderTree(root->right);
}

//后序遍历
void PostOrderTree(struct TreeNode* root) {
	if (root == NULL) {
		return;
	}
	PostOrderTree(root->left);
	PostOrderTree(root->right);
    visit(root-> data);
}
//二叉树的最大深度
int maxDepth(struct TreeNode* root) {
	if (root == NULL) {
		return 0;
	}
	else {
		int maxLeft = maxDepth(root->left), maxRight = maxDepth(root->right);
		if (maxLeft > maxRight) {
			return 1 + maxLeft;
		}
		else {
			return 1 + maxRight;
		}
	}
}

Queue Q;
 
void initilize() { //初始化队列
    Q.front = 0;
    Q.rear = 0;
}
 
void Push(struct TreeNode* root) { //入队
    Q.numQ[++Q.rear] = root;
}
 
struct TreeNode* Pop() { //出队
    return Q.numQ[++Q.front];
}
 
int empty() { //判断对列是否为空
    return Q.rear == Q.front;
}
 
struct TreeNode* creatTree (struct TreeNode* root) {
    int value;
    scanf("%d", &value);
    if (value == -1)
        return NULL;
    root = (struct TreeNode*)malloc(sizeof(struct TreeNode));
    root->data = value;
    printf("请输入%d的左子树：", root->data);
    root->left = creatTree(root->left);
    printf("请输入%d的右子树：", root->data);
    root->right = creatTree(root->right);
    return root;
}
 
void LevelOrderTraversal (struct TreeNode* root) { //二叉树的层次遍历
    struct TreeNode* temp;
    Push(root);
    while (!empty()) {
        temp = Pop();
        printf("%d ", temp->data);  //输出队首结点
        if (temp->left)     //把Pop掉的结点的左子结点加入队列
            Push(temp->left);
        if (temp->right)    //把Pop掉的结点的右子结点加入队列
            Push(temp->right);
    }
}



//叶子结点
int LeafNodeNum(struct TreeNode* root) {
	if (root == NULL) {
		return 0;
	}
 
	if (root->left == NULL&&root->right == NULL) {
		return 1;
	}
	else {
		return LeafNodeNum(root->left) + LeafNodeNum(root->right);
	}
}



int main()
{
   struct TreeNode* root = (struct TreeNode*)malloc(sizeof(struct TreeNode*));  //初始化
   int c = initTree(&(root));  

    //前序遍历
   printf("PreOrder traversal: \n");
   PreOrder(root);
   printf("\n");

    //中序遍历
   printf("InOrderTree traversal: \n");
   InOrderTree(root);
   printf("\n");

    //后序遍历
   printf("PostOrderTreer traversal: \n");
   PostOrderTree(root);
   printf("\n");

    printf("LeafNodeNum: %d\n", LeafNodeNum(root));
    printf("maxDepth: %d\n", maxDepth(root));

    
    initilize();  //初始化队列
    LevelOrderTraversal(root);
    return 0;
} 