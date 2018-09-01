package main

import("fmt"
	"io/ioutil"
	"net/http"
	"encoding/xml")

type Sitemapindex struct {
	Locations []string  `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func main(){
	var s Sitemapindex
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes,&s)
	for _,Location := range s.Locations{
		resp,_ := http.Get(Location)
		bytes, _ :=ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes,&n)
 	} 
}
