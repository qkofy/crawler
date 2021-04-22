cd dpackage main

import (
    "fmt"
    "github.com/qkofy/crawler"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "strconv"
)

const Token = "22d33fea6cca5d10508d9e9b07c45b22"

func action(token string, site string, path string, pg int, no int) (r string) {
    if token == Token {
        return crawler.Run(site, path, pg, no)
    } else {
        return `{"code":"403", "msgs":"Forbidden"}`
    }
}

func spage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    res := action(p.ByName("token"), p.ByName("site"), p.ByName("path"), 0, 0)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    fmt.Fprintf(w, res)
}

func mpage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    pg, _ := strconv.Atoi(p.ByName("page"))
    no, _ := strconv.Atoi(p.ByName("num"))
    res := action(p.ByName("token"), p.ByName("site"), p.ByName("path"), pg, no)
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    fmt.Fprintf(w, res)
}

func main() {
    router := httprouter.New()
    router.GET("/:token/:site/:path", spage)
    router.GET("/:token/:site/:path/:page/:num", mpage)
    log.Fatal(http.ListenAndServe(":10080", router))
}