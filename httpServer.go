// //package CaptureAndStream
package main

import (
	"CaptureAndStream/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func start(port string) {
	r := gin.Default()
	r.LoadHTMLGlob("htmls/*")

	r.GET("/", func(c *gin.Context) {
		//page, _ := os.ReadFile("index.html")
		//c.String(http.StatusOK, string(page))
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ip":   util.GetIP(),
			"port": port,
		})
	})

	r.GET("/newFrame", func(c *gin.Context) {
		frameCode := c.Query("frame")
		c.Header("Cache-control", "no-cache")
		c.Header("Allow-Control-Allow-Origin", "*")

		capture := util.Capture(0)
		compressed := util.Img2CompressedPng(capture)
		//util.Byte2png(compressed, frameCode)
		util.AddData(compressed)
		//c.Data(http.StatusOK, "image/png", compressed)
		c.String(http.StatusOK, string(frameCode))
	})

	r.GET("/frames/:code", func(c *gin.Context) {
		c.Header("Cache-control", "no-cache")
		c.Header("Allow-Control-Allow-Origin", "*")
		code := c.Param("code")
		code = strings.Replace(code, ".png", "", 1)

		data := util.ReadData(code)
		c.Data(http.StatusOK, "image/png", data)
	})

	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func main() {
	fmt.Print("Enter port: ")
	var port string
	fmt.Scan(&port)
	start(port)
}
