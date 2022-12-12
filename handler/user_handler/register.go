package user_handler

import (
	"ucenter/db"
	"ucenter/model"
	"ucenter/util/http_helper"
	"ucenter/util/validation"

	"github.com/kataras/iris/v12"
)

func PostRegisterHandler(ctx iris.Context) {
	sess := http_helper.GetSession(ctx) //get session
	if sess == nil {
		return
	}
	if ctx.Request().Form != nil {
		if err := ctx.Request().ParseForm(); err != nil { //parse form failed return
			ctx.StatusCode(iris.StatusNotAcceptable)
			_, _ = ctx.JSON(model.NewResult(nil, 1003, "invalid form submitted"))
			return
		}
	}
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	if !validation.ValidateUsername(username) || !validation.ValidatePassword(password) { //check username and password format
		ctx.StatusCode(iris.StatusNotAcceptable)
		_, _ = ctx.JSON(model.NewResult(nil, 1001, "invalid username or password"))
		return
	}
	db.NewUser(username, password) //save a new user to the database
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(model.NewResult(nil, 0, "successfully registered"))
}
