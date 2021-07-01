# memoir-notification

Push通知を送信する


## デプロイ

```
$ gcloud functions deploy SendNotification --runtime go113 --trigger-http --region asia-northeast1
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