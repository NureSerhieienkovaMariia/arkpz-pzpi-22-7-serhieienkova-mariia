package handler

import (
	"clinic/server/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userIdCtx           = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	tokenClaims, err := service.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userIdCtx, tokenClaims.Id)
}

//func getUserId(c *gin.Context) (int, error) {
//	id, ok := c.Get(userIdCtx)
//	if !ok {
//		return 0, errors.New("user id not found")
//	}
//
//	idInt, ok := id.(int)
//	if !ok {
//		return 0, errors.New("user id is of invalid type")
//	}
//
//	return idInt, nil
//}

//func (h *Handler) checkIsUserAdmin(c *gin.Context) (bool, error) {
//	id, err := getUserId(c)
//	if err != nil {
//		return false, err
//	}
//
//	user, err := h.services.UserAction.GetById(id)
//	if err != nil {
//		return false, err
//	}
//
//	if user.AccessLevel >= 3 {
//		return true, nil
//	}
//
//	return false, nil
//}
