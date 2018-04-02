package main

import (
	"encoding/json"
	"log"
	"os"
	"path"

	"image/color"

	"github.com/gabstv/go-monero/walletrpc"
	"github.com/olahol/melody"
	qrcode "github.com/skip2/go-qrcode"
)

// Layout is a stuct with all the values needed in the layout template.
type Layout struct {
	Type      string
	Balance   string
	UnBalance string
	Address   string
	CurHeight uint64
	MaxHeight int64
}

func updateLayout(s *melody.Session) {
	data := Layout{Type: "layout"}

	// Get wallet balance.
	b, u, err := wallet.GetBalance()
	if err != nil {
		log.Print(err)
	}
	data.Balance = walletrpc.XMRToDecimal(b)
	data.UnBalance = walletrpc.XMRToDecimal(u)

	// Get wallet address.
	data.Address, err = wallet.GetAddress()
	if err != nil {
		log.Print(err)
	}

	// Get the current and max block height.
	data.CurHeight, err = wallet.GetHeight()
	if err != nil {
		log.Print(err)
	}
	data.MaxHeight, err = daemon.Height()
	if err != nil {
		log.Print(err)
	}

	// Generate QR image if required.
	if _, err := os.Stat(path.Join("static/images/qr", data.Address+
		".png")); os.IsNotExist(err) {
		if err := qrcode.WriteColorFile(data.Address, qrcode.Medium, 226,
			color.Transparent, color.White, path.Join("static/images/qr",
				data.Address+".png")); err != nil {
			log.Print(err)
		}
	}

	msg, err := json.Marshal(data)
	if err != nil {
		log.Print(err)
		return
	}

	s.Write(msg)
}
