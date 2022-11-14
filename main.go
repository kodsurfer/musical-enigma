package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)
	log.Print("Start musical-enigma")

	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatalf("Create window error %s", err)
	}

	window.SetTitle("Musical Enigma MP3 Player")

	window.Connect("quit", func() {
		gtk.MainQuit()
	})

}
