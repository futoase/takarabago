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
