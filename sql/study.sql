CREATE TABLE S
(
    SNO VARCHAR(10),
    SNAME VARCHAR(20),
    STATU VARCHAR(20),
    CITY  VARCHAR(30)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE P 
(
    PNO VARCHAR(10),
    PNAME VARCHAR(20),
    COLOR VARCHAR(30),
    WEIGHTS INT
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE J
(
    JNO VARCHAR(20),
    JNAME VARCHAR(30),
    CITY VARCHAR(30)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE SPJ
(
   SNO VARCHAR(10),
   PNO VARCHAR(10),
   QTY INT
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE Student 
(
    Sno CHAR(9) PRIMARY KEY,
    Sname CHAR(20) UNIQUE, 
    Ssex CHAR(2),
    Sage SMALLINT, 
    Sdept CHAR(20)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE Course
(
    Cno VARCHAR(10),
    Cname VARCHAR(40),    
    Cpno VARCHAR(10),                
    Ccredit VARCHAR(10)                   
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE SC 
(
    Sno VARCHAR(20),
    Cno VARCHAR(20),
    Grade INT
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE UNIQUE INDEX sdeptin on student (Sdept);
CREATE UNIQUE INDEX cnamein on course (Cname);
CREATE ClUSTER INDEX sage on student();


UPDATE ps 
SET COLOR='红'
WHERE COLOR='黄';

DELETE FROM s
WHERE SNO='S2';
DELETE FROM spjs
WHERE SNO='S2';

INSERT INTO spjs 
(sno, pno, jno, qty)
VALUES ("S2", "J6", "P4", 200);
INSERT INTO sc 
(Sno, Cno, Grade)
VALUES("1", "数学", 50);
INSERT INTO sc 
(Sno, Cno, Grade)
VALUES("2", "语文", 80);

UPDATE sc 
SET Grade=40
WHERE Grade=50;

DELETE FROM sc 
WHERE Grade < 60;
UPDATE sps 
SET Qty = 2000
WHERE sno IN 
(
    SELECT sno 
    FROM sses
    WHERE sname="沃尔玛" 
);

DELETE FROM sses 
WHERE sname="国商";
 create view BT_S(Sno, Sname, Sbirth)
    AS
     Select Sno, Sname, 2000-Sage
    From Student;