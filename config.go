package crawler

import (
    "regexp"
    "strconv"
    "strings"
)

type Config struct {
    Filt, Hurl, Mode, Murl, Surl string
    List, Page map[string]string
    Tout int
}

func NewConfig(cname string) *Config {
    ctype := JsonFile
    if cname[0] == '{' {
        ctype = JsonString
    }
    cfg, err := LoadJson(cname, ctype)
    if cfg == nil || err != nil {
        return new(Config)
    }
    m := cfg.(map[string]interface{})
    filt := ParseField("filt", m)
    tout, _ := strconv.Atoi(ParseField("tout", m))
    if ok, _ := regexp.MatchString("^http.+txt$", filt); ok {
        str, err := GetBody(filt, tout)
        filt = str
        if err != nil {
            filt = ""
        }
    } else if filt == "local" {
        jsn, _ := LoadFile(cname, TxtFile)
        filt = jsn.(string)
    }
    list := make(map[string]string)
	for k, v := range m["list"].(map[string]interface{}) {
		list[k] = v.(string)
	}
    page := make(map[string]string)
	for k, v := range m["page"].(map[string]interface{}) {
		page[k] = v.(string)
	}
    return &Config{
        Filt: filt,
        Hurl: ParseField("hurl", m),
        List: list,
        Mode: ParseField("mode", m),
        Murl: ParseField("murl", m),
        Page: page,
        Surl: ParseField("surl", m),
        Tout: tout,
    }
}

func (cfg *Config) Check(str string) bool {
    if len(cfg.List) == 0 {
        if cfg.Mode != "1" {
            if len(cfg.Page) > 0 {
                switch str {
                    case "surl":
                        return cfg.Surl != ""
                    case "murl":
                        return cfg.Murl != ""
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
                if cfg.Surl != "" {
                    if len(cfg.Page) == 0 {
                        return cfg.Mode == "l"
                    }
                    return true
                }
                return false
            case "murl":
                if cfg.Murl != "" {
                    if len(cfg.Page) == 0 {
                        return cfg.Mode == "1"
                    }
                    return true
                }
                return false
            default:
                return false
        }
    }
}

func (cfg *Config) Filter(url string) bool {
    return strings.Contains(cfg.Filt, url)
}