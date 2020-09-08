# 创建镜像

```
docker build -t pengbotao/k8s-go-demo:latest .
docker build -t pengbotao/k8s-go-demo:v1 .
docker build -t pengbotao/k8s-go-demo:v2 .
docker build -t pengbotao/k8s-go-demo:v3 .
```

# 推送

```
docker push pengbotao/k8s-go-demo:v1
```

# 启动

```
docker run -p 38001:38001 --name k8s-go-demo pengbotao/k8s-go-demo
docker run -p 38002:38002 --name k8s-go-demo-v1 pengbotao/k8s-go-demo:v1 -host 0.0.0.0:38002
```

# 输出

```
{
    "ClientIP": "172.17.0.1",
    "Host": "84f82c7f28ad",
    "ServerIP": "172.17.0.3",
    "Time": "2020-09-08 04:00:02",
    "Version": "latest"
}
```





