package main

import (
	"os"

	"github.com/c3sr/go-cupti/clibrary/env"
)

func GetTraceIdEnv() string {
	return os.Getenv(env.TraceIdEnv)
}
