# NAME
kiosk.toml - A configutation file in a toml format for setting kiosk-go 
             parameters

# DESCRIPTION
A configuration file for kiosk-go, by default will be first read from the 
relative path ./kiosk.toml and if not found will fallback to /etc/kiosk.toml

# SYNTAX

## MAIN

- PIDFilePath - String - the absoulute destination path for the PID file 
- MaxRetries - The max amount of retries for launching the app(minimum 1)
- Timeout - Number of ms - the amount of time to wait between each launch of 
    the app 

## KioskTargetApp

- TargetApp - String - the absoulute path to target executable
- TargetAppArgs - List of params - List of params to pass to the target app
- ValidExitCode - List of numbers - List of override valid exit codes
    (by default non-zero exit code is treated as an error)

## Logging

- OutputToStdOut - Boolean - Wheteher the app output will be redirected to 
                              stdout (Default: true)    

# EXAMPLE

\[Main\]
\# Path to PID file
PIDFilePath = "$HOME/.local/run/kiosk/kiosk.pid"
\# The max amount of retries(Minimum 1)
MaxRetries = 3
\# Timeout  between each attempt in ms(Minimum 500)
Timeout = 2000

\[KioskTargetApp\]
\# Target app executable
TargetApp = "/snap/bin/chromium"
\# Target app params
TargetAppArgs = \[ "https://google.com" \]
\# Override valid exit codes
ValidExitCode = \[ 0 \]

\[Logging\]
\# Redirect output of the target program to the kiosk stdout
OutputToStdOut = true

# COPYRIGHT
Copyright 2022 Ds886

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
