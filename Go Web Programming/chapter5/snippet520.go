// 测试 XSS 攻击
package main

import (
	"html/template"
	"net/http"
)

// <script>alert('Pwnd!');</script>
func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("chapter5/form.html")
	t.Execute(w, nil)
}

func processXss(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("chapter5/tmp-xss.html")
	//t.Execute(w, r.FormValue("comment"))
	// 不对 HTML 进行转义
	t.Execute(w, template.HTML(r.FormValue("comment")))
}

func main() {
	http.HandleFunc("/form", form)
	http.HandleFunc("/process-xss", processXss)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
