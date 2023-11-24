package main

import (
	"asqnq-test/pkg/async_server/requests"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		requests.WelcomeRequest("example@gmail.com")
	}
}
