package util

import (
	"bytes"
	"io"
	"log"

	"github.com/longbridgeapp/opencc"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

var s2tw, _ = opencc.New("s2tw")

func ErrorHandler(err error, fatal bool) {
	if err != nil {
		if fatal {
			log.Fatalln(err.Error())
		} else {
			log.Println(err.Error())
		}
	}
}

func GbkToUtf8(b []byte) []byte {
	var result bytes.Buffer
	reader := transform.NewReader(bytes.NewReader(b), simplifiedchinese.GBK.NewDecoder())
	_, err := io.Copy(&result, reader)
	ErrorHandler(err, false)
	return result.Bytes()
}

func Utf8ToBig5(b []byte) []byte {
	var result bytes.Buffer
	reader := transform.NewReader(bytes.NewReader(b), traditionalchinese.Big5.NewEncoder())
	_, err := io.Copy(&result, reader)
	ErrorHandler(err, false)

	return result.Bytes()
}

func Simplified2TW(in string) string {
	out, err := s2tw.Convert(in)
	ErrorHandler(err, false)

	return out
}
