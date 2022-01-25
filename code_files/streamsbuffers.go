package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

type MyStringStruct struct {
	str       string
	readIndex int
}

type SampleStore struct {
	data []byte
}

func (ss *SampleStore) Write(p []byte) (n int, err error) {
	if len(ss.data) == 10 {
		return 0, nil
	}

	//get remaining cap
	remainingCap := 10 - len(ss.data)
	writeLen := len(p)

	if remainingCap <= writeLen {
		writeLen = remainingCap
		err = io.EOF
	}

	ss.data = append(ss.data, p[:writeLen]...)
	n = writeLen
	return
}

func (myStringData *MyStringStruct) Read(p []byte) (n int, err error) {
	strBytes := []byte(myStringData.str)

	if myStringData.readIndex >= len(strBytes) {
		return 0, io.EOF
	}

	nextReadLimit := myStringData.readIndex + len(p)

	if nextReadLimit >= len(strBytes) {
		nextReadLimit = len(strBytes)
		err = io.EOF
	}

	nextBytes := strBytes[myStringData.readIndex:nextReadLimit]
	n = len(nextBytes)

	copy(p, nextBytes)
	myStringData.readIndex = nextReadLimit
	return
}

func main() {
	// src := MyStringStruct{
	// 	str: "hello world",
	// }

	// p := make([]byte, 3)

	// for {
	// 	n, err := src.Read(p)
	// 	fmt.Printf("%d bytes read data from %s \n", n, p[:n])

	// 	if err == io.EOF {
	// 		fmt.Println("end of file occured")
	// 		break
	// 	} else if err != nil {
	// 		fmt.Println("Ooops something went wrong!", err)
	// 		break
	// 	}
	// }

	//we can also use stirng.NewReader for a stream
	// src := strings.NewReader("hello world")
	// p := make([]byte, 3)

	// for {
	// 	n, err := src.Read(p)
	// 	fmt.Printf("%d bytes read data from %s \n", n, p[:n])

	// 	if err == io.EOF {
	// 		fmt.Println("end of file occured")
	// 		break
	// 	} else if err != nil {
	// 		fmt.Println("Ooops something went wrong!", err)
	// 		break
	// 	}
	// }

	//when you read from a stream it just removes every element after reading its like pushing data out of stack
	//reading everything from the source
	// src := strings.NewReader("hello world")
	// data, _ := ioutil.ReadAll(src)
	// fmt.Printf("read data from stream %s\n", data)

	//reading upto some buffer size
	// src := strings.NewReader("hello amazing world!")
	// buff := make([]byte, 14)
	// bytesRead1, err1 := io.ReadFull(src, buff)
	// fmt.Printf("%d, %s, %v\n", bytesRead1, buff[:bytesRead1], err1)

	// bytesRead2, err2 := io.ReadFull(src, buff)
	// fmt.Printf("%d, %s, %v\n", bytesRead1, buff[:bytesRead2], err2)

	// bytesRead3, err3 := io.ReadFull(src, buff)
	// fmt.Printf("%d, %s, %v\n", bytesRead1, buff[:bytesRead3], err3)

	//limiting reader from 20 bytes to 20
	// mainSrc := strings.NewReader("Hello Amazing world!")
	// src := io.LimitReader(mainSrc, 10)
	// p := make([]byte, 3)

	// for {
	// 	n, err := src.Read(p)
	// 	fmt.Printf("%d, %s\n", n, p[:n])

	// 	if err == io.EOF {
	// 		fmt.Println("end of file occured")
	// 		break
	// 	} else if err != nil {
	// 		fmt.Println("oops some error occured", err)
	// 		break
	// 	}
	// }

	// ss := SampleStore{}
	// bytesWritten1, err1 := ss.Write([]byte("hello world"))
	// fmt.Printf("%d, %v", bytesWritten1, err1)
	// fmt.Printf("%s", ss)
	// bytesWritten2, err2 := ss.Write([]byte("hello world"))
	// fmt.Printf("%d, %v", bytesWritten2, err2)
	// fmt.Printf("%s", ss)
	// bytesWritten3, err4 := ss.Write([]byte("hello world"))
	// fmt.Printf("%d, %v", bytesWritten3, err4)
	// fmt.Printf("%s", ss)

	// ss := &SampleStore{}
	// bytesWritten1, err1 := io.WriteString(ss, "hello world")
	// fmt.Printf("%d, %v", bytesWritten1, err1)
	// fmt.Printf("%s", ss)
	// bytesWritten2, err2 := io.WriteString(ss, "hello world")
	// fmt.Printf("%d, %v", bytesWritten2, err2)
	// fmt.Printf("%s", ss)
	// bytesWritten3, err4 := io.WriteString(ss, "hello world")
	// fmt.Printf("%d, %v", bytesWritten3, err4)
	// fmt.Printf("%s", ss)

	// //standard io streams
	// // use `io.WriteString` to write to a `io.Writer`
	// io.WriteString(os.Stdout, "Hello World!\n")
	// // call `Write` method of of a `io.Writer`
	// os.Stdout.Write([]byte("Hello World!\n"))
	// // use `fmt` package function to write to a `io.Writer`
	// fmt.Fprint(os.Stdout, "Hello World!\n")
	// fmt.Fprintln(os.Stdout, "Hello World!") // adds new line
	// fmt.Fprintf(os.Stdout, "%s, World!\n", "Hello")

	//copying data from 1 stream to another
	// stringReader := strings.NewReader("hello world")
	// io.CopyN(os.Stdout, stringReader, 5)
	// io.Copy(os.Stdout, stringReader)

	//pipe is a synchronous process to transfer data and work on runtime
	// src, dst := io.Pipe()

	// go func() {
	// 	dst.Write([]byte("hello line 1"))
	// 	dst.Write([]byte("hello line 1"))
	// 	dst.Close()
	// }()

	// packet := make([]byte, 6)
	// _, err := src.Read(packet)
	// fmt.Printf("error is %v, value is %s\n", err, packet)

	//buffered stream
	buf := bytes.NewBufferString("Hello World!")
	// write some data to the buffer
	fmt.Print("bytes written => ")
	fmt.Println(buf.WriteString("How are you?"))

	// append data from a `io.Reader` to the buffer
	strReader := strings.NewReader(" Doing Well? ")
	fmt.Print("bytes written => ")
	fmt.Println(buf.ReadFrom(strReader))

	// read first `12 bytes` from the buffer
	fmt.Print("bytes read => ")
	fmt.Println(buf.Read(make([]byte, 12)))

	// read all `unread bytes` to STDOUT
	fmt.Print("bytes read => ")
	fmt.Println(buf.WriteTo(os.Stdout))

	// read another `10 bytes` from the buffer
	fmt.Print("bytes read => ")
	fmt.Println(buf.Read(make([]byte, 10))) // EOF

	// write some more data to the buffer
	fmt.Print("bytes written => ")
	fmt.Println(buf.WriteString("Hello!"))

	// read another `10 bytes` from the buffer
	fmt.Print("bytes read => ")
	fmt.Println(buf.Read(make([]byte, 10)))

}
