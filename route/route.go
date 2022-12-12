package route

import (
	"ucenter/handler/api_handler"
	"ucenter/handler/oauth2_handler"
	"ucenter/handler/user_handler"

	"github.com/kataras/iris/v12"
)

func InitRouter(app *iris.Application) {
	/* you can pass this authorization part of the code if you don't need to change it */
	/* if you need to verify accees just focus on the func  "GetSession" (.\http_helper\session.go) */
	/* you can use the func like */
	/*
		  	sess := http_helper.GetSession(ctx)
			if sess == nil {
				return
			}
	*/
	/* and get or delete or save data from the session */
	/*
			sess.Get("something")
			sess.Delete("something")
		    if err := sess.Save(); err != nil {
				log.Println(err)
		    }
	*/
	app.PartyFunc("/oauth2", func(u iris.Party) {
		u.Get("/authorize", oauth2_handler.GetAuthorizeHandler)
		u.Post("/authorize", oauth2_handler.PostAuthorizeHandler)
		u.Get("/auth", oauth2_handler.GetAuthHandler)
		u.Post("/auth", oauth2_handler.PostAuthHandler)
		u.Post("/token", oauth2_handler.PostTokenHandler)
	}) //authorization handler
	app.PartyFunc("/user", func(u iris.Party) {
		u.Post("/register", user_handler.PostRegisterHandler)
		u.Get("/login", user_handler.GetLoginHandler)   //request type "get" so get the session info
		u.Post("/login", user_handler.PostLoginHandler) //request type "post" meaning new login
		u.Get("/logout", user_handler.GetLogoutHandler) //clear login info
		u.Post("/password", user_handler.PostSetPasswordHandler)
	}) //user handler
	app.PartyFunc("/api", func(u iris.Party) {
		u.Get("/data", api_handler.GetDataHandler) //get user info form database
	})
}
