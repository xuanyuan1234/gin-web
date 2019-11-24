package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"work/src/web/controller"
	"work/src/web/dao/db"
)

func main()  {
	router := gin.Default()
	dns := "root:root1234@tcp(localhost:3306)/tag?parseTime=true"
	err := db.Init(dns)

	if err != nil {
		panic(err)
	}

	//加载静态文件
	router.StaticFS("/static", http.Dir("F://2019H2/GoWork/src/web/static"))

	//加载模板
	router.LoadHTMLGlob("F://2019H2/GoWork/src/web/views/*")

	router.GET("/", controller.IndexHandle)

	_ = router.Run()
}