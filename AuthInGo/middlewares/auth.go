package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	env "AuthInGo/config/env"
	db "AuthInGo/config/db"
	repo "AuthInGo/db/repositories"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHandler := r.Header.Get("Authorization")

		if authHandler == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHandler, "Bearer ") {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHandler, "Bearer ")

		if token == "" {
			http.Error(w, "Token is missing", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.GetString("JWT_SECRET", "default_secret")), nil
		})

		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		userId, ok := claims["id"].(float64)
		email, okEmail := claims["email"].(string)

		if !ok || !okEmail {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		fmt.Println("Authenticated user Id:", int64(userId), "Email:", email)

		ctx := context.WithValue(r.Context(), "userId", int64(userId))
		ctx = context.WithValue(ctx, "email", email)

		next.ServeHTTP(w,r.WithContext(ctx))
	})
}


func RequireAllRoles(roles ...string) func(http.Handler) http.Handler{


	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			userId := r.Context().Value("id").(int64)

			dbConn, dbErr := db.SetUpDB()

			if dbErr != nil {
				fmt.Println("Error while setting up Database!")
				return 
			}

			urr := repo.NewUserRoleRepository(dbConn)

			hasAllowedRoles, hasRolesErr := urr.HasRoles(userId,roles)

			if hasRolesErr != nil {
				http.Error(w, "Error checking user roles: "+hasRolesErr.Error(), http.StatusInternalServerError)
				return
			}

			if !hasAllowedRoles {
				http.Error(w, "Forbidden: You do not have the required roles", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w,r)

		})
	}
}	