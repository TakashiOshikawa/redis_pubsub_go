package main

import(
    "log"
    "time"
    "strconv"
    "os"
    "github.com/go-redis/redis"
)


func main() {
    userName := os.Args[1]
    log.Println(userName)

    log.Println("startstart")
    redisClient := ExampleNewClient()
    ExampleClient(redisClient)
    go SendMessageMultiple(redisClient, 3, userName)
    if os.Args[2] == "send" {
        go webSocketOnMsgMock("web socket on msg", redisClient)
    }
    RedisSubscribeExample(redisClient)
    //RedisPublishExample(redisClient, "yesmsg")
    //go AsyncFunc()
    //log.Println("after AsyncFunc")
}

func AsyncFunc() {
    time.Sleep(3*time.Second) // 3秒スリープ
    log.Println("AsycnFunc")
}

func webSocketOnMsgMock(msg string, client *redis.Client) {
    time.Sleep(5*time.Second)
    RedisPublishExample(client, msg)
}

// メッセージを送信する回数を指定して非同期で複数回メッセージを送信する
func SendMessageMultiple(client *redis.Client, number int, userName string) {
    for i := 0; i < number; i++ {
        time.Sleep(5*time.Second)
        RedisPublishExample(client, userName + ": msg"+strconv.Itoa(i))
    }
}

// redis publish
func RedisPublishExample(client *redis.Client, msg string) {
    err := client.Publish("channel1", msg).Err()
    if err != nil {
        panic(err)
    }
}

// redisでのサブスクライブサンプル
func RedisSubscribeExample(client *redis.Client) {
    pubsub := client.Subscribe("channel1")
    ch := pubsub.Channel()

    for msg := range ch {
        log.Println(msg.Channel, msg.Payload)
    }
}


// redisコネクション確認用の関数
func ExampleNewClient() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr: "redis-pubsub-go-redis:6379",
        Password: "",
        DB: 0,
    })

    pong, err := client.Ping().Result()
    log.Println(pong, err)
    log.Println(client)

    return client
}

// redisに値をセットしたりする例
func ExampleClient(client *redis.Client) {
    log.Println("exampleclient")

    err := client.Set("key1", "val1", 0).Err()
    if err != nil {
        panic(err)
    }

    val1, err := client.Get("key1").Result()
    if err != nil {
        panic(err)
    }
    log.Println(val1)

    val2, err := client.Get("key2").Result()
    if err == redis.Nil {
        log.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        log.Println(val2)
    }
}



