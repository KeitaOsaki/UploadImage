package controller

import (
	"encoding/base64"
	"file-server/pkg/model"
	"file-server/pkg/view"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
)

func UploadBase64Image() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
			var maxSize int64
			maxSize =

		*/
		var urls []string

		var request model.UploadRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			log.Println("[ERROR] Faild Bind JSON　\n ", err)
			c.JSON(http.StatusBadRequest, "Request is error")
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"Request is error",
			)
			return
		}

		for _, images := range request.Images {

			uuID, err := uuid.NewRandom()
			if err != nil {
				log.Println("uuid generate is failed")
			}

			// 画像格納用のデイレクトリがない場合ディレクトリを作成
			err = os.MkdirAll("./uploadimages", os.ModePerm)
			if err != nil {
				log.Println("[ERROR] Faild Bind JSON　\n ", err)
				c.JSON(http.StatusInternalServerError, "InternalServerError")
				view.ReturnErrorResponse(
					c,
					http.StatusInternalServerError,
					"InternalServerError",
					"Unable to create a directory to store images",
				)
				return
			}

			data, _ := base64.StdEncoding.DecodeString(images.Image) //[]byte
			// 画像ファイルを作成。
			fileName := fmt.Sprintf("./uploadimages/%s.jpg", uuID)
			file, err := os.Create(fileName)
			if err != nil {
				log.Println("[ERROR] Faild Bind JSON　\n ", err)
				c.JSON(http.StatusInternalServerError, "Request is error")
				view.ReturnErrorResponse(
					c,
					http.StatusInternalServerError,
					"nternalServerError",
					"Unable to create file",
				)
				return

			}
			/*
				fileInter,_ := file.Stat()
				fileSeze := fileInter.Size()


			*/
			//どこかにファイルサイズ制限処理書きたい

			defer file.Close()

			file.Write(data)

			urlName := fmt.Sprintf("/%s.jpg", uuID)

			urls = append(urls, urlName)

		}

		c.JSON(http.StatusOK, view.UploadResponse(urls))
	}
}

func UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		var urls []string
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for _, file := range files {

			uuID, err := uuid.NewRandom()
			if err != nil {
				log.Println("uuid generate is failed")
			}

			fileName := fmt.Sprintf("images/%s%s", file.Filename, uuID)

			err = c.SaveUploadedFile(file, fileName)
			if err != nil {
				log.Println("[ERROR] Faild Bind JSON　\n ", err)
				c.JSON(http.StatusBadRequest, "Request is error")
				view.ReturnErrorResponse(
					c,
					http.StatusBadRequest,
					"Bad Request",
					"Request is error",
				)
				return
			}

			urlName := fmt.Sprintf("/%s%s", file.Filename, uuID)
			urls = append(urls, urlName)
		}

		c.JSON(http.StatusOK, view.UploadResponse(urls))
	}

}
