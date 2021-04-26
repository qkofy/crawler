package crawler

import (
    "regexp"
    "sync"
)

type List struct {
    Urls, Logs []string
}

func NewList() *List {
    return new(List)
}

func (L *List) ParseList(url string, cfg *Config) {
    str, err := GetBody(url, cfg.Tout)
    if err != nil {
        L.Logs = append(L.Logs, url + str)
    } else {
        if lext, ok := cfg.List["extent"]; ok && lext != "" {
            reg := regexp.MustCompile(lext)
            ext := reg.FindStringSubmatch(str)
            str = ext[1]
        }
        regx := regexp.MustCompile(cfg.List["target"])
        urls := regx.FindAllStringSubmatch(str, -1)
        for _, u := range urls {
            if cfg.Filter(u[1]) {
                continue
            }
            L.Urls = append(L.Urls, cfg.Hurl + u[1])
        }
    }
}

func (L *List) BatchParseList(urls []string, cfg *Config) {
    var wg sync.WaitGroup
    for _, url := range urls {
        wg.Add(1)
        go func(url string) {
            L.ParseList(url, cfg)
            wg.Done()
        }(url)
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