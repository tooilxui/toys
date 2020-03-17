package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// 生胎條碼格式
	// 廠(1) + 年(1) + 成代(8) +流水號 (5) +檢查碼(1)
	//
	// 流水號前兩碼32進位, 後三碼10進位
	// EX: 19090 => 0J 090

	currentBarcode := "98K239B0800J090Q"
	fmt.Printf("current barcode          => %s\n", currentBarcode)

	// 用條碼前15取得檢查碼
	checkCode := getCheckCode(currentBarcode[:15])

	//取36進位部分
	prefixSn := currentBarcode[10:12]
	//取10進位部分
	suffixSn := currentBarcode[12:15]

	//36進位轉10進位
	decode, _ := strconv.ParseInt(prefixSn, 32, 10)

	//與10進位結合, 得到原始序號
	realSnInString := fmt.Sprintf("%d%s", decode, suffixSn)
	fmt.Printf("current barcode sequence => %s\n", realSnInString)

	//取下一張條碼序號
	realSn, _ := strconv.Atoi(realSnInString)
	nextSn := realSn + 1
	fmt.Printf("next    barcode sequence => %d\n", nextSn)

	//序號分前兩碼跟後三碼, 並將序號轉為條碼
	prefixSn = strconv.Itoa(nextSn)[:2]
	suffixSn = strconv.Itoa(nextSn)[2:]

	//前兩碼轉36進位
	p, _ := strconv.Atoi(prefixSn)
	prefixSn = strings.ToUpper(strconv.FormatInt(int64(p), 36))
	if len(prefixSn) == 1 {
		prefixSn = fmt.Sprintf("0%s", prefixSn)
	}

	// 用新條碼前15碼取得檢查碼
	barcodePrefix := currentBarcode[:10] + prefixSn + suffixSn
	checkCode = getCheckCode(barcodePrefix)

	//組出新條碼
	newBarcode := barcodePrefix + checkCode
	fmt.Printf("next    barcode          => %s\n", newBarcode)

}

func getCheckCode(s string) string {

	// 將條碼前15碼塞到陣列中
	var arr []string
	for i := 0; i < len(s); i++ {
		ss := s[i:][0:1]
		arr = append(arr, ss)
	}

	// 逐一從36進位轉10進位, 並加總
	var sum int64
	for _, v := range arr {
		decode, _ := strconv.ParseInt(v, 36, 10)
		sum += decode
	}

	// 取 %36 取餘數, 再轉回36進位, 並轉為大寫
	decode := strconv.FormatInt(sum%36, 36)
	return strings.ToUpper(decode)
}
