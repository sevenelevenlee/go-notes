### markdown 语法
*woshi*
---
----
***
****
图片
![图片alt](图片地址 "图片title")

[简书](http://jianshu.com)

```
列表语法
-
+
*

1. wo
   - ai
   + ni
2. ni

```

*golang fmt.Printf 显示类型和值的小技巧*
```$xslt
//如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
fmt.Printf("%+v\n", p) // {x:1 y:2}

//%#v 形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。
fmt.Printf("%#v\n", p) // main.point{x:1, y:2}
```
*golang工程中的文件目录与package*
```
1. 习惯上文件目录的名字与package名字保持一致，
2. 一个文件目录下只能有一个package
3. 如若不一致，import的时候要写目录名，引用的时候要写包名
```
*golang方法指针接收和值接收的区别*
```$xslt
1. 值接收者声明的方法，调用时使用改值的一个副本去执行
2. 指针接收者声明的方法，引用地址->指向接收者的值，可以修改该值
3. 在使用时，值类型的接收者也可以使用指针类型的调用
    原因：对指针进行了转换(*t).method()
```
*声明指针不会开辟内存地址,只是准备要指向内存某个空间,而声明变量会开辟内存地址,准备存放内容.所以指针类型变量都是把一个变量的地址赋值给指针变量*