package webview

import (
	"os/exec"
	"runtime"
)

//WebView holds a url string to navigate to.
type WebView struct {
	URL string
}

//Run gets and runs a command to open a browser session.
func (w *WebView) Run() {
	webview := getOpenCommand(w.URL)
	webview.Run()
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
