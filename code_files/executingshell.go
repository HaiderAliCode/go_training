package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// goFuncMain, err := exec.LookPath("go")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(goFuncMain)

	//to execute a command we need to create a cmd struct

	goExecutable, _ := exec.LookPath("go")
	cmdGoVer := &exec.Cmd{
		Path:   goExecutable,
		Args:   []string{goExecutable, "version"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	//just printing the command
	// fmt.Println(cmdGoVer.String())

	if err := cmdGoVer.Run(); err != nil {
		fmt.Println(err)
	}

	//other way to run cmd command is
	cmdGoVer.Start() // this command runs on background
	cmdGoVer.Wait()  // this command waits for the above command to finish

	//we will mostly use cmd.Command function
	cmd := exec.Command("go", "version")
	// cmd.Stderr = os.Stdout
	// cmd.Stdout = os.Stdout

	// if err := cmd.Run(); err != nil {
	// 	fmt.Println("Error", err)
	// }

	//or we can also get the output from it
	if output, err := cmd.Output(); err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Printf("got following output %s\n", output)
	}
}
