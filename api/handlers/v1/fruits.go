package v1

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/auth/modles"
)

func (h *handlers)GetFruits(ctx *gin.Context){

	var req modles.GetList

	limit:=ctx.Query("limit")
	page:=ctx.Query("page")

	req.Limit,_=strconv.Atoi(limit)
	req.Page,_=strconv.Atoi(page)

	resp,err:=h.storage.FruitsRepo().GetFruits(context.Background(),req)

	if err!= nil{
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(200,resp)

}