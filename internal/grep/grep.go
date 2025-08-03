package grep

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/Komilov31/grep/internal/flags"
)

type Grep struct {
	flags        *flags.Flags
	file         *os.File
	regexp       *regexp.Regexp
	printedNexts map[int]struct{}
}

func New(flags *flags.Flags) *Grep {
	g := Grep{}

	file := initFile(flags)
	reg := initRegexp(flags)

	g.regexp = reg
	g.file = file
	g.flags = flags
	g.printedNexts = make(map[int]struct{})

	return &g
}

func (g *Grep) ProcessProgram() error {
	defer func() {
		err := g.file.Close()
		if err != nil {
			log.Fatal("could not close file: ", err)
		}
	}()

	if g.flags.Flagc {
		fmt.Println(g.countMatches())
		return nil
	}

	n := 1
	for {
		curLine, err := g.readLine()
		if err != nil && err != io.EOF {
			return err
		}

		if g.match(curLine) != g.flags.FlagV {

			if g.flags.FlagB != 0 || g.flags.FlagC != 0 {
				g.printFlagB(n, curLine)
			}

			if g.flags.FlagN {
				fmt.Printf("%d:", n)
			}
			fmt.Print(curLine)

			if g.flags.FlagA != 0 || g.flags.FlagC != 0 {
				g.printFlagA(n)
			}
		}

		if err == io.EOF {
			break
		}
		n++
	}

	return nil
}
