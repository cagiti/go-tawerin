package main

import (
    "github.com/magiconair/properties"
    "github.com/gorilla/mux"
    "html/template"
    "net/http"
    "strings"
    "fmt"
    "os"
)

func getProperties(host string) map[string]string {

    var p map[string]string

    if strings.Contains(host, "wales") {
        p = properties.MustLoadFile("static/wales.properties", properties.UTF8).Map()
    } else {
        p = properties.MustLoadFile("static/cymru.properties", properties.UTF8).Map()
    }

    return p
}

func handler(w http.ResponseWriter, r *http.Request) {
    m := getProperties(r.Host)
    if r.URL.Path == "/" {
        t, _ := template.ParseFiles("templates/index.tmpl")
        t.Execute(w, m)
    } else if r.URL.Path == "/ping" {
        fmt.Fprintf(w, "OK")
    } else {
        t, _ := template.ParseFiles(fmt.Sprintf("templates/%s.tmpl", r.URL.Path))
        t.Execute(w, m)
    }
}

func main() {

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r := mux.NewRouter()
    r.HandleFunc("/", handler)
    r.HandleFunc("/ytim", handler)
    r.HandleFunc("/band", handler)
    r.HandleFunc("/dawnsfeydd", handler)
    r.HandleFunc("/rolau", handler)
    r.HandleFunc("/oriel", handler)
    r.HandleFunc("/perfformiadau", handler)
    r.HandleFunc("/cysylltu", handler)
    r.HandleFunc("/error", handler)
    r.HandleFunc("/result", handler)
    r.HandleFunc("/ping", handler)
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    http.ListenAndServe(fmt.Sprintf(":%s",port), r)
}
