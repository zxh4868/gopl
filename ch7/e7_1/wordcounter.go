package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type DocStats struct {
	words int
	lines int
}

func (d *DocStats) Write(p []byte) (int, error) {
	lines := bufio.NewScanner(bytes.NewReader(p))
	for lines.Scan() {
		d.lines++
	}

	words := bufio.NewScanner(bytes.NewReader(p))
	words.Split(bufio.ScanWords)

	for words.Scan() {
		d.words++
	}
	return len(p), nil
}

const test_text = `Matplotlib produces publication-quality figures in a variety of hardcopy formats and interactive environments across platforms.
Matplotlib can be used in Python scripts, Python/IPython shells, web application servers, and various graphical user interface toolkits.
You've discovered a bug or something else you want to change — excellent!

You've worked out a way to fix it — even better!

You want to tell us about it — best of all!

Start at the contributing guide!`

func main() {
	var d DocStats
	d.Write([]byte(test_text))
	fmt.Println(d)

	d = DocStats{}
	var name = "Goland"
	fmt.Fprintf(&d, "%s\nhtllo\nnihao\nwho are you\ni'm fine and you\ncan I speak english?\n", name)
	fmt.Println(d)

}
