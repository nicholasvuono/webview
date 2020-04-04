package webview

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"reflect"
	"runtime"
)

//webview holds a url string to navigate. Optionally, you can give it a port and router to serve up.
type webview struct {
	url    string
	port   int
	router *http.ServeMux
}

//New is a constructor function for the webview struct.
func New(rawurl string) (*webview, error) {
	_, err := url.Parse(rawurl)
	if err != nil {
		return &webview{}, err
	}
	return &webview{url: rawurl}, nil
}

//SetPort is a setter function for webview's port field.
func (w *webview) SetPort(port int) error {
	if 0 > port || port > 65535 {
		return errors.New("the port specified is not in range of 0-65535")
	}
	w.port = port
	return nil
}

//GetPort is a getter function for webview's port field.
func (w *webview) GetPort() int {
	return w.port
}

//SetRouter is a setter function for webview's router field.
func (w *webview) SetRouter(router *http.ServeMux) error {
	if reflect.TypeOf(router) != reflect.TypeOf(http.NewServeMux()) {
		return errors.New("the router specified is not of type *ServeMux")
	}
	w.router = router
	return nil
}

//GetRouter is a getter function for webview's router field.
func (w *webview) GetRouter() *http.ServeMux {
	return w.router
}

//Run gets and runs a command to open a browser session.
func (w *webview) Run() error {
	command := getOpenCommand(w.url)
	err := command.Run()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//RunAndServe opens a browser in a go routine and starts up a server.
func (w *webview) RunAndServe() error {
	go w.Run()
	err := http.ListenAndServe(":"+string(w.port), w.router)
	if err != nil {
		return err
	}
	return nil
}

//getOpenCommand takes in a url and returns a command to open that url in the default browser.
func getOpenCommand(url string) *exec.Cmd {
	var command *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		command = exec.Command("powershell.exe", "Start", url)
	case "darwin":
		command = exec.Command("open", url)
	default:
		command = exec.Command("xdg-open", url)
	}
	return command
}
