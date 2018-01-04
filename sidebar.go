package main

import (
	"os"
	"path"

	"image/color"

	"github.com/onodera-punpun/sako/monero"
	qrcode "github.com/skip2/go-qrcode"
)

// Sidebar is a stuct with all the values needed in the sidebar templates.
type Sidebar struct {
	Balance   monero.Balance
	Address   string
	CurHeight int64
	MaxHeight int
}

func sidebar() (s Sidebar, err error) {
	// Get wallet balance.
	s.Balance, err = wallet.Balance()
	if err != nil {
		return s, err
	}

	// Get wallet address.
	s.Address, err = wallet.Address()
	if err != nil {
		return s, err
	}

	// Get the current and max block height.
	s.CurHeight, err = wallet.Height()
	if err != nil {
		return s, err
	}
	s.MaxHeight, err = daemonHeight()
	if err != nil {
		return s, err
	}

	// Generate QR image.
	if _, err := os.Stat(path.Join("static/images/qr", s.Address+
		".png")); os.IsNotExist(err) {
		if err := qrcode.WriteColorFile(s.Address, qrcode.Medium, 226,
			color.Transparent, color.White, path.Join("static/images/qr",
				s.Address+".png")); err != nil {
			return s, err
		}
	}

	return s, nil
}
