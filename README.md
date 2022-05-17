# kiosk-go

An implementation in go for enforcing a kiosk mode(running a single program)
This was done since a lot of solutions seems either too complex for a basic watch function or too clunky if you are going with online guides shell script

## Integration path
The project is meant to be used around the launching of the application and doesnt take care of any envrionement hardenning
It is possible to integrate it to any slim down window manager or DE by putting it as an autostart(either through the WM functionality or .xinitrc)  or possibly a user daemon with systemd

## Parameters
| Tag             | Parameter                     | Description                                                                                         | Example                                |
|-----------------|-------------------------------|-----------------------------------------------------------------------------------------------------|----------------------------------------|
| config          | Text                          | Path to the config file(Default: Looks for `./kiosk.toml` if fails  search in `/etc/kiosk.toml`)    | `-config "./config.toml"`              |
| exit-code       | Comma seperated list("1,2,3") | List of valid exit codes(by default every non-zero exit will be treated as an error)                | `-exit-code "0,1"`                     |
| max-retries     | Number                        | The max amount of retries(Default: 5)                                                               | `-max-retries 5`                       |
| pid-path        | Text                          | Path to place the PID file to ensure only one instance running(Default: `/var/run/kiosk/kiosk.pid`) | `-pid-path "/var/run/kiosk/kiosk.pid"` |
| redirect-stdout | `true`/`false`                | Print the program output to stdout of the kiosk(Default: true)                                      | `-redirect-stdout true`                |
| target-app      | Text                          | Path of the target app to launch                                                                    | `-target-app "/usr/bin/firefox"`       |
| timeout         | Time in ms                    | Number in seconds to wait between each launch(Default: 2000)                                        | `-timeout 1000`                           |


## Configuration
An example for Configuration file
```
[Main]
# Path to PID file
PIDFilePath = "$HOME/.local/run/kiosk/kiosk.pid"
# The max amount of retries(Minimum 1)
MaxRetries = 3
# Timeout  between each attempt in ms(Minimum 500)
Timeout = 2000

[KioskTargetApp]
# Target app executable
TargetApp = "/snap/bin/chromium"
# Target app params
TargetAppArgs = [ "https://google.com" ]
# Override valid exit codes
ValidExitCode = [ 0 ]

[Logging]
# Redirect output of the target program to the kiosk stdout
OutputToStdOut = true
```

## Compiling

### Prerequisites
#### Required
 - Go v1.13+
 - go-md2man
 - make

#### Optional(for .deb packaging)
 - dpkg-deb
 - gzip

### Development
1. Run the command `make`

Compiled binaries are found under the bin directory

### Release

1. Run the command `make release`
1. Run the command `sudo make install`

* PREfIX cna be changed using the command `make release prefix=<usr path> DESTDIR=<root path>` by default it will be installed to  `/usr/local` *

### Debian package

1. Run the command `make deb`

## Todo
1. Make an option to skip the config
1. Create a Systemd service file and a more guided install
1. Improve logging features
  1. Add log levels(default only fatal)
  1. Add logging to file
    1. Add custom path
    1. Max file size
    1. logrotate script
1. Improve folder structure to be more go friendly
