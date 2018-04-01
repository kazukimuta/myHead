package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	n         = flag.Int("n", 10, "出力する行数")
	lineCount = flag.Bool("N", false, "行番号を表示する")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, `
使い方 %s:
	 %s [-n count][-N] [file]
	 file指定がない場合は、標準入力を読み取る
Options
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		fmt.Fprint(os.Stderr, "引数の数が不正です\nmyHead [-n count] [file]\n")
		os.Exit(1)
	}

	var file *os.File
	var err error
	if len(args) == 0 {
		file = os.Stdin
	} else {
		file, err = os.Open(args[0])
		defer file.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i++
		if i > *n {
			break
		}
		var printLine string
		if *lineCount {
			printLine = fmt.Sprintf("%v | %s", i, scanner.Text())
		} else {
			printLine = scanner.Text()
		}
		fmt.Fprintln(os.Stdout, printLine)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprint(os.Stderr, "ファイル読み取りエラー", err)
		os.Exit(1)
	}
}
