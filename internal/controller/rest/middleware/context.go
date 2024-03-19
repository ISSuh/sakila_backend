package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func WrapContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("[WrapContext]] begine\n")
		c.Next()
		fmt.Printf("[WrapContext]] end \n")

		// var err error
		// req, err := http.NewAppRequest(c)
		// if err != nil {
		// 	res := http.NewResponseErrWithBody(gohttp.StatusBadRequest, common.InvalidReqeustMessage)
		// 	return c.JSON(gohttp.StatusBadRequest, res)
		// }

		// resp := http.NewResponseOK()

		// appContext := &common.AppContext{c, context.Background()}
		// appContext.Ctx = context.WithValue(appContext.Ctx, common.ReqeustContextKey, req)
		// appContext.Ctx = context.WithValue(appContext.Ctx, common.ResponseContextKey, resp)

		// err = next(appContext)
		// if err != nil {
		// 	res := http.NewResponseErrWithBody(gohttp.StatusBadRequest, common.InvalidReqeustMessage)
		// 	return c.JSON(gohttp.StatusBadRequest, res)
		// }
	}
}
