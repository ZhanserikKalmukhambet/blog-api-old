package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanserikKalmukhambet/blog-api/internal/entity"
	"log"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h *Handler) createUser(ctx *gin.Context) {
	var req entity.User

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	err = h.srvs.CreateUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}
