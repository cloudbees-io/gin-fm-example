# Example Gin application for CloudBees Feature Management

This application demonstrates how you can integrate CloudBees Feature Management with your Gin application.
The application shows how your Gin application can respond live to changes in feature flag settings - when you update flag values in the CloudBees platform the display will update within moments.

## VPC Installations

If you are using Feature Management in a CloudBees-managed VPC environment, you must configure additional settings. There are six endpoints that must be updated in the main.go file under options := server.NewRoxOptions(server.RoxOptionsBuilder{ Your SDK configurations must point to the specific endpoints provided by your CloudBees SRE.

## Insert your SDK Key

Every application using CloudBees Feature Management needs to be configured with an SDK Key that connects it to your Flags & configurations in the [CloudBees Platform](https://cloudbees.io/).
You can retrieve your SDK Key for a particular Environment by visiting _Feature Management -> Installation_.
Then, replace the placeholder in `src/App.js` with your SDK Key:

`sdkKey := "<INSERT YOUR SDK KEY HERE>"`

For example:

`sdkKey := "8993020e-78cd-4ea0-7cfe-e1a64112eddb"`

## Run the application

Run this command:
```
go run main.go
```
...then visit the provided URL.

localhost:8080/demo

## Feature flags

The application uses the following feature flags:

* `showMessage`

A **boolean** flag that turns the message on or off.

* `message`

A **string** flag that sets the message text.

* `fontSize`

A **number** flag that sets the font size. The flag has options for 12, 16, or 24 px text size.

* `fontColor`

A **string** flag that sets the font color. The flag has options for red, green, or blue text color.


## Modifying flag values

Login to the [CloudBees platform](https://cloudbees.io/) and vist the _Feature Management_ section.
If you have configured your SDK Key correctly you should see the above flags have been created.
Change the value of one of these flags then save, ensuring the _Configuration status_ is _On_.
The application's page will automatically update shortly after to reflect the new flag value(s).

For more information on setting flag values, see the [CloudBees Feature Management documentation](https://docs.cloudbees.io/docs/cloudbees-feature-management/latest/).
