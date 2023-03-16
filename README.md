### go-querystruct

[![Badge](https://img.shields.io/badge/link-996.icu-%23FF4D5B.svg?style=flat-square)](https://996.icu/#/en_US)
[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg?style=flat-square)](https://github.com/996icu/996.ICU/blob/master/LICENSE)

go-querystruct is Go library for cast url.Values to struct .

----

### Usage ###

```golang
import "github.com/youkale/go-querystruct"

type User struct {
    UserId  int64   `param:"user_id,100"`
    StoreId int     `param:"store_id"`
    Page    float32 `param:"page"`
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

### Performance ###

```
    goos: linux
    goarch: amd64
    pkg: github.com/youkale/go-querystruct/params
    2000000000	         0.00 ns/op
    PASS
```

## License ##

[![LICENSE](https://img.shields.io/badge/license-Anti%20996-blue.svg)](https://github.com/996icu/996.ICU/blob/master/LICENSE)
