package words

import "os/exec"

func Say(sw string) {
	exec.Command("say", sw).Run()
}
