package ink

func Cors(ctx *Context) {
    ctx.Header().Set("Access-Control-Allow-Origin", "*")
    ctx.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    if ctx.Req.Method == "OPTIONS" {
        ctx.Write([]byte{})
        ctx.Stop()
    }
}
