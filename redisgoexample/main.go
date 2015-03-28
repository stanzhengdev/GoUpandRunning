package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//INIT OMIT
	r, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	//set
	r.Do("SET", "message1", "Hello World")

	for i := 0; i < 100; i++ {
		//messNum := strconv.Itoa(i)
		r.Do("INCR", "foo")
	}

	//get
	world, err := redis.String(r.Do("GET", "foo"))
	if err != nil {
		fmt.Println("key not found")
	}

	fmt.Println(world)
	//ENDINIT OMIT
}
