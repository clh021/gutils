package envutil

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetProcessOwner() string {
	stdout, err := exec.Command("ps", "-o", "user=", "-p", strconv.Itoa(os.Getpid())).Output()
	if err != nil {
		log.Println(err)
	}
	return strings.TrimSuffix(string(stdout), "\n")
}
