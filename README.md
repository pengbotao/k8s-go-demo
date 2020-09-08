docker build -t pengbotao/k8s-go-demo:latest .
docker build -t pengbotao/k8s-go-demo:v1 .
docker build -t pengbotao/k8s-go-demo:v2 .
docker build -t pengbotao/k8s-go-demo:v3 .

docker run -p 38001:38001 --name k8s-go-demo pengbotao/k8s-go-demo
docker run -p 38002:38002 --name k8s-go-demo pengbotao/k8s-go-demo -host 0.0.0.0:38002

docker push pengbotao/k8s-go-demo:v1