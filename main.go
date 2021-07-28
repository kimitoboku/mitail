package main

import (
    "flag"
    "fmt"
	"strings"
	"os/exec"
	"os"
	"time"
)

func main(){
	interval := flag.Int("n", 1000, "milliseconds to wait between updates")
	outPutError := flag.Bool("e", false, "output stderr")
	withTimestamp := flag.Bool("t", false, "output with timestap")
	flag.Parse()
	args := flag.Args()
	flags := []string{"-c"}
	cmd := append(flags, args...)
	for {
		ps, err := exec.Command("sh", cmd...).CombinedOutput()
		output := strings.TrimRight(string(ps), "\n")
		if *withTimestamp{
			fmt.Printf("%s\t", time.Now())
		}
		if *outPutError {
			fmt.Printf("%s\n", output)
			fmt.Fprintf(os.Stderr, "%s", err)
		}else{
			fmt.Printf("%s\n", output)
		}

		time.Sleep(time.Millisecond * time.Duration(*interval))
	}
}