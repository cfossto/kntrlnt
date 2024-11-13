package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func folderContentStringOutput(folderName string) []string {

	entries, err := os.ReadDir(folderName)
	var contentList []string

	if err != nil {

		fmt.Println("Big Yikes")
	}

	for _, entry := range entries {
		contentList = append(contentList, entry.Name())
	}

	return contentList
}

func folderContentsByType(folderName string) map[string]string {

	folderContentMap := make(map[string]string)
	entries, err := os.ReadDir(folderName)

	if err != nil {
		fmt.Println("Big, big, yikes!")
	}

	for _, entry := range entries {

		var fileType string

		if entry.IsDir() {
			fileType = "directory"
		} else if strings.Contains(entry.Name(), "json") {
			fileType = "json"
		} else {
			fileType = "img"
		}

		folderContentMap[folderName+"/"+entry.Name()] = fileType
	}

	return folderContentMap
}

func main() {

	var active_folder string = os.Getenv("KTLFOLDER")

	r := gin.Default()
	r.Static("/static", "static")

	r.GET("/", func(c *gin.Context) {

		var Contents []string = folderContentStringOutput("./" + active_folder)
		r.LoadHTMLGlob("templates/index.tmpl")
		c.HTML(http.StatusOK, "index.tmpl", gin.H{

			"FileList": Contents,
		})
	})

	r.GET("/:foldername", func(ctx *gin.Context) {
		folderName := ctx.Param("foldername")

		var fullname string = active_folder + "/" + folderName
		var folder_content map[string]string = folderContentsByType(fullname)

		r.LoadHTMLGlob("templates/index.tmpl")
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"TypesAndUrl": folder_content,
		})
	})

	r.GET("/sub/*subfolder", func(ctx *gin.Context) {

		folderName := ctx.Param("subfolder")
		var fullname string = folderName
		var folder_content map[string]string = folderContentsByType(fullname)

		r.LoadHTMLGlob("templates/index.tmpl")
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"TypesAndUrl": folder_content,
		})

	})

	r.GET("/image/*imagename", func(ctx *gin.Context) {

		imagename := ctx.Param("imagename")
		fmt.Println(imagename)

		ctx.File("./" + imagename)

	})

	r.Run(":8001")

}
