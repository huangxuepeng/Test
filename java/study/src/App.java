import java.util.Scanner;

public class App {
    static MySQLDemo db = new MySQLDemo();

    static Course cc = new Course();

    public static void main(String[] args) throws Exception {

        System.out.println("*********************************");
        System.out.println("*\t\t\t\t*");
        System.out.println("*\t欢迎使用学生信息管理系统\t*");
        System.out.println("*\t\t\t\t*");
        System.out.println("*********************************");
        while (true) {
            menu();

        }
    }

    static void menu() {
        System.out.println("1、添加学生信息");
        System.out.println("2、删除学生信息");
        System.out.println("3、修改学生信息");
        System.out.println("4、查询学生信息");
        System.out.println("5、添加课程信息");
        System.out.println("6、删除课程信息");
        System.out.println("请输入操作, 以Enter键结束:");
        Scanner scanner = new Scanner(System.in);
        int option = scanner.nextInt();
        switch (option) {
            // 添加
            case 1: {
                System.out.println("请输入学号：");
                String sno = scanner.next();
                String sql = "select sno from student where sno='" + sno + "'";
                if (db.exist(sql)) {
                    System.out.println("该学号已存在，无法添加！");
                    return;
                }
                System.out.println("请输入姓名：");
                String name = scanner.next();
                System.out.println("请输入性别：");
                String sex = scanner.next();
                System.out.println("请输入年龄：");
                String age = scanner.next();
                System.out.println("请输入课程：");
                String dept = scanner.next();
                sql = "insert into student(sno,name,sex, age, dept) values ('" + sno + "','" + name + "','" + sex
                        + "','" + age + "','" + dept + "')";
                if (db.update(sql)) {
                    System.out.println("添加成功！");
                    return;
                }
                System.out.println("系统异常，添加失败");
                break;
            }
            // 删除
            case 2: {
                System.out.println("请输入要删除学生的学号：");
                String sno = scanner.next();
                String sql = "select sno from student where sno='" + sno + "'";
                if (!db.exist(sql)) {
                    System.out.println("该学号不存在，无法删除！");
                    return;
                }
                sql = "delete from student where sno='" + sno + "'";
                if (db.update(sql)) {
                    System.out.println("删除成功！");
                    return;
                }
                System.out.println("系统异常，删除失败！");
                break;
            }
            // 修改
            case 3: {
                System.out.println("请输入要修改学生的学号：");
                String sno = scanner.next();
                String sql = "select sno from student where sno='" + sno + "'";
                if (!db.exist(sql)) {
                    System.out.println("该学号不存在，无法修改！");
                    return;
                }
                System.out.println("请输入新的年龄：");
                String age = scanner.next();
                sql = "update student set age='" + age + "' where sno='" + sno + "'";
                if (db.update(sql)) {
                    System.out.println("修改成功！");
                } else {
                    System.out.println("系统异常，修改失败！");
                }
                break;
            }
            // 查询
            case 4: {
                System.out.println("请输入要查询学生的学号：");
                String sno = scanner.next();
                String sql = "select sno from student where sno='" + sno + "'";
                if (!db.exist(sql)) {
                    System.out.println("该学号不存在，查询无果！");
                    return;
                }
                sql = "select sno,name,sex, age,dept from student where sno='" + sno + "'";
                db.select(sql);
                break;
            }
            // 添加课程
            case 5: {
                System.out.println("请输入课程号：");
                String cno = scanner.next();
                String sql = "select cno from course where cno='" + cno + "'";
                if (cc.exist(sql)) {
                    System.out.println("该课程已存在，无法添加！");
                    return;
                }
                System.out.println("请输入课程名：");
                String name = scanner.next();
                System.out.println("请输入pno：");
                String pno = scanner.next();
                System.out.println("请输入credit：");
                String credit = scanner.next();
                sql = "insert into course(cno,name,pno, credit) values ('" + cno + "','" + name + "','" + pno
                        + "','" + credit + "')";
                if (cc.update(sql)) {
                    System.out.println("添加成功！");
                    return;
                }
                System.out.println("系统异常，添加失败");
                break;
            }
            // 删除课程
            case 6: {
                System.out.println("请输入要删除课程的课程号：");
                String cno = scanner.next();
                String sql = "select cno from course where cno='" + cno + "'";
                if (!cc.exist(sql)) {
                    System.out.println("该学号不存在，无法删除！");
                    return;
                }
                sql = "delete from course where cno='" + cno + "'";
                if (db.update(sql)) {
                    System.out.println("删除成功！");
                    return;
                }
                System.out.println("系统异常，删除失败！");
                break;
            }
            default:
                System.out.println("I'm Sorry,there is not the " + option + " option,please try again.");
        }

    }
}
