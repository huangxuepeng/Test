/*利用顺序队列，层次建立二叉树*/
 
#include <iostream>
using namespace std;
#define MAXSIZE 100
 
struct tnode   //树的数据结构
{
	char data;
	tnode *lchild, *rchild;
};
 
struct queue  //队列的数据结构
{
	tnode *data[MAXSIZE];
	int front, rear;
};
 
void creat(queue &q);   //创建一个空队列
void enqueue(queue &q,tnode *t);   //将t入队
tnode *dequeue(queue &q);   //出队，并返回对头元素
bool isempty(queue &q);    //判断队列是否为空
tnode *creatbintree();    //按层次顺序创建一棵二叉树，并返回根节点
void showbintree(tnode *root);  //层次遍历二叉树
 
int main()
{
	tnode *root = NULL;
	root = creatbintree();
	showbintree(root);
	system("pause");
	return 0;
}
 
void creat(queue &q)
{
	q.front = q.rear = 0;
}
 
void enqueue(queue &q, tnode *t)
{
	if ((q.rear + 1) % MAXSIZE == q.front)
	{
		cout << "栈满，不能进栈" << endl;
		return;
	}
	q.rear = (q.rear + 1) % MAXSIZE;
	q.data[q.rear] = t;
}
 
tnode *dequeue(queue &q)
{
	tnode *t;
	q.front = (q.front + 1) % MAXSIZE;
	t= q.data[q.front];	
	return t;
}
 
bool isempty(queue &q)
{
	return (q.front == q.rear);
}
 
tnode *creatbintree()
{
	//1.先将根节点入队，当队列不为空时，循环执行以下操作：
	//2.输入左子树的值，不为空，将其入队
	//3.输入右子树的值，不为空，将其入队
	char a;
	tnode *root;
	queue Q;
	creat(Q);
	cout << "Please enter '@'  to represent the empty node: " << endl;
	cin >> a;
	if (a == '@')    //如果第一个节点为空，就直接返回空树
		return NULL;
	else
	{
		root = new tnode;
		root->data = a;
		enqueue(Q, root);  //根节点入队
	}
	while (!isempty(Q))   //当队列不为空
	{
		//先输入左孩子的值，再输入右孩子的值
		tnode *p= dequeue(Q);
		cin >> a;
		if (a == '@')   //左孩子为空
			p->lchild = NULL;
		else
		{
			p->lchild = new tnode;
			p->lchild->data = a;
			enqueue(Q, p->lchild);  //左孩子入队
		}
		cin >> a;
		if (a == '@')   //右孩子为空
			p->rchild = NULL;
		else
		{
			p->rchild = new tnode;
			p->rchild->data = a;
			enqueue(Q, p->rchild);   //右孩子入队
		}
 
	}
	return root;
}
 
void showbintree(tnode *root)
{
	//1.先将根节点入队，当队列不为空时，循环执行以下操作：
	//2.出队一个元素，访问它
	//3.若左子树不为空，将其入队
	//4.若右子树不为空，将其入队
	queue Q;
	tnode *p;
	creat(Q);
	if (root == NULL)
		return;
	enqueue(Q, root);
	while (!isempty(Q))
	{
		p = dequeue(Q);
		cout << p->data << ' ';
		if(p->lchild)
		  enqueue(Q, p->lchild);
		if(p->rchild)
		  enqueue(Q, p->rchild);
	}
	cout << endl;
}