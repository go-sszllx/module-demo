package utils

import (
	"fmt"
	"github.com/micro/go-micro"

	"test/utils/http"
)

func TestUtil() {
	fmt.Println("test utils")

	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("test_micro"),
		micro.Version("0.1"),
	)
	srv.Init()

	http.HttpHandle("baidu")
}
