package grep

import (
	"fmt"
	"io"
)

func (g *Grep) processFlagB(curLine string) []string {
	previousLines := []string{}

	cPos, _ := g.file.Seek(0, io.SeekCurrent)
	if cPos == 0 {
		return nil
	}

	curLineCounter := 0
	for i := 0; i < g.flags.FlagB+1; i++ {
		var prevLine string
		var pos int64

		for {
			var err error
			pos, err = g.file.Seek(-2, io.SeekCurrent)
			if err != nil {
				_, err := g.file.Seek(0, io.SeekStart)
				if err != nil {
					continue
				}

				prevLine, _ = g.readLine()
				break
			}

			buf := make([]byte, 1)
			_, _ = g.file.Read(buf)

			if buf[0] == '\n' {
				pos, _ := g.file.Seek(0, io.SeekCurrent)

				prevLine, _ = g.readLine()
				if prevLine == curLine {
					curLineCounter++
				}

				_, err := g.file.Seek(pos, io.SeekStart)
				if err != nil {
					return nil
				}

				break
			}
		}

		if g.flags.FlagV != g.match(prevLine) && curLine != prevLine {
			break
		}

		if curLine == prevLine && curLineCounter > 1 {
			break
		}

		if curLine != prevLine {
			previousLines = append(previousLines, prevLine)
		}

		if pos == 0 {
			break
		}

	}

	g.file.Seek(cPos, io.SeekStart)
	return previousLines
}

func (g *Grep) printFlagB(n int, curLine string) {
	previousLines := g.processFlagB(curLine)

	n = n - len(previousLines)
	for i := len(previousLines) - 1; i >= 0; i-- {
		if _, ok := g.printedNexts[n]; !ok {
			if g.flags.FlagN {
				fmt.Printf("%d-", n)
			}
			fmt.Print(previousLines[i])
		}
		n++
	}
}
