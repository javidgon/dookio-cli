package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestIfAppIsParsed(t *testing.T) {
	app := new(App)
	request := &Request{Params: url.Values{}}
	rawApp := "git/apache"
	parseApp(rawApp, app, request)
	if app.User == "" || app.Repo == "" {
		t.Error("The app was not parsed properly")
	}
	if request.Params.Get("user") != "git" || request.Params.Get("repo") != "apache" {
		t.Error("The request was not filled in properly")
	}
}

func TestIfSimpleCommandIsParsed(t *testing.T) {
	command := new(Command)
	request := &Request{Params: url.Values{}}
	rawCommand := "start"
	parseCommand(rawCommand, command, request)
	if command.Name != "containers" || command.Magnitude != "" {
		t.Error("The command was not parsed properly")
	}
	if request.Params.Get("action") != rawCommand {
		t.Error("The request was not filled in properly")
	}
}

func TestIfComplexCommandIsParsed(t *testing.T) {
	command := new(Command)
	request := &Request{Params: url.Values{}}
	rawCommand := "scale=5"
	parseCommand(rawCommand, command, request)
	if command.Name != "scale" || command.Magnitude != "5" {
		t.Error("The command was not parsed properly")
	}

	if request.Params.Get("multiplicator") != "5" {
		t.Error("The request was not filled in properly")
	}

	if request.Path != "scale" {
		t.Error("The path was not filled in properly")
	}
}

func TestIfTheServerIsContactedUsingTheRightUri_ScaleApp(t *testing.T) {
	command := new(Command)
	app := new(App)
	server := new(Server)
	request := &Request{Params: url.Values{}}
	rawCommand := "scale=5"
	rawApp := "git/apache"
	parseCommand(rawCommand, command, request)
	parseApp(rawApp, app, request)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.RequestURI()))
	}))
	defer ts.Close()

	server.Address = ts.URL[7 : len(ts.URL)-6]
	server.Port = ts.URL[len(ts.URL)-5:]
	content, _ := contactWithDookioServer(command, app, server, request)
	if string(content) != "/scale?multiplicator=5&repo=apache&user=git" {
		t.Error("The response does not match")
	}
}

func TestIfTheServerIsContactedUsingTheRightUri_GetApps(t *testing.T) {
	command := new(Command)
	server := new(Server)
	app := new(App)
	request := &Request{Params: url.Values{}}
	rawCommand := "apps"
	parseCommand(rawCommand, command, request)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.RequestURI()))
	}))
	defer ts.Close()

	server.Address = ts.URL[7 : len(ts.URL)-6]
	server.Port = ts.URL[len(ts.URL)-5:]
	content, _ := contactWithDookioServer(command, app, server, request)
	if string(content) != "/apps" {
		t.Error("The response does not match")
	}
}

func TestIfTheServerIsContactedUsingTheRightUri_StartApp(t *testing.T) {
	command := new(Command)
	server := new(Server)
	app := new(App)
	request := &Request{Params: url.Values{}}
	rawCommand := "start"
	rawApp := "git/apache"
	parseCommand(rawCommand, command, request)
	parseApp(rawApp, app, request)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.RequestURI()))
	}))
	defer ts.Close()

	server.Address = ts.URL[7 : len(ts.URL)-6]
	server.Port = ts.URL[len(ts.URL)-5:]
	content, _ := contactWithDookioServer(command, app, server, request)
	if string(content) != "/containers?action=start&repo=apache&user=git" {
		t.Error("The response does not match")
	}
}

func TestIfTheServerIsContactedUsingTheRightUri_StopApp(t *testing.T) {
	command := new(Command)
	server := new(Server)
	app := new(App)
	request := &Request{Params: url.Values{}}
	rawCommand := "stop"
	rawApp := "git/apache"
	parseCommand(rawCommand, command, request)
	parseApp(rawApp, app, request)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.RequestURI()))
	}))
	defer ts.Close()

	server.Address = ts.URL[7 : len(ts.URL)-6]
	server.Port = ts.URL[len(ts.URL)-5:]
	content, _ := contactWithDookioServer(command, app, server, request)
	if string(content) != "/containers?action=stop&repo=apache&user=git" {
		t.Error("The response does not match")
	}
}

func TestIfTheServerIsContactedUsingTheRightUri_GetContainers(t *testing.T) {
	command := new(Command)
	server := new(Server)
	app := new(App)
	request := &Request{Params: url.Values{}}
	rawCommand := "containers"
	rawApp := "git/apache"
	parseCommand(rawCommand, command, request)
	parseApp(rawApp, app, request)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.RequestURI()))
	}))
	defer ts.Close()

	server.Address = ts.URL[7 : len(ts.URL)-6]
	server.Port = ts.URL[len(ts.URL)-5:]
	content, _ := contactWithDookioServer(command, app, server, request)
	if string(content) != "/containers?repo=apache&user=git" {
		t.Error("The response does not match")
	}
}
