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
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to send http request", err)
		return
	}
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

	// 構造体に格納
	var rsltMap QiitaServices
	json.Unmarshal([]byte(rslt), &rsltMap)

	for _, edge := range rsltMap.Trend.Trend.Edges {
		fmt.Println(edge.Node.Title)
		fmt.Println(edge.Node.LinkURL)
		fmt.Println()
	}

}
