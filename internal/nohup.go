package internal

import "os/exec"

func Nohup(cmd string, args ...string) error {
	_, err := exec.LookPath(cmd)
	if err != nil {
		return err
	}

	c := exec.Command("nohup", append([]string{cmd}, args...)...)

	return c.Start()
}
