package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type App struct {
	User string
	Repo string
}

type Request struct {
	Path   string
	Params url.Values
}

type Command struct {
	Name      string
	Magnitude string
}

type Server struct {
	Address string
	Port    string
}

func showHelp() {
	fmt.Println("Available commands:")
	fmt.Println("--> apps: List all the deployed apps (e.g dookio apps)")
	fmt.Println("--> containers: List all the containers associated to each app (e.g dookio containers git/apache)")
	fmt.Println("--> start: Start a stopped application (e.g dookio start git/apache)")
	fmt.Println("--> stop: Stop a running application (e.g dookio stop git/apache)")
	fmt.Println("--> scale: Scale a running application (e.g dookio scale=5 git/apache)")
}

func parseApp(rawApp string, app *App, request *Request) {
	//Fill in the App struct.
	info := strings.Split(rawApp, "/")
	app.User = info[0]
	app.Repo = info[1]
	request.Params.Add("user", app.User)
	request.Params.Add("repo", app.Repo)
}

func parseCommand(rawCommand string, command *Command, request *Request) {
	//Fill in the Command struct.
	if strings.Contains(rawCommand, "=") {
		info := []string{}
		info = strings.Split(rawCommand, "=")
		command.Name = info[0]
		command.Magnitude = info[1]
		request.Params.Add("multiplicator", command.Magnitude)
	} else {
		if rawCommand != "apps" && rawCommand != "containers" {
			command.Name = "containers"
			request.Params.Add("action", rawCommand)
		} else {
			command.Name = rawCommand
		}
	}
	request.Path = command.Name
}

func contactWithDookioServer(command *Command, app *App, server *Server, request *Request) ([]byte, error) {
	req := url.URL{Scheme: "http",
		Host:     server.Address + ":" + server.Port,
		Path:     request.Path,
		RawQuery: request.Params.Encode()}
	response, err := http.Get(req.String())
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func main() {
	server := &Server{Address: os.Getenv("DOOKIO_SERVER_ADDRESS"), Port: "8000"}
	if server.Address == "" {
		fmt.Println("Please, set first the DOOKIO_SERVER_ADDRESS env var")
	} else if len(os.Args) < 2 || os.Args[1] == "help" {
		showHelp()
	} else {
		command := new(Command)
		app := new(App)
		request := &Request{Params: url.Values{}}
		rawCommand := os.Args[1]
		parseCommand(rawCommand, command, request)
		fmt.Printf("Dookio cli: Running the '%s' command.\n", command.Name)
		if len(os.Args) > 2 {
			rawApp := os.Args[2]
			parseApp(rawApp, app, request)
			fmt.Printf("Dookio cli: For the '%s/%s' app.\n", app.User, app.Repo)
		}
		fmt.Printf("Dookio cli: Connecting with '%s:%s...'\n", server.Address, server.Port)
		content, err := contactWithDookioServer(command, app, server, request)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Printf("Request: %s:%s\n", request.Path, request.Params.Encode())
		} else {
			fmt.Println(string(content))
			fmt.Printf("Request: %s/%s\n", request.Path, request.Params.Encode())
		}
		fmt.Println("Done.")
	}
}
