package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"clinic/server/structures"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(c *gin.Context) {
	var input structures.User

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err.Error())
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.UserAction.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) getUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	user, err := h.services.UserAction.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) getUserByEmail(c *gin.Context) {
	email := c.Param("email")

	//rawAuthToken := readRawAuthToken(c)

	//tokenClaims, err := service.ParseToken(rawAuthToken)
	//if err != nil {
	//	newErrorResponse(c, http.StatusUnauthorized, err.Error())
	//	return
	//}

	user, err := h.services.UserAction.GetByEmail(email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	//if user.Id == tokenClaims.Id {
	//	newErrorResponse(c, http.StatusBadRequest, "Cannot retrieve current user with this method")
	//	return
	//}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input structures.UpdateUserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UserAction.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.UserAction.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
