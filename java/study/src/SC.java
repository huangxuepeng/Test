import java.sql.*;

public class SC {
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
     * 判断一个课程是否存在
     * 本质上还是查询课程信息，只是返回值不同，有则返回true，无则返回false
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
     * 添加、删除 课程
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
        }
        return false;
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
        // close(connection, statement);
    }
}
