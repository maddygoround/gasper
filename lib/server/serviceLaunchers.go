package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/SWS/lib/database"
	"github.com/sdslabs/SWS/lib/docker"
	"github.com/sdslabs/SWS/lib/utils"
	"github.com/sdslabs/SWS/services/dominus"
	"github.com/sdslabs/SWS/services/mysql"
	"github.com/sdslabs/SWS/services/node"
	"github.com/sdslabs/SWS/services/php"
	"github.com/sdslabs/SWS/services/python"

	sshserver "github.com/gliderlabs/ssh"
	"github.com/sdslabs/SWS/services/ssh"
	"github.com/sdslabs/SWS/services/static"
)

// UnivServer is used for handling all types of servers
type UnivServer struct {
	SSHServer  *sshserver.Server
	HTTPServer *http.Server
}

// Bind the services to the launchers
var launcherBindings = map[string]func(string, string) UnivServer{
	"ssh":   startSSHService,
	"mysql": startMySQLService,
	"app":   startAppService,
}

// Bind services to routers here
var serviceBindings = map[string]*gin.Engine{
	"dominus": dominus.Router,
	"static":  static.Router,
	"php":     php.Router,
	"node":    node.Router,
	"python":  python.Router,
	"mysql":   mysql.Router,
}

func initDefaultServer() UnivServer {
	var customServer UnivServer
	customServer.HTTPServer = nil
	customServer.SSHServer = nil
	return customServer
}

func initHTTPServer(service, port string) UnivServer {
	server := &http.Server{
		Addr:         port,
		Handler:      serviceBindings[service],
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	var customServer UnivServer
	customServer.SSHServer = nil
	customServer.HTTPServer = server
	return customServer
}

func startMySQLService(service, port string) UnivServer {
	containers := docker.ListContainers()
	if !utils.Contains(containers, "/mysql") {
		fmt.Printf("No Mysql instance found in host. Building the instance.")
		containerID, err := database.SetupDBInstance()
		if err != nil {
			fmt.Println("There was a problem deploying MySql service.")
			fmt.Printf("ERROR:: %s\n", err.Error())
		} else {
			fmt.Printf("Container has been deployed with ID:\t%s \n", containerID)
		}
	}
	server := initHTTPServer(service, port)
	return server
}

func startSSHService(service, port string) UnivServer {
	var customServer UnivServer
	server, err := ssh.BuildSSHServer(service)
	if err != nil {
		fmt.Println("There was a problem deploying SSH service. Make sure the address of Private Keys is correct in `config.json`.")
		fmt.Printf("ERROR:: %s\n", err.Error())
		defaultServer := initDefaultServer()
		return defaultServer
	}
	customServer.HTTPServer = nil
	customServer.SSHServer = server
	return customServer
}

func startAppService(service, port string) UnivServer {
	server := initHTTPServer(service, port)
	return server
}

// Launcher invokes the respective launcher functions for the services
func Launcher(service, port string) UnivServer {
	if strings.HasPrefix(service, "ssh") {
		return launcherBindings["ssh"](service, port)
	} else if strings.HasPrefix(service, "mysql") {
		return launcherBindings["mysql"](service, port)
	} else if service != "" {
		return launcherBindings["app"](service, port)
	}

	defaultServer := initDefaultServer()
	return defaultServer
}
