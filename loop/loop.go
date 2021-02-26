package loop

import (
	"fmt"
	"github.com/otiai10/copy"
	hook "github.com/robotn/gohook"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

const (
	pathToAppData = "\\AppData\\Roaming"
	ds2Name       = "\\DarkSoulsII"

	pathToDS2    = pathToAppData + ds2Name
	pathToBackup = pathToDS2 + " (backup)"
)

const (
	graphicsConfig = "GraphicsConfig_SOFS.xml"
)

const (
	buffer        = 10
	sleepDuration = 500 * time.Millisecond
)

var (
	functionalKeys = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	//stopKeys = []string{
	//	"ctrl",
	//	"esc",
	//}
)

func Loop() (<-chan error, <-chan interface{}) {
	errCh := make(chan error, buffer)
	stopCh := make(chan interface{}, buffer)

	for i, key := range functionalKeys {
		index := i
		hook.Register(hook.KeyDown, []string{"ctrl", key}, func(event hook.Event) {
			log.Printf("Saving to slot %d...", index)

			err := performAction(index, false)
			if err != nil {
				errCh <- err
			}

			time.Sleep(sleepDuration)
		})
	}

	for i, key := range functionalKeys {
		index := i
		hook.Register(hook.KeyDown, []string{"shift", key}, func(event hook.Event) {
			log.Printf("Loading from slot %d...", index)

			err := performAction(index, true)
			if err != nil {
				errCh <- err
			}

			time.Sleep(sleepDuration)
		})
	}

	//hook.Register(hook.KeyDown, stopKeys, func(event hook.Event) {
	//	log.Printf("Stopping application...")
	//	stopCh <- struct{}{}
	//	hook.End()
	//})

	s := hook.Start()
	<-hook.Process(s)

	return errCh, stopCh
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}

		return home
	}

	return os.Getenv("HOME")
}

func performAction(index int, down bool) error {
	usrDir := userHomeDir()

	var from, to string
	if down {
		from = pathToBackupAt(usrDir, index)
		to = pathToOriginal(usrDir)
	} else {
		from = pathToOriginal(usrDir)
		to = pathToBackupAt(usrDir, index)
	}

	opts := copy.Options{
		Skip: func(src string) (bool, error) {
			return strings.HasSuffix(src, graphicsConfig), nil
		},
	}

	return copy.Copy(from, to, opts)
}

func pathToBackupAt(usrDir string, index int) string {
	return fmt.Sprintf("%s%s%s - %d", usrDir, pathToBackup, ds2Name, index)
}

func pathToOriginal(usrDir string) string {
	return usrDir + pathToDS2
}
