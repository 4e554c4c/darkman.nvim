package main

import (
	"github.com/godbus/dbus/v5"
)

const PORTAL_BUS_NAME = "org.freedesktop.portal.Desktop"
const PORTAL_OBJ_PATH = "/org/freedesktop/portal/desktop"
const PORTAL_INTERFACE = "org.freedesktop.portal.Settings"
const PORTAL_NAMESPACE = "org.freedesktop.appearance"
const PORTAL_KEY = "color-scheme"

type Portal struct {
	*dbus.Conn
}

func setupPortal() (Portal, error) {
	conn, err := dbus.ConnectSessionBus()
	return Portal{conn}, err
}

func (p *Portal) getMode() (uint32, error) {
	dest := p.Object(PORTAL_BUS_NAME, PORTAL_OBJ_PATH)
	var mode uint32
	err := dest.Call(PORTAL_INTERFACE+".Read", 0, PORTAL_NAMESPACE, PORTAL_KEY).Store(&mode)
	if err != nil {
		return 0, err
	}
	return mode, nil
}

func (p *Portal) setupSignal() (<-chan uint32, error) {
	signals := make(chan *dbus.Signal)
	modeChan := make(chan uint32)
	p.Signal(signals)
	err := p.AddMatchSignal(
		dbus.WithMatchSender(PORTAL_BUS_NAME),
		dbus.WithMatchObjectPath(PORTAL_OBJ_PATH),
		dbus.WithMatchInterface(PORTAL_INTERFACE),
		dbus.WithMatchMember("SettingChanged"),
		dbus.WithMatchArg0Namespace(PORTAL_NAMESPACE),
		dbus.WithMatchArg(1, PORTAL_KEY),
	)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			sig := <-signals
			if len(sig.Body) != 3 {
				continue
			}
			val, ok := sig.Body[2].(dbus.Variant).Value().(uint32)
			if ok {
				modeChan <- val
			}
		}
	}()

	return modeChan, nil
}
