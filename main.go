package main

import (
	"log"
	"ucenter/config"
	"ucenter/route"

	"github.com/go-session/session"
	"github.com/kataras/iris/v12"
)

func init() {
	session.InitManager(
		session.SetCookieName(config.GetConfig().Session.Name),
		session.SetSign([]byte(config.GetConfig().Session.SecretKey)),
	)
} //unfinished function about session initialization

func main() {
	app := newApp()
	route.InitRouter(app)                                    //bind routes
	err := app.Run(iris.Addr(":" + config.GetConfig().Port)) //start server
	if err != nil {
		log.Fatalln(err)
	}
}

func newApp() *iris.Application {
	app := iris.New()
	app.RegisterView(iris.HTML("./static", ".html"))
	app.Configure(iris.WithOptimizations)
	app.AllowMethods(iris.MethodOptions)
	return app
} //new a app using iris
