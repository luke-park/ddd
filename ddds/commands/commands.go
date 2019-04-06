package commands

import (
	"fmt"
	"os/exec"
)

// RestartDocker does exactly what you might imagine it would do.
func RestartDocker() error {
	cmd := exec.Command("bash", "-c", "sudo docker-compose down && sudo docker-compose pull && sudo docker-compose up -d")
	output, err := cmd.CombinedOutput()
	if err == nil {
		return nil
	}

	if len(output) == 0 {
		return err
	}

	return fmt.Errorf("stderr/stdout: %v (underlying: %v)", string(output), err)
}
