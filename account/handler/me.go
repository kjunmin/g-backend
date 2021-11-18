package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kjunmin/g-backend/model"
	"github.com/kjunmin/g-backend/model/apperrors"
)

func (h *Handler) Me(c *gin.Context) {

	user, exists := c.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reasons %v", c)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
	}

	uid := user.(*model.User).UID

	u, err := h.UserService.Get(c, uid)

	if err != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := apperrors.NewNotFound("user", uid.String())

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
