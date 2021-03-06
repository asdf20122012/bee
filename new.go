package main

import (
	"fmt"
	"os"
	path "path/filepath"
	"strings"
)

var cmdNew = &Command{
	UsageLine: "new [appname]",
	Short:     "create an application base on beego framework",
	Long: `
create an application base on beego framework,

which in the current path with folder named [appname].

The [appname] folder has following structure:

    |- main.go
    |- conf
        |-  app.conf
    |- controllers
         |- default.go
    |- models
    |- static
         |- js
         |- css
         |- img             
    |- views
        index.tpl                   

`,
}

func init() {
	cmdNew.Run = createApp
}

func createApp(cmd *Command, args []string) {
	curpath, _ := os.Getwd()
	if len(args) != 1 {
		fmt.Println("[ERRO] Argument [appname] is missing")
		os.Exit(2)
	}

	gopath := os.Getenv("GOPATH")
	Debugf("gopath:%s", gopath)
	if gopath == "" {
		fmt.Printf("[ERRO] $GOPATH not found\n")
		fmt.Printf("[HINT] Set $GOPATH in your environment vairables\n")
		os.Exit(2)
	}
	haspath := false
	appsrcpath := ""

	wgopath := path.SplitList(gopath)
	for _, wg := range wgopath {
		wg = path.Join(wg, "src")

		if path.HasPrefix(strings.ToLower(curpath), strings.ToLower(wg)) {
			haspath = true
			appsrcpath = wg
			break
		}
	}

	if !haspath {
		fmt.Printf("[ERRO] Unable to create an application outside of $GOPATH(%s)\n", gopath)
		fmt.Printf("[HINT] Change your work directory by `cd $GOPATH%ssrc`\n", string(path.Separator))
		os.Exit(2)
	}

	apppath := path.Join(curpath, args[0])

	if _, err := os.Stat(apppath); os.IsNotExist(err) == false {
		fmt.Printf("[ERRO] Path(%s) has alreay existed\n", apppath)
		os.Exit(2)
	}

	fmt.Println("[INFO] Creating application...")

	os.MkdirAll(apppath, 0755)
	fmt.Println(apppath + string(path.Separator))
	os.Mkdir(path.Join(apppath, "conf"), 0755)
	fmt.Println(path.Join(apppath, "conf") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "controllers"), 0755)
	fmt.Println(path.Join(apppath, "controllers") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "models"), 0755)
	fmt.Println(path.Join(apppath, "models") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static"), 0755)
	fmt.Println(path.Join(apppath, "static") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static", "js"), 0755)
	fmt.Println(path.Join(apppath, "static", "js") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static", "css"), 0755)
	fmt.Println(path.Join(apppath, "static", "css") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "static", "img"), 0755)
	fmt.Println(path.Join(apppath, "static", "img") + string(path.Separator))
	fmt.Println(path.Join(apppath, "views") + string(path.Separator))
	os.Mkdir(path.Join(apppath, "views"), 0755)
	fmt.Println(path.Join(apppath, "conf", "app.conf"))
	writetofile(path.Join(apppath, "conf", "app.conf"), strings.Replace(appconf, "{{.Appname}}", args[0], -1))

	fmt.Println(path.Join(apppath, "controllers", "default.go"))
	writetofile(path.Join(apppath, "controllers", "default.go"), controllers)

	fmt.Println(path.Join(apppath, "views", "index.tpl"))
	writetofile(path.Join(apppath, "views", "index.tpl"), indextpl)

	fmt.Println(path.Join(apppath, "main.go"))
	writetofile(path.Join(apppath, "main.go"), strings.Replace(maingo, "{{.Appname}}", strings.Join(strings.Split(apppath[len(appsrcpath)+1:], string(path.Separator)), string(path.Separator)), -1))

	fmt.Println("[SUCC] New application successfully created!")
}

var appconf = `appname = {{.Appname}}
httpport = 8080
runmode = dev
`

var maingo = `package main

import (
	"{{.Appname}}/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}

`
var controllers = `package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}
`

var indextpl = `<!DOCTYPE html>

<html>
  	<head>
    	<title>Beego</title>
    	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  	</head>
	
	<style type="text/css">
		body {
			margin: 0px;
			font-family: "Helvetica Neue",Helvetica,Arial,sans-serif;
			font-size: 14px;
			line-height: 20px;
			color: rgb(51, 51, 51);
			background-color: rgb(255, 255, 255);
		}

		.hero-unit {
			padding: 60px;
			margin-bottom: 30px;
			border-radius: 6px 6px 6px 6px;
		}

		.container {
			width: 940px;
			margin-right: auto;
			margin-left: auto;
		}

		.row {
			margin-left: -20px;
		}

		h1 {
			margin: 10px 0px;
			font-family: inherit;
			font-weight: bold;
			text-rendering: optimizelegibility;
		}

		.hero-unit h1 {
			margin-bottom: 0px;
			font-size: 60px;
			line-height: 1;
			letter-spacing: -1px;
			color: inherit;
		}

		.description {
		    padding-top: 5px;
		    padding-left: 5px;
		    font-size: 18px;
		    font-weight: 200;
		    line-height: 30px;
		    color: inherit;
		}

		p {
		    margin: 0px 0px 10px;
		}
	</style>
  	
  	<body>
  		<header class="hero-unit" style="background-color:#A9F16C">
			<div class="container">
			<div class="row">
			  <div class="hero-text">
			    <h1>Welcome to Beego!</h1>
			    <p class="description">
			    	Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra.
			    <br />
			    	Official website: <a href="http://{{.Website}}">{{.Website}}</a>
			    <br />
			    	Contact me: {{.Email}}</a>
			    </p>
			  </div>
			</div>
			</div>
		</header>
	</body>
</html>
`

func writetofile(filename, content string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(content)
}
