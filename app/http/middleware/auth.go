package middleware

import (
	"goravel/app/helpers"

	"github.com/golang-jwt/jwt/v5"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Auth() http.Middleware {
	return func(ctx http.Context) {
		token := ctx.Request().Header("Authorization")
		if token == "" {
			ctx.Request().AbortWithStatusJson(401, helpers.JsonResponse{
				StatusCode: 401,
				Message:    "Unauthorized - Token required",
			})
			return
		}

		// Remove "Bearer " prefix if present
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		// Parse and validate JWT token manually
		jwtSecret := facades.Config().GetString("jwt.secret")
		_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil {
			ctx.Request().AbortWithStatusJson(401, helpers.JsonResponse{
				StatusCode: 401,
				Message:    "Unauthorized - Invalid token",
			})
			return
		}

		ctx.Request().Next()
	}
}
