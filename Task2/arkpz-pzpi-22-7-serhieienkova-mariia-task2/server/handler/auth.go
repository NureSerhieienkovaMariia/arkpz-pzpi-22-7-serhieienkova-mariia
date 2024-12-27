package handler

import (
	"clinic/server/service"
	"fmt"
	"net/http"
	"strings"

	"clinic/server/structures"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input structures.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Generate JWT access token for the new user
	accessToken, err := h.services.Authorization.GenerateTokenByUserId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Generate JWT refresh token for the new user
	refreshToken, err := h.services.Authorization.GenerateTokenByUserId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":           id,
		"access_jwt_token":  accessToken.Token,
		"refresh_jwt_token": refreshToken.Token,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(fmt.Sprintf("input: %v", input))

	userToken, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Generate JWT refresh token for the authenticated user
	refreshToken, err := h.services.Authorization.GenerateTokenByUserId(userToken.UserId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":           userToken.UserId,
		"access_jwt_token":  userToken.Token,
		"refresh_jwt_token": refreshToken.Token,
	})
}

type refreshTokenInput struct {
	RefreshToken string `json:"refresh_jwt_token" binding:"required"`
}

func (h *Handler) refreshToken(c *gin.Context) {
	var input refreshTokenInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newAccessToken, newRefreshToken, err := h.services.Authorization.RefreshToken(input.RefreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":           newAccessToken.UserId,
		"access_jwt_token":  newAccessToken.Token,
		"refresh_jwt_token": newRefreshToken.Token,
	})
}

func (h *Handler) currentUser(c *gin.Context) {
	rawAuthToken := readRawAuthToken(c)

	tokenClaims, err := service.ParseToken(rawAuthToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.services.Authorization.GetUserById(tokenClaims.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":                      user.Id,
		"email":                   user.Email,
		"name":                    user.Name,
		"surname":                 user.Surname,
		"premium_expiration_date": user.PremiumExpirationDate,
	})
}

func readRawAuthToken(c *gin.Context) string {
	header := c.GetHeader("Authorization")
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
	}
	return headerParts[1]
}

//var basicUserActions = []string{GetUserByIdAction, GetTeamByIdAction}
//var managerActions = append([]string{CreateTeamAction}, basicUserActions...)
//var adminActions = append([]string{DeleteUserAction}, managerActions...)

//var actionsForLevelAccess = map[int][]string{
//	1: basicUserActions,
//	2: managerActions,
//	3: adminActions,
//}

//func BasicAuthPermission(permission string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		header := c.GetHeader(authorizationHeader)
//		if header == "" {
//			newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
//			return
//		}
//
//		headerParts := strings.Split(header, " ")
//		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
//			newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
//			return
//		}
//
//		if len(headerParts[1]) == 0 {
//			newErrorResponse(c, http.StatusUnauthorized, "token is empty")
//			return
//		}
//
//		tokenClaims, err := service.ParseToken(headerParts[1])
//		if err != nil {
//			newErrorResponse(c, http.StatusUnauthorized, err.Error())
//			return
//		}
//
//		if !userHasPermission(permission, tokenClaims.AccessLevel) {
//			newErrorResponse(c, http.StatusUnauthorized, "Action not available for this access level")
//		}
//
//		c.Set(userIdCtx, tokenClaims.Id)
//		c.Next()
//	}
//}
//
//func userHasPermission(permission string, accessLevel int) bool {
//	availableActions := actionsForLevelAccess[accessLevel]
//
//	for _, action := range availableActions {
//		if action == permission {
//			return true
//		}
//	}
//	return false
//}
