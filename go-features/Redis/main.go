package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

/*

- start redis server `$ redis-server /usr/local/etc/redis.conf`
- default port is 6379
- to shutdown `$ redis-cli shutdown`

- alternatively, you can connect to redis server with `$ redis-cli` (after starting the server)
- try `$ SET foo "hello"` ==> `$ GET foo` ==> "hello"
- other simple commands here: https://codeburst.io/redis-what-and-why-d52b6829813

*/

type Podcast struct {
	Title   string  `redis:"title"`
	Creator string  `redis:"creator"`
	Fee     float64 `redis:"fee"`
}

func main() {

	// default port for redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal("could not connect to redis:", err)
	}

	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// the main 'compartment' will be 'podcast:1'
	_, err = conn.Do(
		"HMSET",
		"podcast:1",
		"title",
		"Tech Over Tea",
		"creator",
		"Mackie",
		"fee",
		9.99,
	)

	if err != nil {
		log.Fatal("could not set in redis:", err)
	}



	// get a single value
	title, err := redis.String(conn.Do("HGET", "podcast:1", "title"))
	if err != nil {
		log.Fatal("could not get from redis:", err)
	}
	fmt.Println("Title:", title)



	// get multiple values; returns a map
	values, err := redis.StringMap(conn.Do("HGETALL", "podcast:1"))
	if err != nil {
		log.Fatal("could not get all values from redis:", err)
	}
	fmt.Println(values)
	for k, v := range values {
		fmt.Println(k, ":", v)
	}



	// get values in the form of a struct
	var podcast Podcast
	values2, err := redis.Values(conn.Do("HGETALL", "podcast:1"))
	if err != nil {
		log.Fatal("could not get all values from redis:", err)
	}
	err = redis.ScanStruct(values2, &podcast)
	if err != nil {
		log.Fatal("could not map to struct:", err)
	}
	fmt.Printf("Podcast: %+v\n", podcast)
}
