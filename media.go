package main

import (
	"fmt"


	"os"

  "go-media-controller/roku"
  "go-media-controller/denon"
)

const (
	VERSION = "0.4.1"
	USAGE   = `usage: groku [--version] [--help] <command> [<args>]
CLI remote for your Roku
Commands:
  home            Return to the home screen
  rev             Reverse
  fwd             Fast Forward
  select          Select
  left            Left
  right           Right
  up              Up
  down            Down
  back            Back
  info            Info
  backspace       Backspace
  enter           Enter
  search          Search
  replay          Replay
  play            Play
  pause           Pause
  discover        Discover a roku on your local network
  text            Send text to the Roku
  apps            List installed apps on your Roku
  app             Launch specified app
`
)

func main() {

	switch os.Args[1] {
	case "home", "rev", "fwd", "select", "left", "right", "down", "up",
		"back", "info", "backspace", "enter", "search", "replay", "play", "pause", "apps", "app":
    if len(os.Args) > 2 {
		  roku.RokuDo(os.Args[1], os.Args[2])
    } else {
      roku.RokuDo(os.Args[1], "")
    }
		os.Exit(0)
  case "on", "off", "vol", "volup", "voldown", "src", "muteon","muteoff":
		if len(os.Args) > 2 {
			denon.DenonDo(os.Args[1], os.Args[2])
		} else {
			denon.DenonDo(os.Args[1], "")
		}
    os.Exit(0)
	case "status":
		denon.PrintStatus()
		roku.PrintStatus()
	default:
		fmt.Println(USAGE)
		os.Exit(1)
	}
}
