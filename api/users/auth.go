package users

import (
	"context"
	"errors"
	"net/http"

	"internetBanking/api/models"
)

const (
	userkey = "user"
)

// SetUserToRequest ...
func SetUserToRequest(req *http.Request, user *models.User) *http.Request {
	ctx := context.WithValue(req.Context(), userkey, user)
	return req.WithContext(ctx)
}

// UserFromRequest ...
func UserFromRequest(req *http.Request) (*models.User, error) {
	user, ok := req.Context().Value(userkey).(*models.User)
	if !ok {
		return nil, errors.New("user is not set")
	}
	return user, nil
}

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

		r = SetUserToRequest(r, user)

		h.ServeHTTP(w, r)
	})
}
