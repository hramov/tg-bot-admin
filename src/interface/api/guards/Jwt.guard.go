package guards

import (
	"github.com/gin-gonic/gin"
	"github.com/hramov/tg-bot-admin/src/interface/api/utils"
	appError "github.com/hramov/tg-bot-admin/src/interface/error"
	"github.com/hramov/tg-bot-admin/src/modules/jwt"
	"net/http"
	"strconv"
)

func JwtGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.GetTokenFromContext(c)
		if err != nil {
			utils.SendError(http.StatusInternalServerError, err.Error(), c)
			return
		}
		data, err := jwt.TokenValid(token, jwt.AccessToken)
		if err != nil {
			utils.SendError(http.StatusUnauthorized, err.Error(), c)
			return
		}
		rawId := data["jti"].(string)
		id, err := strconv.Atoi(rawId)
		if err != nil {
			utils.SendError(http.StatusInternalServerError, err.Error(), c)
			return
		}
		if id != 0 {
			c.Set("id", id)
			c.Next()
			return
		}
		utils.SendError(appError.CannotGetIdError().Status(), appError.CannotGetIdError().Error(), c)
	}
}
