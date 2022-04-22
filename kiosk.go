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
	"fmt"
	"github.com/shirou/gopsutil/v3/process"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func fncVerifyTargetAppRunnig(strTatgetApp string) {

	//Cleaning to just the executable
	arrTatgetAppClean := strings.Split(strTatgetApp, "/")
	strTatgetAppClean := arrTatgetAppClean[len(arrTatgetAppClean)-1]

	arrProcess, errProcessList := process.Processes()
	if errProcessList != nil {
		fmt.Println(errProcessList)
	}

	for _, itemProcess := range arrProcess {
		// Get name
		strCurrName, err := itemProcess.Name()
		if err != nil {
		}

		// Check if kioks alread running
		reTargetApp, errRegex := regexp.MatchString(strTatgetAppClean, strCurrName)
		if errRegex != nil {
			log.Print("Error in parsing regex")
		}

		if reTargetApp == true {
			log.Print(strTatgetAppClean, " is already running with pid:", itemProcess.Pid, " exiting")
			err := itemProcess.Kill()
			if err != nil {
				log.Fatalln("Failed in killing the process, exiting due to assuming insufficient permission")
			}
		}
	}
}

func fncRunKiosk(cfg cfgKiosk) bool {
	var nCurrentRetries int = 0
	// To first initialize the variable
	cmdCommand := exec.Command(cfg.KioskTargetApp.TargetApp, "")

	for {
		if nCurrentRetries > cfg.Main.MaxRetries {
			log.Println("Max retries reached exiting")
			os.Exit(-1)
		}

		fncVerifyTargetAppRunnig(cfg.KioskTargetApp.TargetApp)

		// Running target app
		cmdCommand = exec.Command(cfg.KioskTargetApp.TargetApp, cfg.KioskTargetApp.TargetAppArgs...)
		if cfg.Logging.OutputToStdOut == true {
			cmdCommand.Stdout = os.Stdout
		}
		err := cmdCommand.Start()

		if err != nil {
			log.Println("Failed in running the process:", err.Error())
		}

		log.Println("Prorgam run with PID:", cmdCommand.Process.Pid)
		cmdCommand.Wait()

		if cmdCommand.ProcessState.ExitCode() != 0 {
			log.Println("Processes exited with non-zero exit code")
			nCurrentRetries++
		}

		time.Sleep(time.Duration(cfg.Main.Timeout) * time.Second)
	}
}
