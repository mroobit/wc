# wc
<p>
  <a href="https://golang.org/doc/go1.20">
    <img alt="Go" src="https://img.shields.io/github/go-mod/go-version/mroobit/untitled-sidescroller?color=00ADD8&style=flat"
  </a> 
</p>
    
*A Go implementation of the command line tool `wc`, which prints newline, word, and byte counts for each file*

From inside the directory, run:
```
$ go build
$ ./ccwc --help
```
Alternately, you can view the man page:
```
$ man ./ccwc-man.0.gz
```

## Options
`-c`, `--bytes` print the byte count

`-m`, `--chars` print the character count

`-l`, `--lines` print the line count

`-w`, `--words` print the word count

`--help` display options

`--version` display version
