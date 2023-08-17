#include<iostream>
#include<string.h>

using namespace std;

class Person
{
public:
	Person(char *nam,char s,long int num)
	{
		strcpy(name,nam);
		sex=s;
		PersonId=num;
	}
	void Print()
	{
		cout<<"姓名："<<name<<"\n性别："<<sex<<"\n身份证号："<<PersonId<<endl;
	}
private:
	char name[20];
	char sex;
	long int PersonId;
};
class Student:public Person
{
public:
	Student(char *nam,char s,long int num,int nu,char *exp):Person(nam,s,num)
	{
		strcpy(stuexp,exp);
		stunum=nu;
	}
	void Print()
	{
		Person::Print();
		cout<<"学号："<<stunum<<"专业："<<stuexp<<endl;
	}
private:
	char stuexp[20];
	int stunum;
};
int main()
{
	Student s1("小明",'n',15210220000327,120408,"ruanjiangongcheng");
	s1.Print();
	return 0;
}