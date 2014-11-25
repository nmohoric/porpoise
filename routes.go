package main

import (
	"github.com/fsouza/go-dockerclient"
    "github.com/martini-contrib/render"
)

func HandleContainers(r render.Render){
	endpoint        := "unix:///var/run/docker.sock"
	client, _       := docker.NewClient(endpoint)

	containers, _   := client.ListContainers(docker.ListContainersOptions{false,false,10,"",""})

    r.JSON(200, containers)
}

func HandlePost(r render.Render){
}

type Result struct {
    FirstName string `json:"first"`
    LastName  string `json:"last"`
}

func HandleJSON(r render.Render){
//	r.JSON(200, gin.H{"first":"tee", "last":"dub"})
}
