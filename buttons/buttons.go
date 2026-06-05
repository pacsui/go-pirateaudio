package buttons

import (
	"log"
	"time"

	"github.com/warthog618/go-gpiocdev"
	"periph.io/x/host/v3"
)

func init() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
}

func OnButtonAPressed(fn func()) {
	onButtonPressed(5, fn)
}

func OnButtonBPressed(fn func()) {
	onButtonPressed(6, fn)
}

func OnButtonXPressed(fn func()) {
	onButtonPressed(16, fn)
}

func OnButtonYPressed(fn func()) {
	onButtonPressed(24, fn)
}

func onButtonPressed(n int, fn func()) {
	go func() {
		c, err := gpiocdev.NewChip("/dev/gpiochip0")
		if err != nil {
			log.Fatal(err)
		}
		defer c.Close()
		line, err := c.RequestLine(n,
			gpiocdev.WithPullUp,
			gpiocdev.WithFallingEdge,
			gpiocdev.WithDebounce(100*time.Millisecond),
			gpiocdev.WithEventHandler(func(evt gpiocdev.LineEvent) {
				fn()
			}),
		)
		if err != nil {
			log.Fatal(err)
		}
		defer line.Close()
		for {
		}
	}()
}
