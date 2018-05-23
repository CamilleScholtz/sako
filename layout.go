package main

import (
	"log"
	"os"
	"path"

	"image/color"

	"github.com/gabstv/go-monero/walletrpc"
	qrcode "github.com/skip2/go-qrcode"
)

func updateSidebar() {
	msg := struct {
		Balance   string
		UnBalance string
		Address   string
		CurHeight uint64
		MaxHeight int64
	}{}

	b, u, err := wallet.GetBalance()
	if err != nil {
		log.Print(err)
	}
	msg.Balance = walletrpc.XMRToDecimal(b)
	msg.UnBalance = walletrpc.XMRToDecimal(u)

	msg.Address, err = wallet.GetAddress()
	if err != nil {
		log.Print(err)
	}

	msg.CurHeight, err = wallet.GetHeight()
	if err != nil {
		log.Print(err)
	}
	msg.MaxHeight, err = daemon.Height()
	if err != nil {
		log.Print(err)
	}

	// Generate QR image if required.
	if _, err := os.Stat(path.Join("static/images/qr", msg.Address+".png")); os.
		IsNotExist(err) {
		if err := generateQR(msg.Address); err != nil {
			log.Print(err)
		}
	}

	event <- Event{"sidebar", &msg}
}

func updatePrice() {
	msg, err := cryptoComparePrice("XMR")
	if err != nil {
		log.Print(err)
	}

	event <- Event{"price", msg}

}

func generateQR(address string) error {
	return qrcode.WriteColorFile(address, qrcode.Medium, 226, color.Transparent,
		color.RGBA{0xED, 0xE4, 0xE1, 0xFF}, path.Join("static/images/qr",
			address+".png"))
}
