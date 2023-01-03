package handlers

import (
	"fmt"
	"github.com/extrame/xls"
	"html/template"
	"net/http"
)

func TestHttp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}

var templates = template.Must(template.ParseFiles("templates/index.html"))

func GetUpload(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func PostUpload(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	openFile, _ := handler.Open()
	defer openFile.Close()

	xlsFile, _ := xls.OpenReader(openFile, "utf-8")

	fmt.Println(xlsFile.NumSheets())

	fmt.Println(handler.Filename)
	fmt.Println(handler.Header)
	fmt.Println(handler.Size)

	w.WriteHeader(http.StatusOK)
}
