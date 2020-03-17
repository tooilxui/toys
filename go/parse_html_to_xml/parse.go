package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type snMapping struct {
	StdSpecSn string
	BomSn     []string
}

type ItemInfo struct {
	// CCC, Name string
	CCC     string
	ENNames []string `xml:"英文名稱"`
	Details []string `xml:"中文名稱"`
}

type Material struct {
	SID      string   `xml:"原料項次" json:"sid,omitempty"`
	Info     ItemInfo `json:"info,omitempty"`
	Quantity string   `xml:"數量" json:"quantity,omitempty"`
	Unit     string   `xml:"單位" json:"unit,omitempty"`
}

// 每個退稅標準文號有99個核退文號(即BOM表)
type Record struct {
	BomSn     string `xml:"申請核退文號"`
	Product   ItemInfo
	Materials []Material
}

type Stdspc struct {
	StdSpecSn string `xml:"核退標準文號"`
	Records   []Record
}

type Output struct {
	Stdspc []Stdspc
}

func parseOneRecord(i int, s *goquery.Selection) Record {
	var r Record

	r.BomSn = snMap[get_sSort(specID)].BomSn[i]

	iter := 0
	assignNextIfMatch := func(sel *goquery.Selection, s string, assign func(*goquery.Selection)) bool {

		if strings.TrimSpace(sel.Eq(iter).Text()) == s {
			iter++
			assign(sel.Eq(iter))
			return true
		}
		return false
	}
	getItems := func(s *goquery.Selection) []string {
		return s.Find("li").Map(func(i int, s *goquery.Selection) string {
			//html較長的名稱都會有"\n                          "，取代後才是原有正常的值
			return strings.Replace(strings.TrimSpace(s.Text()), "\n                          ", "", -1)
		})
	}

	// Four such tr sections per product record: 外銷品使用原料數量計算表（M327）, 外銷成品, 應用原料, 原料應用數量換算公式說明
	sections := s.Children().Filter("td").Children().Filter("table").Children().Filter("tbody").Children().Filter("tr")
	productInfo := sections.Eq(1).Children().Filter("td").Children().Filter("table").Children().Filter("tbody").Children().Filter("tr").Children().Filter("td")
	for iter = 0; iter < productInfo.Length()-1; iter++ {
		if assignNextIfMatch(productInfo, "C.C.C. Code", func(s *goquery.Selection) { r.Product.CCC = s.Text() }) {
			continue
		}
		if assignNextIfMatch(productInfo, "英文名稱", func(s *goquery.Selection) { r.Product.ENNames = getItems(s) }) {
			continue
		}
		if assignNextIfMatch(productInfo, "貨品規格", func(s *goquery.Selection) { r.Product.Details = getItems(s) }) {
			continue
		}
	}

	materialsInfo := sections.Eq(2).Children().Filter("td").Children().Filter("table").Children().Filter("tbody").Children().Filter("tr").Children().Filter("td")

	// materialsInfoStrings := materialsInfo.Map(func(i int, s *goquery.Selection) string { return strings.TrimSpace(s.Text()) })
	for iter = 0; iter < materialsInfo.Length()-1; iter++ {
		if assignNextIfMatch(materialsInfo, "原料項次", func(s *goquery.Selection) { r.Materials = append(r.Materials, Material{SID: s.Text()}) }) {
			continue
		}
		if assignNextIfMatch(materialsInfo, "C.C.C. Code", func(s *goquery.Selection) { r.Materials[len(r.Materials)-1].Info.CCC = s.Text() }) {
			continue
		}
		if assignNextIfMatch(materialsInfo, "應用數量", func(s *goquery.Selection) { r.Materials[len(r.Materials)-1].Quantity = s.Text() }) {
			continue
		}
		if assignNextIfMatch(materialsInfo, "英文名稱", func(s *goquery.Selection) { r.Materials[len(r.Materials)-1].Info.ENNames = getItems(s) }) {
			continue
		}
		if assignNextIfMatch(materialsInfo, "數量單位", func(s *goquery.Selection) { r.Materials[len(r.Materials)-1].Unit = strings.TrimSpace(s.Text()) }) {
			continue
		}
	}

	return r
}

func parse(r io.Reader) ([]Record, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	// func parseOneRecord(i int, s *goquery.Selection)
	recordSelections := doc.Find("#viewCaseTable").Find("#dataTR")
	records := make([]Record, recordSelections.Length())
	// fmt.Println(recordSelections.Html())
	fmt.Println(specID)

	recordSelections.Each(func(i int, s *goquery.Selection) {
		// fmt.Printf("%v\n", i)
		records[i] = parseOneRecord(i, s)
	})

	return records, nil
}

var specID string
var snMap = []snMapping{}

func main() {

	var Opt = Output{}

	snMap = make([]snMapping, 13)
	mapListLoad()

	files, err := ioutil.ReadDir("./Input")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.Contains(f.Name(), ".htm") {

			specID = strings.Replace(f.Name(), ".htm", "", -1)

			f, err := os.Open("./Input/" + f.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			records, err := parse(f)
			if err != nil {
				log.Fatal(err)
			}

			var Stdspc = Stdspc{}
			Stdspc.StdSpecSn = specID
			Stdspc.Records = records

			Opt.Stdspc = append(Opt.Stdspc, Stdspc)
		}
	}

	fo, err := os.Create("./output/Output.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	b, err := xml.Marshal(Opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(fo, "%v\r\n", string(b))

}

func get_sSort(matchString string) int {
	for i := 0; i < 13; i++ {
		if snMap[i].StdSpecSn == matchString {
			return i
		}
	}
	return 99
}

// this list is from smhuang
func mapListLoad() {
	var fstring string
	var bstring string
	bstring = "3P7E"

	for i := 0; i < 13; i++ {
		switch i {
		case 0:
			fstring = "05012193"
			snMap[i].StdSpecSn = "105R381217"
		case 1:
			fstring = "07037903"
			snMap[i].StdSpecSn = "107R446799"
		case 2:
			fstring = "07037907"
			snMap[i].StdSpecSn = "107R447186"
		case 3:
			fstring = "07037904"
			snMap[i].StdSpecSn = "107R447783"
		case 4:
			fstring = "07037906"
			snMap[i].StdSpecSn = "107R448204"
		case 5:
			fstring = "07037908"
			snMap[i].StdSpecSn = "107R449499"
		case 6:
			fstring = "07037900"
			snMap[i].StdSpecSn = "107R450047"
		case 7:
			fstring = "07037901"
			snMap[i].StdSpecSn = "107R451590"
		case 8:
			fstring = "07037899"
			snMap[i].StdSpecSn = "107R452129"
		case 9:
			fstring = "07037913"
			snMap[i].StdSpecSn = "107R452131"
		case 10:
			fstring = "07037902"
			snMap[i].StdSpecSn = "107R453224"
		case 11:
			fstring = "07077310"
			snMap[i].StdSpecSn = "107R464627"
		case 12:
			fstring = "07085726"
			snMap[i].StdSpecSn = "107R466239"
		}
		snMap[i].BomSn = make([]string, 99)
		for j := 0; j < 99; j++ {
			var sn string
			if j < 9 {
				sn = "0" + strconv.Itoa(j+1)
			} else {
				sn = strconv.Itoa(j + 1)
			}
			snMap[i].BomSn[j] = fstring + sn + bstring
		}
	}

}
