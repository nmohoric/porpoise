package main

import (
	"github.com/fsouza/go-dockerclient"
    "github.com/gin-gonic/gin"
)

func HandleIndex(c *gin.Context){
    obj := gin.H{"title": "Main website"}
    c.HTML(200, "layout.html", obj)
}

func HandleContainers(c *gin.Context){
	endpoint        := "unix:///var/run/docker.sock"
	client, _       := docker.NewClient(endpoint)

	containers, _   := client.ListContainers(docker.ListContainersOptions{false,false,10,"",""})

    c.JSON(200, containers)
}

func HandlePost(c *gin.Context){
}

type Result struct {
    FirstName string `json:"first"`
    LastName  string `json:"last"`
}

func HandleJSON(c *gin.Context){
	c.JSON(200, gin.H{"first":"tee", "last":"dub"})
}
