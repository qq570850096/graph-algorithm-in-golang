# golang图论算法
本项目的初衷是通过golang实现图论常用算法，熟悉golang语言的一个开源项目。


[TOC]

## 图的表示方法

在Adj文件夹中保存了三种图的表示方法

1. Matrix：邻接矩阵
2. Table：邻接表
3. Hash：也是邻接表，不过使用了哈希表的存储方式加快速度，本项目中后续算法实现都是基于此结构。