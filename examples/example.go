package main

import (
	"github.com/robowealth-mutual-fund/stdlog"
	"github.com/robowealth-mutual-fund/stdlog/constant/errorCode/otp"
)

func main() {
	err := otp.INCORRECT
	stdlog.Info().Err(&err)
}
