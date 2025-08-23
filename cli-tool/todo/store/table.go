package store

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func PrintTable(taskList []Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Titulo", "Completado"})

	var data [][]string

	for _, item := range taskList {
		data = append(data, []string{
			strconv.Itoa(item.ID),
			item.Title,
			checkDone(item.Done)})
	}
	table.AppendBulk(data)
	table.Render()
}

func checkDone(b bool) string {
	if b {
		return "SI"
	}
	return "NO"
}
