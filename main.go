package main

import (
	"log"

	"os"
	"time"

	"github.com/bogem/id3v2/v2"
	"github.com/gotk3/gotk3/gtk"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

var loaded, playing bool
var player oto.Player

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

	mp3 := os.Args[1]
	id3tag, err := id3v2.Open(mp3, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal("id3 tag parse error")
	}

	userInterface(window, id3tag)

}

func userInterface(window *gtk.Window, tag *id3v2.Tag) {
	grid, _ := gtk.GridNew()

	title, _ := gtk.LabelNew(tag.Title())
	grid.Attach(title, 0, 0, 1, 1)

	btn, _ := gtk.ButtonNew()
	btn.Connect("click", func() {
		if loaded {
			if playing {
				log.Println("Pause")
				player.Pause()
			} else {
				log.Println("Play")
				player.Play()
			}
		}
	})

	go func() {
		playing = player.IsPlaying()
		for {
			if playing != player.IsPlaying() {
				playing = player.IsPlaying()
				if playing {
					btn.SetLabel("Play")
				} else {
					btn.SetLabel("Pause")
				}
			}
			time.Sleep(16 * time.Millisecond)
		}
	}()

	btn_label, _ := gtk.LabelNew("Play/Pause")
	btn.Add(btn_label)

	grid.AttachNextTo(btn, title, gtk.POS_BOTTOM, 1, 1)
	window.Add(btn)
}

func initOto(mp3file string) {
	file, err := os.Open(mp3file)
	if err != nil {
		log.Fatalln(err.Error())
	}

	mp3.NewDecoder(file)

	oto.NewContext(44100, 2, 2)
}
