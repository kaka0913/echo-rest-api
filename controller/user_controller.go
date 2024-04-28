package controller

import (
	"go-rest-api/usecase"
	"go-rest-api/model"
	"net/http"
	"time"
	"os"
	"github.com/labstack/echo/v4"
)

type IUserController interface{
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct{
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController{//コンストラクタでDIを実現している
	return &userController{uu}									//usecaseを受け取りuserControllerを返す
}

func (uc *userController) SignUp(c echo.Context) error{
	user := model.User{}
	if err := c.Bind(&user); err != nil {//リクエストのボディをuserにバインド
		return c.JSON(http.StatusBadRequest, err.Error()) 
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) Login(c echo.Context) error{
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error()) 
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)//有効期限
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true //HTTPSでのみ送信されるようにする
	cookie.HttpOnly = true //クライアントのJavaScriptからアクセスできないようにする
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie) //クッキーをセットしてレスポンスに含める
	return c.NoContent(http.StatusOK)
}

func (uc *userController) LogOut(c echo.Context) error{
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true 
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie) 
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error{
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}