package main

import (
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
		t.Error("The request was not filled in")
	}

	if request.Path != "scale" {
		t.Error("The path was not filled in properly")
	}
}
