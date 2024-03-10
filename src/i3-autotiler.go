package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"slices"
	"time"
)

type Msg struct {
	Change    string    `json:"change"`
	Container Container `json:"container"`
}

type Container struct {
	WindowRect WindowRect `json:"window_rect"`
}

type WindowRect struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func waitForI3() {
	cmd := exec.Command("ps", "-C", "i3", "-o", "pid=")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			output, _ := cmd.Output()
			if len(output) > 0 {
				println("Process found for i3:", output)
				return
			} else {
				println("Waiting for i3 to start...")
			}
		}
	}

}

func main() {
	waitForI3()

	debug := slices.Contains(os.Args, "--debug")
	msgEvent := exec.Command("i3-msg", "-t", "subscribe", "-m", "[ \"window\" ]")

	stdout, err := msgEvent.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = msgEvent.Start()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		var msg Msg

		m := scanner.Text()
		err = json.Unmarshal([]byte(m), &msg)
		if err != nil {
			log.Fatal(err)
		}

		if msg.Change == "focus" {
			if msg.Container.WindowRect.Width > msg.Container.WindowRect.Height {
				if debug {
					println("Width", msg.Container.WindowRect.Width, ", Height", msg.Container.WindowRect.Height, " ,Next split will be horizontal")
				}
				exec.Command("i3-msg", "split", "horizontal").Run()
			} else {
				if debug {
					println("Width", msg.Container.WindowRect.Width, ", Height", msg.Container.WindowRect.Height, " ,Next split will be vertical")
				}
				exec.Command("i3-msg", "split", "vertical").Run()
			}
		}
	}

	msgEvent.Wait()
}
