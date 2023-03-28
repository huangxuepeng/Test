#include <iostream>
#include <cmath>
 
using namespace std;

#define p 3.14

 
// 给三角形边长，求三角形面积
inline double area(double a, double b, double c, double s) {
    //海伦公式求三角形面积
    
    return sqrt(s * (s-a) * (s-b) * (s-c));
}
 
// 给梯形的上下底和高, 求面积
inline double area(double on,double down, double height) {
    return (on + down)*height/2;
}
 
// 给圆半径，求圆面积
inline double area(double r) {
    return 2 * p * r * r;
}
 
int main() {
    cout<<"边长为 3, 4, 5 的三角形面积为 "<<area(3,4,5,6)<<endl
        <<"上下底分别是4,5, 高为3的梯形面积为"<<area(4,5,3)<<endl
        <<"半径为 2 的圆形面积为 "<<area(2)<<endl;
 
    return 0;
}

// #include <iostream>
// using namespace std;
// inline int cube(int);
// int main(){
//     for(int i=1;i<=10;i++){
//         int p=cube(i);
//         cout<<i<<'*'<<i<<'*'<<i<<'='<<p<<endl;
//     }
// }
// inline int cube(int n){
//     return n*n*n;
// }