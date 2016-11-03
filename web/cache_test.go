package web_test

import (
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/kapmahc/lotus/web"
)

type S struct {
	IV int
	SV string
}

func TestCache(t *testing.T) {

	c := web.Cache{Redis: &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				"localhost:6379",
			)
		},
	}}

	s1 := S{IV: 100, SV: "hello, lotus!"}
	if err := c.Set("hello", &s1, 60); err != nil {
		t.Fatal(err)
	}
	var s2 S
	if err := c.Get("hello", &s2); err == nil {
		if s1.IV != s2.IV || s1.SV != s2.SV {
			t.Fatalf("wang %v get %v", s1, s2)
		}
	} else {
		t.Fatal(err)
	}
	if keys, err := c.Keys(); err == nil {
		t.Logf("keys: %v", keys)
	} else {
		t.Fatal(err)
	}

	if err := c.Flush(); err != nil {
		t.Fatal(err)
	}
}
