package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	command string
	options = struct {
		Bytes   bool
		Chars   bool
		Words   bool
		Lines   bool
		Help    bool
		Version bool
	}{
		Bytes:   false,
		Chars:   false,
		Words:   false,
		Lines:   false,
		Help:    false,
		Version: false,
	}

	optionFlags = map[string]string{
		"-c":        "Bytes",
		"--bytes":   "Bytes",
		"-m":        "Chars",
		"--chars":   "Chars",
		"-l":        "Lines",
		"--lines":   "Lines",
		"-w":        "Words",
		"--words":   "Words",
		"-h":        "Help",
		"--help":    "Help",
		"--version": "Version",
	}
)

func main() {
	command = os.Args[0]

	flags := make([]string, 0)
	files := make([]*File, 0)

	for i := range os.Args {
		switch {
		case os.Args[i][0] == '-':
			if _, ok := optionFlags[os.Args[i]]; !ok {
				// separate unrecognized flag into individual character flags
				moreOpts := strings.Split(strings.TrimLeft(os.Args[i], "-"), "")
				for _, o := range moreOpts {
					o = "-" + o
					flags = append(flags, o)
				}
			} else {
				flags = append(flags, os.Args[i])
			}
		case i > 0:
			nf := NewFile(os.Args[i])
			files = append(files, nf)
		}
	}

	for _, flag := range flags {
		if _, ok := optionFlags[flag]; !ok {
			// unrecognized flag, display help
			options.Help = true
		} else {
			switch {
			case optionFlags[flag] == "Bytes":
				options.Bytes = true
			case optionFlags[flag] == "Chars":
				options.Chars = true
			case optionFlags[flag] == "Words":
				options.Words = true
			case optionFlags[flag] == "Lines":
				options.Lines = true
			case optionFlags[flag] == "Help":
				options.Help = true
			case optionFlags[flag] == "Version":
				options.Version = true
			}
		}
	}

	if len(flags) == 0 {
		options.Bytes = true
		options.Lines = true
		options.Words = true
	}

	if options.Help {
		Help(command)
		return
	}

	if options.Version {
		Version(command)
		return
	}

	total := NewFile("Total")

	for _, f := range files {

		byteStream, err := os.ReadFile(f.Name)
		if err != nil {
			fmt.Println("Error reading in whole file: ", err)
			return
		}

		f.Bytes = len(byteStream)

		strFile := fmt.Sprintf("%s", byteStream)

		f.Chars = len([]rune(strFile))
		f.Words = len(strings.Fields(strFile))
		f.Lines = len(strings.Split(strFile, "\n")) - 1

		if options.Lines {
			fmt.Printf(" %d", f.Lines)
		}
		if options.Words {
			fmt.Printf(" %d", f.Words)
		}
		if options.Bytes {
			fmt.Printf(" %d", f.Bytes)
		}
		if options.Chars {
			fmt.Printf(" %d", f.Chars)
		}
		fmt.Printf("\t%s\n", f.Name)

		total.Bytes += f.Bytes
		total.Chars += f.Chars
		total.Lines += f.Lines
		total.Words += f.Words
	}

	if len(files) > 1 {
		if options.Lines {
			fmt.Printf(" %d", total.Lines)
		}
		if options.Words {
			fmt.Printf(" %d", total.Words)
		}
		if options.Bytes {
			fmt.Printf(" %d", total.Bytes)
		}
		if options.Chars {
			fmt.Printf(" %d", total.Chars)
		}
		fmt.Printf("\t%s\n", total.Name)
	}

}

type File struct {
	Name  string
	Bytes int
	Words int
	Lines int
	Chars int
}

func NewFile(f string) *File {
	file := &File{
		Name: f,
	}
	return file
}

func Help(c string) {
	byteStream, err := os.ReadFile("./help.txt")
	if err != nil {
		fmt.Println("Help file not found: ", err)
		return
	}

	strFile := fmt.Sprintf("%s", byteStream)
	strFile = strings.Replace(strFile, "commandName", c, -1)

	fmt.Println(strFile)
	return
}

func Version(c string) {
	c = strings.TrimLeft(c, "./")
	byteStream, err := os.ReadFile("./version.txt")
	if err != nil {
		fmt.Println("Version information not found: ", err)
		return
	}

	strFile := fmt.Sprintf("%s", byteStream)
	strFile = strings.Replace(strFile, "commandName", c, -1)

	fmt.Println(strFile)
	return
}
