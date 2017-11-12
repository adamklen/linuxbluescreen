package main

import (
	"time"

	tb "github.com/nsf/termbox-go"
)

// Prints the message to screen using termbox.
func print(x, y int, msg string, fg, bg tb.Attribute) {
	for _, r := range msg {
		if r == '\n' {
			x = 0
			y++
			continue
		}
		tb.SetCell(x, y, r, fg, bg)
		x++
	}
}

func main() {
	const text = `A problem has been detected and Windows has been shut down to prevent damage
to your computer.

The problem seems to be caused by the following file: TRBLHCKX.SYS

MEME_FAULT_IN_NONMEME_AREA

If this is the first time you've seen this stop error screen,
restart your computer. If this screen appears again, follow
these steps:

Check to make sure any new hardware or software is properly installed.
If this is a new installation, ask your hardware or software manufacturer
 for any Windows updates you might need.

If problems continue, disable or remove any newly installed hardware
or software. Disable BIOS memory options such as caching or shadowing.
If you need to use Safe Mode to remove or disable components, restart
your computer, press F8 to select Advanced Startup Options, and then
select Safe Mode.

Technical information:

*** STOP: 0x00000050 (0xFD3094C2,0x00000001,0xFBFE7617,0x00000000)


***  SPCMDCON.SYS - Address FBFE7617 base at FBFE5000, DateStamp 3d6dd67`

	err := tb.Init()
	if err != nil {
		panic(err)
	}
	defer tb.Close()
	tb.HideCursor()

	const fg tb.Attribute = tb.ColorWhite | tb.AttrBold
	const bg tb.Attribute = tb.ColorBlue
	time.Sleep(1 * time.Millisecond)
	tb.Sync()
	w, h := tb.Size()
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			tb.SetCell(x, y, ' ', fg, bg)
		}
	}
	print(0, 0, text, fg, bg)
	tb.Flush()
	for {
		go time.Sleep(1 * time.Second)
	}
}
