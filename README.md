takarabago
==========

inspired by [takarabako](https://github.com/willnet/takarabako).

How to install
--------------

```
> go get github.com/futoase/takarabago
```

How to use
----------

### on command-line

```
> takarabago
ブリザードチュニック
> takarabago
リザードブレイド
```

### on library

```go
package main

import (
        "fmt"
        tg "github.com/futoase/takarabago/libs"
)

func main() {
        fmt.Println(tg.Takarabako())
}
```


License
-------

MIT


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/futoase/takarabago/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

