package grep

import (
	"fmt"
	"io"
)

func (g *Grep) processFlagA() ([]string, error) {

	cPos, err := g.file.Seek(0, io.SeekCurrent)
	if err != nil {
		return nil, err
	}
	nextLines := make([]string, 0)

	for i := 0; i < g.flags.FlagA; i++ {
		nextLine, err := g.readLine()
		if err != nil && err != io.EOF {
			break
		}

		if g.flags.FlagV != g.match(nextLine) {
			break
		}
		nextLines = append(nextLines, nextLine)

		if err == io.EOF {
			return nextLines, err
		}
	}

	g.file.Seek(cPos, io.SeekStart)

	return nextLines, nil
}

func (g *Grep) printFlagA(n int) {
	nextLines, _ := g.processFlagA()

	n++
	for i := 0; i < len(nextLines); i++ {
		g.printedNexts[n] = struct{}{}

		if g.flags.FlagN {
			fmt.Printf("%d-", n)
		}

		fmt.Print(nextLines[i])
		n++
	}
}
