package middleware

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		files := form.File["files"]
		fmt.Println(files, "file")
		allowedExtensions := map[string]bool{
			".txt": true,
			".pdf": true,
		}
		var dataFiles []string
		for _, file := range files {
			ext := filepath.Ext(file.Filename)
			if !allowedExtensions[ext] {
				return c.JSON(http.StatusBadRequest, "The file extension is wrong. Allowed file extensions are .txt and .pdf")
			}

			src, err := file.Open()
			fmt.Println(src)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer src.Close()

			tempFile, err := ioutil.TempFile("uploads", "files-*.txt")
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer tempFile.Close()

			if _, err = io.Copy(tempFile, src); err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}

			data := tempFile.Name()
			filename := data[8:]

			// c.Set("dataFile", filename)
			dataFiles = append(dataFiles, filename)
		}
		c.Set("dataFiles", dataFiles)
		return next(c)
	}
}
