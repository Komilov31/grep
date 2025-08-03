package grep

import (
	"io"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/Komilov31/grep/internal/flags"
)

func initFile(flags *flags.Flags) *os.File {
	file := os.Stdin

	if flags.FileName != "" {
		var err error
		file, err = os.Open(flags.FileName)
		if err != nil {
			log.Fatal(err)
		}
	}

	return file
}

func initRegexp(flags *flags.Flags) *regexp.Regexp {
	reg, err := regexp.Compile(flags.Pattern)
	if err != nil {
		log.Fatal(err)
	}

	return reg
}

func (g *Grep) readLine() (string, error) {
	line := strings.Builder{}

	for {
		buf := make([]byte, 1)

		_, err := g.file.Read(buf)
		if err != nil && err != io.EOF {
			return "", err
		}

		if buf[0] != 0 {
			line.WriteByte(buf[0])
		}

		if buf[0] == '\n' {
			break
		}

		if err == io.EOF {
			if line.Len() != 0 {
				line.WriteByte('\n')
			}
			return line.String(), err
		}
	}

	return line.String(), nil
}

func (g *Grep) countMatches() int {
	count := 0

	for {
		nextLine, err := g.readLine()
		if err != nil && err != io.EOF {
			log.Fatal("could not read nextLine from file")
		}
		if g.flags.FlagV != g.match(nextLine) {
			if err == io.EOF && nextLine == "" {
				break
			}
			count++
		}

		if err == io.EOF {
			break
		}
	}

	return count
}

func (g *Grep) match(nextLine string) bool {
	if g.flags.FlagI {
		nextLine = strings.ToLower(nextLine)
		g.flags.Pattern = strings.ToLower(g.flags.Pattern)
	}

	if g.flags.FlagF {
		return strings.Contains(nextLine, g.flags.Pattern)
	}

	matched := g.regexp.MatchString(nextLine)
	return matched
}
