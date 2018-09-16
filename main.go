package main

import (
	"github.com/astaxie/beego"
	//"github.com/sylvesterchiang/cohesion"
	"github.com/tkanos/gonfig"
	//"github.com/zmb3/spotify"
)

type MainController struct {
	beego.Controller
}

func getConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/config.dev.json", &configuration)
	return configuration
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
	configuration := getConfig()
	this.Ctx.WriteString(configuration.ClientID)
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
