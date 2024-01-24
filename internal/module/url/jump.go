package url

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/http"
	"short-url/internal/global/database"
	"short-url/internal/global/errs"
)

func Jump(c *gin.Context) {
	req := struct {
		ShortUrl string `uri:"shortUrl" binding:"required"`
	}{}

	if err := c.ShouldBindUri(&req); err != nil {
		errs.Fail(c, errs.InvalidRequest.WithTips(err.Error()))
		return
	}

	url, err := database.Query.Url.Where(database.Query.Url.ShortURL.Eq(req.ShortUrl)).First()
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		errs.Fail(c, errs.NotFound)
		return
	case err != nil:
		errs.Fail(c, errs.DatabaseError.WithOrigin(err))
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.LongURL)
}
