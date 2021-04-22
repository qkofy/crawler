package crawler

import "encoding/json"

type Result struct {
    Code string      `json:"code"`
    Data interface{} `json:"data"`
    Msgs string      `json:"msgs"`
}

func (R *Result) ToJson() string {
    str, err := json.Marshal(&R)
    if err != nil {
        return `{"code":"500", "msgs":"data error"}`
    }
    return string(str)
}