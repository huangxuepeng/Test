import java.sql.*;

public class MySQLDemo {
    /**
     * 连接MySQL数据库实例test
     * 这里都需要引入java.sql.···jar包
     * 因为一下三个方法都要连接，所以直接定义一个连接的方法，为了简化代码
     * 
     * @return
     */
    private Statement createStatement() {
        try {
            // 加载驱动
            Class.forName("com.mysql.cj.jdbc.Driver");

            // 连接MySQL数据库，并返回Connection类型的对象
            String url = "jdbc:mysql://127.0.0.1:3306/dbname?useUnicode=true&characterEncoding=UTF-8&userSSL=false&serverTimezone=GMT%2B8";
            Connection connection = DriverManager.getConnection(url, "root", "123456");// 用户名和密码都为root（因计算机有所不同）

            // 创建Statement对象，相当于 Navicat Premium 中 New Query，并返回
            return connection.createStatement();

        } catch (Exception e) {
            e.printStackTrace();
        }

        // 若连接异常，则返回null
        return null;

    }

    /**
     * 添加、删除、修改学生信息
     * 
     * @param sql 要执行的SQL语句
     */
    public boolean update(String sql) {
        Connection connection = null;
        Statement statement = null;
        try {
            // 调用createStatement()方法
            statement = createStatement();
            // 执行SQL语句，并返回影响的行数
            int affect = statement.executeUpdate(sql);
            // 无论怎样都要先执行finally里的语句之后，再执行return;
            return affect > 0;

        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            // 释放资源
            close(connection, statement);
        }
        return false;
    }

    /**
     * 查询学生信息
     * 
     * @param sql
     */
    public void select(String sql) {
        Connection connection = null;
        Statement statement = null;
        ResultSet resultset = null;
        try {

            statement = createStatement();
            // 执行sql语句，并把学生信息存储在resultset资源中
            resultset = statement.executeQuery(sql);
            if (resultset.next()) {// next指向一行数据，执行完后跳到下一行，你可以近似把它看C语言中的“指针”
                String sno = resultset.getString("sno");
                String name = resultset.getString("name");
                String sex = resultset.getString("sex");
                String age = resultset.getString("age");
                String dept = resultset.getString("dept");
                System.out.println("学号：" + sno + ",姓名：" + name + ",性别：" + sex + ",年龄：" + age + ",选课" + dept);
            } else {
                System.out.println("系统异常，无法查询！");
            }

        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            close(resultset, connection, statement);
        }
    }

    /**
     * 判断一个学生是否存在
     * 本质上还是查询学生信息，只是返回值不同，有则返回true，无则返回false
     * 
     * @param sql
     * @return
     */
    public boolean exist(String sql) {
        Connection connection = null;
        Statement statement = null;
        ResultSet resultset = null;
        try {
            statement = createStatement();
            resultset = statement.executeQuery(sql);
            return resultset.next();// 若存在则返回学生信息
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            // 要及时释放资源
            close(resultset, connection, statement);
        }
        return false;// 若不存在则返回false
    }

    /**
     * 定义一个close方法，主要是为了简化代码，
     * 因为上面三个都要释放资源
     * 
     * @param connection
     * @param statement
     */
    private static void close(Connection connection, Statement statement) {
        try {
            if (connection != null) {
                connection.close();
            }
        } catch (SQLException e) {
            e.printStackTrace();
        }

        try {
            if (statement != null) {
                statement.close();
            }
        } catch (SQLException e) {
            e.printStackTrace();
        }
    }

    /**
     * 与上面close()方法属于重载
     * 
     * @param resultset
     * @param connection
     * @param statement
     */
    private static void close(ResultSet resultset, Connection connection, Statement statement) {
        try {
            if (resultset != null) {
                resultset.close();
            }
        } catch (SQLException e) {
            e.printStackTrace();
        }
        close(connection, statement);
    }
}
