package main

import (
	"errors"
	"fmt"

	"github.com/neovim/go-client/nvim"
	"github.com/neovim/go-client/nvim/plugin"
)

const (
	NO_PREFERENCE = iota
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
	case LIGHT, NO_PREFERENCE:
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

func setup(v *nvim.Nvim, args setupArgs) {
	var err error
	var p Portal
	var ch <-chan uint32
	if currentMode != UNINITIALIZED {
		err = errors.New("setup() already called")
		goto error
	}
	args.v = v
	if p, err = setupPortal(); err != nil {
		goto error
	}
	if currentMode, err = p.getMode(); err != nil {
		goto error
	}
	if err = args.handleNewMode(); err != nil {
		goto error
	}

	if ch, err = p.setupSignal(); err != nil {
		goto error
	}
	go func() {
		for {
			if newMode := <-ch; newMode != currentMode {
				currentMode = newMode
				args.handleNewMode()
			}
		}
	}()
	return

error:
	v.WriteErr(fmt.Sprintf("darkman: %v\n", err))
	return
}

func main() {
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "DarkmanGetMode"}, getMode)
		p.HandleFunction(&plugin.FunctionOptions{Name: "DarkmanSetup"}, setup)
		return nil
	})
}
