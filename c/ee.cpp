#include<bits/stdc++.h>
using namespace std;

#define MAX 100
int edge[MAX][MAX];  //邻接矩阵
int vexnum;   //顶点数
int visited[MAX]; //标记数组

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
    cout<<v;
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
    queue<int> q;
    cout<<v;
    visited[v] = 1;
    q.push(v);
    while(!q.empty()){
        int u = q.front(); //获取队首元素
        q.pop(); //出队
        for(int w=getFirstNeighbor(u); w!=-1; w=getNextNeighbor(u,w)){
            if(!visited[w]){
                cout<<w;
                visited[w] = 1;
                q.push(w);
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
    cin>>vexnum>>m; //输入顶点数、边数
    for(int i=0; i<m; i++){  //输入有向边的两个顶点以及权值
        int x,y,cost;
        cin>>x>>y>>cost;
        edge[x][y] = cost;
    }
    DFSTraverse();
    cout<<endl;
    BFSTraverse();
    return 0;
}
