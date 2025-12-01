# Gmail API
Google Cloud Gmail API のサンプルコード

## 環境変数

| 名前                             | 説明 　　　　　　                                                            | 必須 | 備考  |
| -------------------------------- | -------------------------------------------------------------------------- | ---- | ------------ |
| `GOOGLE_CLOUD_CREDENTIAL_PATH`   | Google Cloud Gmail API の OAuth 2.0 クライアント ID（JSON ファイル）のパス    | ✔    |              |
| `GOOGLE_CLOUD_ACCESS_TOKEN_PATH` | Google Cloud Gmail API のアクセストークン（JSON ファイル）のパス              |     | 自動で生成される（手順を参照）   |
| `GOOGLE_CLOUD_PROJECT_NAME`      | Google Cloud のプロジェクト名 　　　　　　　　　　　                          |      | `watch` APIを利用する際に必須 |
| `GOOGLE_CLOUD_WATCH_TOPIC_NAME`  | Google Cloud Pub/Sub の Gmail 通知先のトピック名 　　　　　　　　　　　        |     |  `watch` APIを利用する際に必須 |

---

## 実行方法

### 前提

- Google Cloud Gmail API の OAuth 2.0 クライアント ID（JSON ファイル）を用意すること（https://console.cloud.google.com/apis/credentials）。
- Google Cloud API & Services の OAuth 同意画面からテストユーザーとして自身のメールアドレスを追加していること（https://console.cloud.google.com/auth/audience）。

### 手順

1. `docker-compose.example.yml`をコピーし、`docker-compose.yml`を作成。
2. `docker-compose.yml`に上記の必須環境変数を設定。
3. Google Cloud Gmail API の OAuth 2.0 クライアント ID を`GOOGLE_CLOUD_CREDENTIAL_PATH`で設定したパスに配置。
4. `docker compose up`でアプリケーションを実行。
5. ターミナル出力されるURLを踏むと、OAuth同意画面へ遷移するため、自身のGoogleアカウントで許可し、認可コードをターミナルに貼り付けてEnter。
6. アクセストークン（`access_token.json`）が生成される（次回からはこのアクセストークンを利用するため、`5.`の作業は不要）。

## 参考
- [GCP の Pub/Sub と Gmail APIを使って、メールをトリガーにイベントを実行する](https://qiita.com/Kentaro91011/items/e8f9c5764ff6792c1d98)
- [Go クイックスタート](https://developers.google.com/workspace/gmail/api/quickstart/go?hl=ja)


