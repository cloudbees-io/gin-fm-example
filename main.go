package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rollout/rox-go/v5/core/client"
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
	sdkKey := "<INSERT YOUR SDK KEY HERE>"
	options := server.NewRoxOptions(server.RoxOptionsBuilder{
		DisableSignatureVerification: true,
		NetworkConfigurationsOptions: client.NewNetworkConfigurationsOptions(client.NetworkConfigurationsBuilder{
                        GetConfigApiEndpoint: "https://api.vpc-install-test.saas-tools.beescloud.com/device/get_configuration",
                        GetConfigCloudEndpoint: "https://rox-conf.vpc-install-test.saas-tools.beescloud.com",
                        SendStateApiEndpoint: "https://api.vpc-install-test.saas-tools.beescloud.com/device/update_state_store/",
                        SendStateCloudEndpoint: "https://rox-state.vpc-install-test.saas-tools.beescloud.com",
                        AnalyticsEndpoint: "https://fm-analytics.vpc-install-test.saas-tools.beescloud.com",
                        PushNotificationEndpoint: "https://sdk-notification-service.vpc-install-test.saas-tools.beescloud.com/sse",
                })})
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
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello, You Created a Web App!"})
}

func demo(c *gin.Context) {
	msg := ""
	if flags.ShowMessage.IsEnabled(nil) {
		msg = flags.Message.GetValue(nil)
	} else {
		msg = "Flag Message hidden. Enable the flag in Cloudbees Platform to see it."
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": msg, "fontColor": flags.FontColor.GetValue(nil), "fontSize": flags.FontSize.GetValue(nil)})
}
