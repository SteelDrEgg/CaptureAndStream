// //package CaptureAndStream
package main

import (
	"CaptureAndStream/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var frameNum = 0
var quality int

func start(port string, resolution string, fps int) {
	//Creating new router
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.LoadHTMLGlob("htmls/*")

	//return web viewer
	r.GET("/", func(c *gin.Context) {
		var w, h int
		if resolution == "1" {
			w = 480
			h = 640
		} else if resolution == "2" {
			w = 720
			h = 1280
		} else if resolution == "3" {
			w = 1080
			h = 1920
		} else {
			w = 480
			h = 640
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
		compressed := util.Img2CompressedPng(capture, quality)
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
	var port, resolution string
	var fps int
	fmt.Print("Enter port: ")
	fmt.Scan(&port)
	fmt.Println("Quality")
	fmt.Println("1. 480p\t2. 720p\t3. 1080p")
	fmt.Scan(&resolution)
	fmt.Print("Enter quality (lo 0-100 hi): ")
	fmt.Scan(&quality)
	fmt.Print("Enter fps: ")
	fmt.Scan(&fps)
	start(port, resolution, fps)
}
