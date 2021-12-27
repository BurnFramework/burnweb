# burnweb
Burn Web Framework

## About burn:
- burn is a golang framework for building restapi
- It's main aim is to be fast, productivity, and easy 

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

- for more tutorials check out the burn official[docs]()