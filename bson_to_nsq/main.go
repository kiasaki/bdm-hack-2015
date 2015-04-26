package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/bitly/go-nsq"
	"github.com/kiasaki/batbelt/bson"
	"github.com/kiasaki/bdm-hack-2015/data"
	"github.com/kiasaki/bdm-hack-2015/pool"
)

var fileLocation = flag.String("file", "", "location of bson dump to import")
var nsqTopic = flag.String("nsq-topic", "", "nsq topic to produce to")
var nsqHosts = flag.String("nsq-hosts", "127.0.0.1:4060", "nsqd hosts to produce to, comma separated")

func acquireFileHandle(location string) *os.File {
	if location == "" {
		fmt.Println("File to import location is required")
		os.Exit(1)
	} else {
		fmt.Println("Importing: " + location)
	}

	fileHandle, err := os.Open(location)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}

	return fileHandle
}

func clearTypeTable(dbSession *mgo.Session, importType string) {
	var collection string
	if importType == "user" {
		collection = "users"
	} else if importType == "business" {
		collection = "businesses"
	} else if importType == "review" {
		collection = "reviews"
	} else {
		fmt.Println("Import type didn't match user, business or review")
		os.Exit(1)
	}
	// Empty db
	if _, err := dbSession.DB("").C(collection).RemoveAll(nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleFatalError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	dbSession := dialMongo(*dbUrl)
	reader := acquireFileHandle(*fileLocation)
	defer reader.Close()
	bsonStream := bsonutil.NewBSONStream(reader)
	clearTypeTable(dbSession, *importType)

	mu := sync.Mutex{}
	nsqdHostCount := 0
	nsqdHosts := strings.Split(*nsqdHosts, ",")
	factory := func() (net.Conn, error) {
		host := nsqdHosts[nsqdHostCount%len(nsqdHosts)]
		mu.Lock()
		nsqdHostCount += 1
		mu.Unlock()
		return nsq.NewProducer(host, nsq.NewConfig())
	}
	nsqPool, err := pool.NewChannelPool(5, 30, factory)
	defer nsqPool.Close()

	i := 0
	fan := make(chan bool)
	errChannel := make(chan error)

	lineRequestChan := make(chan bool)
	lineFeedChan := make(chan []byte)

	// document reader
	go func() {
		docBytes := make([]byte, bson.MaxBSONSize)
		for {
			select {
			case <-lineRequestChan:
				if success, docSize := bsonStream.ReadNext(docBytes); !success {
					errChannel <- bsonStream.Err()
					break
				} else if docSize == 0 {
					errChannel <- io.EOF
					break
				} else {
					lineFeedChan <- docBytes[o:docSize]
				}
			}
		}
	}()

	// document parsers and publishers
	for count := 0; count < 30; count++ {
		go func() {
			for {
				var err error
				lineRequestChan <- true
				line := <-lineFeedChan

				producer, err := nsqPool.Get()
				if err != nil {
					errChannel <- err
					break
				}

				producer.Publish(*nsqTopic)

				if err = data.Save(workerSession.DB(""), model); err != nil {
					errChannel <- err
					break
				}

				fan <- true
			}
		}()
	}

	for {
		select {
		case <-fan:
			i++
			if i%1000 == 0 {
				fmt.Printf("Processed %d\n", i)
			}
		case err := <-errChannel:
			if err.Error() == "EOF" {
				fmt.Println("\nDone!")
				os.Exit(0)
				break
			} else {
				panic(err)
				break
			}
		}
	}
}
