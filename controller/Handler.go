package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"work/src/web/service"
)

//访问主页的控制器
func IndexHandle(c *gin.Context)  {
	prototypeList, err := service.GetAllProtoType()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.html", nil)
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"prototype_list" : prototypeList,
	})
}