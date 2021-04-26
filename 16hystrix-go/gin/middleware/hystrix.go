package middleware

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

func HystrixDo(ctx *gin.Context) {
	commandName := ctx.Request.Method + "-" + ctx.Request.RequestURI
	hystrix.Do(
		commandName,
		func() error {
			ctx.Next()
			// do what to do
			return nil
		},
		func(err error) error {
			ctx.Set("hystrixretmsg", "setmsg")
			return err
		},
	)
}
