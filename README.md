# webview
[![Go Report Card](https://goreportcard.com/badge/github.com/nicholasvuono/webview)](https://goreportcard.com/report/github.com/nicholasvuono/webview)

A simple cross-platform webview library for the Go programming language. 

The aim of this project is to create an easy-to-use API for generating a UI layer using the system's default browser.

## Features

* Pure Go library
* Simple API
* Less than 100 lines of code
* Cross-platform (Windows, Mac, Linux)
* Use any frontend web technology

## How it works

Behind the scenes, webview uses a single command (depending on the operating system) to open up a URL in the system's default browser. In essence, this project has been designed to allow access to the frontend of a local server by opening up a browser instance and pointing to the route at which the frontend is being served up. Optionally, if you dont want to start your own server separately within your code, you can provide a router (a.k.a http.ServeMux) and the webview library will do it for you. This may make a bit more sense after reading the examples below.

## Simplest Working Example

```go
package main

import ( 
        "fmt"
        "log"
        
        "github.com/nicholasvuono/webview"
 )

func main() {
        webview, err := webview.New("https://github.com/nicholasvuono/webview")
        if err != nil {
                fmt.Println(err)
        }
        log.Fatal(webview.Run())
}
```
