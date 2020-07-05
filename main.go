package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/volume/:vol", SetVolume)
	r.POST("/video/:videoUrl", PlayYoutubeVideo)


	r.Run(":8080")
}


func SetVolume(c *gin.Context){
	volume := c.Param("vol")

	cmd := fmt.Sprintf("amixer -q -M sset %s%", volume)

	volCmd := exec.Command("Set Loudness", cmd)
	fmt.Printf("LOG: Trying to set volume to %s%", volume)
	err := volCmd.Run()
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "exec successful"})
	}
}


func PlayYoutubeVideo(c *gin.Context) {
	videoUrl := c.Param("videoUrl")

	cmd := fmt.Sprintf("mpv %s --no-video", videoUrl)
	vidCmd := exec.Command("Set Video", cmd)

	err := vidCmd.Run()
	fmt.Printf("LOG: Trying to play video at %s...", videoUrl)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
	} else {
		c.JSON(200, gin.H{"msg": "video loaded"})
	}
}
