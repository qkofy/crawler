package gobot

import (
    "regexp"
    "sync"
)

type List struct {
    Urls []string
    Logs []string
}

func NewList() *List {
    return new(List)
}

func (L *List) ParseList(url string, C *Config) {
    str, err := Crawler(url, C.Tout)
    if err != nil {
        L.Logs = append(L.Logs, url + str)
    } else {
        regx := regexp.MustCompile(C.List)
        urls := regx.FindAllStringSubmatch(str, -1)
        for _, u := range urls {
            if C.Filter(u[1]) {
                continue
            }
            L.Urls = append(L.Urls, C.Hurl + u[1])
        }
    }
}

func (L *List) BatchParseList(urls []string, cfg *Config) {
    var wg sync.WaitGroup
    for _, url := range urls {
        wg.Add(1)
        go func(url string, wg *sync.WaitGroup) {
            L.ParseList(url, cfg)
            wg.Done()
        }(url, &wg)
    }
    wg.Wait()
}

func (L *List) NewResult() *Result {
    if len(L.Urls) == 0 && len(L.Logs) > 0 {
        return &Result{Code: "404", Data: L.Logs, Msgs: "Not Found"}
    } else {
        return &Result{Code: "200", Data: L.Urls}
    }
}