package httpdcounter

import (
	"fmt"
	"testing"
	"time"
)

func FullData() *Counter {
	c := New()
	c.Set(time.Now().Format("2006-01-02"))
	c.Set("2017-01-02")
	c.Set("2017-01-02")
	c.Set("2017-01-02")
	c.Set("2017-01-03")
	c.Set("2017-01-03")
	return c
}

func TestSet(t *testing.T) {
	c := FullData()
	fmt.Println(c.String())
}

func TestGetMap(t *testing.T) {
	c := FullData()
	m := c.GetMap()
	fmt.Print(m)
}
