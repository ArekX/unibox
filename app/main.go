package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new label widget to show in the window.
	l, err := gtk.LabelNew("Hello, gotk3! man!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	// Create a new label widget to show in the window.
	lx, err := gtk.LabelNew("Hello, gotk3adwdad! man!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	p, _ := gtk.PanedNew(gtk.ORIENTATION_HORIZONTAL)

	p.Add(l)
	p.Add(lx)

	// Add the label to the window.
	win.Add(p)

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
