package main

import (
	"ds2-tool/loop"
	"fmt"
	"log"
)

func main() {
	fmt.Println("|-----------------------------------------------------|")
	fmt.Println("|:::::::::::::::::::DS2 Backup Tool:::::::::::::::::::|")
	fmt.Println("|-----------------------------------------------------|")

	fmt.Println("| Press Ctrl + 1 to backup DS2 saves to slot 1        |")
	fmt.Println("| Press Shift + 1 to load DS2 backup from slot 1      |")
	fmt.Println("| Each of the keys 0 to 9 corresponds to a save slot  |")
	//fmt.Println("Press Ctrl + Escape to quit")

	fmt.Println("|-----------------------------------------------------|")
	fmt.Println()

	errCh, stopCh := loop.Loop()

feedbackLoop:
	for {
		select {
		case <-stopCh:
			break feedbackLoop
		case err := <-errCh:
			log.Printf("Error: %s", err)
		}

		fmt.Println("DS2 Backup Tool Stopped.")
	}
}
