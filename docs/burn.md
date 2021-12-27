## Burn Docs:

## Installation:
```
$ go get -u github.com/BurnFramework/burn
```

## Hello World:

```golang
package main

import (
	burn "github.com/BurnFramework/burnweb"
)

func main() {
	br := burn.New()

	br.Get("/hello", func(ctx burn.Context) {
		ctx.SendString("Hello")
	})

	br.Start(":8080")
}
```