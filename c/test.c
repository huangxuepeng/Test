#include <stdio.h>
#include <string.h>
int main(){
    struct{
        char *name;  //姓名
        int num;  //学号
        int age;  //年龄
        char group;  //所在小组
        float score;  //成绩
    } stu1;
    //给结构体成员赋值
   scanf("%s%d%d%c%lf", stu1.name, &stu1.num, &stu1.age, &stu1.group, &stu1.score);
    //读取结构体成员的值
    printf("%s的学号是%d，年龄是%d，在%c组，今年的成绩是%.1f！\n", stu1.name, stu1.num, stu1.age, stu1.group, stu1.score);
    return 0;
}