package main

import (
	"context"
	"net/url"
	"os"
	"strings"

	"net/http"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

func main() {
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

	// Case1 normal put object
	name := "test/objectPut.go"
	f := strings.NewReader("test")

	_, err := c.Object.Put(context.Background(), name, f, nil)
	if err != nil {
		panic(err)
	}

	// Case2 put object with the options
	name = "test/put_option.go"
	f = strings.NewReader("test xxx")
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "text/html",
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			//XCosACL: "public-read",
			XCosACL: "private",
		},
	}
	_, err = c.Object.Put(context.Background(), name, f, opt)
	if err != nil {
		panic(err)
	}

	// Case3 put object by local file path
	_, err = c.Object.PutFromFile(context.Background(), name, "./test", nil)
	if err != nil {
		panic(err)
	}

}
