package url

import (
	"github.com/gin-gonic/gin"
	"short-url/internal/global/database"
	"short-url/internal/global/errs"
	"short-url/internal/model"
	"short-url/service/shorten"
	"short-url/tools"
)

func Shorten(c *gin.Context) {
	req := struct {
		Url string `json:"url" binding:"required"`
	}{}

	if err := c.ShouldBindJSON(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithTips(err.Error()))
		return
	}

	newURL := &model.Url{
		ShortURL: shorten.ShortenURL(req.Url),
		LongURL:  req.Url,
	}

	err := database.Query.Url.Create(newURL)

	for tools.IsDuplicateKeyError(err) {
		newURL.ShortURL = newURL.ShortURL + tools.RandString(1)
		err = database.Query.Url.Create(newURL)
	}

	if err != nil {
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}

	errs.Success(c, baseURL+newURL.ShortURL)
}
