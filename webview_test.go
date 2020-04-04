package webview

import (
	"net/http"
	"strconv"
	"testing"
)

var tests = []func(t *testing.T){
	TestNew,
	TestSetPort,
	TestSetRouter,
	TestRun,
	TestRunAndServe,
}

func TestNew(t *testing.T) {
	_, err := New("https://github.com/nicholasvuono/webview")
	if err != nil {
		t.Error(err)
	}
}

func TestSetPort(t *testing.T) {
	webview, _ := New("https://github.com/nicholasvuono/webview")
	err := webview.SetPort(3000)
	if err != nil {
		t.Error(err)
	}
	err = webview.SetPort(-1)
	if err == nil {
		t.Error(err)
	}
}

func TestSetRouter(t *testing.T) {
	webview, _ := New("https://github.com/nicholasvuono/webview")
	err := webview.SetRouter(http.NewServeMux())
	if err != nil {
		t.Error(err)
	}
}

func TestRun(t *testing.T) {
	webview, _ := New("https://github.com/nicholasvuono/webview")
	err := webview.Run()
	if err != nil {
		t.Error(err)
	}
}

func TestRunAndServe(t *testing.T) {
	webview, _ := New("https://github.com/nicholasvuono/webview")
	webview.SetPort(3000)
	webview.SetRouter(http.NewServeMux())
	err := webview.RunAndServe()
	if err == nil {
		t.Error(err)
	}
}

func TestAll(t *testing.T) {
	for i, test := range tests {
		t.Run(strconv.Itoa(i), test)
	}
}
