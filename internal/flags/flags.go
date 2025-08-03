package flags

import (
	"github.com/pborman/getopt/v2"
)

type Flags struct {
	FlagA    int
	FlagB    int
	FlagC    int
	Flagc    bool
	FlagI    bool
	FlagV    bool
	FlagF    bool
	FlagN    bool
	FileName string
	Pattern  string
}

func Parse() *Flags {
	flagA := getopt.Int('A', 0, "print NUM lines of trailing context")
	flagB := getopt.Int('B', 0, "print NUM lines of leading context")
	flagC := getopt.Int('C', 0, "print NUM lines of output context")
	flagc := getopt.Bool('c', "print only a count of selected lines")
	flagI := getopt.Bool('i', "ignore case distinctions in patterns and data")
	flagV := getopt.Bool('v', "select non-matching lines")
	flagF := getopt.Bool('F', "PATTERNS are strings")
	flagN := getopt.Bool('n', "print line number with output lines")
	getopt.Parse()

	flags := Flags{
		FlagA: *flagA,
		FlagB: *flagB,
		FlagC: *flagC,
		Flagc: *flagc,
		FlagI: *flagI,
		FlagV: *flagV,
		FlagF: *flagF,
		FlagN: *flagN,
	}

	if flags.FlagC != 0 && flags.FlagA == 0 {
		flags.FlagA = flags.FlagC
	}

	if flags.FlagC != 0 && flags.FlagB == 0 {
		flags.FlagB = flags.FlagC
	}

	args := getopt.Args()
	if len(args) >= 1 {
		flags.Pattern = args[0]
	}

	if len(args) >= 2 {
		flags.FileName = args[1]
	}

	return &flags
}
