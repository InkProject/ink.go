package ink

import (
    "os"
    "fmt"
    "net/http"
    "path/filepath"
)

func Static(root string) func(ctx *Context) {
    return func(ctx *Context) {
        if ctx.Req.URL.Path == "" || ctx.Req.URL.Path == "/" {
            http.ServeFile(ctx.Res, ctx.Req, filepath.Join(root, "index.html"))
        } else {
            fileName := root + ctx.Req.URL.Path
            f, err := os.Stat(fileName)
            if err == nil {
                if f.IsDir() {
                    http.ServeFile(ctx.Res, ctx.Req, filepath.Join(fileName, "index.html"))
                } else {
                    http.ServeFile(ctx.Res, ctx.Req, fileName)
                }
                ctx.Stop()
            } else {
                ctx.Write([]byte("404 Not found"))
                fmt.Println("Not found: " + fileName)
            }
        }
    }
}
