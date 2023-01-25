package main

import (
	"github.com/godbus/dbus/v5"
)

func portalGetMode() (uint32, error) {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return 0, err
	}
	dest := conn.Object("org.freedesktop.portal.Desktop", "/org/freedesktop/portal/desktop")
	var mode uint32
	err = dest.Call("org.freedesktop.portal.Settings.Read", 0, "org.freedesktop.appearance", "color-scheme").Store(&mode)
	if err != nil {
		return 0, err
	}
	return mode, nil
}

func setupSignal() (<-chan uint32, error) {
	var err error
	return nil, err
}
