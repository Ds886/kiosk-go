/*
Copyright Ds886 2022

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func fncMainKiosk(cfg cfgKiosk) {
	fncDispalyConfig(cfg)
	fncPIDCheck(cfg.Main.PIDFilePath)
	fncRunKiosk(cfg)
}

func fncConvertExitCodeToInt(strValidExitCodeString string) ([]int, error) {
	var arrFlagRawValidExitCode []string
	var arrOutput []int
	arrFlagRawValidExitCode = strings.Split(strValidExitCodeString, ",")

	for _, itemExitCode := range arrFlagRawValidExitCode {
		nCurrExitCode, err := strconv.Atoi(itemExitCode)
		if err != nil {
			return nil, fmt.Errorf("Error in parsing exit code: %s", itemExitCode)
		}
		arrOutput = append(arrOutput, nCurrExitCode)
	}
	return arrOutput, nil
}

func main() {
	// Init flags
	// config path
	var strFlagConfigPath string
	flag.StringVar(&strFlagConfigPath, "config", "", "Specify the config path")

	// General
	var strFlagPIDPath string
	flag.StringVar(&strFlagPIDPath, "pid-path", "", "Specify the path to create the PID path")

	var nFlagMaxRetries int
	flag.IntVar(&nFlagMaxRetries, "max-retries", -1, "Speiify the max amount of retries")

	var nFlagTimeout int
	flag.IntVar(&nFlagTimeout, "timeout", -1, "Speiify the timeout inbetween running apps in case of crash")

	// Target app
	var strFlagTargetApp string
	flag.StringVar(&strFlagTargetApp, "target-app", "", "Specify the runtime of the running app")

	var strFlagRawValidExitCode string
	var err error
	var arrFlagValidExitCode []int
	flag.StringVar(&strFlagRawValidExitCode, "exit-code", "", "override the valid exit codes [Param need to be comma-seprated i.e. '-exit-code \"0,2\"']")

	if strFlagRawValidExitCode != "" {
		arrFlagValidExitCode, err = fncConvertExitCodeToInt(strFlagRawValidExitCode)
	}

	// Target logging
	var bFlagOutputToStdOut bool
	flag.BoolVar(&bFlagOutputToStdOut, "redirect-stdout", true, "Overrides redirection of target app output to stdout")

	// Base config variable
	var cfg cfgKiosk

	flag.Parse()

	//var strPIDFile string = "/var/run/kiosk/kiosk.pid"
	//var strTargetApp string = "/usr/bin/firefox"
	strConfigPath, err := fncFindConfig(strFlagConfigPath)
	if err != nil {
		log.Fatalln(err.Error())
	}

	cfg = fncTOMLRead(strConfigPath)

	//Overriding config from params

	if strFlagPIDPath != "" {
		log.Println("Overriding PID file with parameter: ", strFlagPIDPath)
		cfg.Main.PIDFilePath = strFlagPIDPath
	}

	if nFlagMaxRetries != -1 {
		cfg.Main.MaxRetries = nFlagMaxRetries
	}

	if nFlagTimeout != -1 {
		log.Println("Overriding Timeout between attempts with: ", strFlagTargetApp)
		cfg.Main.Timeout = nFlagTimeout
	}

	if strFlagTargetApp != "" {
		log.Println("Overriding Target app with parameter: ", strFlagTargetApp)
		cfg.KioskTargetApp.TargetApp = strFlagTargetApp
	}

	if strFlagRawValidExitCode != "" {
		log.Println("Overriding valid exit codes from parameter: ", strFlagRawValidExitCode)
		cfg.KioskTargetApp.ValidExitCode = arrFlagValidExitCode
	}

	errConfig := fncVerifyConfig(cfg)
	if errConfig != nil {
		log.Fatalf("Error in config: \"%s\"", errConfig)
	}

	fncMainKiosk(cfg)

	chanSignal := make(chan os.Signal, 1)
	signal.Notify(
		chanSignal,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL)

	chanExit := make(chan int)
	go func() {
		for {
			sig := <-chanSignal
			switch sig {
			case syscall.SIGINT:
				log.Println("Exiting due to request from outside")
				fncPIDClear(cfg.Main.PIDFilePath)
				os.Exit(0)

			case syscall.SIGTERM:
				log.Println("Exiting due to request from outside")
				fncPIDClear(cfg.Main.PIDFilePath)
				os.Exit(0)

			case syscall.SIGQUIT:
				log.Println("Exiting due to request from outside")
				fncPIDClear(cfg.Main.PIDFilePath)
				os.Exit(0)

			case syscall.SIGKILL:
				log.Println("Exiting due to request from outside")
				fncPIDClear(cfg.Main.PIDFilePath)
				os.Exit(0)
			}
		}
	}()
	nExitCode := <-chanExit
	os.Exit(nExitCode)
}
