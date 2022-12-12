package api_handler

import (
	"strings"
	"ucenter/db"
	"ucenter/handler/oauth2_handler"
	"ucenter/model"

	"github.com/kataras/iris/v12"
)

func GetDataHandler(ctx iris.Context) {
	_ = ctx.Request().ParseForm()
	ok, res := oauth2_handler.VerifyTokenHandler(ctx) //verify token
	if !ok {
		ctx.StatusCode(iris.StatusNotAcceptable)
		_, _ = ctx.JSON(res)
		return
	}
	scopes := strings.Split(ctx.FormValue("scope"), ",") //scope is the name of the fields of the user(用户某项字段的信息)
	result := make(map[string]interface{})
	for i := range scopes {
		result[scopes[i]] = db.GetDataRecordByUser(ctx.FormValue("username"), scopes[i])
	}
	_, _ = ctx.JSON(model.NewResult(result, 0, "success"))
}
