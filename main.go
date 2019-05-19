package main

import (
	"context"
	"log"

	"github.com/astaxie/beego"
	//"github.com/sylvesterchiang/cohesion"
	"github.com/tkanos/gonfig"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

const redirectURI = "http://localhohst:8080/callback"

var (
	client spotify.Client
)

type MainController struct {
	beego.Controller
}

type PlaylistController struct {
	beego.Controller
}

type TestController struct {
	beego.Controller
}

type RespJson struct {
	Value1 string `json:"value1"`
	Value2 string `json:"value2"`
}

type TrackJson struct {
	Track         []string
	AudioFeatures []*spotify.AudioFeatures
}

func getConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/config.dev.json", &configuration)
	return configuration
}

func createClientToken() {
	configuration := getConfig()
	//this.Ctx.WriteString(configuration.ClientID)
	config := &clientcredentials.Config{
		ClientID:     configuration.ClientID,
		ClientSecret: configuration.ClientSecret,
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client = spotify.Authenticator{}.NewClient(token)
	log.Print("new client token")
	log.Print(client)
}

func (this *MainController) Get() {
	//this.Ctx.WriteString("hello world")
	configuration := getConfig()
	this.TplName = "playlist.tpl"
	this.Data["clientid"] = configuration.ClientID

	createClientToken()

	log.Print("missing client token")
	log.Print(client)

	//plStr := "https://open.spotify.com/user/helenprejean/playlist/1knWcCeNV7zX1YLaDpLzzW?si=3lEQ6xd_RJG71n7kLqF-Aw"
	plStr := "1knWcCeNV7zX1YLaDpLzzW"
	//plTrackPage is type PlayListTrackPage, error
	id := spotify.ID(plStr)
	log.Print(id)
	plTrackPage, err := client.GetPlaylistTracks(id)
	if err != nil {
		log.Print(err)
	}
	log.Print("WHATEVER FUCK OFF")
	log.Print(plTrackPage.Tracks[0].Track)
}

func (this *TestController) Get() {
	responseJSON := RespJson{
		Value1: "get out",
		Value2: "of here",
	}
	this.Data["json"] = &responseJSON
	this.ServeJSON()
}

func (this *PlaylistController) Get() {
	this.Ctx.WriteString("/playlist get call")
}

func (this *PlaylistController) Post() {

	idString := this.GetString("id")
	if idString == "" {
		log.Print("undefined id string")
		this.Ctx.WriteString("EMPTY STRING FUCKER")
	}

	createClientToken()

	plTrackPage, err := client.GetPlaylistTracks(spotify.ID(idString))
	if err != nil {
		log.Print(err)
	}

	//check to see how long the playlist is
	if len(plTrackPage.Tracks) > 10 {
		this.Ctx.WriteString("Your playlist is too long, it has be to less than 10 songs.")
	}
	var tempArray [10]string
	//var idArray [10]spotify.ID
	idArray := make([]spotify.ID, 0, 10) //length = 0, capactiy = 30

	for i := 0; i < 10; i++ {
		log.Print(plTrackPage.Tracks[i])
		tempArray[i] = plTrackPage.Tracks[i].Track.String()
		idArray = append(idArray, plTrackPage.Tracks[i].Track.ID)
	}

	features, err := client.GetAudioFeatures(idArray...) //turns array into variadic input

	if err != nil {
		log.Print(err)
	}

	responseJSON := TrackJson{
		Track:         tempArray[:],
		AudioFeatures: features,
	}
	this.Data["json"] = &responseJSON
	this.ServeJSON()
}

func main() {
	beego.Router("/", &MainController{})
	beego.Router("/playlist/", &PlaylistController{})
	beego.Run()
}
