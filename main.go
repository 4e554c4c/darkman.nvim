package main

import (
	"errors"
	"fmt"

	"github.com/neovim/go-client/nvim"
	"github.com/neovim/go-client/nvim/plugin"
)

const (
	UNKNOWN = iota
	DARK
	LIGHT
	UNINITIALIZED
)

var currentMode uint32 = UNINITIALIZED

type setupArgs struct {
	v                *nvim.Nvim `msgpack:"-"`
	ChangeBackground bool       `msgpack:",array"`
	SendUserEvent    bool
	Colorscheme      *struct {
		Dark  string `msgpack:"dark"`
		Light string `msgpack:"light"`
	}
}

func getMode(args []string) (int, error) {
	if currentMode == UNINITIALIZED {
		return 0, errors.New("Mode not yet initialized, call `Setup`")
	}
	return int(currentMode), nil
}

func (args *setupArgs) handleNewMode() error {
	var err error
	var background, colorscheme, event string
	switch currentMode {
	case DARK:
		background, event = "dark", "DarkMode"
		if args.Colorscheme != nil {
			colorscheme = args.Colorscheme.Dark
		}
	case LIGHT:
		background, event = "light", "LightMode"
		if args.Colorscheme != nil {
			colorscheme = args.Colorscheme.Light
		}
	default:
		return errors.New(fmt.Sprintf("Unexpected mode: %d", currentMode))
	}
	if c := args.Colorscheme; c != nil {
		err = args.v.Command("colorscheme " + colorscheme)
		if err != nil {
			return err
		}
	}
	if args.ChangeBackground {
		err = args.v.SetOption("background", background)
		if err != nil {
			return err
		}
	}
	if args.SendUserEvent {
		err = args.v.Command("doautocmd User " + event)
		if err != nil {
			return err
		}
	}
	return err
}

func setup(v *nvim.Nvim, args setupArgs) error {
	if currentMode != UNINITIALIZED {
		return errors.New("setup() already called")
	}
	args.v = v
	p, err := setupPortal()
	if err != nil {
		return err
	}
	currentMode, err = p.getMode()
	if err != nil {
		return err
	}
	args.handleNewMode()

	ch, err := p.setupSignal()
	if err != nil {
		return err
	}
	go func() {
		for {
			newMode := <-ch
			if newMode == currentMode {
				continue
			}
			currentMode = newMode
			args.handleNewMode()
		}
	}()

	return nil
}

func main() {
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "DarkmanGetMode"}, getMode)
		p.HandleFunction(&plugin.FunctionOptions{Name: "DarkmanSetup"}, setup)
		return nil
	})
}
