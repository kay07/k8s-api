FROM golang

RUN git config --global http.sslVerify false \
    && go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct

workdir /app
add . .
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o main main.go

from scratch
copy --from=0 /app/main /
copy --from=0 /app/.env /.env
cmd ["/main"]