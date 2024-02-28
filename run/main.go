package main

import (
	"github.com/gin-gonic/gin"
	kakao "github.com/tae2089/kakao-bot-go"
)

func main() {
	g := gin.Default()
	g.POST("/demo", func(c *gin.Context) {
		button := kakao.NewButton("label", func(value string) kakao.ButtonType {
			return kakao.WithWebLink("https://www.naver.com")
		})
		textCard := kakao.NewTextCardBlock("Title", "Description", []kakao.Button{*button})
		response := kakao.Response{}
		response.Version = "2.0"
		response.Template = kakao.Template{
			Blocks: []kakao.Block{
				textCard,
			},
		}
		c.JSON(200, response)
	})

	g.Run(":8080")
}
