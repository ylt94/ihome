package utils

import redis2 "github.com/garyburd/redigo/redis"

func Redis() redis2.Conn {
	conn, err := redis2.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	return conn
}
