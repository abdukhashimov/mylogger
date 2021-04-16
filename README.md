# MyLogger (mylogger)
## It is the package that generates config for `uber-go/zap`'s logger

## How to use it ?

```
package main

import (
	"github.com/abdukhashimov/mylogger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	loggerConfig := mylogger.LoadZapConfig(true)
	logger, _ := loggerConfig.Build()
	logger.Info("Hello World", zap.Bool("abc", false))

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}

```
