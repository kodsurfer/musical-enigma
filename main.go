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

	userInterface(window)
}

func userInterface(window *gtk.Window) {
	btn, _ := gtk.ButtonNew()
	btn.Connect("click", func() {
		//
	})

	btn_label, _ := gtk.LabelNew("Play/Pause")
	btn.Add(btn_label)

	window.Add(btn)
}
