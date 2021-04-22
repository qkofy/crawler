package crawler

import (
    "fmt"
    "testing"
)

func TestList(t *testing.T) {
    site := "geekotg"
    cfg := NewConfig(site)
    
    list := NewList()
    
    url := "http://m.geekotg.com/games"
    list.ParseList(url, cfg)
    fmt.Println("list = ", list)
    
    urls := []string{"http://m.geekotg.com/games/page/2","http://m.geekotg.com/page/3"}
    list.BatchParseList(urls, cfg)
    fmt.Println("list = ", list)
    
    fmt.Println("list.result = ", list.NewResult())
}