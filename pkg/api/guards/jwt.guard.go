package guards

import (
	"fmt"
	"github.com/hramov/tg-bot-admin/internal/config"
	appError "github.com/hramov/tg-bot-admin/internal/error"
	"github.com/hramov/tg-bot-admin/pkg/api/utils"
	"github.com/hramov/tg-bot-admin/pkg/jwt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func checkRoles(roles []string, id int, params httprouter.Params) error {
	for _, role := range roles {
		if role == "equal_id" {
			idPathRaw := params.ByName("user_id")
			idPath, err := strconv.Atoi(idPathRaw)
			if err != nil {
				return err
			}
			if id != idPath {
				return fmt.Errorf("id not match")
			}
		}
	}
	return nil
}

func JwtGuard(h httprouter.Handle, roles []string) httprouter.Handle {
	cfg := config.GetConfig()
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		token, err := utils.GetTokenFromRequest(r)
		if err != nil {
			utils.SendError(http.StatusUnauthorized, "cannot get auth token", w)
			return
		}

		data, err := jwt.TokenValid(token, jwt.AccessToken, cfg.Jwt.Secret)
		if err != nil {
			utils.SendError(http.StatusUnauthorized, err.Error(), w)
			return
		}

		rawId := data["jti"].(string)
		id, err := strconv.Atoi(rawId)
		if err != nil {
			utils.SendError(http.StatusInternalServerError, err.Error(), w)
			return
		}

		if id != 0 {
			err = checkRoles(roles, id, params)
			if err != nil {
				utils.SendError(http.StatusUnauthorized, err.Error(), w)
				return
			}
			h(w, r, params)
			return
		}

		utils.SendError(appError.CannotGetIdError().Status(), appError.CannotGetIdError().Error(), w)
	}
}
