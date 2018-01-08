package main

import (
	"encoding/json"
	"os"
	"path"

	"image/color"

	"github.com/olahol/melody"
	"github.com/onodera-punpun/sako/monero"
	qrcode "github.com/skip2/go-qrcode"
)

// Sidebar is a stuct with all the values needed in the sidebar templates.
type Sidebar struct {
	Type      string
	Balance   monero.Balance
	Address   string
	CurHeight int64
	MaxHeight int64
}

func updateSidebar(s *melody.Session) (err error) {
	sb := Sidebar{Type: "sidebar"}

	// Get wallet balance.
	sb.Balance, err = wallet.Balance()
	if err != nil {
		return err
	}

	// Get wallet address.
	sb.Address, err = wallet.Address()
	if err != nil {
		return err
	}

	// Get the current and max block height.
	// TODO: Can I use this this to increase load times, as in compare the two
	// and return if they are not equal?
	sb.CurHeight, err = wallet.Height()
	if err != nil {
		return err
	}
	sb.MaxHeight, err = daemon.Height()
	if err != nil {
		return err
	}

	// Generate QR image if required.
	if _, err := os.Stat(path.Join("static/images/qr", sb.Address+
		".png")); os.IsNotExist(err) {
		if err := qrcode.WriteColorFile(sb.Address, qrcode.Medium, 226,
			color.Transparent, color.White, path.Join("static/images/qr",
				sb.Address+".png")); err != nil {
			return err
		}
	}

	msg, err := json.Marshal(sb)
	if err != nil {
		return err
	}

	return s.Write(msg)
}
