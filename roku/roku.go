package roku

import (
  "encoding/xml"
  "fmt"
  "net/http"
)

const rokuAddress = "http://192.168.107.10:8060/"


type dictonary struct {
	XMLName xml.Name `xml:"apps"`
	Apps    []app    `xml:"app"`
}

type activeapp struct {
  XMLName xml.Name `xml:"active-app"`
	App    string    `xml:"app"`
}

type app struct {
	Name string `xml:",chardata"`
	ID   string `xml:"id,attr"`
}

type grokuConfig struct {
	Address   string `json:"address"`
	Timestamp int64  `json:"timestamp"`
}

func RokuDo(cmd string, arg string) {

	switch cmd {
	case "home", "rev", "fwd", "select", "left", "right", "down", "up",
		"back", "info", "backspace", "enter", "search":
		http.PostForm(fmt.Sprintf("%vkeypress/%v", rokuAddress, cmd), nil)
		return
	case "replay":
		http.PostForm(fmt.Sprintf("%vkeypress/%v", rokuAddress, "InstantReplay"), nil)
		return
	case "play", "pause":
		http.PostForm(fmt.Sprintf("%vkeypress/%v", rokuAddress, "Play"), nil)
	   return
	case "apps":
		dict := queryApps()
		for _, a := range dict.Apps {
			fmt.Println(a.Name)
		}
		return
	case "app":
		dict := queryApps()

		for _, a := range dict.Apps {
			if a.Name == arg {
				http.PostForm(fmt.Sprintf("%vlaunch/%v", rokuAddress, a.ID), nil)
				return
			}
		}
		fmt.Printf("App %q not found\n", arg)
		return
	default:
		return
	}
}

func PrintStatus() {

  resp, err := http.Get(fmt.Sprintf("%squery/active-app", rokuAddress))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

  var app activeapp
  if err := xml.NewDecoder(resp.Body).Decode(&app); err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println("Roku App: " + app.App)
}

func queryApps() dictonary {
	resp, err := http.Get(fmt.Sprintf("%squery/apps", rokuAddress))
	if err != nil {
		fmt.Println(err)
		return dictonary{}
	}

	defer resp.Body.Close()

	var dict dictonary
	if err := xml.NewDecoder(resp.Body).Decode(&dict); err != nil {
		fmt.Println(err)
		return dictonary{}
	}

	return dict
}
