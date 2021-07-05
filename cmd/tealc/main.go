// Copyright (C) 2019 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/algorand/go-algorand/data/transactions/logic"
)

const (
	stdinFilename = "-"
)

var fname = flag.String("input", stdinFilename, "input file to assemble (if unspecified, use stdin)")
var disassemble = flag.Bool("disassemble", false, "set to disassemble instead of assemble")
var disassembleShort = flag.Bool("D", false, "set to disassemble instead of assemble (alias for --disassemble)")

func readFile(filename string) ([]byte, error) {
	if filename == stdinFilename {
		return ioutil.ReadAll(os.Stdin)
	}
	return ioutil.ReadFile(filename)
}

func panicf(fmtstr string, a ...interface{}) {
	s := fmt.Sprintf(fmtstr, a...)
	panic(s)
}

func assembleFile(fname string) (program []byte) {
	text, err := readFile(fname)
	if err != nil {
		panicf("%s: %s\n", fname, err)
	}
	program, err = logic.AssembleString(string(text))
	if err != nil {
		panicf("%s: %s\n", fname, err)
	}
	return program
}

func disassembleFile(fname string) string {
	program, err := readFile(fname)
	if err != nil {
		panicf("%s: %s\n", fname, err)
	}
	text, err := logic.Disassemble(program)
	if err != nil {
		panicf("%s: %s\n", fname, err)
	}
	return text
}

func main() {
	flag.Parse()

	if *disassemble || *disassembleShort {
		os.Stdout.Write([]byte(disassembleFile(*fname)))
		return
	}
	os.Stdout.Write(assembleFile(*fname))
}
