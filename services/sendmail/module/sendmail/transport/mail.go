package mailtransport

import (
	"fmt"
	"net/http"
	"shopbee/common"
	"shopbee/component/appctx"
	mailbiz "shopbee/module/sendmail/biz"

	"github.com/gin-gonic/gin"
)

func SendMail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		type Mail struct {
			Receiver string `json:"receiver"`
			Subject  string `json:"subject"`
			Body     string `json:"body"`
		}

		var mail Mail
		if err := c.ShouldBind(&mail); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fmt.Print(mail)
		mailbiz.SendMail(mail.Receiver, mail.Subject, mail.Body)

		c.JSON(http.StatusOK, gin.H{
			"data": "send mail success",
		})
	}
}
