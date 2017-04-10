package main

import (
	"fmt"

	"./gfile"
)

func main() {
	musicP := "/Music/Dir/Path"
	flist := gfile.Getfile(musicP)
	for _, f := range flist {
		fmt.Println(f)

		//err := exec.Command("afplay", musicP+"/"+f).Run()
		//if err != nil {
		//	fmt.Printf("%v", err)
		//	os.Exit(1)
		//}
	}
}
