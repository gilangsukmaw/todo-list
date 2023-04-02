package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	cfg "gitlab.com/todo-list-app1/todo-list-backend/cfg/env"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/consts"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/entity"
	"gitlab.com/todo-list-app1/todo-list-backend/internal/server"
	"gitlab.com/todo-list-app1/todo-list-backend/lib/logger"
	"strings"
)

func ValidateJWT(dctx *fiber.Ctx, cfg *cfg.Config) server.Response {
	var (
		headers = entity.Headers{}
		err     = dctx.ReqHeaderParser(&headers)
		// Initialize a new instance of `Claims`
		claims = &entity.Claims{}

		//logger
		lf = logger.Field{
			EventName: "middleware_validating_bearer_token",
		}
		log = logger.NewLoggerField(lf)
	)

	logrus.WithField("event", log).Info("middleware_validating_bearer_token")

	if err != nil {
		logrus.WithField("event",
			lf.Append("cast headers got error", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 500, Message: consts.InternalServerError}
	}

	authorization := strings.TrimSpace(headers.Authorization)

	if authorization == "" {
		logrus.WithField("event",
			lf.Append("authorization", fmt.Sprintf(`got empty string`))).Error()
		return server.Response{Code: 401, Message: consts.Unauthorized}
	}

	value := strings.Split(authorization, " ")

	token := value[1]

	if len(value) < 2 {
		logrus.WithField("event",
			lf.Append("authorization", fmt.Sprintf(`got invalid token format`))).Error()
		return server.Response{Code: 401, Message: consts.InvalidTokenFormat}
	}

	if len(value[1]) == 0 {
		logrus.WithField("event",
			lf.Append("authorization", fmt.Sprintf(`token got empty string or len is 0`))).Error()
		return server.Response{Code: 401, Message: consts.Unauthorized}
	}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		jwtKey := cfg.JwtToken
		return []byte(jwtKey), nil
	})
	if err != nil {
		if strings.Contains(fmt.Sprintf(`%s`, err), consts.SignatureInvalid) {
			logrus.WithField("event",
				lf.Append("verify token got", fmt.Sprintf(`%s`, err))).Error()
			return server.Response{Code: 401, Message: consts.SignatureInvalid}
		}

		if strings.Contains(fmt.Sprintf(`%s`, err), consts.TokenExpired) {
			logrus.WithField("event",
				lf.Append("verify token got", fmt.Sprintf(`%s`, err))).Error()
			return server.Response{Code: 401, Message: consts.TokenExpired}
		}

		logrus.WithField("event",
			lf.Append("verify token got", fmt.Sprintf(`%s`, err))).Error()
		return server.Response{Code: 401, Message: consts.Unauthorized}
	}

	c := tkn.Claims.(*entity.Claims)
	dctx.Locals("user_id", c.UserId)
	dctx.Locals("username", c.Username)
	dctx.Locals("role", c.Role)

	return server.Response{Code: 200, Message: "success"}
}
