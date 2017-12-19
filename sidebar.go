package main

import (
	"os"
	"path"

	qrcode "github.com/skip2/go-qrcode"
)

type sidebar struct {
	Balance   float64
	UnBalance float64
	Address   string
}

func sidebarValues() (sidebar, error) {
	// Get wallet balance.
	b, err := wallet.GetBalance()
	if err != nil {
		return sidebar{}, err
	}

	// Get wallet address.
	a, err := wallet.GetAddress()
	if err != nil {
		return sidebar{}, err
	}

	// Generate QR image.
	if _, err := os.Stat(path.Join("assets/images/qr", a.Address+
		".png")); os.IsNotExist(err) {
		if err := qrcode.WriteFile(a.Address, qrcode.Medium, 175, path.Join(
			"assets/images/qr", a.Address+".png")); err != nil {
			return sidebar{}, err
		}
	}

	return sidebar{
		float64(b.Balance) / 1.e+12,
		float64(b.UnBalance) / 1.e+12,
		a.Address,
	}, nil
}
