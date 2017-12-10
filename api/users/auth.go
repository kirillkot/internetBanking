package users

import (
	"net/http"

	"internetBanking/api/models"
)

// AuthMiddleware ...
func (v *View) AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			v.Failure(w, "Invalid basic auth header", http.StatusUnauthorized)
			return
		}

		user, where := &models.User{}, &models.User{Name: username}
		if err := v.DB().Find(user, where).Error; err != nil {
			v.Failure(w, "auth middl: get user: "+err.Error(), http.StatusForbidden)
			return
		}
		if user.Password != password {
			v.Failure(w, "auth middl: invalid password", http.StatusForbidden)
			return
		}

		r = models.SetUserToRequest(r, user)

		h.ServeHTTP(w, r)
	})
}
