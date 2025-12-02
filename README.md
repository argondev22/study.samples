# Samples

様々な技術のサンプルコード集です。各種APIの統合、並行処理、インフラストラクチャのコード例を含んでいます。

## サンプル一覧

### API統合

- **[Gmail API](src/gmail-api/)** - Google Cloud Gmail APIのサンプル実装
- **[Twilio](src/twilio/)** - Twilioを使った電話通話のサンプル
- **[gRPC](src/grpc/)** - gRPCサーバー・クライアントの実装例

### Webアプリケーション

- **[Flask-Gunicorn](src/flask-gunicorn/)** - FlaskとGunicornを使用したAPIサーバー

### インフラストラクチャ

- **[AWS Lambda](src/aws-lambda/)** - AWS Lambda関数のサンプル（Go/Python）とTerraform設定

### アルゴリズム・並行処理

- **[Parallelism](src/parallelism/)** - Goでの並行処理のサンプルコード
- **[Algorithm](src/algorithm/)** - アルゴリズムのサンプル実装

## 開発環境

### 必要なツール

- Docker
- Node.js (開発ツール用)
- Git LFS

### セットアップ

1. **リポジトリのクローン**

   ```bash
   git clone <repository-url>
   cd study.samples
   ```

2. **開発依存関係のインストール**

   ```bash
   npm install
   ```

   このコマンドでHusky、Prettier、Commitlintなどの開発ツールがインストールされます。

3. **Git LFSのセットアップ**

   ```bash
   # Git LFSのインストール（初回のみ）
   brew install git-lfs  # macOS
   # または https://git-lfs.github.com/ からダウンロード

   # Git LFSの初期化（初回のみ）
   git lfs install

   # LFSファイルの取得
   git lfs pull
   ```

## プロジェクト構成

```text
.
├── .github/                  # GitHubの設定（Issue/PRテンプレート）
├── .vscode/                  # VS Code設定
├── src/                      # サンプルコードのディレクトリ
│   ├── algorithm/            # アルゴリズムのサンプル
│   ├── aws-lambda/           # AWS Lambdaのサンプル
│   ├── flask-gunicorn/       # Flask-GunicornのAPIサーバー
│   ├── gmail-api/            # Gmail APIのサンプル
│   ├── grpc/                 # gRPCのサンプル
│   ├── parallelism/          # Goの並行処理サンプル
│   └── twilio/               # Twilioのサンプル
├── .editorconfig             # エディタ設定
├── .prettierrc.json          # Prettierの設定
├── .commitlintrc.json        # Commitlintの設定
└── package.json              # 開発依存関係
```

## コード品質管理

このリポジトリは以下のツールを使用してコード品質を維持しています：

- **Prettier** - コードフォーマッター
- **Commitlint** - コミットメッセージの規約チェック
- **Markdownlint** - Markdownファイルのリント
- **Husky** - Git hooksの管理

### 利用可能なコマンド

```bash
# コードフォーマット
npm run format

# フォーマットチェック
npm run format:check

# Markdownのリント
npm run lint:markdown
```

## コントリビューション

1. リポジトリをフォーク
2. フィーチャーブランチを作成
3. 変更を加える
4. プルリクエストを作成

コミットメッセージは[Conventional Commits](https://www.conventionalcommits.org/)の規約に従ってください。

## ライセンス

このプロジェクトのライセンスについては[LICENSE](LICENSE)を参照してください。
