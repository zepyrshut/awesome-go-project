package core

import (
	"fmt"
	"github.com/extrame/xls"
	"os"
	"text/tabwriter"
)

// PrintXlsFile prints the contents of a xls file to the console.
// File from handler is passed to this function.
func PrintXlsFile(file *xls.WorkBook) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 3, 2, ' ', 0)

	for i := 0; i < file.NumSheets(); i++ {
		sheet := file.GetSheet(i)
		for j := 0; j <= (int(sheet.MaxRow)); j++ {
			lastCol := sheet.Row(j).LastCol()
			var row []string
			for k := 0; k <= (int(lastCol)); k++ {
				cell := sheet.Row(j).Col(k)
				row = append(row, fmt.Sprintf("%s\t", cell))
			}

			_, err := fmt.Fprintln(w, row)
			if err != nil {
				return err
			}
		}

		err := w.Flush()
		if err != nil {
			return err
		}
	}

	return nil
}
