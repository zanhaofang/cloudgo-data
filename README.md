# 使用xorm构建数据服务
## 1、任务要求
使用 xorm 或 gorm 实现本文的程序，从编程效率、程序结构、服务性能等角度对比 database/sql 与 orm 实现的异同！

## 2、实验过程
a、启动容器
![image](https://github.com/zanhaofang/cloudgo-data/blob/master/pics/pic1.png)

b、运行mysql

![image](https://github.com/zanhaofang/cloudgo-data/blob/master/pics/pic2.png)

c、创建数据库

![image](https://github.com/zanhaofang/cloudgo-data/blob/master/pics/pic3.png)

d、测试web服务

运行服务器

![image](https://github.com/zanhaofang/cloudgo-data/blob/master/pics/pic4.png)

插入数据
![image]()

查询数据
![image]()

d、性能测试

database/sql性能测试
![image]()

xorm性能测试
![image]()

## 3、比较
xorm通过连写操作，可以通过很少的语句完成数据库操作，编程效率更高。
