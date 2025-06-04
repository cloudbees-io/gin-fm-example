package main

import (
	"fmt"
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

	notFlags := NewNotFlags("notflags message", "Red", 16, true)
	fmt.Printf("notflags message: %s, fontColor: %s, fontSize: %d, showMessage: %t\n",
		notFlags.GetMessage(), notFlags.GetFontColor(), notFlags.GetFontSize(), notFlags.IsShowMessage())

	router := gin.Default()
	router.GET("/", homePage)
	router.GET("/demo", demo)
	router.GET("/wrapper", wrapperDemo)

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

func wrapperDemo(c *gin.Context) {
	msg := ""
	showMessageValue, err := getFlagValue("ShowMessage")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve flag value: " + err.Error()})
		return
	}

	if showMessageValue == "true" {
		messageValue, err := getFlagValue("Message")
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve flag value: " + err.Error()})
			return
		}
		msg = messageValue
	} else {
		msg = "Flag message hidden. Enable the flag in the Cloudbees platform to display it."
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": msg, "fontColor": flags.FontColor.GetValue(nil), "fontSize": flags.FontSize.GetValue(nil)})
}
