package main
import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/pedroalbanese/whirlpool"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

	var check = flag.String("c", "", "Check hashsum file.")
	var recursive = flag.Bool("r", false, "Process directories recursively.")
	var target = flag.String("t", "", "Target file/wildcard to generate hashsum list.")
	var verbose = flag.Bool("v", false, "Verbose mode. (The exit code is always 0 in this mode)")

func main() {
    flag.Parse()

        if (len(os.Args) < 2) {
	fmt.Println("Whirlpool Hashsum Tool - ALBANESE Lab (c) 2020-2021\n")
	fmt.Println("Usage of",os.Args[0]+":")
        fmt.Printf("%s [-v] [-c <hash.ext>] [-r] -t <file.ext>\n\n", os.Args[0])
        flag.PrintDefaults()
        os.Exit(1)
        }


        if *target != "" && *recursive == false {
	files, err := filepath.Glob(*target)
	if err != nil {
	    log.Fatal(err)
	}

	for _, match := range files {
	h := whirlpool.New()
        f, err := os.Open(match)
        if err != nil {
            log.Fatal(err)
        }
        if _, err := io.Copy(h, f); err != nil {
            log.Fatal(err)
        }
    	fmt.Println(hex.EncodeToString(h.Sum(nil)), "*" + f.Name())
	}
	}


        if *target != "" && *recursive == true {
		err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			file, err := os.Stat(path)
			if file.IsDir() {
			} else {
			h := whirlpool.New()
			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			if _, err := io.Copy(h, f); err != nil {
				log.Fatal(err)
			}
			fmt.Println(hex.EncodeToString(h.Sum(nil)), "*" + f.Name())
			}
			return nil
		})
		if err != nil {
			log.Println(err)
		}
	}


        if *check != "" {
	file, err := os.Open(*check)
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
	file.Close()
 

	for _, eachline := range txtlines {
	lines := strings.Split(string(eachline), " *")

	if strings.Contains(string(eachline), " *") {

		h := whirlpool.New()

		_, err := os.Stat(lines[1])
		if err == nil {
			f, err := os.Open(lines[1])
			if err != nil {
			     log.Fatal(err)
			}
			io.Copy(h, f)
			
			if *verbose {
				if hex.EncodeToString(h.Sum(nil)) == lines[0] {
					fmt.Println(lines[1] + "\t", "OK")
				} else {
					fmt.Println(lines[1] + "\t", "FAILED")
				}
			} else {
				if hex.EncodeToString(h.Sum(nil)) == lines[0] {
				} else {
					os.Exit(1)
				}
			}
		} else {
			if *verbose {
				fmt.Println(lines[1] + "\t", "Not found!")
			} else {
				os.Exit(1)	
			}	
		}
		}
		}
	}
}
