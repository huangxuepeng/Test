-- CREATE TABLE S
-- (
--     SNO VARCHAR(10),
--     SNAME VARCHAR(20),
--     STATU VARCHAR(20),
--     CITY  VARCHAR(30)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- CREATE TABLE P 
-- (
--     PNO VARCHAR(10),
--     PNAME VARCHAR(20),
--     COLOR VARCHAR(30),
--     WEIGHTS INT
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- CREATE TABLE J
-- (
--     JNO VARCHAR(20),
--     JNAME VARCHAR(30),
--     CITY VARCHAR(30)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- CREATE TABLE SPJ
-- (
--    SNO VARCHAR(10),
--    PNO VARCHAR(10),
--    QTY INT
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- CREATE TABLE Student 
-- (
--     Sno CHAR(9) PRIMARY KEY,
--     Sname CHAR(20) UNIQUE, 
--     Ssex CHAR(2),
--     Sage SMALLINT, 
--     Sdept CHAR(20)
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- CREATE TABLE Course
-- (
--     Cno VARCHAR(10),
--     Cname VARCHAR(40),    
--     Cpno VARCHAR(10),                
--     Ccredit VARCHAR(10)                   
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- CREATE TABLE SC 
-- (
--     Sno VARCHAR(20),
--     Cno VARCHAR(20),
--     Grade INT
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- CREATE UNIQUE INDEX sdeptin on student (Sdept);
-- CREATE UNIQUE INDEX cnamein on course (Cname);
-- CREATE ClUSTER INDEX sage on student();


-- UPDATE ps 
-- SET COLOR='红'
-- WHERE COLOR='黄';

-- DELETE FROM s
-- WHERE SNO='S2';
-- DELETE FROM spjs
-- WHERE SNO='S2';

-- INSERT INTO spjs 
-- (sno, pno, jno, qty)
-- VALUES ("S2", "J6", "P4", 200);
-- INSERT INTO sc 
-- (Sno, Cno, Grade)
-- VALUES("1", "数学", 50);
-- INSERT INTO sc 
-- (Sno, Cno, Grade)
-- VALUES("2", "语文", 80);

-- UPDATE sc 
-- SET Grade=40
-- WHERE Grade=50;

-- DELETE FROM sc 
-- WHERE Grade < 60;
-- UPDATE sps 
-- SET Qty = 2000
-- WHERE sno IN 
-- (
--     SELECT sno 
--     FROM sses
--     WHERE sname="沃尔玛" 
-- );

-- DELETE FROM sses 
-- WHERE sname="国商";
--  create view BT_S(Sno, Sname, Sbirth)
--     AS
--      Select Sno, Sname, 2000-Sage
--     From Student;



SELECT d.*, e.* from dept d left outer join emp e on e.dept_id = d.id;
create TABLE score (
    学号 CHAR(8) primary key,
    姓名 CHAR(6) not null, 
    sqls float check(sqls between 0 and 100), 
    asp float check(asp between 0 and 100), 
    vb float check(vb between 0 and 100)
);
create table 图书信息表 (
    书号 VARCHAR(10) primary key, 
    书名 VARCHAR(20) not null, 
    作者 VARCHAR(20) not null, 
    单价 int null, 
    库存量 int default '10'
);

insert into 图书信息表 VALUES ('1006', '数据库原理', '王', 25, 20);
select 书号, 书名, 单价*库存量 as 总金额 from 图书信息表;
UPDATE 图书信息表 set 库存量 = 38 WHERE 书名 = '数据库原理';
create view  AA as select 书号, 书名, 单价 from 图书信息表;
DELETE from 图书信息表 where;
create procedure xxx as select *from 图书信息表;
select *from 商品表;
select top 2 * from 商品表;
select distinct 商品名, 单价 from 商品表;
select * from 商品表 where 数量 < 10;
select * from 商品表 where 单价 between 1000 and 3000; 
select top 1 * from 商品表 order by 单价 desc;
select 商品名, 单价 from 商品表 order by 单价 desc;
select * from 商品表 where 商品名 like '电%';
INSERT into  商品表 values('007', '电冰箱', 4560, 56);
update 商品表 set 单价 = 280 where 商品名 = '自行车';
delete from 商品表 where ;
create view aa as select 编号, 商品名, 单价 * 0.8 from 商品表;
insert into Card Values('201105', '张三', '软工407');
update borrow set Rdate = Rdate + 3 where Cno in(Select
Cno from Card where Class ='网络401');
Delete from borrow where Quantity = 0;
Create view zm as select *from Book where Bname = '赵敏';
select Bno, Bname, Author, Price, Quantity from Book where Bname = '网络编程';
select Bno, Bname, Author from Book where Bname like '%安全%';
select Name, Class from Card where Cno in(
    select Borrow.Cno from Borrow, Book where Book.Bno = Borrow.Bno and Book.Bname = 'llll';
)
Select Cno, Count(Bno) as '借书的数量' from Borrow group by Cno having Count(Bno) > 5;
SELECT Cno, Name, Class from Card where Cno in (
    Select Cno from Borrow where Rdate < 20190623;

)
SELECT Card.Cno, Name, Class from Card, Borrow Where Borrow.Cno = Card.Cno AND Rdate < '201999';
select Cno, Rdate from Borrow, Book where Book.Bno = Borrow.Bno AND Book.Price > 100 order by cno desc;

create table score (
    学号 CHAR(8) primary key,
    姓名 CHAR(6) not null,
    SQL float check(sql between 0 and 100),
    asp float check(asp between 0 and 100),
    vb float  check(vb between 0 and 100)
);
select * from score;
select 姓名,sql from score order by sql desc;
select * from score where 姓名 like '王%'; 
select 学号, 姓名, sql+asp+vb as 总分 from score;
insert into score values(1005,'赵强', 64, 82, 69);
update score set vb = 85 where 姓名 = '王英';
create view xsl as select 学号, 姓名, sql from score;
select avg(sql), avg(asp), avg(vb) from score;
delete from score where 姓名 like '王%';
drop database test;

create table 图书 (
    书号 VARCHAR(10) primary key,
    书名 varchar(20) not null,
    作者 varchar(20) not null,
    单价 int null,
    库存量 int default '10'
);
insert into 图书 values('','');
select * from 图书;
select max(单价) as 单价的最大值 from 图书;
select * from 图书 where 单价 > 30 and 库存量 < 10;
select 书号, 书名, 单价*库存量 as 总金额 from 图书;
update 图书 set 库存量 = 38 from 图书 where 书名 = '多媒体技术';
create view AA as select 书名, 书号, 单价 from 图书;
delete from 图书 where 作者 like '张%';
drop table 图书;
create procedure xxx as select count(*) from 图书 where 出版社 = '高等教育出版社';
execute xxx;

select * from 商品;
select top 2 * from 商品表;   
select distinct 商品名, 单价 from 商品表;
select top 1 商品 from 商品表 order by 单价 desc;
select 商品名, 单价 from 商品表 order by 单价 desc;
select * from 商品表 where 商品名 like '电%';