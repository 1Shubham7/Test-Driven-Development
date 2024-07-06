package adt

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	// "github.com/1shubham7/testing-e2e/adt"
)

var _ = Describe("Set", func() {
	Describe("Emptiness", func() {
		Context("When the set does not contain item", func() {
			It("should be empty", func(){
				set := NewSet()

				Expect(set.IsEmpty()).To(BeTrue())
			})
		})

		Context("When the set contains items", func(){
			It("Should not be empty", func(){
				set := NewSet()
				set.Add("Red")
				Expect(set.IsEmpty()).To(BeFalse())
			})
		})
	})

	Describe("Size", func(){
		Context("As Items are added", func(){
			It("It should return size of set", func(){
				set := NewSet()

				By("Empty set will return 0", func() {
					Expect(set.Size()).To(BeZero())
				})

				By("Adding one element", func() {
					set.Add("Kutta")

					Expect(set.Size()).To(Equal(1))
				})

				By("Adding another element", func() {
					set.Add("Billi")
					Expect(set.Size()).To(Equal(2))
				})
			})
		})
	})
},
)