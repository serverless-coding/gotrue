package api

import (
	"fmt"
	"net/http"

	"github.com/netlify/gotrue/models"
	"github.com/netlify/gotrue/storage"
	"github.com/sirupsen/logrus"
)

// Logout is the endpoint for logging out a user and thereby revoking any refresh tokens
func (a *API) Logout(w http.ResponseWriter, r *http.Request) error {
	defer func() {
		if r := recover(); r != nil {
			logrus.Error("panic,err:", r)
		}
	}()
	ctx := r.Context()
	instanceID := getInstanceID(ctx)
	fmt.Println("-------instanceId------")

	a.clearCookieToken(ctx, w)

	fmt.Println("-------cookie------")

	u, err := getUserFromClaims(ctx, a.db)
	if err != nil {
		return unauthorizedError("Invalid user").WithInternalError(err)
	}
	fmt.Println("-------claims------")

	err = a.db.Transaction(func(tx *storage.Connection) error {
		if terr := models.NewAuditLogEntry(tx, instanceID, u, models.LogoutAction, nil); terr != nil {
			return terr
		}
		return models.Logout(tx, instanceID, u.ID)
	})
	if err != nil {
		return internalServerError("Error logging out user").WithInternalError(err)
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
