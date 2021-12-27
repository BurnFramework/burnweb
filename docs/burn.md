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

## Static files:

- serving html 
```main.go```
```golang
package main

import (
	burn "github.com/BurnFramework/burnweb"
)

func main() {
    br := godzilla.New()

    br.Static("/index", "./index.html")

    br.Start(":8080")
}
```

```index.html```
```html
<html>
<head>
<body>
	<p>WELCOME TO THE PARTY</p>
</body>
</head>
</html>
```