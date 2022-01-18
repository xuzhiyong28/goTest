package main

import (
	"example/project/study/go_session"
	_ "example/project/study/go_session/memory"
	"html/template"
	"net/http"
)

var globalSessions, _ = go_session.NewManager("memory", "gosessionid", 3600)

func main(){
	//tcp_demo.PortOpenTest()

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		sess := globalSessions.SessionStart(w, r)
		r.ParseForm()
		if r.Method == "GET" {
			t, _ := template.ParseFiles("go_session/login.gtpl")
			w.Header().Set("Content-Type", "text/html")
			t.Execute(w, sess.Get("username"))
		} else {
			sess.Set("username", r.Form["username"])
			http.Redirect(w, r, "/", 302)
		}
	})
	http.ListenAndServe("127.0.0.0:8080", nil)

}

