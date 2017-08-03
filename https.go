package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

// Flags
var fs *flag.FlagSet
var port *string

// Init
const (
	DefaultServerName = "https"
	DefaultPortNumber = "8080"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// Parse firn
	r.ParseForm()
	query := strings.TrimSpace(r.FormValue("Query"))
	body := strings.TrimSpace(r.FormValue("Body"))

	var response string
	spew.Dump()
	if len(query) > 0 {
		// Send request
		log.Println(query)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{
			Transport: tr,
			Timeout:   time.Second * 6,
		}
		res, err := client.Post(query, "text/plain", bytes.NewBufferString(body))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "http://127.0.0.1:"+*port, 301)
			return
		}
		defer res.Body.Close()

		// Get response
		content, _ := ioutil.ReadAll(res.Body)
		response = strings.TrimSpace(string(content))
		response = html.EscapeString(response)
	} else {
		query = "https://"
	}

	// https://10.0.7.209:4000/sniper.atx?major=query&minor=select&name=rule&length=0&delim=
	form := `%s
        <form name="Query" method="post" action="/">
            <p>
                <label>Query</label>
                <input type='text' class="console" name='Query' value='%s' style="height: 20px"></html>
            </p>

            <p>
                <label>Body</label>
                <textarea name="Body" rows="5" class="console">%s</textarea>
                <input type="submit" value="Send" style="width: 100px; height:45px"/>
            </p>

            <p>
                <label>Response</label>
                <pre id="response" class="response console">%s</p[re>
            </p>
        </form>
%s
    `
	form = fmt.Sprintf(form, getHeader(), query, body, response, getFooter())
	fmt.Fprint(w, form)

}

func getHeader() string {
	return `<!doctype html>
<!--[if lt IE 8]><html class="no-js lt-ie8"> <![endif]-->
<html class="no-js">
<head>
    <meta charset="utf-8">
    <title>HTTPS Query 1.03</title>
    <link rel='stylesheet' href='http://fonts.googleapis.com/earlyaccess/nanumgothic.css'>
    <style>
        body { font-family: "Nanum Gothic", "맑은 고딕", sans-serif;}
        textarea {width: 100%%; border: 1px dashed #acacac;}
        table {width: 100%%; }
        .console {font-family: consolas; width: 100%; }
        .response { height: 250px; overflow-y: scroll;  border: 1px dashed #acacac; }
        input {width: 100%%; border: 1px dashed #acacac; font-size: 14px; }

    </style>
</head>
<body>
    `
}

func getFooter() string {
	return `
    </body>
</html>`
}

func main() {
	log.SetPrefix("[" + DefaultServerName + "] ")

	// Set flags
	fs = flag.NewFlagSet("", flag.ExitOnError)
	port = fs.String("p", DefaultPortNumber, "HTTP port number")
	fs.Usage = printHelp
	fs.Parse(os.Args[1:])

	// Open browser
	openBrowser("http://127.0.0.1:" + *port)

	// Run http server
	log.Printf("HTTP server listening to %s", *port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+*port, nil)

}

func printHelp() {
	fmt.Println(DefaultServerName + " [options]")
	fs.PrintDefaults()
}

func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
