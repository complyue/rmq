// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"tutorial"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  void ping()")
	fmt.Fprintln(os.Stderr, "  i32 add(i32 num1, i32 num2)")
	fmt.Fprintln(os.Stderr, "  i32 calculate(i32 logid, Work w)")
	fmt.Fprintln(os.Stderr, "  void zip()")
	fmt.Fprintln(os.Stderr, "  string RmqCall(string bin)")
	fmt.Fprintln(os.Stderr, "  SharedStruct getStruct(i32 key)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := tutorial.NewCalculatorClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "ping":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Ping requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Ping())
		fmt.Print("\n")
		break
	case "add":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Add requires 2 args")
			flag.Usage()
		}
		tmp0, err9 := (strconv.Atoi(flag.Arg(1)))
		if err9 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		tmp1, err10 := (strconv.Atoi(flag.Arg(2)))
		if err10 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		fmt.Print(client.Add(value0, value1))
		fmt.Print("\n")
		break
	case "calculate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "Calculate requires 2 args")
			flag.Usage()
		}
		tmp0, err11 := (strconv.Atoi(flag.Arg(1)))
		if err11 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		arg12 := flag.Arg(2)
		mbTrans13 := thrift.NewTMemoryBufferLen(len(arg12))
		defer mbTrans13.Close()
		_, err14 := mbTrans13.WriteString(arg12)
		if err14 != nil {
			Usage()
			return
		}
		factory15 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt16 := factory15.GetProtocol(mbTrans13)
		argvalue1 := tutorial.NewWork()
		err17 := argvalue1.Read(jsProt16)
		if err17 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.Calculate(value0, value1))
		fmt.Print("\n")
		break
	case "zip":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Zip requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Zip())
		fmt.Print("\n")
		break
	case "RmqCall":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "RmqCall requires 1 args")
			flag.Usage()
		}
		argvalue0 := []byte(flag.Arg(1))
		value0 := argvalue0
		fmt.Print(client.RmqCall(value0))
		fmt.Print("\n")
		break
	case "getStruct":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetStruct requires 1 args")
			flag.Usage()
		}
		tmp0, err19 := (strconv.Atoi(flag.Arg(1)))
		if err19 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetStruct(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
