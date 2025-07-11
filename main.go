package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rollout/rox-go/v5/server"
)

type Flags struct {
	ShowMessage server.RoxFlag
	Message     server.RoxString
	FontColor   server.RoxString
	FontSize    server.RoxInt
}

var flags = &Flags{
	ShowMessage: server.NewRoxFlag(false),
	Message:     server.NewRoxString("This is the default message; try changing some flag values!", []string{}),
	FontColor:   server.NewRoxString("Black", []string{"Red", "Green", "Blue", "Black"}),
	FontSize:    server.NewRoxInt(12, []int{12, 16, 24}),
}

func initFlags() {
	sdkKey := "<YOUR-SDK-KEY>"
	options := server.NewRoxOptions(server.RoxOptionsBuilder{DisableSignatureVerification: true})
	rox := server.NewRox()

	// Register the flags container with the CloudBees platform
	rox.RegisterWithEmptyNamespace(flags)

	// Setup the feature management environment key
	<-rox.Setup(sdkKey, options)
}

func main() {
	initFlags()

	router := gin.Default()
	router.GET("/", homePage)
	router.GET("/demo", demo)

	router.Run(":8080")
}

func homePage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello, You created a web app!"})
}

func demo(c *gin.Context) {
	msg := ""
	if flags.ShowMessage.IsEnabled(nil) {
		msg = flags.Message.GetValue(nil)
	} else {
		msg = "Flag message hidden. Enable the flag in the Cloudbees platform to display it."
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": msg, "fontColor": flags.FontColor.GetValue(nil), "fontSize": flags.FontSize.GetValue(nil)})
}
