package pkg

import (
	"FilesDIR/config"
	"FilesDIR/loger"
	"FilesDIR/rgb"
	"fmt"
	"time"
)

const (
	start = `
		███████╗██╗██╗     ███████╗██████╗ ██╗██████╗ 
		██╔════╝██║██║     ██╔════╝██╔══██╗██║██╔══██╗
		█████╗  ██║██║     █████╗  ██║  ██║██║██████╔╝
		██╔══╝  ██║██║     ██╔══╝  ██║  ██║██║██╔══██╗
		██║     ██║███████╗███████╗██████╔╝██║██║  ██║
		╚═╝     ╚═╝╚══════╝╚══════╝╚═════╝ ╚═╝╚═╝  ╚═╝`
	sep = `~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~`

	author  = `Auteur:  `
	version = `Version: `
)

// ...
// FilesDIR
func DrawStart() {
	defer time.Sleep(1 * time.Second)

	loger.Ui(start)
	loger.Ui("\t\t", author+config.Author, "\n", "\t\t", version+config.Version)
	loger.Ui("\n")
	loger.Ui(sep)
	loger.Ui("\n")

	rgb.Green.Println(start)
	fmt.Print("\t\t", author+rgb.Green.Sprint(config.Author), "\n", "\t\t", version+rgb.Green.Sprint(config.Version))
	fmt.Print("\n\n")
	rgb.Gray.Println(sep)
	fmt.Print("\n")
}

func DrawEnd() {
	defer time.Sleep(10 * time.Second)

	loger.Ui(sep)
	loger.Ui(author+config.Author, "\n", version+config.Version)

	rgb.Gray.Println(sep)
	fmt.Print(author+rgb.Green.Sprint(config.Author), "\n", version+rgb.Green.Sprint(config.Version))
	fmt.Print("\n\n")
}

// ...
// Search
func DrawParam(txt string) {
	defer time.Sleep(600 * time.Millisecond)

	_pre := "          "
	_txt := fmt.Sprintf(" %s \n", txt)

	loger.Ui(_pre, _txt)
	loger.Ui("\n")

	fmt.Printf("%s%s", rgb.BgYellow.Sprint(_pre), rgb.Yellow.Sprint(_txt))
}
