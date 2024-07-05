package adt_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	Describe("Emptiness", func() {
		Context("When the set does not contain item", func() {
			It("should be empty", func(){
				set := NewSet()

				Expect(set.IsEmpty()).To(BeTrue())
			})
		})
	})
}
)