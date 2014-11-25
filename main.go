package main

import (
    "github.com/gin-gonic/gin"
)

func main(){
    r := gin.Default()
    r.LoadHTMLTemplates("templates/*.html")

    r.GET("/", func(c *gin.Context){
        obj := gin.H{".Title": "Main website"}
        c.HTML(200, "index.html", obj)
    })

    r.GET("/containers", HandleContainers)
    r.POST("/post", HandlePost)
    r.GET("/json", HandleJSON)

    r.Run(":8080")
}
