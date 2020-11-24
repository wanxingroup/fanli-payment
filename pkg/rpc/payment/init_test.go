package payment_test

import (
	"os"
	"testing"

	"dev-gitlab.wanxingrowth.com/fanli/payment/pkg/utils/test"
)

func TestMain(m *testing.M) {

	test.Init()

	code := m.Run()

	test.Release()

	os.Exit(code)
}
