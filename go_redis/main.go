package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main(){
	client := redis.NewClient(&redis.Options{
		Addr:               "127.0.0.1:6379",
		Password:           "",
		DB:                 0,
	})
	//pongTest(client)
	//listKey(client)
	//hashKey(client)
	setKey(client)
}

func pongTest(client *redis.Client){
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("test_key", 11, 0).Err()
	if err != nil{
		fmt.Println(err.Error())
	}
}

func stringKey(client *redis.Client){
	val, err := client.Get("test_key").Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(val)

	val, err = client.Get("test_key_2").Result()
	if err == redis.Nil {
		fmt.Println(err.Error())
	} else if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(val)
	}

	client.Del("test_key")
}

func listKey(client *redis.Client){
	err := client.LPush("test_list", 1,2,3).Err()
	if err != nil {
		fmt.Println(err.Error())
	}
	lrange_res, _ := client.LRange("test_list", 0, -1).Result()
	fmt.Println(lrange_res)

	lindex_res, _ := client.LIndex("test_list", 1).Result()
	fmt.Println(lindex_res)

	llen_res := client.LLen("test_list")
	fmt.Println(llen_res)

	lpop_res := client.LPop("test_list")
	fmt.Println(lpop_res)

	client.LRem("test_list", 10, 3)
	lrange_res, _ = client.LRange("test_list", 0, -1).Result()
	fmt.Println(lrange_res)
}

func hashKey(client *redis.Client){
	err := client.HSet("test_hash", "name", "sunqi").Err()
	if err != nil {
		fmt.Println(err.Error())
	}

	client.HMSet("test_hash", map[string]interface{}{"age": 22, "sex": "female"})

	res := client.HExists("test_hash", "name")
	fmt.Println(res.Val())

	hget_res, _ := client.HGet("test_hash", "name").Result()
	fmt.Println(hget_res)

	hgetall_res, _ := client.HGetAll("test_hash").Result()
	fmt.Println(hgetall_res)

	client.HIncrBy("test_hash", "age", 1)

	hkeys_res, _ := client.HKeys("test_hash").Result()
	fmt.Println(hkeys_res)

	hmget_res, _ := client.HMGet("test_hash", "name", "age", "sex").Result()
	fmt.Println(hmget_res)

	hvals_res, _ := client.HVals("test_hash").Result()
	fmt.Println(hvals_res)
}

func setKey(client *redis.Client){
    add_count, _ := client.SAdd("test_set", 1,2,3,4).Result()
    fmt.Println(add_count)

    count, _ := client.SCard("test_set").Result()
    fmt.Println(count)

    res, _ := client.SIsMember("test_set", 3).Result()
    fmt.Println(res)

    sm_res, _ := client.SMembers("test_set").Result()
    fmt.Println(sm_res)

    spop_res, _ := client.SPop("test_set").Result()
    fmt.Println(spop_res)

    srem_res, _ := client.SRem("test_set", 1,2,3).Result()
    fmt.Println(srem_res )

    test := string(2)
    fmt.Println(test)
}
