# go-querystruct #

go-querystruct is Go library for convert URL query parameters to struct .



## Usage ##

```
import "github.com/mokeoo/go-querystruct/params"

type User struct {
	UserId  int64   `param:"user_id,100"`
	StoreId int     `param:"store_id"`
	Page    float32 `param:"store_id"`
	Name    string  `param:"name"`
	Age     uint8   `param:"age,18"`
	Enable  bool    `param:"enable,false"`
}

o := Order{}
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
		e := params.Unmarshal(want, &o)
		if e == nil {
			if o.StoreId != storeId || o.UserId != userId || o.Page != page {
				b.Error("has error ")
			}
		} else {
			b.Error(e)
		}

```

## Performance ##
```
goos: linux
goarch: amd64
pkg: github.com/mokeoo/go-querystruct/params
2000000000	         0.00 ns/op
PASS
```

## License ##

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.
