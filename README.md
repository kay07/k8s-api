# k8s-api

1.创建deployment

deployment采用结构体和结构体切片进行组合的模式，不断的嵌套，嵌套成一个json格式的对象，然后将json格式对象转换成字节数据然后转成yaml，完成部署。

2.创建service

service采用字节数组直接封装，由于service格式比较固定，所以偷了个懒，封装后将字节数组转成yaml完成部署。

3.样例

封装后的yaml格式如下：

![image](https://user-images.githubusercontent.com/40294040/126454019-1506df67-1671-41d7-b933-797f9f95cf5b.png)

![image](https://user-images.githubusercontent.com/40294040/126454087-c4c2d665-1e82-4d53-bb3c-eecb83336289.png)
