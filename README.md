# k8s-api

1.创建deployment

deployment采用结构体和结构体切片进行组合的模式，不断的嵌套，嵌套成一个json格式的对象，然后将json格式对象转换成字节数据然后转成yaml，完成部署。

2.创建service

service采用字节数组直接封装，由于service格式比较固定，所以偷了个懒，封装后将字节数组转成yaml完成部署。

3.样例

封装后的yaml格式如下：

apiVersion: apps/v1
kind: Deployment
metadata:
  name: test
spec:
  replicas: 2
  selector:
    matchLabels:
      app: test
  template:
    metadata:
      labels:
        app: test
    spec:
      containers:
      - env:
        - name: server.port
          value: "40021"
        - name: spring.cloud.nacos.server-addr
          value: 192.168.0.10:3722
        - name: spring.cloud.nacos.discovery.namespace
          value: 673abc66-7092-4eb1-a315-cfccf8eb1aff
        - name: spring.cloud.nacos.config.namespace
          value: 673abc66-7092-4eb1-a315-cfccf8eb1aff
        - name: spring.cloud.nacos.config.file-extension
          value: yaml
        - name: spring.profiles.active
          value: dev
        image: 192.168.0.10:8802/api/mp:dev
        imagePullPolicy: Always
        name: test
        volumeMounts:
        - mountPath: /logs
          name: a
        - mountPath: /etc/localtime
          name: b
      volumes:
      - hostPath:
          path: /home/logs/test
        name: a
      - hostPath:
          path: /etc/localtime
        name: b
---
apiVersion: v1
kind: Service
metadata:
  labels:
    app: test
  name: test
spec:
  ports:
  - name: test
    port: 40021
    targetPort: 40021
  selector:
    app: test
  type: NodePort
