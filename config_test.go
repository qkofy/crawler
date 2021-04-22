package crawler

import (
    "fmt"
    "testing"
)

func TestConfig(t *testing.T) {
    site := "geekotg"
    cfg := NewConfig(site)
    fmt.Println("config = ", cfg)
    
    fmt.Printf("cfg.Surl is %v\n", cfg.Check("surl"))
    
    url := "http://www.geekotg.com"
    fmt.Printf("url: %s is filter %v\n", url, cfg.Filter(url))
}

