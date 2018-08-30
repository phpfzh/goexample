package main

import ("fmt" 
	"net/http")

func index_handel(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"<h1>hey there</h1>")
	fmt.Fprintf(w,"<h1>输出信息了</h1>")
	fmt.Fprintf(w,"<p>输出%s信息%s了</p>","can","<strong>2323</strong")
}

func main(){
   http.HandleFunc("/",index_handel)
   http.ListenAndServe(":8000",nil)
}