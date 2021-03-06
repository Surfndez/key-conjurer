package main

import (
	"fmt"

	"github.com/riotgames/key-conjurer/api/consts"
	"github.com/riotgames/key-conjurer/api/keyconjurer"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	fmt.Printf(`Starting GetUserData Lambda
	Version: %v
	`, consts.Version)
	lambda.Start(keyconjurer.GetUserDataEventHandler)
}
