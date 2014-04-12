package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("End")
}

func Main(){
	main()
}