package main

import (
	"context"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/kardianos/service"
	"github.com/robfig/cron"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

type services struct {
	log service.Logger
	srv *http.Server
	cfg *service.Config
}

func (srv *services) Start(s service.Service) error {
	if srv.log != nil {
		srv.log.Info("Start run mydump")
	}

	lis, err := net.Listen("tcp", ":3308")
	if err != nil {
		return err
	}

	go srv.srv.Serve(lis)

	var cronObject = cron.New()
	var cronSpec = iniFile.Section("").Key("cron").String()
	cronObject.AddFunc(cronSpec, func() {
		dumpDatabase()
		deleteRedundantFiles()
	})

	go cronObject.Start()

	return nil
}

func (srv *services) Stop(s service.Service) error {
	if srv.log != nil {
		srv.log.Info("Start stop mydump")
	}
	return srv.srv.Shutdown(context.Background())
}

/*
func main() {
	deleteRedundantFiles()
}
*/

var iniFile *ini.File

func main() {
	// 日志的设置，放在程序最开始
	logFile, err := os.Create("mydump.log")
	if err != nil {
		logFile = os.Stdout
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	exeFileName, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	workDir, err := filepath.Abs(filepath.Dir(exeFileName))
	if err != nil {
		log.Fatal(err)
	}

	var iniFileName = "mydump.ini"
	var iniFileFullName = workDir + string(os.PathSeparator) + iniFileName

	if !pathExists(iniFileFullName) {
		var goPath = os.Getenv("GOPATH")
		iniFileFullName = goPath + string(os.PathSeparator) + "src" + string(os.PathSeparator) + "mydump" + string(os.PathSeparator) + iniFileName
	}

	if !pathExists(iniFileFullName) {
		log.Fatal("mydump.ini file not exist!")
	}

	iniFile, err = ini.Load(iniFileFullName)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.URL.Path)
	})

	var s = &services{srv: &http.Server{Handler: http.DefaultServeMux}, cfg: &service.Config{
		Name:        "mydump",
		DisplayName: "mydump",
		Description: "mysql dump service",
	}}

	sys := service.ChosenSystem()
	srv, err := sys.New(s, s.cfg)
	if err != nil {
		log.Fatalf("Init service error:%s\n", err.Error())
	}

	s.log, err = srv.SystemLogger(nil)
	if err != nil {
		log.Printf("Set logger error:%s\n", err.Error())
	}

	if len(os.Args) != 1 && len(os.Args) != 3 {
		// 参数格式不对
		return
	}

	if len(os.Args) == 3 {
		if os.Args[1] != "service" {
			// 参数格式不对
			return
		}

		var serviceOperation = os.Args[2]

		if serviceOperation == "install" {
			err := srv.Install()
			if err != nil {
				log.Fatalf("Install service error:%s\n", err.Error())
			}
		} else if serviceOperation == "uninstall" {
			err := srv.Uninstall()
			if err != nil {
				log.Fatalf("Uninstall service error:%s\n", err.Error())
			}
		} else {
			// 参数格式不对
		}

		return
	}

	err = srv.Run()

	if err != nil {
		log.Fatalf("Run programe error:%s\n", err.Error())
	}
}

/*

https://ini.unknwon.io/docs/intro/getting_started

*/
