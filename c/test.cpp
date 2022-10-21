// pubuse. h 是公共使用的常量定义和系统函数调用声明。
#include<string.h>
#include<ctype.h>
#include<malloc.h> /* malloc()等*/
#include<limits.h> /* INT_MAX 等*/
#include<stdio.h> /* EOF(=^Z 或F6),NULL */
#include<stdlib.h> /* atoi() */
#include<io.h> /* eof() */
#include<math.h> /* floor(),ceil(),abs() */
#include<process.h> /* exit() */
/* 函数结果状态代码*/
#define TRUE 1
#define FALSE 0
#define OK 1
#define ERROR 0
#define INFEASIBLE -1
typedef int Status; /* Status 是函数的类型,其值是函数结果状态代码，如OK 等*/
typedef int Boolean; /* Boolean 是布尔类型,其值是TRUE 或FALSE */ //文件LinkQueue.h 中实现单链队列－－队列的链式存储结构的表示。
typedef struct QNode
{
QElemType data;
struct QNode *next;
}QNode,*QueuePtr;
typedef struct
{
QueuePtr front,rear; /* 队头、队尾指针*/
}LinkQueue;

// (2)文件LinkQueueAlgo.h 中实现的链队列的基本算法，其存储结构由LinkQueueDef.h 定义。
Status InitQueue(LinkQueue &Q)
{ /* 构造一个空队列Q */
Q.front=Q.rear=(QueuePtr)malloc(sizeof(QNode));
if(!Q.front)
exit(OVERFLOW);
Q.front->next=NULL;
return OK;
}
int QueueLength(LinkQueue Q)
{ /* 求队列的长度*/
int i=0;
QueuePtr  p;
p=Q.front;
while(Q.rear!=p)
{
i++;
p=p->next;
}
return i;
}
Status EnQueue(LinkQueue &Q,QElemType e)
{ /* 插入元素e 为Q 的新的队尾元素*/
QueuePtr p=(QueuePtr)malloc(sizeof(QNode));
if(!p) /* 存储分配失败*/
exit(OVERFLOW);
p->data=e;
p->next=NULL;
Q.rear->next=p;
Q.rear=p;
return OK;
}
Status DeQueue(LinkQueue &Q,QElemType &e)
{ /* 若队列不空,删除Q 的队头元素,用e 返回其值,并返回OK,否则返回ERROR */
QueuePtr p;
if(Q.front==Q.rear)
return ERROR;
p=Q.front->next;
e= p->data;
Q.front->next=p->next;
if(Q.rear==p)
Q.rear=Q.front;
free(p);
return OK;
}
// (3)文件LinkQueueUse.cpp 中包含检验LinkQueueAlgo.h 中关于链式队列基本操作的声明、测试数据和主函数。
// #include "pubuse.h" /* 与实验一的意义相同*/
// typedef  int QElemType; /* 假设链式队列中的结点是一组整数*/
// #include "linkqueue.h"
// #include "linkqueuealgo.h"
// void main()
// {   //编写主函数，完成建立队列、进队、出队、求队列长度的函数的调用测试。
// LinkQueue  Q;
//  InitQueue(&Q);
// QElemType e;
// printf("请输入入队元素：")；
// While(e!=0)
//   {scanf("%d",&e);
//    EnQueue(&Q, e);
//   }
// printf("输出队列长度为：%d/n", QueueLength(Q));
// printf("输出队列元素/n");
// while(int i<= QueueLength(Q))
//  { DeQueue(&Q, &e)
//   printf("%d",e);



// }