package main
import ("fmt"
	"net/http"
	"html/template"
	"io/ioutil"
	"sync"
	"encoding/xml")
 
var wg sync.WaitGroup
	type Sitemapindex struct {
	Locations []string  `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct{
   Keyword string
   Location string
}

type NewsAgePage struct {
	Title string
	Message map[string]NewsMap
}

func newsRoutine(c chan News,Location string){
	defer wg.Done()
	var n News
	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes,&n)
	resp.Body.Close()
	c <- n
}
func newsAgeHandler(w http.ResponseWriter,r *http.Request){
	var s Sitemapindex
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes,&s)
	resp.Body.Close()
	news_map := make(map[string]NewsMap)
	queue := make(chan News,30)
	for _, Location := range s.Locations{
		wg.Add(1)
		go newsRoutine(queue,Location)
        }
	wg.Wait()
	close(queue)
	for elem := range queue {
		for idx, _ := range elem.Keywords {
			news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx],elem.Locations[idx]}
		}
	}
 
	p := NewsAgePage{Title:"Html 测试Title",Message:news_map}
	t, _ := template.ParseFiles("htmlgo.html")
	fmt.Println(t.Execute(w,p))
}

func indexHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"<h1>欢迎进入首页</h1>")
}

func main(){
	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/new",newsAgeHandler)
	http.ListenAndServe(":8000",nil)
}