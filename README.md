
redisに接続するコンテナ3つ
docker-compose start redis-pubsub-go-1 // 1~3までそれぞれstart
docker-compose exec redis-pubsub-go-1 ash
> go run main

redisからpublishしたい場合コンテナに繋いでやる
docker-compose redis-pubsub-go-redis ash
redis-cli
> subscribe channel1
> publish channel1 msg1


websocketとpub/subを利用してスケーラブルな構成にする
それぞれのユーザデバイスが別のサーバーにwebsocketで繋いでいた場合、双方向通信は出来ない
そこでそれぞれのserverにメッセージを送信するためにpub/subを利用する

device1 device2 device3
   |       |       |   (websocket)
server1 server2 server3
   \       |      /
    redis(pub/sub)

上記のような接続状態でdevice1とdevice3のユーザが異なるサーバーにwebsocketで通信している場合、device1から送信されたメッセージがdevice3に届かない
そこで各serverはredis(pub/sub)に接続して、server1で受け取ったメッセージをpub/subを通してserver(1,2,3)全てに送信する
その後に受け取ったメッセージを処理してwebsocketで返すようにする


## 簡易コード
// websocketからメッセージを受信する
ws.on(message => {
  redisClient.publish(channel, message)
})
// pub/subメッセージ受信
val pubsub = redisClient.subscribe(channel)
pubsub.map { msg =>
  (msg.channel, msg.payload)
}









