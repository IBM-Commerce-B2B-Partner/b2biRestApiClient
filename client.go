package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/go-resty/resty"
	"github.com/tealeg/xlsx"
)

type CodeList struct {
	CodeListName  string `json:"codeListName"`
	VersionNumber int    `jdon:"versionNumber"`
	Codes         []struct {
		SenderCode   string `json:"senderCode"`
		ReceiverCode string `json:"receiverCode"`
		Description  string `json:"Description"`
	} `json:"codes"`
}

func main() {
	var username string
	var pword string

	var host string
	var port string
	var codelistname string
	var codelistversion string
	var allValuesSet bool = true

	flag.StringVar(&username, "username", "", "specify username for login")
	flag.StringVar(&pword, "password", "", "specify password for login")
	flag.StringVar(&host, "host", "", "specify host")
	flag.StringVar(&port, "port", "", "specify port")
	flag.StringVar(&codelistname, "codelistname", "", "specify codelistname")
	flag.StringVar(&codelistversion, "codelistversion", "", "specify codelistversion")

	flag.Parse()
	flag.VisitAll(func(arg1 *flag.Flag) {
		if len(arg1.Value.String()) == 0 {
			allValuesSet = false
		}
	})
	if allValuesSet == false {
		fmt.Printf("Please use with these Parameters\n")
		flag.PrintDefaults()
		return
	}
	var url string
	url = fmt.Sprintf("http://%v:%v/B2BAPIs/svc/codelists/%v:||%v", host, port, codelistname, codelistversion)
	fmt.Printf("Download from %v\n", url)
	resp, err := resty.R().
		SetBasicAuth(username, pword).
		Get(url)

	if err != nil {
		fmt.Printf("\nError %v", err)
		return
	}

	var clVar CodeList

	err2 := json.Unmarshal([]byte(resp.Body()), &clVar)

	if err2 != nil {
		fmt.Printf("Unmarshalling error %v\n", err2)
	}
	fmt.Printf("Successfully downloaded %v entries", len(clVar.Codes))
	writeToXls(clVar)
}

func typeOf(v interface{}) string {
	return fmt.Sprintf("%T\n", v)
}

func writeToXls(cl CodeList) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	// var cell *xlsx.Cell
	var err error

	filename := fmt.Sprintf("%v-%v.xlsx", cl.CodeListName, cl.VersionNumber)

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	// First add header
	row = sheet.AddRow()
	sendercell := row.AddCell()
	receivercell := row.AddCell()
	descriptioncell := row.AddCell()
	sendercell.Value = "SenderCode"
	receivercell.Value = "ReceiverCode"
	descriptioncell.Value = "CodesDescription"

	// Then add rows

	for zaehler := 0; zaehler < len(cl.Codes); zaehler++ {
		row = sheet.AddRow()
		sendercell := row.AddCell()
		receivercell := row.AddCell()
		descriptioncell := row.AddCell()
		sendercell.Value = cl.Codes[zaehler].SenderCode
		receivercell.Value = cl.Codes[zaehler].ReceiverCode
		descriptioncell.Value = cl.Codes[zaehler].Description
	}

	err = file.Save(filename)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
