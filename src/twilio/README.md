# sample.twilio

## 環境変数

| 環境変数 | 説明 | 必須 |
|----------|------|------|
| TWILIO_ACCOUNT_SID | Twilioのアカウント識別子。Twilioコンソールから取得できます。 | ○ |
| TWILIO_AUTH_TOKEN | Twilioの認証トークン。Twilioコンソールから取得できます。 | ○ |
| TWILIO_PHONE_NUMBER | 発信元の電話番号。Twilioで購入した電話番号を設定します。 | ○ |
| TWIML_URL | 通話時に再生される音声コンテンツのURL。TwiMLを提供するエンドポイントを指定します。 | ○ |

## 参考

- [twilio/twilio-go](https://github.com/twilio/twilio-go)
