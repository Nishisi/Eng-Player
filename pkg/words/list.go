package words

import (
	"bufio"
	"os"
)

// WordList用のインターフェース
type List interface {
	GetList() ([]wordInfo, error)
	LoadWordList(s string) error
	SetupMeanAndFile() error
}

// 学習する単語のリスト
type WordList struct {
	lists []wordInfo
}

// 単語の情報
type wordInfo struct {
	word string
	mean string
	file string
}

// 読み込み
// file から学習する単語を読み込む
func (wl *WordList) LoadWordList(filename string) error {
	var lists []string

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lists = append(lists, sc.Text())
	}

	// 検証
	for i := 0; i < len(lists); i++ {
		wi := wordInfo{lists[i], "", ""}
		wl.lists = append(wl.lists, wi)
	}

	return nil
}

func (wl *WordList) SetupMeanAndFile() error {
	return nil
}

func (wl *WordList) GetList() ([]wordInfo, error) {
	return wl.lists, nil
}
