// //package CaptureAndStream
package main

import (
	"CaptureAndStream/util"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var frameNum = 0
var mutex sync.Mutex

func start(port string, quality string) {
	//Creating new router
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.LoadHTMLGlob("htmls/*")

	//return web viewer
	r.GET("/", func(c *gin.Context) {
		var w, h, fps int
		if quality == "1" {
			w = 480
			h = 640
			fps = 19
		} else if quality == "2" {
			w = 720
			h = 1280
			fps = 15
		} else if quality == "3" {
			w = 1080
			h = 1920
			fps = 10
		} else {
			w = 480
			h = 640
			fps = 10
		}
		util.W = w
		util.H = h

		c.HTML(http.StatusOK, "index.html", gin.H{
			"ip":   util.GetIP(),
			"port": port,
			"w":    w,
			"h":    h,
			"fps":  fps,
		})
	})

	//create new frame
	r.GET("/newFrame", func(c *gin.Context) {
		c.Header("Cache-control", "no-cache")
		c.Header("Allow-Control-Allow-Origin", "*")

		frameCode := frameNum
		frameNum++

		capture := util.Capture(0)
		compressed := util.Img2CompressedPng(capture)
		util.AddData(compressed)
		c.JSON(http.StatusOK, gin.H{"code": fmt.Sprintf("%d", frameCode)})
	})

	//return frame (png)
	r.GET("/frames/:code", func(c *gin.Context) {
		c.Header("Cache-control", "no-cache")
		c.Header("Allow-Control-Allow-Origin", "*")
		code := c.Param("code")
		code = strings.Replace(code, ".png", "", 1)

		data := util.ReadData(code)
		c.Data(http.StatusOK, "image/png", data)
	})

	//run server
	r.Run(":" + port)
}

func main() {
	var port, quality string
	fmt.Print("Enter port: ")
	fmt.Scan(&port)
	fmt.Println("Quality")
	fmt.Println("1. 480p\t2. 720p\t3. 1080p")
	fmt.Scan(&quality)
	start(port, quality)
}
