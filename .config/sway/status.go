package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func currentTime() string {
	t := time.Now().Local()
	return t.Format("2006-01-02 15:04:05")
}

func battery() (string, string) {
	data, err := os.ReadFile("/sys/class/power_supply/BAT0/status")
	if err != nil {
		fmt.Println(err)
	}
	status := strings.TrimSpace(string(data))

	data, err = os.ReadFile("/sys/class/power_supply/BAT0/capacity")
	if err != nil {
		fmt.Println(err)
	}
	capacity := strings.TrimSpace(string(data))

	return status, capacity
}

func volume() string {
	cmdstr := "amixer sget Master | grep 'Right:' | awk -F'[][]' '{print $4}'"
	cmd := exec.Command("sh", "-c", cmdstr)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	isMuted := strings.TrimSpace(string(out)) == "off"
	if isMuted {
		return "<span color='yellow'>muted</span>"
	}

	cmdstr = "amixer sget Master | grep 'Right:' | awk -F'[][]' '{print $2}'"
	cmd = exec.Command("sh", "-c", cmdstr)
	out, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	volume := strings.TrimSpace(string(out))
	return "VOL " + volume
}

func wifi() string {
	// get wifi
	cmdstr := "nmcli -t d | grep wlan0 |  grep :connected: | awk -F'[:]' '{print $4}'"
	cmd := exec.Command("sh", "-c", cmdstr)
	out, err := cmd.Output()
	if err != nil || len(out) < 1 {
		return "<span color='red'>W: down</span>"
	}
	wifi := strings.TrimSpace(string(out))

	cmdstr = fmt.Sprintf("nmcli d w | grep %v | awk '{print $8}'", wifi)
	cmd = exec.Command("sh", "-c", cmdstr)
	out, _ = cmd.Output()
	signal := strings.TrimSpace(string(out))

	return fmt.Sprintf("<span color='green'>W: %v %v%%</span>", wifi, signal)
}

func ethernet() string {
	// get ethernet
	cmdstr := "nmcli -t d | grep ethernet | grep :connected:"
	cmd := exec.Command("sh", "-c", cmdstr)
	_, err := cmd.Output()
	if err != nil {
		return "<span color='red'>E: down</span>"
	}
	return "<span color='green'>E: up</span>"
}

func main() {
	currentTime := currentTime()
	status, capacity := battery()
	volume := volume()
	wifi := wifi()
	ethernet := ethernet()

	fmt.Printf("%v | %v | %v |  %v %v%% | %v", wifi, ethernet, volume, status, capacity, currentTime)
}
