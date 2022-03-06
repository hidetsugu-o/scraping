package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

func main() {
	url := "https://qiita.com/"

	// Getリクエスト
	res, _ := http.Get(url)
	defer res.Body.Close()

	// 読み取り
	buf, _ := ioutil.ReadAll(res.Body)

	// 文字コード判定
	det := chardet.NewTextDetector()
	detRslt, _ := det.DetectBest(buf)
	fmt.Println(detRslt.Charset)

	// 文字コード変換
	bReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

	// HTMLパース
	doc, _ := goquery.NewDocumentFromReader(reader)

	// titleを抜き出し
	rslt := doc.Find(`script[data-component-name="HomeIndexPage"]`).Text()
	// fmt.Println(rslt)

	// mapに格納
	var rsltMap map[string]map[string]interface{}
	b := []byte(rslt)
	json.Unmarshal(b, &rsltMap)

	for _, v := range rsltMap["trend"] {
		fmt.Println(v)
	}
}
