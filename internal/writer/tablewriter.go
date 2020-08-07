package writer

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
)

type Table struct {
	table *tablewriter.Table
}

func TableWriter(writer io.Writer) Writer {
	table := tablewriter.NewWriter(writer)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	return &Table{ table}
}

func (t *Table) WriteOne(data interface{}) error{
	return 	JSONWriter(os.Stdout).WriteOne(data)
}

func (t *Table) WriteTable(data interface{}) error {

	allHeaders := make(map[string]bool, 1)
	for rowData := range data.([]*interface{}) {
		for _, s := range getHeaders("", rowData) {
			allHeaders[s] = true
		}
	}
	var header []string
	for key, _ := range allHeaders {
		header = append(header, key)
	}
	t.table.SetHeader(header)


	t.table.Render()

	return nil
}

func (t *Table) WriteRow(data interface{}) error {



	data2 := [][]string{
		{"node1.example.com", "Ready", "compute", "1.11"},
	}

	t.table.AppendBulk(data2) // Add Bulk Data

	return nil
}


//TODO: this is not tested yet
func getHeaders(prefix string, data interface{}) []string{
	var resultHeaders []string
	var jsonToMap map[string]interface{}

	if prefix != "" {
		prefix = prefix + "."
	}

	dataJsonString, _ := json.Marshal(data)

	err := json.Unmarshal(dataJsonString, &jsonToMap)

	if err != nil {
		fmt.Printf("%v", err)
	}

	for key, value := range jsonToMap {
		if _, ok := value.(string); ok {
			resultHeaders = append(resultHeaders, prefix+key)
		} else {
			resultHeaders = append(resultHeaders, getHeaders(key, value)...)
		}
	}

	return resultHeaders
}