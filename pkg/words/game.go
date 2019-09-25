package words

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/dobuzora/learn-words-cl/pkg/aruku"
)

// Game struct
type Game struct {
	// Game Reader
	reader io.Reader
	// Game Writer
	writer io.Writer
	// word.List interface
	wordList List
}

// New
func New(r io.Reader, w io.Writer, filename string) *Game {
	wl := new(WordList)
	err := wl.LoadWordList(filename)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
		os.Exit(1)
	}
	return &Game{
		reader:   r,
		writer:   w,
		wordList: wl,
	}
}

// Play Game.
func (gm *Game) Do() {
	gm.displayRule()
	gm.playGame()
}

func (gm *Game) playGame() {
	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, 30*time.Second)
	defer cancel()

	ch := gm.input()

	wl, err := gm.wordList.GetList()
	if err != nil {

	}

	// aruku
	a := aruku.New("")

	// Clear
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

gameLoop:
	// for i := 0; i < len(gm.questions); i++ {
	for i, v := range wl {
		question := v.word
		fmt.Fprintf(gm.writer, "Number %d, ", i+1)
		fmt.Fprintf(gm.writer, "\x1b[31m%s\x1b[0m\n", question)
		err = Say(question)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error : %v", err)
		}
		fmt.Fprintf(gm.writer, "> ")

		var in string
		var ok bool

		select {
		case in, ok = <-ch:
			if !ok {
				break gameLoop
			}
		case <-ctx.Done():
			break gameLoop
		}

		ans, _ := a.GetMeanAndFile(question)
		if in == question {
			// fmt.Fprintf(gm.writer, "\n\n %q \n\n", getDictionary(question))
			fmt.Fprintf(gm.writer, "\n\n %s \n\n", ans)
		} else {
			// fmt.Fprintf(gm.writer, "Uncorrect..\n\n %q\n\n", v.Translation)
			fmt.Fprintf(gm.writer, "\n\n %s \n\n", ans)
		}
		// cmd.Run()
	}
}

func (gm *Game) displayRule() {
	fmt.Fprintf(gm.writer, "Please type word displaying windows screen\n")
}

func (gm *Game) input() <-chan string {
	ch := make(chan string)
	go func() {
		sc := bufio.NewScanner(gm.reader)
		for sc.Scan() {
			ch <- sc.Text()
		}
		close(ch)
	}()
	return ch
}
