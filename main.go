package main

import (
	"fmt"
	"os"

	"github.com/rikez/ipset-collector/ipset"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	// fm, err := ipset.TraverseIPSetDir("/tmp/update-iplist")
	// if err != nil {
	// 	panic(err)
	// }

	// ipset.ReadIPSetFiles(fm)

	strs, err := ipset.CIDRParser("192.188.0.1/50")
	if err != nil {
		panic(err)
	}

	for _, str := range strs {
		fmt.Printf(">>> %s", str)
	}
}
