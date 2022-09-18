package commons

import (
	"github.com/kranzuft/stoc/cmd/com/nodlim/stoc/commons"
	"github.com/kranzuft/stoc/cmd/com/nodlim/stoc/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

/**
 * Init function
 */
func TestArrayUtil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CommonsArrayUtil")
}

var T_LBR = types.Token{Typ: types.LBR, Exp: "("}
var T_RBR = types.Token{Typ: types.RBR, Exp: ")"}
var T_AND = types.Token{Typ: types.AND, Exp: "&"}
var T_OR = types.Token{Typ: types.OR, Exp: "|"}

var _ = Describe("ArrayUtil tests", func() {
	// LastIndexOf tests
	Describe("LastIndexOf tests", func() {
		// small unique items integer array tests last index
		Describe("an int value item", func() {
			Context("against a small integer array of unique items", func() {
				It("should return the index of the integer in the array", func() {
					Expect(commons.LastIndexOf([]int{1, 2, 3, 4, 5}, func(t int) bool { return t == 1 })).To(Equal(0))
					Expect(commons.LastIndexOf([]int{1, 2, 3, 4, 5}, func(t int) bool { return t == 5 })).To(Equal(4))
					Expect(commons.LastIndexOf([]int{1, 2, 3, 4, 5}, func(t int) bool { return t == 3 })).To(Equal(2))
					Expect(commons.LastIndexOf([]int{1, 2, 3, 4, 5}, func(t int) bool { return t == 10 })).To(Equal(-1))
				})
			})
		})

		// small repeated items integer array tests last index
		Describe("an int value item", func() {
			Context("against a small integer array of repeated items", func() {
				It("should return the last index of the integer in the array", func() {
					Expect(commons.LastIndexOf([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, func(t int) bool { return t == 1 })).To(Equal(5))
					Expect(commons.LastIndexOf([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, func(t int) bool { return t == 5 })).To(Equal(9))
					Expect(commons.LastIndexOf([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, func(t int) bool { return t == 3 })).To(Equal(7))
					Expect(commons.LastIndexOf([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, func(t int) bool { return t == 10 })).To(Equal(-1))
				})
			})
		})

		// small repeated items string array tests last index
		Describe("a string value item", func() {
			Context("against a small string array of repeated items", func() {
				It("should return the last index of the string in the array", func() {
					Expect(commons.LastIndexOf([]string{"a", "b", "c", "d", "e", "a", "b", "c", "d", "e"}, func(t string) bool { return t == "a" })).To(Equal(5))
					Expect(commons.LastIndexOf([]string{"a", "b", "c", "d", "e", "a", "b", "c", "d", "e"}, func(t string) bool { return t == "e" })).To(Equal(9))
					Expect(commons.LastIndexOf([]string{"a", "b", "c", "d", "e", "a", "b", "c", "d", "e"}, func(t string) bool { return t == "f" })).To(Equal(-1))
				})
			})
		})

		// small repeated items token array tests last index
		Describe("a token value item", func() {
			Context("against a small token array of repeated items", func() {
				It("should return the last index of the token in the array", func() {
					Expect(commons.LastIndexOf([]types.Token{T_LBR, T_RBR, T_AND, T_OR, T_AND, T_OR}, func(t types.Token) bool { return t.Typ == types.AND })).To(Equal(4))
					Expect(commons.LastIndexOf([]types.Token{T_LBR, T_RBR, T_AND, T_OR, T_AND, T_OR}, func(t types.Token) bool { return t.Typ == types.OR })).To(Equal(5))
				})
			})
		})
	})

	Describe("StartsWith test", func() {
		Context("Test StartsWith with index 0", func() {
			It("should return true", func() {
				Expect(commons.StartsWith([]rune("12345"), 0, []rune("12345"))).To(Equal(true))
			})
			It("should return false", func() {
				Expect(commons.StartsWith([]rune("123"), 0, []rune("12345"))).To(Equal(false))
				Expect(commons.StartsWith([]rune("123567"), 0, []rune("12346"))).To(Equal(false))
			})
		})

		Context("Test StartsWith with index middle", func() {
			It("should return true", func() {
				Expect(commons.StartsWith([]rune("12345"), 3, []rune("45"))).To(Equal(true))
			})
			It("should return false", func() {
				Expect(commons.StartsWith([]rune("ABC123"), 3, []rune("12345"))).To(Equal(false))
				Expect(commons.StartsWith([]rune("123567"), 3, []rune("467"))).To(Equal(false))
			})
		})

		Context("Test StartsWith with index end", func() {
			It("should return true", func() {
				Expect(commons.StartsWith([]rune("12345"), 4, []rune("5"))).To(Equal(true))
			})
			It("should return false", func() {
				Expect(commons.StartsWith([]rune("ABC123"), 6, []rune("3"))).To(Equal(false))
				Expect(commons.StartsWith([]rune("123567"), 6, []rune("3"))).To(Equal(false))
			})
		})

		Context("Test multiple search terms", func() {
			It("should return true", func() {
				Expect(commons.StartsWith([]rune("ABC123"), 0, []rune("A"), []rune("AB"), []rune("C1"), []rune("23"), []rune("ABC123"))).To(Equal(true))
				Expect(commons.StartsWith([]rune("ABC123"), 2, []rune("A"), []rune("AB"), []rune("C1"), []rune("23"), []rune("ABC123"))).To(Equal(true))
				Expect(commons.StartsWith([]rune("ABC123"), 4, []rune("A"), []rune("AB"), []rune("C1"), []rune("23"), []rune("ABC123"))).To(Equal(true))
				Expect(commons.StartsWith([]rune("'test'"), 0, []rune("\""), []rune("'"))).To(Equal(true))
				Expect(commons.StartsWith([]rune("\"test\""), 0, []rune("\""), []rune("'"))).To(Equal(true))
			})
			It("should return false", func() {
				Expect(commons.StartsWith([]rune("ABC123"), 1, []rune("A"), []rune("AB"), []rune("C1"), []rune("23"), []rune("ABC123"))).To(Equal(false))
				Expect(commons.StartsWith([]rune("ABC123"), 3, []rune("A"), []rune("AB"), []rune("C1"), []rune("23"), []rune("ABC123"))).To(Equal(false))
				Expect(commons.StartsWith([]rune("ABC123"), 5, []rune("A"), []rune("AB"), []rune("C1"), []rune("23"), []rune("ABC123"))).To(Equal(false))
			})
		})
	})
})
