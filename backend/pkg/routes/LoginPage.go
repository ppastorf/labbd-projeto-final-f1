package routes

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Login    string `form:"login"`
	Password string `form:"password"`
}

func Login(ctx echo.Context) error {
	request := &LoginRequest{}

	err := ctx.Bind(request)
	if err != nil {
		ctx.Echo().Logger.Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	// checar se usuario existe no banco

	// checar se senha está certa

	// logar acesso

	// redireciona pra página principal

	// if err = request.Validate(ctx); err != nil {
	// 	ctx.Echo().Logger.Error(err)
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	return ctx.JSON(http.StatusCreated, struct{}{})
}
