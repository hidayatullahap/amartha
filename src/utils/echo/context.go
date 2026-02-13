package echo

import (
	"amartha/src/utils/constants"

	e "github.com/labstack/echo/v5"
)

type AuthenticatedUser struct {
	ID   int64
	Role string
}

func GetAuthUser(c *e.Context) AuthenticatedUser {
	id, _ := c.Get(constants.UserId).(*int64)
	role, _ := c.Get(constants.UserRole).(string)

	user := AuthenticatedUser{
		Role: role,
	}

	if id != nil {
		user.ID = *id
	}

	return user
}
