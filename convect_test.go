package params

import (
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"testing"
)

type User struct {
	UserId  int64   `param:"user_id,100"`
	StoreId int     `param:"store_id"`
	Page    float32 `param:"page"`
	Name    string  `param:"name"`
	Age     uint8   `param:"age,18"`
	Enable  bool    `param:"enable,false"`
}

func TestReflect(t *testing.T) {
	o := User{}
	want := url.Values{
		"store_id": {"3"},
	}
	e := Convert(want, &o)
	if e != nil {
		t.Error(e)
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i > b.N; i++ {
		o := User{}
		userId := rand.Int63()
		storeId := rand.Int()
		page := rand.Float32()
		age := rand.Intn(8)
		want := url.Values{
			"store_id": {fmt.Sprintf("%v", storeId)},
			"user_id":  {strconv.FormatInt(userId, 64)},
			"page":     {fmt.Sprintf("%v", page)},
			"name":     {"sdfdsfs"},
			"age":      {fmt.Sprintf("%v", age)},
		}
		e := Convert(want, &o)
		if e == nil {
			if o.StoreId != storeId || o.UserId != userId || o.Page != page {
				b.Error("has error ")
			}
		} else {
			b.Error(e)
		}
	}
}
