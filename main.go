package main

import (
	"fmt"
	"net/http"
	"os"

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
		
		folderContentMap[entry.Name()] = entry.Type().String()
	}

	return folderContentMap
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		var Contents []string = folderContentStringOutput("./example-folder")
		r.LoadHTMLGlob("templates/index.tmpl")
		c.HTML(http.StatusOK, "index.tmpl", gin.H{

			"FileList": Contents,
		})
	})

	r.GET("/:foldername", func(ctx *gin.Context) {
		folderName := ctx.Param("foldername")

		var fullname string = "example-folder"+"/"+folderName

		var folder_content map[string]string =  folderContentsByType(fullname)

		r.LoadHTMLGlob("templates/index.tmpl")
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"TypesAndUrl": folder_content,
		})

	})

	r.Run(":8001")

}
