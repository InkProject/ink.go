package ink

import (
    "fmt"
    "time"
    "crypto/rand"
)

var tokenMap map[string]typeToken

type typeToken struct {
    time time.Time
    data map[string]interface{}
}

func GUID() string {
    b := make([]byte, 16)
    rand.Read(b)
    return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func (ctx *Context) TokenGet(key string) interface{} {
    tokenId := ctx.Req.Header.Get("Token")
    if token, ok := tokenMap[tokenId]; ok {
        if value, ok := token.data[key]; ok {
            return value
        }
    }
    return nil
}

func (ctx *Context) TokenSet(key string, value interface{}) {
    tokenId := ctx.Req.Header.Get("Token")
    token, ok := tokenMap[tokenId]
    if ok {
        token.data[key] = value
    }
}

func (ctx *Context) TokenNew() string {
    if len(tokenMap) == 0 {
        tokenMap = make(map[string]typeToken)
    }
    tokenId := GUID()
    token := new(typeToken)
    token.data = make(map[string]interface{})
    tokenMap[tokenId] = *token
    ctx.Req.Header.Set("Token", tokenId)
    return tokenId
}
