package go_cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"testing"
	"time"
)

func TestBaseDemo(t *testing.T) {
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set("foo", "bar", cache.DefaultExpiration)
	c.Set("baz", 42, cache.NoExpiration)
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}
}
