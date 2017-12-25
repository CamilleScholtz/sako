package main

import (
	"os"
	"path"

	qrcode "github.com/skip2/go-qrcode"
)

// Sidebar is a stuct with all sidebar values.
type Sidebar struct {
	Balance   float64
	UnBalance float64
	Address   string
}

func sidebar() (Sidebar, error) {
	s := Sidebar{}

	// Get wallet balance.
	b, err := wallet.GetBalance()
	if err != nil {
		return s, err
	}

	// Get wallet address.
	a, err := wallet.GetAddress()
	if err != nil {
		return s, err
	}

	// Generate QR image.
	if _, err := os.Stat(path.Join("static/images/qr", a.Address+
		".png")); os.IsNotExist(err) {
		if err := qrcode.WriteFile(a.Address, qrcode.Medium, 175, path.Join(
			"static/images/qr", a.Address+".png")); err != nil {
			return s, err
		}
	}

	s.Balance = float64(b.Balance) / 1.e+12
	s.UnBalance = float64(b.UnBalance) / 1.e+12
	s.Address = a.Address

	return s, nil
}
