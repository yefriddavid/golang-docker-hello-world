```
docker run --rm -v "$PWD":/app -w /app golang:1.8 ./app
docker run --rm -v "$PWD":/app -w /app golang:1.8 go build -v



docker run --rm -e GOPATH=/home/go -v "$PWD/libs":/home/go -v "$PWD":/app -w /app golang:1.8 go get -u github.com/gorilla/mux
docker run --rm -p 1528:1528 -v "$PWD":/app -w /app golang:1.8 ./app
```

```
http://gorm.io/docs/


```


