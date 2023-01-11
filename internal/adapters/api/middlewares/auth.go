package middlewares

import (
	"context"
	"github.com/hramov/tg-bot-admin/internal/config"
	"github.com/hramov/tg-bot-admin/pkg/jwt"
	"github.com/hramov/tg-bot-admin/pkg/utils"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"strconv"
)

func checkRoles(roles []string, id int, perm jwt.Permissions) error {
	if perm.Admin {
		return nil
	}
	return nil
}

func Auth(h http.HandlerFunc, roles []string) http.HandlerFunc {
	cfg := config.GetConfig()
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.GetTokenFromRequest(r)
		if err != nil {
			utils.SendError(http.StatusUnauthorized, "cannot get auth token", w)
			return
		}

		data, err := jwt.GetClaims(token, cfg.Jwt.AccessSecret)
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
			err = checkRoles(roles, id, userPerm)
			if err != nil {
				utils.SendError(http.StatusUnauthorized, err.Error(), w)
				return
			}

			ctx := context.WithValue(r.Context(), "user_id", id)
			h(w, r.WithContext(ctx))
			return
		}

		utils.SendError(http.StatusBadRequest, "cannot get user id", w)
	}
}
