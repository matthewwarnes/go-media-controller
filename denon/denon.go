package denon

// Simple program to send commands to Denon AVR and get their result

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const denonip = "192.168.107.9"

var sources = map[string]string{
  "roku": "SAT/CBL",
	"switch": "GAME",
	"bluray": "BD",
	"steam": "AUX2",
	"laptop": "AUX1",
	"record": "PHONO",
	"wii": "DVD",
	"ps3": "MPLAY",
}

var (
	conn    net.Conn    // global network connection to the AVR
)

func sendCmd(cmd string) {
	fmt.Println("Sending: ", cmd)
	cmd = strings.ToUpper(cmd)
	cmd = cmd + "\r"
	fmt.Fprintf(conn, cmd)

}

func sendSrc(source string) {
	sendCmd("SI" + sources[source])
}

func receiver() {
	for { // There must be more information..keep reading.
		status, err := bufio.NewReader(conn).ReadString('\r')
		if status == "" {
			return
		}
		fmt.Println("received: ", status)
		if err != nil {
			fmt.Println("ERROR: ", err)
			return
		}
	}
}


func init() {
	lconn, err := net.Dial("tcp", denonip + "23")

	if err != nil {
		fmt.Println("Connection failed")
		os.Exit(1)
	}

	conn = lconn

	go receiver()
}

func DenonDo(cmd string, arg string) {
	defer conn.Close()


  switch cmd {
  case "on":
  	{
  		sendCmd("PWON")

  	}
    case "off":
  	{
  		sendCmd("PWSTANDBY")
  	}
	case "src":
		{
			sendSrc(arg)
		}
  case "volup":
    {
      sendCmd("MVUP")
    }
	case "vol":
		{
			sendCmd("MV" + arg)
		}
  case "voldown":
    {
      sendCmd("MVDOWN")
    }
	case "muteon":
		{
			sendCmd("MUON")
		}
	case "muteoff":
		{
			sendCmd("MUOFF")
		}
  }

	time.Sleep(1000 * time.Millisecond)


}


func PrintStatus() {
	sendCmd("SI?")
	sendCmd("PW?")
	sendCmd("MV?")
	sendCmd("MU?")
	time.Sleep(1000 * time.Millisecond)
}
