package main

import (
	"github.com/ActiveState/tail"
	"net"
	"flag"
//	"bufio"
	"fmt"
)

type Sawdust struct {
	Filepath string
	Ship bool
	Host string
}


func (sd *Sawdust) Shipline(line string) {
	fmt.Println("shipping")
  //p :=  make([]byte, 2048)
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Error %v", err)
		return 
	}

	fmt.Fprintf(conn, line)
	/*_, err = bufio.NewReader(conn).Read(p)
  if err == nil {
      fmt.Printf("%s\n", p)
  } else {
      fmt.Printf("Error %v\n", err)
  }
  */
  conn.Close()
  fmt.Println("shipped")
}

func (sd *Sawdust) TailFile() {
	fmt.Println("in Tail")

	t, err := tail.TailFile(sd.Filepath, tail.Config{
	    Follow: true,
	    ReOpen: true})
		for line := range t.Lines {
		    fmt.Println(line.Text)
		    // do build buffer and filter for multiline
		    if sd.Ship {
		        sd.Shipline(line.Text)
		    }
		}

		if err != nil {
		    fmt.Println(err)
		}
}


func RunSawdust(filepath string, ship bool, host string) {

	sd := new(Sawdust)
	sd.Filepath = filepath
	sd.Ship = ship
	sd.TailFile()

}

func main() {

  var filename = flag.String("file", "test.log", "name of file to parse")	
  var ship = flag.Bool("ship", false, "ship the logs",)
  var host = flag.String("host", "127.0.0.1:5000", "host to ship the logs to")	

  flag.Parse()

	RunSawdust(*filename,*ship,*host)
}
