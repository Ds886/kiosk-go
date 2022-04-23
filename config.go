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
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/sys/unix"
	"io/ioutil"
	"log"
	"os"
)

func fncVerifyConfig(cfg cfgKiosk) error {

	if cfg.Main.MaxRetries < CFG_MIN_RETRIES {
		return fmt.Errorf("MaxRetries too low, got %d. please make sure it is > %d", cfg.Main.MaxRetries, CFG_MIN_RETRIES)
	}

	if cfg.Main.Timeout < CFG_MIN_TIMEOUT {
		return fmt.Errorf("Timeout too low, got %d. please make sure it is >%d", cfg.Main.Timeout, CFG_MIN_TIMEOUT)
	}

	fileTargetApp, err := os.Stat(cfg.KioskTargetApp.TargetApp)
	if err != nil {
		return fmt.Errorf("Target application path: \"%s\"", cfg.KioskTargetApp.TargetApp)
	} else {
		if unix.Access(cfg.KioskTargetApp.TargetApp, unix.X_OK) != nil {
			return fmt.Errorf("Target application is not executable by the user. Application \"%s\" mode is: \"%s\"", cfg.KioskTargetApp.TargetApp, fileTargetApp.Mode().String())
		}
	}

	return nil
}

func fncValidateConfigPath(strPath string) (string, error) {
	filePath, err := os.Stat(strPath)
	if err != nil {
		return "", fmt.Errorf("Couldn't find config in: %s", strPath)
	} else {
		log.Println("Found file  ", filePath.Name(), " in path: ", strPath)
		return strPath, err
	}
}

func fncFindConfig(strConfigPath string) (string, error) {
	// Check provoide path
	var strFinalPath string
	var err error

	if strConfigPath != "" {
		log.Println("Trying the user provided path:", strConfigPath)
		strFinalPath, err = fncValidateConfigPath(strConfigPath)
		if err == nil {
			return strFinalPath, nil
		} else {
			log.Println("Couldn't find file in tagged file path")
		}
	} else {
		log.Println("Skipping user provided path(not provided)")
	}

	log.Println("Trying local path \"./kiosk.toml\"")
	strFinalPath, err = fncValidateConfigPath("./kiosk.toml")
	if err == nil {
		return strFinalPath, nil
	} else {
		log.Println(err.Error())
	}

	log.Println("Trying global path \"/etc/kiosk.toml\"")
	strFinalPath, err = fncValidateConfigPath("/etc/kiosk.toml")
	if err == nil {
		return strFinalPath, nil
	} else {
		return "", fmt.Errorf("Couldn't find any config")
	}
}

func fncTOMLRead(strFilePath string) cfgKiosk {
	var cfg cfgKiosk
	log.Println("Start reading file: ", strFilePath)

	fileConfig, err := ioutil.ReadFile(strFilePath)
	if err != nil {
		log.Fatalf("Error in openning toml config: %s", err)
	}

	err = toml.Unmarshal(fileConfig, &cfg)
	if err != nil {
		log.Fatalf("Error in parsing toml config: %s", err)
	}

	return cfg
}

func fncDispalyConfig(cfg cfgKiosk) {
	log.Println("Got the commands:")
	log.Println("	General:")
	log.Println("		PID file path: ", cfg.Main.PIDFilePath)
	log.Println("		Max retries: ", cfg.Main.MaxRetries)
	log.Println("		Timeout before restart attempt: ", cfg.Main.Timeout)
	log.Println("	Target App:")
	log.Println("		Target app path: ", cfg.KioskTargetApp.TargetApp)
	log.Println("		Valid exit codes: ", cfg.KioskTargetApp.ValidExitCode)
	log.Println("	Logging:")
	log.Println("		Redirect Target app output to stdout: ", cfg.Logging.OutputToStdOut)
}
