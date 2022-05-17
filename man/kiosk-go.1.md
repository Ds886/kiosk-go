# kiosk-go

# NAME
  An application enforcing a single app running


# SYNOPSIS
kiosk-go \[-target-app <path>\] \[-config <path>\] \[-exit-code "N,N,N"\]
         \[-max-retries N\] \[-pid-path <path>\] \[-redirect-stdout <true/false>\]
         \[-timeout N\]


# DESCRiPTION
An application enforcing a single app running by default it will recieve it's
parameters from a config file either from the file kiosk.toml in the local
directory or from /etc/kiosk.toml


# PARAMETERS
-target-app 
   Absoulute path to the target app

-config 
  Absoulute path for the config file if fails will searchin the following
  order:
    1. "./kiosk.toml"
    2. "/etc/kiosk.toml"

-exit-code "N,N,N"
  Overrides the list of what considered valid exit codesi
  (Default: any non-zero exit code)

-max-retries N
.br
  The max amount of retries (Default: 5)

-pid-path 
  Absoulute path to  place the PID (Default: /var/run/kiosk/kiosk.pid)

-redirect-stdout 
  Enables redirection to stdout (Default: true)

-timeout N(ms)
  Number of ms to wait between each launch


# EXAMPLES
- Running using default paths:
  kiosk-go

- Running with specified app:
  kiosk-go -target-app "/usr/bin/firefox"

- Full command
    kiosk -target-app "/usr/bin/firefox" -config "/opt/kiosk/config.toml"
      -exit-code "5,7" -max-retries 10 -pid-path "/opt/kiosk.pid"
      -redirect-stdout true -timeout 500


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
