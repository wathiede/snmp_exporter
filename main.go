package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/soniah/gosnmp"
)

var (
	target    = flag.String("target", "127.0.0.1", "target host to scrape")
	community = flag.String("community", "public", "SNMP community")
)

func getTable(g *gosnmp.GoSNMP, oid string) {
	results, err := g.WalkAll(oid)
	if err != nil {
		log.Fatalf("Get() err: %v", err)
	}

	for _, variable := range results {
		fmt.Printf("oid: %s ", variable.Name[len(oid)+1:])

		// the Value of each variable returned by Get() implements
		// interface{}. You could do a type switch...
		switch variable.Type {
		case gosnmp.OctetString:
			fmt.Printf("string: %s\n", string(variable.Value.([]byte)))
		default:
			// ... or often you're just interested in numeric values.
			// ToBigInt() will return the Value as a BigInt, for plugging
			// into your calculations.
			fmt.Printf("number: %d\n", gosnmp.ToBigInt(variable.Value))
		}
	}
}

func main() {
	flag.Parse()
	g := &gosnmp.GoSNMP{
		Port:      161,
		Target:    *target,
		Community: *community,
		Version:   gosnmp.Version1,
		Timeout:   time.Duration(2) * time.Second,
		Retries:   3,
		//Logger:    log.New(os.Stdout, "", 0),
	}
	err := g.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer g.Conn.Close()

	oid := ".1.3.6.1.2.1.2.2.1.10"
	getTable(g, oid)

}
