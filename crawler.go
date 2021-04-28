package crawler

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
    "strconv"
    "time"
)

const (
    TxtFile    = "txt"
    JsonFile   = "json"
    JsonString = "string"
)

func IsNum(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

func GetBody(url string, sec int) (string, error) {
    t := time.Second * time.Duration(sec)
    client := &http.Client{Timeout: t}
    resp, err := client.Get(url)
    if err != nil {
        return "request error", err
    }
    defer func() {
        _ = resp.Body.Close()
    }()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "response error", err
    }
    return string(body), err
}

func LoadFile(cname, ctype string) (cfg interface{}, err error) {
    jsn, err := ioutil.ReadFile("./config/"+cname+"."+ctype)
    if err != nil {
        return
    }
    switch ctype {
        case JsonFile:
            err = json.Unmarshal(jsn, &cfg)
        case TxtFile:
            cfg = string(jsn)
        default:
            err = errors.New("load error")
    }
    return
}

func LoadJson(cname, ctype string) (cfg interface{}, err error) {
    if ctype == JsonString {
        err = json.Unmarshal([]byte(cname), &cfg)
    } else {
        cfg, err = LoadFile(cname, ctype)
    }
    return
}

func ParseField(k string, m map[string]interface{}) string {
    if v, ok := m[k]; ok {
        if k == "tout" && !IsNum(v.(string)) {
            return "3"
        }
        return v.(string)
    } else {
        if k == "tout" {
            return "3"
        }
        return ""
    }
}

type Resulter interface {
    NewResult() *Result
}

func GetResult(r Resulter) *Result {
    return r.NewResult()
}

func Run(site string, path string, pg int, no int) string {
    cfg := NewConfig(site)
    list := NewList()
    if pg == 0 {
        if !cfg.Check("surl") {
            return `{"code":"500", "msgs":"surl config error"}`
        }
        url := fmt.Sprintf(cfg.Surl, path)
        if cfg.Mode == "2" {
            list.Urls = append(list.Urls, url)
        } else {
            list.ParseList(url, cfg)
        }
    } else {
        if !cfg.Check("murl") {
            return `{"code":"500", "msgs":"murl config error"}`
        }
        var urls []string
        for i := pg; i <= no; i++ {
            urls = append(urls, fmt.Sprintf(cfg.Murl, path, i))
        }
        if cfg.Mode == "2" {
            list.Urls = append(list.Urls, urls...)
        } else {
            list.BatchParseList(urls, cfg)
        }
    }
    if cfg.Mode == "1" {
        return GetResult(list).ToJson()
    } else {
        page := NewPage()
        if cfg.Mode == "2" {
            page.ParsePage(list.Urls[0], cfg)
        } else {
            page.BatchParsePage(list.Urls, cfg)
        }
        return GetResult(page).ToJson()
    }
}