package user_handler

import (
	"ucenter/db"
	"ucenter/model"
	"ucenter/util/http_helper"
	"ucenter/util/validation"

	"github.com/kataras/iris/v12"
)

func PostSetPasswordHandler(ctx iris.Context) {
	sess := http_helper.GetSession(ctx) //get session
	if sess == nil {
		return
	}
	if ctx.Request().Form == nil { //empty request
		if err := ctx.Request().ParseForm(); err != nil {
			ctx.StatusCode(iris.StatusNotAcceptable)
			_, _ = ctx.JSON(model.NewResult(nil, 1003, "invalid form submitted"))
			return
		}
	}
	username := ctx.FormValue("username")
	oldPassword := ctx.FormValue("old_password")
	newPassword := ctx.FormValue("new_password")
	if !validation.ValidateUsername(username) || !validation.ValidatePassword(oldPassword) || !validation.ValidatePassword(newPassword) { //check format
		ctx.StatusCode(iris.StatusNotAcceptable)
		_, _ = ctx.JSON(model.NewResult(nil, 1001, "invalid username or password"))
		return
	}
	ok, res := db.VerifyPassword(username, oldPassword) //verify access
	if !ok {
		ctx.StatusCode(iris.StatusNotAcceptable)
		_, _ = ctx.JSON(res)
		return
	}
	db.SetPassword(username, newPassword) //reset password
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(model.NewResult(nil, 0, "successfully set password"))
}
