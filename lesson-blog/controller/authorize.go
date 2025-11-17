package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) (uint, error) {
	user_id_value, exist := c.Get("userID")
	if !exist {
		return 0, errors.New("userID not found")
	}

	user_id, err := strconv.Atoi(user_id_value.(string))

	return uint(user_id), err
}
