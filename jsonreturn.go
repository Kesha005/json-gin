package jsonreturn

import "github.com/gin-gonic/gin"

var ctx *gin.Context

type ErrorMessage struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type SuccessMsg struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func (err *ErrorMessage)ReturnError() {
	ctx.JSON(400,gin.H{
		"status":err.Status,
		"message":err.Message,
	})
}

func (su *SuccessMsg)ReturnSuccess(){
	ctx.JSON(400,gin.H{
		"status":su.Status,
		"data":su.Data,
	})
}