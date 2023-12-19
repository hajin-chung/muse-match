package utils

import qrcode "github.com/skip2/go-qrcode"

func GetQrCode(url string) ([]byte, error) {
	var q *qrcode.QRCode
	q, err := qrcode.NewWithForcedVersion(url, 3, qrcode.Low)
	if err != nil {
		return nil, err
	}
	png, err := q.PNG(256)
	if err != nil {
		return nil, err
	}

	return png, nil
}
