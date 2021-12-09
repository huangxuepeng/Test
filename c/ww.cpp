/*给出n个学生的考试成绩表，每条记录由学号、姓名和分数和名次组成，设计算法完成下列操作：
设计一个显示对学生信息操作的菜单函数如下所示：
*************************
	1、录入学生基本信息
	2、直接插入排序
	3、冒泡排序
	4、快速排序
	5、简单选择排序
	6、堆排序
	7、2-路归并排序
	8、输出学生信息
	0、退出
*************************
请选择：
算法设计要求：按分数从高到低的顺序进行排序，分数相同的为同一名次。
输入的学生信息存入文件中，每选择一种排序方法，必须从文件中取出数据。*/
#include<iostream>
#define N 20
using namespace std;
typedef struct{
	int num;
	char name[N];
	float score;
	int mc;
}ElemType;
typedef struct{
	ElemType elem[N];
	int length;
}Table;
void Get(Table &S,int n){
	int i;
	freopen("stu.txt","r",stdin);
	for(i=1;i<=n;i++){
		cin>>S.elem[i].num;
		cin>>S.elem[i].name;
		cin>>S.elem[i].score;
	}
	S.length=n;
}
void InsertSort(Table &S){
	//直接插入排序
	int i,j;
	for(i=2;i<=S.length;i++)
		if(S.elem[i].score>S.elem[i-1].score){
			S.elem[0]=S.elem[i];
			S.elem[i]=S.elem[i-1];
			for(j=i-2;S.elem[0].score>S.elem[j].score;j--)
				S.elem[j+1]=S.elem[j];
			S.elem[j+1]=S.elem[0];
		} 
}
void BubbleSort(Table &S){
	//冒泡排序
	int i,j;
	for(i=1;i<=S.length;i++){
		for(j=1;j<=S.length-i;j++){
			if(S.elem[j].score<S.elem[j+1].score){
				S.elem[0]=S.elem[j];
				S.elem[j]=S.elem[j+1];
				S.elem[j+1]=S.elem[0];
			}
		}
	} 
}
 
int Partition(Table &S,int low,int high){
	float pivotkey;
	S.elem[0]=S.elem[low]; //第一个枢纽
	pivotkey=S.elem[low].score;
	while(low<high){
		while(low<high&&S.elem[high].score<=pivotkey) --high;
		S.elem[low]=S.elem[high];
		while(low<high&&S.elem[low].score>=pivotkey) ++low;
		S.elem[high]=S.elem[low];
	} 
	S.elem[low]=S.elem[0];
	return low;
} 
void QSort(Table &S,int low,int high){
	float pivotloc; 
	if(low<high){  //长度大于1 
		pivotloc=Partition(S,low,high);
		QSort(S,low,pivotloc-1);
		QSort(S,pivotloc+1,high);
	}
}
void QuickSort(Table &S){
	//快速排序
	QSort(S,1,S.length); 
}
 
void SelectSort(Table &S){
	//简单选择排序
	int i,j,k;
	for(i=1;i<S.length;i++){
		k=i;
		for(j=i+1;j<=S.length;j++){
			if(S.elem[j].score>S.elem[j-1].score){
				k=j;
			}
		}
		if(i!=k){
			S.elem[0]=S.elem[i];
			S.elem[i]=S.elem[k];
			S.elem[k]=S.elem[0];
		}
	} 
}
 
void HeapAdjust(Table &S,int s,int m){
	//调整S.elem[s]  成为小顶堆 
	int j;
	ElemType rc;
	rc=S.elem[s];
	for(j=2*s;j<=m;j=j*2){
		if(j<m&&S.elem[j].score>S.elem[j+1].score)  ++j;
		if(rc.score<=S.elem[j].score)  break;
		S.elem[s]=S.elem[j];   s=j;
	}
	 S.elem[s]=rc;
} 
void HeapSort(Table &S){
	//堆排序
	int i;
	ElemType t;
	for(i=S.length/2;i>0;i--) 
		HeapAdjust(S,i,S.length);
	for(i=S.length;i>1;i--){
		t=S.elem[i];
		S.elem[i]=S.elem[1];
		S.elem[1]=t;
		HeapAdjust(S,1,i-1);  //去除堆顶元素后重新调整 
	}
}
 
void Merge(ElemType SR[],ElemType TR[],int i,int m,int n){
	//将有序的SR[i...m]和SR[m+1...n]归并为有序的TR[i...n]
	int j,k;
	for(j=m+1,k=i;i<=m&&j<=n;k++){
		if(SR[i].score>SR[j].score)  TR[k]=SR[i++];
		else TR[k]=SR[j++];
	} 
	if(i<=m){
		TR[k++]=SR[i++];
	}
	if(j<=n){
		TR[k++]=SR[j++];
	}
}
void MSort(ElemType SR[],ElemType TR1[],int s,int t){
	int m;
	ElemType TR2[N];
	if(s==t) TR1[s]=SR[s];
	else{
		m=(s+t)/2;
		MSort(SR,TR2,s,m);
		MSort(SR,TR2,m+1,t);
		Merge(TR2,TR1,s,m,t);
	}
}
void MergeSort(Table &S){
	//二路归并排序
	MSort(S.elem,S.elem,1,S.length); 
}
void Mc(Table &S){
	//名次 
	int i;
	for(i=1;i<=S.length;i++){
		S.elem[i].mc=i;
		if(S.elem[i].score==S.elem[i+1].score){
			S.elem[i+1].mc=i; 
			i++;
		}
	}
}
void Output(Table S){
	int i;
	for(i=1;i<=S.length;i++){
		printf("%d  %s  %.2f  %d\n",S.elem[i].num,S.elem[i].name,S.elem[i].score,S.elem[i].mc);
	}
}
int main(){
	int n;
	Table S;
	printf("*************************\n");
	printf("1、录入学生基本信息\n");
	printf("2、直接插入排序\n");
	printf("3、冒泡排序\n");
	printf("4、快速排序\n");
	printf("5、简单选择排序\n");
	printf("6、堆排序\n");
	printf("7、2-路归并排序\n");
	printf("8、输出学生信息\n");
	printf("0、退出\n"); 
	printf("*************************\n");
				printf("输入学生个数:");
				scanf("%d",&n);
				Get(S,n);
				InsertSort(S);
				Get(S,n);
				BubbleSort(S);
				Get(S,n);
				QuickSort(S);
				Get(S,n);
				SelectSort(S);
				Get(S,n);
				HeapSort(S);
				Get(S,n);
				MergeSort(S);
				Mc(S);
				Output(S);
	return 0;	
} 