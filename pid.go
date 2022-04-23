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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func fncPIDClear(strPidFile string) {
	errDeletePID := os.Remove(strPidFile)
	if errDeletePID != nil {
		log.Print("Failed to remove file")
	}
}

func fncPIDCreate(strPidFile string) {
	strPIDFolder := filepath.Dir(strPidFile)
	filePIDFolder, err := os.Stat(strPIDFolder)
	if err != nil {
		err := os.MkdirAll(strPIDFolder, os.ModeSticky|os.ModePerm)
		if err != nil {
			log.Fatalf("Folder: \"%s\" is not writeable by the user. Error: %s", strPIDFolder, err)
		}
	} else {
		log.Printf("Passed check if PID folder: \"%s\" is accesible", filePIDFolder.Name())
	}

	filePid, errFile := os.Create(strPidFile)
	if errFile != nil {
		log.Fatal("Can't create pid file in" + strPidFile)
	}
	defer filePid.Close()
	os.Truncate(strPidFile, 0)
	fmt.Fprint(filePid, os.Getpid())
}

func fncPIDCheck(strPidFile string) {
	if _, err := os.Stat(strPidFile); os.IsNotExist(err) {
		fncPIDCreate(strPidFile)
	} else {
		log.Println("kiosk already running - terminating existing instance")

		datPid, errDatPid := ioutil.ReadFile(strPidFile)
		if errDatPid != nil {
			log.Fatal("Error reading the file: ", errDatPid)
		}

		nCurrentPID, errConvertPID := strconv.Atoi(string(datPid))
		if errConvertPID != nil {
			log.Fatal("Failed to convert string: ", errConvertPID)
		}

		pActive, errGetActiveProcess := os.FindProcess(nCurrentPID)
		if errGetActiveProcess != nil {
			log.Print("Failed getting active process creating new one")
		} else {
			pActive.Kill()
		}
		fncPIDCreate(strPidFile)

	}
}
