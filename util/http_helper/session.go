package http_helper

import (
	"log"

	"github.com/go-session/session"
	"github.com/kataras/iris/v12"
)

func GetSession(ctx iris.Context) session.Store {
	sess, err := session.Start(ctx.Request().Context(), ctx.ResponseWriter(), ctx.Request()) // Start a session and return to session storage
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		if _, err = ctx.WriteString("session internal error"); err != nil {
			log.Println(err)
		}
		return nil
	}
	return sess
}
