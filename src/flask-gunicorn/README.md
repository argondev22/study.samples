# Flask-Gunicorn APIサーバー

シンプルなFlaskとGunicornを使用したAPIサーバーです。

## セットアップ

### 通常の実行方法

1. 依存関係のインストール:

```bash
pip install -r requirements.txt
```

### Dockerを使用した実行方法

1. Dockerイメージのビルドと起動:

```bash
docker-compose up --build
```

1. バックグラウンドで実行する場合:

```bash
docker-compose up -d
```

1. コンテナの停止:

```bash
docker-compose down
```

## 実行方法

開発モード:

```bash
python app.py
```

本番モード（Gunicorn）:

```bash
gunicorn wsgi:app
```

## 利用可能なエンドポイント

- `GET /api/health` - ヘルスチェック
- `GET /api/hello` - 挨拶メッセージ

## レスポンス例

### ヘルスチェック

```json
{
  "status": "healthy",
  "message": "APIサーバーは正常に動作しています"
}
```

### 挨拶

```json
{
  "message": "こんにちは！"
}
```

## Docker環境について

- ポート: 8000番を使用
- 開発モードで実行
- ホットリロード対応（ソースコードの変更が即座に反映）
- コンテナが予期せず停止した場合は自動再起動
