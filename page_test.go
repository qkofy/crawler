package crawler

import (
    "fmt"
    "testing"
)

func TestPage(t *testing.T) {
    site := "geekotg"
    cfg := NewConfig(site)
    
    page := NewPage()
    
    url := "http://m.geekotg.com/games/47177.html"
    page.ParsePage(url, cfg)
    fmt.Println("page = ", page)
    
    urls := []string{"http://m.geekotg.com/games/47212.html","http://m.geekotg.com/games/47206.html"}
    page.BatchParsePage(urls, cfg)
    fmt.Println("page = ", page)
    
    fmt.Println("page.result = ", page.NewResult())
}