package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	flags "github.com/jessevdk/go-flags"
	"github.com/shopspring/decimal"
	qrcode "github.com/skip2/go-qrcode"
)

func SVG(content string, level qrcode.RecoveryLevel, size int) (map[string]interface{}, error) {
	var q *qrcode.QRCode

	q, err := qrcode.New(content, level)

	if err != nil {
		return nil, err
	}

	var bitmap []interface{}
	qbitmap := q.Bitmap()
	pixWidth, err := decimal.NewFromString("5.0")
	if err != nil {
		return nil, err
	}
	pixHeight, err := decimal.NewFromString("5.0")
	if err != nil {
		return nil, err
	}
	offsetX, err := decimal.NewFromString("2.0")
	if err != nil {
		return nil, err
	}
	offsetY, err := decimal.NewFromString("2.0")
	if err != nil {
		return nil, err
	}
	idx := 0
	for x, line := range qbitmap {
		// Skip wide security border
		if x < 3 || x >= len(qbitmap)-3 {
			continue
		}
		for y, e := range line {
			// Skip wide security border
			if y < 3 || y >= len(qbitmap)-3 {
				continue
			}
			if e {
				bitmap = append(bitmap, map[string]interface{}{
					"Index":     idx,
					"X":         offsetX.Add(pixWidth.Mul(decimal.New(int64(x-3), 0))),
					"Y":         offsetY.Add(pixHeight.Mul(decimal.New(int64(y-3), 0))),
					"PixWidth":  pixWidth,
					"PixHeight": pixHeight,
				})
			}
			idx++
		}
	}
	// Tighten security border
	Width := pixWidth.Mul(decimal.New(int64(len(qbitmap)-5), 0))      // X
	Height := pixHeight.Mul(decimal.New(int64(len(qbitmap[0])-5), 0)) // Y
	return map[string]interface{}{
		"PixWidth":  pixWidth,
		"PixHeight": pixHeight,
		"Width":     Width,
		"Height":    Height,
		"Bitmap":    bitmap,
	}, nil
}

func byte2Color(c byte) string {
	colourScheme := []string{"E11DAA", "FF3D74", "E5EBDE", "00CAFF", "A720FF", "007EFF", "D59D54", "008742"}
	idx := int(c) >> 5
	if idx >= len(colourScheme) {
		return "000000"
	}
	return colourScheme[idx]
}

func AddrToHex(acc map[string]interface{}, addr common.Address) error {
	addrBytes := addr.Bytes()
	for i := 0; i < 13; i++ {
		bits := addrBytes[4+i]
		color := byte2Color(bits)
		blockIdx := fmt.Sprintf("Block%d", i)
		acc[blockIdx] = color
	}
	return nil
}

var opts struct {
	PrivateKey string `long:"key" required:"true" description:"Private key used to sign transaction"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	receiptIndex := 0
	pkey, err := crypto.HexToECDSA(opts.PrivateKey)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	key := fmt.Sprintf("%064x", pkey.D)
	addr := crypto.PubkeyToAddress(pkey.PublicKey)
	url := fmt.Sprintf("https://token.6120.eu/%v/%v", key, receiptIndex)
	log.Println("Address:", addr.Hex(), "URL:", url)
	svgMeta, err := SVG(url, qrcode.Medium, 256)
	if err = AddrToHex(svgMeta, addr); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	templ, err := template.ParseFiles("business_card.tpl")
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	var b bytes.Buffer
	templ.Execute(&b, svgMeta)
	fmt.Println(b.String())
}
