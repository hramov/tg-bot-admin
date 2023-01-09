package guards

import (
	"github.com/hramov/tg-bot-admin/internal/config"
	"github.com/hramov/tg-bot-admin/pkg/jwt"
	"github.com/hramov/tg-bot-admin/pkg/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"strconv"
)

func checkRoles(roles []string, id int, perm jwt.Permissions, params httprouter.Params) error {
	//for _, role := range roles {
	//	if role == "equal_id" {
	//		idPathRaw := params.ByName("user_id")
	//		idPath, err := strconv.Atoi(idPathRaw)
	//		if err != nil {
	//			return err
	//		}
	//		if id != idPath {
	//			return fmt.Errorf("id not match")
	//		}
	//	}
	//}
	if perm.Admin {
		return nil
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

		data, err := jwt.TokenValid(token, cfg.Jwt.AccessSecret)
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
			userPermRaw := data["permissions"].(map[string]interface{})
			var userPerm jwt.Permissions
			err = mapstructure.Decode(userPermRaw, &userPerm)
			if err != nil {
				utils.SendError(http.StatusUnauthorized, err.Error(), w)
				return
			}
			err = checkRoles(roles, id, userPerm, params)
			if err != nil {
				utils.SendError(http.StatusUnauthorized, err.Error(), w)
				return
			}
			h(w, r, params)
			return
		}

		utils.SendError(http.StatusBadRequest, "cannot get user id", w)
	}
}
