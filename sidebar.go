package main

import (
	"os"
	"path"

	"image/color"

	qrcode "github.com/onodera-punpun/go-qrcode"
)

// Sidebar is a stuct with all sidebar values.
type Sidebar struct {
	Balance   float64
	UnBalance float64
	Address   string
	Height    int64
}

func sidebar() (Sidebar, error) {
	s := Sidebar{}

	// Get wallet balance.
	// TODO: Make this always use the same amount of decimals.
	b, err := wallet.GetBalance()
	if err != nil {
		return s, err
	}

	// Get wallet address.
	a, err := wallet.GetAddress()
	if err != nil {
		return s, err
	}

	// Get the wallet's current block height.
	h, err := wallet.GetHeight()
	if err != nil {
		return s, err
	}

	// Generate QR image.
	if _, err := os.Stat(path.Join("static/images/qr", a.Address+
		".png")); os.IsNotExist(err) {
		if err := qrcode.WriteColorFile(a.Address, qrcode.Medium, 226, color.Transparent, color.White,
			path.Join("static/images/qr", a.Address+".png")); err != nil {
			return s, err
		}
	}

	s.Balance = float64(b.Balance) / 1.e+12
	s.UnBalance = float64(b.UnBalance) / 1.e+12
	s.Address = a.Address
	s.Height = h.Height

	return s, nil
}
