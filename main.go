package main

import (
    "github.com/gorilla/mux"
    "github.com/newrelic/go-agent"
    "html/template"
    "net/http"
    "fmt"
    "log"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/index.tmpl")
    t.Execute(w, map[string] string {"title": "simple-go-webapp"})
}

func main() {
    newRelicLicenseKey := os.Getenv("NEWRELIC_LICENSE_KEY")
    newRelicApplicationName := os.Getenv("NEWRELIC_APP_NAME")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    config := newrelic.NewConfig(newRelicApplicationName, newRelicLicenseKey)
    app, err := newrelic.NewApplication(config)

    if err != nil {
        log.Fatal("Unable to create new relic application: ", err)
        os.Exit(1)
    }

    r := mux.NewRouter()
    r.HandleFunc(newrelic.WrapHandleFunc(app,"/", handler))
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    http.ListenAndServe(fmt.Sprintf(":%s",port), r)

}
