package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rollout/rox-go/v5/core/context"
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

func main() {
	router := gin.Default()
	initFlags()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, You Created a Web App!"})
	})

	router.GET("/demo", func(c *gin.Context) {
		ctx := context.NewContext(map[string]interface{}{"email": "[email protected]"})
		msg := ""
		if flags.ShowMessage.IsEnabled(ctx) {
			msg = flags.Message.GetValue(ctx)
		} else {
			msg = "Flag Message hidden. Enable the flag in Cloudbees Platform to see it."
		}
		c.JSON(http.StatusOK, gin.H{"message": msg, "fontColor": flags.FontColor.GetValue(ctx), "fontSize": flags.FontSize.GetValue(ctx)})
	})

	router.Run(":8080")
}

func initFlags() {
	sdkKey := "<INSERT YOUR SDK KEY HERE>"
	options := server.NewRoxOptions(server.RoxOptionsBuilder{DisableSignatureVerification: true})
	rox := server.NewRox()

	// Register the flags container with the CloudBees platform
	rox.RegisterWithEmptyNamespace(flags)

	// Setup the feature management environment key
	<-rox.Setup(sdkKey, options)
}
