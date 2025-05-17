package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wenlng/go-captcha-assets/bindata/chars"
	"github.com/wenlng/go-captcha/v2/base/option"
	"image/png"
	"log"
	"net/http"

	"github.com/golang/freetype/truetype"
	"github.com/wenlng/go-captcha-assets/resources/fonts/fzshengsksjw"
	"github.com/wenlng/go-captcha-assets/resources/images_v2"
	"github.com/wenlng/go-captcha/v2/click"
)

var textCapt click.Captcha

func init() {
	builder := click.NewBuilder()

	fonts, err := fzshengsksjw.GetFont()
	if err != nil {
		log.Fatalln(err)
	}

	img, err := images.GetImages()
	if err != nil {
		log.Fatalln(err)
	}

	builder.SetResources(
		click.WithChars(chars.GetChineseChars()),
		click.WithFonts([]*truetype.Font{fonts}),
		click.WithBackgrounds(img),
	)

	builder.SetOptions(
		click.WithImageAlpha(0.5),
		click.WithShadowColor("#fffafa"),
		click.WithRangeVerifyLen(option.RangeVal{Min: 3, Max: 3}),
		click.WithRangeThumbImageSize(option.Size{Width: 200, Height: 40}),
	)

	textCapt = builder.Make()
}

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/img", getImg)
	}
	r.Run(":8089")
}

func getImg(c *gin.Context) {
	captData, err := textCapt.Generate()
	if err != nil {
		log.Fatalln(err)
	}

	dotData := captData.GetData()
	if dotData == nil {
		log.Fatalln(">>>>> generate err")
	}

	dots, _ := json.Marshal(dotData)
	fmt.Println(">>>>> ", string(dots))

	//err = captData.GetMasterImage().SaveToFile("master.jpg", option.QualityNone)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = captData.GetThumbImage().SaveToFile("thumb.png")
	//if err != nil {
	//	fmt.Println(err)
	//}

	image := captData.GetMasterImage().Get()

	var buf bytes.Buffer
	if err := png.Encode(&buf, image); err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "图片编码失败"})
		return
	}

	// 返回 PNG 图片给浏览器
	c.Data(http.StatusOK, "image/png", buf.Bytes())

}
