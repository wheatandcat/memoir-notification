# memoir-notification

Push通知を送信する


## 本番デプロイ

```
$ git checkout main
$ git pull --ff-only origin main
$ git tag -a v1.0.0 -m 'リリース内容'
$ git push origin v1.0.0
```


## Cloud TasksにQueueを登録


```
$ gcloud tasks queues create sendNotification
```


```
$ gcloud tasks queues update sendNotification --max-attempts 6 --min-backoff 5s --max-doublings 3
```

## Cloud TasksのQueueを確認

```
$ gcloud tasks queues describe sendNotification
```

## CI環境

### レビュー環境

```
$ base64 -i serviceAccount.review.json | pbcopy
```

### 本番環境

```
$ base64 -i serviceAccount.production.json | pbcopy
```
