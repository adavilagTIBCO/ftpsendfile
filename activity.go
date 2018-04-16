package ftpsendfile

import (
	"fmt"
	"os"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/dutchcoders/goftp"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	//Input declaration
	server := context.GetInput("server").(string)
	port := context.GetInput("port").(int)
	username := context.GetInput("username").(string)
	password := context.GetInput("password").(string)
	pathsrc := context.GetInput("pathsrc").(string)
	filesrc := context.GetInput("filesrc").(string)
	pathdest := context.GetInput("pathdest").(string)
	filedest := context.GetInput("filedest").(string)
	url := server + ":" + strconv.Itoa(port)

	// do eval
	var ftp *goftp.FTP

	// For debug messages: goftp.ConnectDbg("ftp.server.com:21")
	if ftp, err = goftp.Connect(url); err != nil {
		panic(err)
	}

	defer ftp.Close()
	fmt.Println("Successfully connected to ", server)

	// Username / password authentication
	if err = ftp.Login(username, password); err != nil {
		panic(err)
	}

	if err = ftp.Cwd("/"); err != nil {
		panic(err)
	}

	var curpath string
	if curpath, err = ftp.Pwd(); err != nil {
		panic(err)
	}

	fmt.Printf("Current path: %s", curpath)

	// Get directory listing
	var files []string
	if files, err = ftp.List("/"); err != nil {
		panic(err)
	}
	fmt.Println("Directory listing:", files)

	// Upload a file
	var file *os.File
	if file, err = os.Open(pathsrc + filesrc); err != nil {
		panic(err)
	}

	if err := ftp.Stor(pathdest+filedest, file); err != nil {
		panic(err)
	}

	context.SetOutput("output", "Successfully connected to "+url)

	return true, nil
}
