package words

import "os/exec"

// Say is to speak a word.
func Say(sw string) error {
	err := exec.Command("say", sw).Run()
	if err != nil {
		return err
	}
	return nil
}
