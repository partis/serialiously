package main

import (
	"fmt"
	"os"
	"flag"
	"bufio"
	"github.com/golang/glog"
	"github.com/tarm/serial"
)

func usage() {
  //print usage to iniate logging
  fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n", )
  flag.PrintDefaults()
  os.Exit(2)
}

func init() {
  //set the usage to the above func
  flag.Usage = usage
  //parse the flags from the command line to configre logging
  flag.Parse()
}

func main() {
	glog.Info("Reading in config")

	config := ReadConfig("serialiously.cfg")
	fmt.Println("Config loaded: " + toJson(config))
	
	glog.Info("Reading in serial commands")
	commands := ReadCommands(config.CommandFile)
	glog.V(1).Info("Commands loaded")

	glog.Flush()

	c := &serial.Config{Name: config.ComPort, Baud: 115200}
        s, err := serial.OpenPort(c)
        if err != nil {
                glog.Fatal(err)
        }

	buf := make([]byte, 128)
        n, err := s.Read(buf)
        if err != nil {
	        glog.Fatal(err)
        }
	glog.Info("Read this from the port:")
        glog.Info(string(buf[:n]))

	for _,command := range commands {
		glog.Info("Running command: " + string(command))
	        n, err = s.Write([]byte(string(command)))
        	if err != nil {
                	glog.Fatal(err)
        	}

        	buf := make([]byte, 128)
        	n, err = s.Read(buf)
        	if err != nil {
                	glog.Fatal(err)
       		}
		glog.Info("Read this from the port:")
        	glog.Info(string(buf[:n]))

		glog.Flush()
	}
}

func ReadCommands(filename string) (commands []string) {
	file, err := os.Open(filename)
	if err != nil {
        	glog.Fatal(err)
    	}
    	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		commands = append(commands, scanner.Text())
    	}

    	if err := scanner.Err(); err != nil {
        	glog.Fatal(err)
    	}
	return
}
