package main

import (
    "log"
    "net/http"
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
)

func main(){
    m := martini.Classic()
    m.Use(render.Renderer(render.Options{
        Layout: "layout",
        IndentJSON: true,
    }))

    m.Get("/", func(r render.Render){
        r.HTML(200, "index", map[string]interface{}{"Title": "home"})
    })

    m.Get("/containers", HandleContainers)
    m.Post("/post", HandlePost)
    m.Get("/json", HandleJSON)

    log.Fatal(http.ListenAndServe(":8080", m))
}
