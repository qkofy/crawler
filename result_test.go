package crawler

import (
    "fmt"
    "testing"
)

func TestResult(t *testing.T) {
    res := new(Result)
    fmt.Println("result = ", res.ToJson())
}