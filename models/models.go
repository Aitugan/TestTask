package models

import "encoding/xml"

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Chan    Channel  `xml:"channel"`
}

type Channel struct {
	XMLName xml.Name   `xml:"channel"`
	Curr    []Currency `xml:"item"`
}

type Currency struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	PubDate     string   `xml:"pubDate"`
	Description string   `xml:"description"`
	Quant       string   `xml:"quant"`
	Index       string   `xml:"index"`
	Change      string   `xml:"change"`
}

type Conversion struct {
	ConvertFrom string  `json:"convertFrom"`
	Amount      float64 `json:"amount"`
	ConvertTo   string  `json:"convertTo"`
}

type GetRHistoryResult struct {
	Date  string
	Value string
}

type History struct {
	Title       string
	PubDate     string
	Description float64
	Change      float64
}
