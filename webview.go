package webview

import (
	"errors"
	"net/http"
	"net/url"
	"os/exec"
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
	err := errors.New("the port specified is not in range of 0-65535")
	if 0 <= port && port <= 65535 {
		return err
	}
	w.port = port
	return nil
}

//GetPort is a getter function for webview's port field.
func (w *webview) GetPort() int {
	return w.port
}

//SetRouter is a setter function for webview's router field.
func (w *webview) SetRouter(router *http.ServeMux) {
	w.router = router
}

//GetRouter is a getter function for webview's router field.
func (w *webview) GetRouter() *http.ServeMux {
	return w.router
}

//Run gets and runs a command to open a browser session.
func (w *webview) Run() {
	command := getOpenCommand(w.url)
	command.Run()
}

//RunAndServe opens a browser in a go routine and starts up a server.
func (w *webview) RunAndServe() {
	go w.Run()
	http.ListenAndServe(":"+string(w.port), w.router)
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
