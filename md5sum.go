/*
Copyright © 2020 Tim_Paik <timpaik@163.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		if err := printHelp(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	} else if os.Args[1] == "-h" || os.Args[1] == "--help" {
		if err := printHelp(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	} else if os.Args[1] == "-v" || os.Args[1] == "--version" {
		if err := printVersion(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	for i := 1; i < len(os.Args); i++ {
		if MD5, err := md5Sum(os.Args[i]); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(MD5 + " " + os.Args[i])
		}
	}
	os.Exit(0)
}

func printVersion() (err error) {
	if _, err := fmt.Println(`md5sum 0.0.1
Copyright (C) 2020 Tim_Paik <timpaik@163.com>`); err != nil {
		return err
	}
	return nil
}

func printHelp() (err error) {
	if _, err := fmt.Println(`Usage: md5sum [OPTION]... [FILE]...
Print or check MD5 (128-bit) checksums.

      --help     display this help and exit
      --version  output version information and exit

Copyright (C) 2020 Tim_Paik <timpaik@163.com>`); err != nil {
		return err
	}
	return nil
}

func md5Sum(filePath string) (MD5string string, err error) {
	checkMD5 := md5.New()

	if file, err := os.Open(filePath); err != nil {
		return "", err
	} else {
		if _, err := io.Copy(checkMD5, file); err != nil {
			return "", err
		}
	}

	MD5string = hex.EncodeToString(checkMD5.Sum(nil))
	return MD5string, nil
}
