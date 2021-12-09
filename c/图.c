#include<stdio.h>

typedef struct 
{
    int elem[100];
    int front;
    int rear;
} SeqQueue;

//初始化操作
void InitQueue(SeqQueue *Q)
{
    Q -> front = Q -> rear = 0;
}

int IsEmpty(SeqQueue Q) 
{
    if (Q.rear == Q.front) {
        return 0;
    }
    return 1;
}

//入队的操作
void  EnterQueue(SeqQueue *Q, int x)
{
    if ((Q -> rear + 1)%100 == Q -> front) 
    {
         printf("queue have full\n");    //队列已满
         return ;
    }
    Q -> elem[Q -> rear] = x;
    Q -> rear = (Q -> rear + 1) % 100;
   
}

//出队操作
int  DeleteQueue(SeqQueue *Q)
{
    int x;
    if (Q -> front == Q -> rear) 
    {
        printf("Queue is empty.\n");
        return 0;
    }
    x = Q -> elem[Q -> front];
    Q -> front = (Q -> front + 1)% 100;
    return x;

}
int edge[100][100];  //邻接矩阵
int vexnum;   //顶点数
int visited[100]; //标记数组

//返回第一个邻接点
int getFirstNeighbor(int v){
    if(v!=-1){
    	for(int i=0; i<vexnum; i++){
        	if(edge[v][i] > 0){
            	return i;
        	}
   	 	}
   	}
    return -1;
}

//返回下一个邻接点
int getNextNeighbor(int v, int w){
	if(v!=-1 && w!=-1){
    	for(int i=w+1; i<vexnum; i++){
        	if(edge[v][i] > 0){
            	return i;
        	}
    	}
    }
    return -1;
}

//深度优先遍历
void DFS(int v){
    printf("%d", v);
    visited[v] = 1;
    for(int w=getFirstNeighbor(v); w!=-1; w=getNextNeighbor(v,w)){
        if(!visited[w]){
            DFS(w);
        }
    }
}
void DFSTraverse(){
	for(int i=1; i<=vexnum; i++){
		visited[i] = 0;  //初始化访问标记数组
	}
	for(int i=1; i<=vexnum; i++){
		if(!visited[i]){
			DFS(i);
		}
	}
}

//广度优先遍历
void BFS(int v){
     SeqQueue q;
    //初始化队列
    InitQueue(&q);

   printf("%d", v);
    visited[v] = 1;
    EnterQueue(&q, v);
    while(!IsEmpty(q)){
        int u = DeleteQueue(&q); //获取队首元素
        for(int w=getFirstNeighbor(u); w!=-1; w=getNextNeighbor(u,w)){
            if(!visited[w]){
                printf("%d", w);
                visited[w] = 1;
                EnterQueue(&q, w);;
            }
        }
    }
}
void BFSTraverse(){
    for(int i=1; i<=vexnum; i++){
		visited[i] = 0;  //初始化访问标记数组
	}
	for(int i=1; i<=vexnum; i++){
		if(!visited[i]){
			BFS(i);
		}
	}
}

int main(){
    int m; //边数
    scanf("%d %d", &vexnum, &m);  //输入顶点数、边数
    for(int i=0; i<m; i++){  //输入有向边的两个顶点以及权值
        int x,y,cost;
        scanf("%d %d %d", &x, &y, &cost);
        edge[x][y] = cost;
    }
    DFSTraverse();
    printf("\n");
    BFSTraverse();
    return 0;
}
