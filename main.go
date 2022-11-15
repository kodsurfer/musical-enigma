package main

import (
	"log"

	"os"

	"github.com/bogem/id3v2/v2"
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

	mp3 := os.Args[1]
	_, err = id3v2.Open(mp3, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("id3 tag parse error")
	}

}

func userInterface(window *gtk.Window) {
	btn, _ := gtk.ButtonNew()
	btn.Connect("click", func() {
		//click state
	})

	btn_label, _ := gtk.LabelNew("Play/Pause")
	btn.Add(btn_label)

	window.Add(btn)
}
