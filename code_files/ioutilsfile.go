package main

import (
	"fmt"
	"io/ioutil"
)

const tempDir = "./"

func main() {
	// files, _ := ioutil.ReadDir(tempDir)

	// for _, file := range files {

	// 	fmt.Printf("Name %v, Size %v, Mode: %v, IsDir %v\n",
	// 		file.Name(),
	// 		file.Size(),
	// 		file.Mode(),
	// 		file.IsDir(),
	// 	)
	// }

	//glob file search
	// pdfFiles, _ := filepath.Glob("./*go")
	// for _, pdfFile := range pdfFiles {
	// 	file, _ := os.Stat(pdfFile)

	// 	fmt.Printf("Name %v, Size %v, Mode: %v, IsDir %v\n",
	// 		file.Name(),
	// 		file.Size(),
	// 		file.Mode(),
	// 		file.IsDir(),
	// 	)
	// }

	//reading a file
	// html, _ := ioutil.ReadFile("./index.html")
	// fmt.Println("raw file", html)
	// fmt.Println("string of file", string(html))

	//open file and write into file
	data := []byte("hello friend")
	path := "./filewritint.txt"
	err := ioutil.WriteFile(path, data, 0777)
	if err != nil {
		fmt.Println(err)
	}

	//can make nested folder
	//os.MkdirAll

	/*The ioutil package also provides TempDir and TempFile functions to work with temporary files which can be useful if you want to process some files before putting them on the disk where users can see and use them.*/

}
