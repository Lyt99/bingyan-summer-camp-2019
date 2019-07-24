package models

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func PutPic(name string, f io.Reader) error {
	u, _ := url.Parse("https://demo-1258020847.cos.ap-chengdu.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("AKIDWnpbt74XrwlLeiDOt88uQl3IKtFXlTDi"),
			SecretKey: os.Getenv("PbZIsEdniOq24rPltmiRvAoF2dvHNEn9"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	//insert pic
	_, err := c.Object.Put(context.Background(), name, f, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DelPic(name string) error {
	u, _ := url.Parse("https://demo-1258020847.cos.ap-chengdu.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("AKIDWnpbt74XrwlLeiDOt88uQl3IKtFXlTDi"),
			SecretKey: os.Getenv("PbZIsEdniOq24rPltmiRvAoF2dvHNEn9"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})
	//delete pic
	_, err := c.Object.Delete(context.Background(), name)
	if err != nil {
		return err
	}
	return nil
}
