package cookie

import (
	"net/http"
	"os"
)

func SetAuthCookie(
	w http.ResponseWriter,
	token string,
) {
	isProd := os.Getenv("ENV") == "production"

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   isProd,
		SameSite: func() http.SameSite {
			if isProd {
				return http.SameSiteStrictMode
			}
			return http.SameSiteLaxMode
		}(),
	})
}
