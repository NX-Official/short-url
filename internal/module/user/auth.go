package user

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"short-url/internal/global/database"
	"short-url/internal/global/errs"
	"short-url/internal/global/jwt"
	"short-url/tools"
)

func Login(c *gin.Context) {
	var req User
	if err := c.BindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithTips(err.Error()))
		return
	}
	u := database.Query.User
	userInfo, err := u.Where(u.Email.Eq(req.Email)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		errs.Fail(c, errs.NotFound)
		return
	case err != nil:
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}

	if tools.Compare(req.Password, userInfo.Password) {
		errs.Fail(c, errs.InvalidPassword)
		return
	}
	errs.Success(c, map[string]string{
		"token": jwt.CreateToken(jwt.Payload{UserId: userInfo.ID}),
	})
}
