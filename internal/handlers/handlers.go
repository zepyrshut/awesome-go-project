package handlers

import (
	"awesome-go-project/internal/configuration"
	"awesome-go-project/internal/core"
	"fmt"
	"github.com/extrame/xls"
	"html/template"
	"mime/multipart"
	"net/http"
)

var app *configuration.Application

func NewHandlers(a *configuration.Application) {
	app = a
}

// Basic template for the upload page.
var templates = template.Must(template.ParseFiles("templates/index.html"))

// GetUpload shows the upload page.
func GetUpload(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		app.Logger.Error("Error executing template", err)
		JSON(http.StatusInternalServerError, w, `{"error": "Error executing template"}`)
		return
	}
}

// PostUpload handles the upload of the xls file and prints the contents to the console.
func PostUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		app.Logger.Error("Error parsing multipart form", err)
		JSON(http.StatusInternalServerError, w, `{"error": "Error parsing multipart form"}`)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		app.Logger.Error("Error retrieving the file", err)
		JSON(http.StatusBadRequest, w, `{"error": "Error retrieving the file"}`)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			app.Logger.Error("Error closing file", err)
			JSON(http.StatusInternalServerError, w, `{"error": "Error closing file"}`)
			return
		}
	}(file)

	app.Logger.Info(fmt.Sprintf("Uploaded File: %+v", handler.Filename))

	openFile, err := handler.Open()
	if err != nil {
		app.Logger.Error("Error opening file", err)
		JSON(http.StatusInternalServerError, w, `{"error": "Error opening file"}`)
		return
	}
	defer func(openFile multipart.File) {
		err := openFile.Close()
		if err != nil {
			app.Logger.Error("Error closing file", err)
			JSON(http.StatusInternalServerError, w, `{"error": "Error closing file"}`)
			return
		}
	}(openFile)

	xlsFile, err := xls.OpenReader(openFile, "utf-8")
	if err != nil {
		app.Logger.Error("Error opening xls file", err)
		JSON(http.StatusInternalServerError, w, `{"error": "Error opening xls file"}`)
		return
	}
	err = core.PrintXlsFile(xlsFile)
	if err != nil {
		app.Logger.Error("Error printing xls file", err)
		JSON(http.StatusInternalServerError, w, `{"error": "Error printing xls file"}`)
		return
	}

	JSON(http.StatusOK, w, `{"message": "success"}`)
}

func JSON(code int, w http.ResponseWriter, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(data))
}
