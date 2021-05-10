package main

import (
	"fmt"
	"os/exec"

	"github.com/cupsprintmonitor/logging"
)

func main() {
	logger := logging.New()
	if logger == nil {
		fmt.Println("unable to start logger")

		return
	}

	out, err := exec.Command("/usr/bin/lpstat", "-p").CombinedOutput()
	if err != nil {
		logger.Log.Error().Msgf("an error occurred executing /usr/bin/lpstat -p: %s %s", err.Error(), string(out))

		return
	}

	logger.Log.Info().Msgf("%s", string(out))
}
