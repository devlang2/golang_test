package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	//	"html"
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

func view(w http.ResponseWriter, r *http.Request) {
	page := `
<!doctype html>
<!--[if lt IE 8]><html class="no-js lt-ie8"> <![endif]-->
<html class="no-js">
<head>
    <meta charset="utf-8">
    <title>HTTPS Query 1.03</title>
    <link rel="stylesheet" href="http://fonts.googleapis.com/earlyaccess/nanumgothic.css">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
    <style>
        body {
            margin: 10px;
        }
        .mt5 {
            margin-top: 5px;
        }
        .p0 {
            padding: 0px;
        }
    </style>
</head>
<body>
    <form id="req" method="post" action="/">
        <div class="row">
            <div class="col-xs-10">
                <label>Query</label>
                <input type="text" class="form-control " name="query" value="https://google.com"/>
            </div>
            <div class="col-xs-2">
                <label>&nbsp;</label>
                <button type="submit" class="btn btn-danger btn-block">Send</button>
            </div>
        </div>
        <p></p>

        <div class="row">
            <div class="col-lg-12">
                <label>Body</label>
                <textarea name="body" rows="5" class="form-control"></textarea>
            </div>
        </div>
        <p></p>

        <label>Response</label>
        <button type="button" class="btn btn-primary btn-xs btn-sm btn-copy">Copy</button>
        <div class="panel">
            <div class="panel-body p0">
                <pre id="response"></pre>
            </div>
        </div>
    </form>

    <script src="http://code.jquery.com/jquery-3.2.1.min.js" integrity="sha256-hwg4gsxgFZhOsEEamdOYGBf13FyQuiTwlAQgxVSNgt4=" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
    <script>
        $("#req").submit(function(e) {
            e.preventDefault();
            console.log("###");

            $.ajax({
                type: "POST",
                async: true,
                url: "/req",
                data: $("#req").serialize()
            }).done(function(data) {
                $("#response").text(data);
            });
        });

        $(".btn-copy").click(function(e) {
            CopyToClipboard("response");
        });
        function CopyToClipboard(containerid) {
            if (document.selection) {
                var range = document.body.createTextRange();
                range.moveToElementText(document.getElementById(containerid));
                range.select().createTextRange();
                document.execCommand("Copy");

            } else if (window.getSelection) {
                var range = document.createRange();
                range.selectNode(document.getElementById(containerid));
                window.getSelection().addRange(range);
                document.execCommand("Copy");
                alert("Copied");
            }
        }
    </script>
</body>
</html>
    `

	fmt.Fprintf(w, page)
}

func request(w http.ResponseWriter, r *http.Request) {
	// Parse forn
	r.ParseForm()
	query := strings.TrimSpace(r.FormValue("query"))
	body := strings.TrimSpace(r.FormValue("body"))

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
			return
		}
		defer res.Body.Close()

		// Get response
		content, _ := ioutil.ReadAll(res.Body)
		response = strings.TrimSpace(string(content))
		//		response = html.EscapeString(response)

		//		fmt.Fprintf(w, response)
		fmt.Fprint(w, response)
	} else {

	}

	//	// https://10.0.7.209:4000/sniper.atx?major=query&minor=select&name=rule&length=0&delim=
	//	form := `%s

	//%s
	//    `
	//	form = fmt.Sprintf(form, getHeader(), query, body, response, getFooter())
	//	fmt.Fprintf(w, form)

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
	http.HandleFunc("/", view)
	http.HandleFunc("/req", request)
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
