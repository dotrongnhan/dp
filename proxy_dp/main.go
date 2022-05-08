package main

import (
	"fmt"
	"proxy/proxy"
)

func main() {
	nginxServer := proxy.NewNginxServer()
	appStatusURL := "app/status"
	createUserURL := "create/user"
	httpCode, message := nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nMessage: %s\n", appStatusURL, httpCode, message)

	httpCode, message = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nMessage: %s\n", appStatusURL, httpCode, message)

	httpCode, message = nginxServer.HandleRequest(appStatusURL, "GET")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nMessage: %s\n", appStatusURL, httpCode, message)

	httpCode, message = nginxServer.HandleRequest(createUserURL, "POST")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nMessage: %s\n", appStatusURL, httpCode, message)

	httpCode, message = nginxServer.HandleRequest(createUserURL, "POST")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nMessage: %s\n", appStatusURL, httpCode, message)

	httpCode, message = nginxServer.HandleRequest(createUserURL, "POST")
	fmt.Printf("\nURL: %s\nHttpCode: %d\nMessage: %s\n", appStatusURL, httpCode, message)
}
