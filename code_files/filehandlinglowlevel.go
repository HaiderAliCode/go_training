package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//opening missing file
	// _, err := os.Open("hello_world.txt")
	// isMissing := os.IsNotExist(err)

	// fmt.Println("isMissing", isMissing)

	//working with files
	file, err := os.OpenFile("info.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0744)
	if err != nil {
		fmt.Println("error opening file")
	}
	fmt.Println("Done working with file", file)

	//different methods of writing into a file
	// n, err := file.Write([]byte("Hello World!"))
	// n, err := file.WriteString("Hello World!")
	// n, err := io.WriteString(file, "Hello World!")

	//reading all data from file
	data, _ := ioutil.ReadAll(file)
	fmt.Printf("%s\n", data)

	//changing file modes
	//getting admin group number

	// adminGroup, _ := user.LookupGroup("admin")
	// adminGroupID, _ := strconv.ParseInt(adminGroup.Gid, 10, 64)
	// fmt.Println("group id =>", adminGroupID)

	// //getting user struct from users
	// myuser, _ := user.Lookup("haider ali")
	// myuserID, _ := strconv.ParseInt(myuser.Uid, 10, 64)

	// file.Chown(int(myuserID), int(adminGroupID))
	// file.Chmod(0755)

	//after using the file we must close it
	// file.Close()

	//other important file methods
	//file.Fd() // file descriptor
	//file.Name() //file name
	//file.Truncate() // truncates the content of file
	// file.Sync() //commits any changes make to file to file system, flushes any data saved in file into file system
	/*don’t need to call this method since any data written to the file is synched to the file-system in real-time. But for added safety, you can defer file.Sync() call after the defer file.Close() call.*/

	//file info object
	// type FileInfo interface {
	// 	Name() string       // base name of the file
	// 	Size() int64        // file length in bytes
	// 	Mode() FileMode     // file mode bits
	// 	ModTime() time.Time // modification time
	// 	IsDir() bool        // abbreviation for Mode().IsDir()
	// 	Sys() interface{}   // underlying data source (can return nil)
	// }

	//can get fileinfo object by either calling state(name string or lstate(name string)

	// fileinfo.Mode gives us fileMode object
	//working with directories

	//openining nested directory
	dir, _ := os.Open("./")

	//read on zero index
	content, _ := dir.ReadDir(0)

	fmt.Println(content)

	for _, fileInfo := range content {
		fmt.Printf("Name: %s\n", fileInfo.Name())
	}

	//similarly renaming
	// os.Rename("old path", "newpath")

	//removing file
	// os.Remove("file path")

	// os.RemoveAll("directory path")
}
