package main

func main() {
	// fs := http.FileServer(http.Dir("/Users/mp-haidera-pyse-403/Desktop/go_code"))
	// // fs := http.FileServer(http.Dir("/Users/mp-haidera-pyse-403/Desktop/go_code/goroutines.go"))

	// //setted hello.html and got the view of html there instead of file sharing
	// http.ListenAndServe(":8080", fs)

	// fs := http.FileServer(http.Dir("/Users/mp-haidera-pyse-403/Desktop/go_code"))
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.Header().Set("Content-Type", "text/html")
	// 	fmt.Fprint(rw, "<h1>Golang!</h1>")
	// })

	//will throw error because we have / at the end
	// http.Handle("/static", fs)
	//will still throw error because our static word is still in url
	// http.Handle("/static/", fs)
	// http.Handle("/static/", http.StripPrefix("/static", fs))
	// http.ListenAndServe(":8080", nil)

	//handlefunc mechanism
	// http.HandleFunc("/pdf", func(rw http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(rw, r, "/Users/mp-haidera-pyse-403/Desktop/docs/personal/haider_nov.pdf")
	// })
	// http.ListenAndServe(":8080", nil)

	//filepath.FromSlash replaces / \ according to os system
	// filepath.join() just like os.path.join in python

	/*We need to honor the fact that Go is a compiled language. Out of multiple program files located in multiple locations, we can create a binary executable. Any relative file paths used in these files hold no meaning in runtime.*/

	//just some random directory operations
	// exePath, err := os.Executable()   => /usr/share/program.exe
	// exeDir := filepath.Dir(exePath) => /usr/share/

	get working directory
	// os.Getwd()
}
