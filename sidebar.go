package main

import (
	"os"
	"path"

	"image/color"

	qrcode "github.com/skip2/go-qrcode"
)

// Sidebar is a stuct with all the values needed in the sidebar templates.
type Sidebar struct {
	Balance   float64
	UnBalance float64
	Address   string
	CurHeight int64
	MaxHeight int
}

func sidebar() (s Sidebar, err error) {
	// Get wallet balance.
	s.Balance, s.UnBalance, err = walletBalance()
	if err != nil {
		return s, err
	}

	// Get wallet address.
	s.Address, err = walletAddress()
	if err != nil {
		return s, err
	}

	// Get the current and max block height.
	s.CurHeight, err = walletHeight()
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
