package users

import (
	"errors"
	"net/http"
	"strconv"

	"internetBanking/api/models"
)

const (
	authCookieKey = "auth"
)

func setAuthCookie(w http.ResponseWriter, user *models.User) {
	cookie := &http.Cookie{
		Name:  authCookieKey,
		Value: strconv.FormatUint(uint64(user.ID), 10),
		Path:  "/",
	}
	http.SetCookie(w, cookie)
}

func userIDFromCookie(r *http.Request) (uint, error) {
	cookie, err := r.Cookie(authCookieKey)
	if err != nil {
		return 0, errors.New("get cookie: " + err.Error())
	}

	id, err := strconv.ParseUint(cookie.Value, 10, 32)
	if err != nil {
		return 0, errors.New("unquote: " + err.Error())
	}

	return uint(id), nil
}

// AuthMiddleware ...
func (v *View) AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := userIDFromCookie(r)
		if err != nil {
			v.Failure(w, "Invalid auth cookie: "+err.Error(), http.StatusUnauthorized)
			return
		}

		user := &models.User{}
		if err := v.DB().Where("id = ?", id).Find(user).Error; err != nil {
			v.Failure(w, "auth middl: get user: "+err.Error(), http.StatusForbidden)
			return
		}

		r = models.SetUserToRequest(r, user)
		setAuthCookie(w, user)

		h.ServeHTTP(w, r)
	})
}
