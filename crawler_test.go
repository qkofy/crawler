package crawler

import (
    "fmt"
    "testing"
)

func TestIsNum(t *testing.T) {
    fmt.Printf("aaa is number: %v\n", IsNum("aaa"))

    fmt.Printf("1as is number: %v\v", IsNum("1as"))
    
    fmt.Printf("123 is number: %v\n", IsNum("123"))
}

func TestGetBody(t *testing.T) {
    fmt.Println(GetBody("http://www.baidu.com", 3))
}

func TestLoadFile(t *testing.T) {
    fmt.Println(LoadFile("geekotg", "json"))
}

func TestLoadJson(t *testing.T) {
    fmt.Println(LoadJson("geekotg", "json"))
}

func TestParseField(t *testing.T) {
    m, err := LoadJson("geekotg", "json")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("tout = ", ParseField("tout", m.(map[string]interface{})))
}

func TestGetResult(t *testing.T) {
    list := new(List)
    fmt.Println("list = ", GetResult(list))
    
    page := new(Page)
    fmt.Println("page = ", GetResult(page))
}

func TestRun(t *testing.T) {
    site := "geekotg"
    path := "games"
    pg, no := 0, 0
    res := Run(site, path, pg, no)
    fmt.Println("run result : ", res)
}