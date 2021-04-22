package crawler

import (
    "regexp"
    "sync"
)

type Page struct {
    Data []map[string]string
    Logs []string
}

func NewPage() *Page {
    return new(Page)
}

func (P *Page) ParsePage(url string, cfg *Config) {
    str, err := GetBody(url, cfg.Tout)
    if err != nil {
        P.Logs = append(P.Logs, url + str)
    } else {
        result := make(map[string]string)
        for k, v := range cfg.Page {
            regx := regexp.MustCompile(v)
            info := regx.FindStringSubmatch(str)
            if len(info) > 1 {
                result[k] = info[1]
            } else {
                result[k] = ""
            }
        }
        P.Data = append(P.Data, result)
    }
}

func (P *Page) BatchParsePage(urls []string, cfg *Config) {
    var wg sync.WaitGroup
    for _, url := range urls {
        wg.Add(1)
        go func() {
            P.ParsePage(url, cfg)
            wg.Done()
        }()
    }
    wg.Wait()
}

func (P *Page) NewResult() *Result {
    if len(P.Data) == 0 && len(P.Logs) > 0 {
        return &Result{Code: "404", Data: P.Logs, Msgs: "Not Found"}
    } else {
        return &Result{Code: "200", Data: P.Data}
    }
}