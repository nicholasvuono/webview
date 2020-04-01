package webview

import (
	"net/http"
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
func New(url string, port int, router *http.ServeMux) *webview {
	return &webview{
		url:    url,
		port:   port,
		router: router,
	}
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
