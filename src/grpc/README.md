# test_grpc-communication

## Quick Start

### 前提条件

- Dockerがインストールされていること

### 1. リポジトリのクローン

```bash
git clone git@github.com:<username>/test.grpc-communication.git test_grpc-communication
cd test_grpc-communication
```

### 2. コンテナの起動

```bash
docker compose up -d
```

### 3. コンテナにログイン

```bash
docker compose exec -it src bash
```

### 4. サーバーの起動

```bash
cd /go/src/cmd/server
go run main.go
```

### 5. クライアントの起動

```bash
cd /go/src/cmd/client
go run main.go
```
