package adt_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAdt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Adt Suite")
}
