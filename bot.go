package gobot

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "regexp"
    "strconv"
    "strings"
    "time"
)

func Crawler(url string, sec int) (string, error) {
    t := time.Second * time.Duration(sec)
    client := &http.Client{Timeout: time.Duration(t)}
    resp, err := client.Get(url)
    if err != nil {
        return "request error", err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "response error", err
    }
    return string(body), err
}

func Reader(file string) (cfg interface{}, err error) {
    jsn, err := ioutil.ReadFile(file)
    if err != nil {
        return
    }
    err = json.Unmarshal(jsn, &cfg)
    return
}

func IsNum(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
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

/*
array(
    'filt' => 'http://test.haozhuodao.com/data/caiji.txt',
    'hurl' => '',
    'list' => '<a href="(http://m.geekotg.com/[a-z]+/\d+.html)">',
    'mode' => 'page',
    'murl' => 'http://m.geekotg.com/%s/page/%s.html',
    'page' => array(
        'description' => '<meta name="description" content="([^"]+)">',
		'imgurls' => '(?sU)<div class="scoll-soft-pic">\s*<div class="box">(.*)</div>',
		'introduce' => '(?s)<div class="article-con  padlf yyb_app">(.+)\s</div>\s</div>\s<div class="part-box">',
		'keywords' => '<meta name="keywords" content="([^"]+)">',
		'litpic' => '<div class="link-img"> <img src="([^"]+)" alt="[^"]*"></div>',
		'seotitle' => '<title>([^|]+) [^<]*</title>',
		'softlinks' => '<a class="and" href=\'([^\']*)\' target=\'_blank\'>[^<]*</a>',
		'title' => '<h1 class="link-name">([^<]+)</h1>',
        'typename' => '<p>类别：([^<]*)</p>',
		'versionnumber' => '<title>[^ ]* ([^ ]*) [^<]*</title>',
    ),
    'surl' => 'http://m.geekotg.com/%s',
    'tout' => '6'
)
*/
type Config struct {
    Filt string
    Hurl string
    List string
    Mode string
    Murl string
    Page map[string]string
    Surl string
    Tout int
}

func NewConfig(site string) *Config {
    cfg, err := Reader("./"+site+"/config.txt")
    if cfg == nil || err != nil {
        return new(Config)
    }
    m := cfg.(map[string]interface{})
    filt := ParseField("filt", m)
    tout, _ := strconv.Atoi(ParseField("tout", m))
    if ok, _ := regexp.MatchString("^http.+txt$", filt); ok {
        str, err := Crawler(filt, tout)
        if err != nil {
            filt = ""
        }
        filt = str
    }
    page := make(map[string]string)
	for k, v := range m["page"].(map[string]interface{}) {
		page[k] = v.(string)
	}
    return &Config{
        Filt: filt,
        Hurl: ParseField("hurl", m),
        List: ParseField("list", m),
        Mode: ParseField("mode", m),
        Murl: ParseField("murl", m),
        Page: page,
        Surl: ParseField("surl", m),
        Tout: tout,
    }
}

func (C *Config) Check(str string) bool {
    if C.List == "" {
        if C.Mode != "1" {
            if len(C.Page) > 0 {
                switch str {
                    case "surl":
                        return C.Surl != ""
                    case "murl":
                        return C.Murl != ""
                    default:
                        
                }
            } else {
                return false
            }
        }
        return false
    } else {
        switch str {
            case "surl":
                if C.Surl != "" {
                    if len(C.Page) == 0 {
                        return C.Mode == "l"
                    }
                    return true
                }
                return false
            case "murl":
                if C.Murl != "" {
                    if len(C.Page) == 0 {
                        return C.Mode == "1"
                    }
                    return true
                }
                return false
            default:
                return false
        }
    }
}

func (C *Config) Filter(url string) bool {
    return strings.Contains(C.Filt, url)
}

type Result struct {
    Code string      `json:"code"`
    Data interface{} `json:"data"`
    Msgs string      `json:"msgs"`
}

type Resulter interface {
    NewResult() *Result
}

func GetResult(r Resulter) *Result {
    return r.NewResult()
}

func (R *Result) ToJson() string {
    str, err := json.Marshal(&R)
    if err != nil {
        return `{"code":"500", "msgs":"data error"}`
    }
    return string(str)
}