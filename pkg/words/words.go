package words

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

type Game struct {
	reader   io.Reader
	writer   io.Writer
	wordList *WordList
}

func New(r io.Reader, w io.Writer, filename string) *Game {
	wl := new(WordList)
	err := wl.loadListFile(filename)
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

func (gm *Game) Play() {
	gm.displayRule()
	gm.playGame()
}

func (gm *Game) playGame() {
	bc := context.Background()
	ctx, cancel := context.WithTimeout(bc, 30*time.Second)
	defer cancel()

	ch := gm.input()

gameLoop:
	// for i := 0; i < len(gm.questions); i++ {
	for i, v := range gm.wordList.Lists {
		question := v.Word
		fmt.Fprintf(gm.writer, "Number %d, %s\n", i+1, question)
		Say(question)
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

		if in == question {
			// fmt.Fprintf(gm.writer, "Correct!\n\n %q\n\n", v.Translation)
			fmt.Fprintf(gm.writer, "\n\n %q \n\n", getDictionary(question))
		} else {
			fmt.Fprintf(gm.writer, "Uncorrect..\n\n %q\n\n", v.Translation)
		}
	}
}

func (gm *Game) displayRule() {
	fmt.Fprintf(gm.writer, "Please type word displaying windows screen")
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
