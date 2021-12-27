# burnweb
Burn Web Framework

## QuickInstallation:
```
$ go get -u github.com/BurnFramework/burnweb
```

## QuickStart:
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
