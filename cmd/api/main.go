package main

import (
	"awesome-go-project/internal/router"
	"fmt"
	"net/http"
)

func main() {

	//f, _ := xls.Open("example.xls", "utf-8")
	//
	//sheet := f.GetSheet(0)
	//
	//fmt.Println(sheet.Row(1).Col(2))
	//
	//if sheet1 := f.GetSheet(0); sheet1 != nil {
	//	fmt.Println("total lines", sheet1.MaxRow, sheet1.Name)
	//	for i := 0; i <= (int(sheet1.MaxRow)); i++ {
	//		row1 := sheet1.Row(i)
	//		lastCol := row1.LastCol()
	//		for j := 0; j <= (int(lastCol)); j++ {
	//			col := row1.Col(j)
	//			fmt.Print(col)
	//		}
	//	}
	//}

	srv := &http.Server{
		Addr:    "192.168.1.45:8080",
		Handler: router.Router(),
	}

	fmt.Println("Server is running on port 8080")

	srv.ListenAndServe()
}
