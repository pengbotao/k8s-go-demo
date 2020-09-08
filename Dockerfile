FROM golang:1.14 AS build

WORKDIR /go

COPY . /go/k8s-go-demo/
RUN cd /go/k8s-go-demo/ \
&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o k8s-go-demo .


FROM alpine
COPY --from=build /go/k8s-go-demo/k8s-go-demo /bin/
EXPOSE 38001

CMD [ "-host", "0.0.0.0:38001" ]
ENTRYPOINT [ "k8s-go-demo" ]