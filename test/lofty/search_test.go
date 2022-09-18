package lofty

import (
	"github.com/kranzuft/stoc/cmd/com/nodlim/stoc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

// Test data
const shortTargetProse = "The lazy fox jumped over the fence"
const longTargetProse = "Programmers are often encouraged to use long variable names " +
	"regardless of context.\nThat is a mistake: clarity is often achieved through brevity. ∆√∫"

/**
 * Init function
 */
func TestSearch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Search")
}

// Get rid of error, we just want to no if success or not
func SearchString(command string, target string) bool {
	success, _ := stoc.SearchString(command, target)
	return success
}

/**
 * BDD Tests
 */
var _ = Describe("Search a few lines", func() {
	//
	// !(((A & !B))) | ((((A & B)))) | ((((!A & !B)))) | ((((!A & B))))
	//
	Describe("complex example", func() {
		Context(" ", func() {
			It(" ", func() {
				Expect(SearchString("!(((lazy & !dog))) | ((((lazy & dog)))) | !((((lazy | dog)))) | ((((!lazy & dog))))", shortTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo) => (foo)
	//
	Describe("singular expression (A)", func() {
		Context("in line with A", func() {
			It("should be true", func() {
				Expect(SearchString("The", shortTargetProse)).To(Equal(true))
				Expect(SearchString("jumped", shortTargetProse)).To(Equal(true))
				Expect(SearchString("fence", shortTargetProse)).To(Equal(true))
				Expect(SearchString("Programmers", longTargetProse)).To(Equal(true))
				Expect(SearchString("a", longTargetProse)).To(Equal(true))
				Expect(SearchString("variable", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line without A", func() {
			It("should be false", func() {
				Expect(SearchString("Thee", shortTargetProse)).To(Equal(false))
				Expect(SearchString("jumper", shortTargetProse)).To(Equal(false))
				Expect(SearchString("felt", shortTargetProse)).To(Equal(false))
				Expect(SearchString("Programmed", longTargetProse)).To(Equal(false))
				Expect(SearchString("alt", longTargetProse)).To(Equal(false))
				Expect(SearchString("variety", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo) => (foo)
	//
	Describe("singular expression (A̅)", func() {
		Context("in line with A", func() {
			It("should be false", func() {
				Expect(SearchString("!The", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!jumped", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!Programmers", longTargetProse)).To(Equal(false))
				Expect(SearchString("!a", longTargetProse)).To(Equal(false))
				Expect(SearchString("!variable", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line without A", func() {
			It("should be true", func() {
				Expect(SearchString("!Thee", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!jumper", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!felt", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!Programmed", longTargetProse)).To(Equal(true))
				Expect(SearchString("!alt", longTargetProse)).To(Equal(true))
				Expect(SearchString("!variety", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// foo | bar => foo or bar
	//
	Describe("expression with or operator (A+B)", func() {
		Context("in line with A and B", func() {
			It("should be true", func() {
				Expect(SearchString("lazy | fox", shortTargetProse)).To(Equal(true))
				Expect(SearchString("the | fence", shortTargetProse)).To(Equal(true))
				Expect(SearchString("over | the", shortTargetProse)).To(Equal(true))
				Expect(SearchString("are | brevity", longTargetProse)).To(Equal(true))
				Expect(SearchString("is | is", longTargetProse)).To(Equal(true))
				Expect(SearchString("encouraged | regardless", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("lazy | abc", shortTargetProse)).To(Equal(true))
				Expect(SearchString("the | fend", shortTargetProse)).To(Equal(true))
				Expect(SearchString("over | they", shortTargetProse)).To(Equal(true))
				Expect(SearchString("are | bravado", longTargetProse)).To(Equal(true))
				Expect(SearchString("is | si", longTargetProse)).To(Equal(true))
				Expect(SearchString("encouraged | regarding", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("laser | fox", shortTargetProse)).To(Equal(true))
				Expect(SearchString("than | fence", shortTargetProse)).To(Equal(true))
				Expect(SearchString("overtly | the", shortTargetProse)).To(Equal(true))
				Expect(SearchString("arsenal | brevity", longTargetProse)).To(Equal(true))
				Expect(SearchString("island | is", longTargetProse)).To(Equal(true))
				Expect(SearchString("encouraging | regardless", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not B and not A", func() {
			It("should be false", func() {
				Expect(SearchString("laser | abc", shortTargetProse)).To(Equal(false))
				Expect(SearchString("than | fend", shortTargetProse)).To(Equal(false))
				Expect(SearchString("overtly | they", shortTargetProse)).To(Equal(false))
				Expect(SearchString("arsenal | brevado", longTargetProse)).To(Equal(false))
				Expect(SearchString("island | si", longTargetProse)).To(Equal(false))
				Expect(SearchString("encouraging | regarding", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// foo & bar => foo and bar
	//
	Describe("expression with and operator (A.B) ", func() {
		Context("in line with A and B", func() {
			It("should be true", func() {
				Expect(SearchString("lazy & fox", shortTargetProse)).To(Equal(true))
				Expect(SearchString("The & over", shortTargetProse)).To(Equal(true))
				Expect(SearchString("jumped & fence", shortTargetProse)).To(Equal(true))
				Expect(SearchString("Programmers & mistake", longTargetProse)).To(Equal(true))
				Expect(SearchString("∆√∫ & brevity", longTargetProse)).To(Equal(true))
				Expect(SearchString("mistake & clarity", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("lazy & foxy", shortTargetProse)).To(Equal(false))
				Expect(SearchString("The & overt", shortTargetProse)).To(Equal(false))
				Expect(SearchString("jumped & fencing", shortTargetProse)).To(Equal(false))
				Expect(SearchString("Programmers & mistook", longTargetProse)).To(Equal(false))
				Expect(SearchString("∆√∫ & brave", longTargetProse)).To(Equal(false))
				Expect(SearchString("mistake & claire", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and B", func() {
			It("should be false", func() {
				Expect(SearchString("lazyish & fox", shortTargetProse)).To(Equal(false))
				Expect(SearchString("Then & over", shortTargetProse)).To(Equal(false))
				Expect(SearchString("jumpy & fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("Programmered & mistake", longTargetProse)).To(Equal(false))
				Expect(SearchString("through-fair & brevity", longTargetProse)).To(Equal(false))
				Expect(SearchString("mistakes & clarity", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("ladle & foxy", shortTargetProse)).To(Equal(false))
				Expect(SearchString("Thorn & overt", shortTargetProse)).To(Equal(false))
				Expect(SearchString("jam & fencing", shortTargetProse)).To(Equal(false))
				Expect(SearchString("Frog & mistook", longTargetProse)).To(Equal(false))
				Expect(SearchString("Clark & brave", longTargetProse)).To(Equal(false))
				Expect(SearchString("Park & claire", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// foo|(bar|baz) => foo or (bar or baz)
	//
	Describe("expression with and operator A+(B+C) ", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("lazy|(the|jumped)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("fence|(fence|fence)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("fox|(fence|The)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("Programmers|(mistake|variable)", longTargetProse)).To(Equal(true))
				Expect(SearchString("clarity|(often|achieved)", longTargetProse)).To(Equal(true))
				Expect(SearchString("use|(long|are)", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("lazy|(the|zumped)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("fence|(fence|zence)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("fox|(fence|zhe)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("Programmers|(mistake|zariable)", longTargetProse)).To(Equal(true))
				Expect(SearchString("∆√∫|(often|zchieved)", longTargetProse)).To(Equal(true))
				Expect(SearchString("use|(long|zre)", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("lazy|(th¬kje|jumped)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("fence|(fkjence|fence)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("fox|(fenjce|the)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("Programmers|(mis˚take|variable)", longTargetProse)).To(Equal(true))
				Expect(SearchString("clarity|(ofte¬n|achieved)", longTargetProse)).To(Equal(true))
				Expect(SearchString("use|(loßng|are)", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("ladπle|(the|jumped)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("deønse|(fence|fence)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("bo´´x|(fence|the)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("Grammers|(mistake|variable)", longTargetProse)).To(Equal(true))
				Expect(SearchString("Cl®arity|(often|achieved)", longTargetProse)).To(Equal(true))
				Expect(SearchString("US†A|(long|are)", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("lazys|(these|jumpeds)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("fencåes|(fe∂nces|fßences)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("foxs|(fences|thes)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("Proƒgrammering|(mßistakes|variabl∂es)", longTargetProse)).To(Equal(false))
				Expect(SearchString("claritys|(oftens|achi∫eveds)", longTargetProse)).To(Equal(false))
				Expect(SearchString("uses|(longs|ares)", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo|bar)|baz => (foo or bar) or baz
	//
	Describe("expression with and operator (A+B)+C ", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(the|jumped)|lazy", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fence|fence)|fence", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fence|The)|fox", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(mistake|variable)|Programmers", longTargetProse)).To(Equal(true))
				Expect(SearchString("(often|achieved)|clarity", longTargetProse)).To(Equal(true))
				Expect(SearchString("(long|are)|use", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(the|zumped)|lazy", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fence|zence)|fence", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fence|zhe)|fox", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(mistake|zariable)|Programmers", longTargetProse)).To(Equal(true))
				Expect(SearchString("(often|zchieved)|∆√∫", longTargetProse)).To(Equal(true))
				Expect(SearchString("(long|zre)|use", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(th¬kje|jumped)|lazy", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fkjence|fence)|fence", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fenjce|the)|fox", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(mis˚take|variable)|Programmers", longTargetProse)).To(Equal(true))
				Expect(SearchString("(ofte¬n|achieved)|clarity", longTargetProse)).To(Equal(true))
				Expect(SearchString("(loßng|are)|use", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(the|jumped)|ladπle", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fence|fence)|deønse", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fence|the)|bo´´x", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(mistake|variable)|Grammers", longTargetProse)).To(Equal(true))
				Expect(SearchString("(often|achieved)|Cl®arity", longTargetProse)).To(Equal(true))
				Expect(SearchString("(long|are)|US†A", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(these|jumpeds)|lazys", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fe∂nces|fßences)|fencåes", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fences|thes)|foxs", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(mßistakes|variabl∂es)|Proƒgrammering", longTargetProse)).To(Equal(false))
				Expect(SearchString("(oftens|achi∫eveds)|claritys", longTargetProse)).To(Equal(false))
				Expect(SearchString("(longs|ares)|uses", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo&bar)&baz => (foo and bar) and baz
	//
	Describe("expression with and operator (A.B).C ", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(the&jumped)&lazy", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fence&fence)&fence", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fence&The)&fox", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(mistake&variable)&Programmers", longTargetProse)).To(Equal(true))
				Expect(SearchString("(often&achieved)&clarity", longTargetProse)).To(Equal(true))
				Expect(SearchString("(long&are)&use", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(the&zumped)&lazy", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fence&zence)&fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fence&zhe)&fox", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(mistake&zariable)&Programmers", longTargetProse)).To(Equal(false))
				Expect(SearchString("(often&zchieved)&∆√∫", longTargetProse)).To(Equal(false))
				Expect(SearchString("(long&zre)&use", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(th¬kje&jumped)&lazy", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fkjence&fence)&fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fenjce&the)&fox", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(mis˚take&variable)&Programmers", longTargetProse)).To(Equal(false))
				Expect(SearchString("(ofte¬n&achieved)&clarity", longTargetProse)).To(Equal(false))
				Expect(SearchString("(loßng&are)&use", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(the&jumped)&ladπle", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fence&fence)&deønse", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fence&the)&bo´´x", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(mistake&variable)&Grammers", longTargetProse)).To(Equal(false))
				Expect(SearchString("(often&achieved)&Cl®arity", longTargetProse)).To(Equal(false))
				Expect(SearchString("(long&are)&US†A", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(these&jumpeds)&lazys", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fe∂nces&fßences)&fencåes", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fences&thes)&foxs", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(mßistakes&variabl∂es)&Proƒgrammering", longTargetProse)).To(Equal(false))
				Expect(SearchString("(oftens&achi∫eveds)&claritys", longTargetProse)).To(Equal(false))
				Expect(SearchString("(longs&ares)&uses", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// ((foo&bar)&baz) => ((foo and bar) and baz)
	//
	Describe("expression with and operator ((A.B).C)", func() {
		Context("in line with A and B and C)", func() {
			It("should be false)", func() {
				Expect(SearchString("((fence&fence)&fence)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("((the&jumped)&lazy)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("((fence&The)&fox)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("((mistake&variable)&Programmers)", longTargetProse)).To(Equal(true))
				Expect(SearchString("((often&achieved)&clarity)", longTargetProse)).To(Equal(true))
				Expect(SearchString("((long&are)&use)", longTargetProse)).To(Equal(true))
			})
		})

		Context("(in line with A and B and not C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((the&zumped)&lazy)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fence&zence)&fence)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fence&zhe)&fox)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((mistake&zariable)&Programmers)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((often&zchieved)&∆√∫)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((long&zre)&use)", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with A and not B and C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((th¬kje&jumped)&lazy)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fkjence&fence)&fence)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fenjce&the)&fox)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((mis˚take&variable)&Programmers)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((ofte¬n&achieved)&clarity)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((loßng&are)&use)", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with not A and B and C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((the&jumped)&ladπle)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fence&fence)&deønse)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fence&the)&bo´´x)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((mistake&variable)&Grammers)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((often&achieved)&Cl®arity)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((long&are)&US†A)", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with not A and not B and not C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((these&jumpeds)&lazys)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fe∂nces&fßences)&fencåes)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fences&thes)&foxs)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((mßistakes&variabl∂es)&Proƒgrammering)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((oftens&achi∫eveds)&claritys)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((longs&ares)&uses)", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// foo|!(bar|!baz) => foo or not (bar or not baz)
	//
	Describe("expression with or-not operator A+(B̅+C̅) ", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("lazy|!(the|!jumped)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("fence|!(fence|!fence)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("fox|!(fence|!The)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("Programmers|!(mistake|!variable)", longTargetProse)).To(Equal(true))
				Expect(SearchString("clarity|!(often|!achieved)", longTargetProse)).To(Equal(true))
				Expect(SearchString("use|!(long|!are)", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("lazy|!(the|!zumped)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("fence|!(fence|!zence)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("fox|!(fence|!zhe)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("Programmers|!(mistake|!zariable)", longTargetProse)).To(Equal(true))
				Expect(SearchString("∆√∫|!(often|!zchieved)", longTargetProse)).To(Equal(true))
				Expect(SearchString("use|!(long|!zre)", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("lazy|!(th¬kje|!jumped)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("fence|!(fkjence|!fence)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("fox|!(fenjce|!the)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("Programmers|!(mis˚take|!variable)", longTargetProse)).To(Equal(true))
				Expect(SearchString("clarity|!(ofte¬n|!achieved)", longTargetProse)).To(Equal(true))
				Expect(SearchString("use|!(loßng|!are)", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("ladπle|!(the|!jumped)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("deønse|!(fence|!fence)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("bo´´x|!(fence|!the)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("Grammers|!(mistake|!variable)", longTargetProse)).To(Equal(false))
				Expect(SearchString("Cl®arity|!(often|!achieved)", longTargetProse)).To(Equal(false))
				Expect(SearchString("US†A|!(long|!are)", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("lazys|!(these|!jumpeds)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("fencåes|!(fe∂nces|!fßences)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("foxs|!(fences|!thes)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("Proƒgrammering|!(mßistakes|!variabl∂es)", longTargetProse)).To(Equal(false))
				Expect(SearchString("claritys|!(oftens|!achi∫eveds)", longTargetProse)).To(Equal(false))
				Expect(SearchString("uses|!(longs|!ares)", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo|!bar)|!baz => (foo or not bar) or not baz
	//
	Describe("expression with or not operator (A+B̅)+C̅ ", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(the|!jumped)|!lazy", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fence|!fence)|!fence", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fence|!The)|!fox", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(mistake|!variable)|!Programmers", longTargetProse)).To(Equal(true))
				Expect(SearchString("(often|!achieved)|!clarity", longTargetProse)).To(Equal(true))
				Expect(SearchString("(long|!are)|!use", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(the|!jumped)|!zazy", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fence|!fence)|!zence", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fence|!the)|!zox", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(mistake|!variable)|!zrogrammers", longTargetProse)).To(Equal(true))
				Expect(SearchString("(often|!achieved)|!zlap", longTargetProse)).To(Equal(true))
				Expect(SearchString("(long|!are)|!zuse", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(the|!jomped)|!lazy", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fence|!fonce)|!fence", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fence|!tho)|!fox", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(mistake|!voriable)|!Programmers", longTargetProse)).To(Equal(true))
				Expect(SearchString("(often|!ochieved)|!clarity", longTargetProse)).To(Equal(true))
				Expect(SearchString("(long|!ore)|!use", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(thez|!jumped)|!lazy", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fencez|!fence)|!fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fencez|!the)|!fox", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(mistakez|!variable)|!Programmers", longTargetProse)).To(Equal(false))
				Expect(SearchString("(oftenz|!achieved)|!clarity", longTargetProse)).To(Equal(false))
				Expect(SearchString("(longz|!are)|!use", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(these|!jumpeds)|!lazys", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fe∂nces|!fßences)|!fencåes", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(fences|!thes)|!foxs", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(mßistakes|!variabl∂es)|!Proƒgrammering", longTargetProse)).To(Equal(true))
				Expect(SearchString("(oftens|!achi∫eveds)|!claritys", longTargetProse)).To(Equal(true))
				Expect(SearchString("(longs|!ares)|!uses", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// (foo&!bar)&!baz => (foo and not bar) and not baz
	//
	Describe("expression with and-not operator (A.B̅).C̅ ", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(the&!jumped)&!lazy", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fence&!fence)&!fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fence&!The)&!fox", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(mistake&!variable)&!Programmers", longTargetProse)).To(Equal(false))
				Expect(SearchString("(often&!achieved)&!clarity", longTargetProse)).To(Equal(false))
				Expect(SearchString("(long&!are)&!use", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(the&!zumped)&!lazy", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fence&!zence)&!fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fence&!zhe)&!fox", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(mistake&!zariable)&!Programmers", longTargetProse)).To(Equal(false))
				Expect(SearchString("(often&!zchieved)&!∆√∫", longTargetProse)).To(Equal(false))
				Expect(SearchString("(long&!zre)&!use", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(th¬kje&!jumped)&!lazy", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fkjence&!fence)&!fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fenjce&!the)&!fox", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(mis˚take&!variable)&!Programmers", longTargetProse)).To(Equal(false))
				Expect(SearchString("(ofte¬n&!achieved)&!clarity", longTargetProse)).To(Equal(false))
				Expect(SearchString("(loßng&!are)&!use", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(the&!jumped)&!ladπle", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fence&!fence)&!deønse", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fence&!the)&!bo´´x", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(mistake&!variable)&!Grammers", longTargetProse)).To(Equal(false))
				Expect(SearchString("(often&!achieved)&!Cl®arity", longTargetProse)).To(Equal(false))
				Expect(SearchString("(long&!are)&!US†A", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(these&!jumpeds)&!lazys", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fe∂nces&!fßences)&!fencåes", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(fences&!thes)&!foxs", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(mßistakes&!variabl∂es)&!Proƒgrammering", longTargetProse)).To(Equal(false))
				Expect(SearchString("(oftens&!achi∫eveds)&!claritys", longTargetProse)).To(Equal(false))
				Expect(SearchString("(longs&!ares)&!uses", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// ((foo&!bar)&!baz) => ((foo and not bar) and not baz)
	//
	Describe("expression with and-not operator ((A.B̅).C̅)", func() {
		Context("in line with A and B and C)", func() {
			It("should be false)", func() {
				Expect(SearchString("((fence&!fence)&!fence)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((the&!jumped)&!lazy)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fence&!The)&!fox)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((mistake&!variable)&!Programmers)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((often&!achieved)&!clarity)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((long&!are)&!use)", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with A and B and not C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((the&!zumped)&!lazy)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fence&!zence)&!fence)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fence&!zhe)&!fox)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((mistake&!zariable)&!Programmers)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((often&!zchieved)&!∆√∫)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((long&!zre)&!use)", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with A and not B and C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((th¬kje&!jumped)&!lazy)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fkjence&!fence)&!fence)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fenjce&!the)&!fox)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((mis˚take&!variable)&!Programmers)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((ofte¬n&!achieved)&!clarity)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((loßng&!are)&!use)", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with not A and B and C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((the&!jumped)&!ladπle)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fence&!fence)&!deønse)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fence&!the)&!bo´´x)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((mistake&!variable)&!Grammers)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((often&!achieved)&!Cl®arity)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((long&!are)&!US†A)", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with not A and not B and not C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((these&!jumpeds)&!lazys)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fe∂nces&!fßences)&!fencåes)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((fences&!thes)&!foxs)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((mßistakes&!variabl∂es)&!Proƒgrammering)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((oftens&!achi∫eveds)&!claritys)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((longs&!ares)&!uses)", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// "foo &! bar" => foo and not bar
	//
	Describe("expression with and-not operator (A.B̅) ", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("lazy &! fox", shortTargetProse)).To(Equal(false))
				Expect(SearchString("The &! over", shortTargetProse)).To(Equal(false))
				Expect(SearchString("jumped &! fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("Programmers &! mistake", longTargetProse)).To(Equal(false))
				Expect(SearchString("∆√∫ &! brevity", longTargetProse)).To(Equal(false))
				Expect(SearchString("mistake &! clarity", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("lazy &! foxy", shortTargetProse)).To(Equal(true))
				Expect(SearchString("The &! overt", shortTargetProse)).To(Equal(true))
				Expect(SearchString("jumped &! fencing", shortTargetProse)).To(Equal(true))
				Expect(SearchString("Programmers &! mistook", longTargetProse)).To(Equal(true))
				Expect(SearchString("∆√∫ &! brave", longTargetProse)).To(Equal(true))
				Expect(SearchString("mistake &! claire", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with not A and B", func() {
			It("should be false", func() {
				Expect(SearchString("lazyish &! fox", shortTargetProse)).To(Equal(false))
				Expect(SearchString("Then &! over", shortTargetProse)).To(Equal(false))
				Expect(SearchString("jumpy &! fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("Prop &! mistake", longTargetProse)).To(Equal(false))
				Expect(SearchString("through-fair &! brevity", longTargetProse)).To(Equal(false))
				Expect(SearchString("mistakes &! clarity", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("ladle &    ! foxy", shortTargetProse)).To(Equal(false))
				Expect(SearchString("Thorn &   ! overt", shortTargetProse)).To(Equal(false))
				Expect(SearchString("jam & ! fencing", shortTargetProse)).To(Equal(false))
				Expect(SearchString("Frog &      ! mistook", longTargetProse)).To(Equal(false))
				Expect(SearchString("Clark &    ! brave", longTargetProse)).To(Equal(false))
				Expect(SearchString("Park &   ! claire", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// "foo |! bar" => foo or not bar
	//
	Describe("expression with 'or not' operator (A+B̅)", func() {
		Context("in line with A and B", func() {
			It("should be true", func() {
				Expect(SearchString("lazy |! fox", shortTargetProse)).To(Equal(true))
				Expect(SearchString("the |! fence", shortTargetProse)).To(Equal(true))
				Expect(SearchString("over |! the", shortTargetProse)).To(Equal(true))
				Expect(SearchString("are |! brevity", longTargetProse)).To(Equal(true))
				Expect(SearchString("is |! is", longTargetProse)).To(Equal(true))
				Expect(SearchString("encouraged |! regardless", longTargetProse)).To(Equal(true))

				Expect(SearchString("lazy |  ! fox", shortTargetProse)).To(Equal(true))
				Expect(SearchString("the | ! fence", shortTargetProse)).To(Equal(true))
				Expect(SearchString("over |  ! the", shortTargetProse)).To(Equal(true))
				Expect(SearchString("are | ! brevity", longTargetProse)).To(Equal(true))
				Expect(SearchString("is | ! is", longTargetProse)).To(Equal(true))
				Expect(SearchString("encouraged |     ! regardless", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("lazy |! abc", shortTargetProse)).To(Equal(true))
				Expect(SearchString("the |! fend", shortTargetProse)).To(Equal(true))
				Expect(SearchString("over |! they", shortTargetProse)).To(Equal(true))
				Expect(SearchString("are |! bravado", longTargetProse)).To(Equal(true))
				Expect(SearchString("is |! si", longTargetProse)).To(Equal(true))
				Expect(SearchString("encouraged | regarding", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with B and not A", func() {
			It("should be false", func() {
				Expect(SearchString("laser |! fox", shortTargetProse)).To(Equal(false))
				Expect(SearchString("than |! fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("overtly |! the", shortTargetProse)).To(Equal(false))
				Expect(SearchString("arsenal |! brevity", longTargetProse)).To(Equal(false))
				Expect(SearchString("island |! is", longTargetProse)).To(Equal(false))
				Expect(SearchString("encouraging |! regardless", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("laser |! abc", shortTargetProse)).To(Equal(true))
				Expect(SearchString("than |! fend", shortTargetProse)).To(Equal(true))
				Expect(SearchString("overtly |! they", shortTargetProse)).To(Equal(true))
				Expect(SearchString("arsenal |! brevado", longTargetProse)).To(Equal(true))
				Expect(SearchString("island |! si", longTargetProse)).To(Equal(true))
				Expect(SearchString("encouraging |! regarding", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo &! bar" => not foo and not bar
	//
	Describe("expression with and-not operator (A̅.B̅) ", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("!lazy &! fox", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!The &! over", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!jumped &! fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!Programmers &! mistake", longTargetProse)).To(Equal(false))
				Expect(SearchString("!∆√∫ &! brevity", longTargetProse)).To(Equal(false))
				Expect(SearchString("!mistake &! clarity", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("!lazy &! foxy", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!The &! overt", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!jumped &! fencing", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!Programmers &! mistook", longTargetProse)).To(Equal(false))
				Expect(SearchString("!∆√∫ &! brave", longTargetProse)).To(Equal(false))
				Expect(SearchString("!mistake &! claire", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and B", func() {
			It("should be false", func() {
				Expect(SearchString("!lazyish &! fox", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!Then &! over", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!jumpy &! fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!Prop &! mistake", longTargetProse)).To(Equal(false))
				Expect(SearchString("!through-fair &! brevity", longTargetProse)).To(Equal(false))
				Expect(SearchString("!mistakes &! clarity", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("!ladle &    ! foxy", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!Thorn &   ! overt", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!jam & ! fencing", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!Frog &      ! mistook", longTargetProse)).To(Equal(true))
				Expect(SearchString("!Clark &    ! brave", longTargetProse)).To(Equal(true))
				Expect(SearchString("!Park &   ! claire", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo |! bar" => not foo or not bar
	//
	Describe("expression with 'or not' operator (A̅+B̅)", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("!lazy |! fox", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!the |! fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!over |! the", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!are |! brevity", longTargetProse)).To(Equal(false))
				Expect(SearchString("!is |! is", longTargetProse)).To(Equal(false))
				Expect(SearchString("!encouraged |! regardless", longTargetProse)).To(Equal(false))

				// with extra spaces
				Expect(SearchString("!lazy |  ! fox", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!the | ! fence", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!over |  ! the", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!are | ! brevity", longTargetProse)).To(Equal(false))
				Expect(SearchString("!is | ! is", longTargetProse)).To(Equal(false))
				Expect(SearchString("!encouraged |     ! regardless", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("!lazy |! abc", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!the |! fend", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!over |! they", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!are |! bravado", longTargetProse)).To(Equal(true))
				Expect(SearchString("!is |! si", longTargetProse)).To(Equal(true))
				Expect(SearchString("!encouraged |! regarding", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("!laser |! fox", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!than |! fence", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!overtly |! the", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!arsenal |! brevity", longTargetProse)).To(Equal(true))
				Expect(SearchString("!island |! is", longTargetProse)).To(Equal(true))
				Expect(SearchString("!encouraging |! regardless", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("!laser |! abc", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!than |! fend", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!overtly |! they", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!arsenal |! brevado", longTargetProse)).To(Equal(true))
				Expect(SearchString("!island |! si", longTargetProse)).To(Equal(true))
				Expect(SearchString("!encouraging |! regarding", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo &! bar" => not foo and not bar
	//
	Describe("expression with and-not operator (A̅.B̅) ", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("(!lazy &! fox))", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!The &! over)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!jumped &! fence)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!Programmers &! mistake)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!∆√∫ &! brevity)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!mistake &! clarity)", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("(!lazy &! foxy)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!The &! overt)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!jumped &! fencing)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!Programmers &! mistook)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!∆√∫ &! brave)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!mistake &! claire)", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and B", func() {
			It("should be false", func() {
				Expect(SearchString("(!lazyish &! fox)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!Then &! over)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!jumpy &! fence)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!Prop &! mistake)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!through-fair &! brevity)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!mistakes &! clarity)", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("(!ladle &    ! foxy)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!Thorn &   ! overt)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!jam & ! fencing)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!Frog &      ! mistook)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!Clark &    ! brave)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!Park &   ! claire)", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo |! bar" => not foo or not bar
	//
	Describe("expression with 'or not' operator (A̅+B̅)", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("(!lazy |! fox)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!the |! fence)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!over |! the)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!are |! brevity)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!is |! is)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!encouraged |! regardless)", longTargetProse)).To(Equal(false))

				// with extra spaces
				Expect(SearchString("(!lazy |  ! fox)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!the | ! fence)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!over |  ! the)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!are | ! brevity)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!is | ! is)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!encouraged |     ! regardless)", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("(!lazy |! abc)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!the |! fend)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!over |! they)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!are |! bravado)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!is |! si)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!encouraged |! regarding)", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("(!laser |! fox)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!than |! fence)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!overtly |! the)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!arsenal |! brevity)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!island |! is)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!encouraging |! regardless)", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("(!laser |! abc)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!than |! fend)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!overtly |! they)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!arsenal |! brevado)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!island |! si)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!encouraging |! regarding)", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo &bar" => not foo and bar
	//
	Describe("expression with and-not operator (A̅.B) ", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("(!lazy &fox))", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!The &over)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!jumped &fence)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!Programmers &mistake)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!∆√∫ &brevity)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!mistake &clarity)", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("(!lazy &foxy)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!The &overt)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!jumped &fencing)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!Programmers &mistook)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!∆√∫ &brave)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!mistake &claire)", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and B", func() {
			It("should be true", func() {
				Expect(SearchString("(!lazyish &fox)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!Then &over)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!jumpy &fence)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!Prop &mistake)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!through-fair &brevity)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!mistakes &clarity)", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("(!ladle &    foxy)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!Thorn &   overt)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!jam &  fencing)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!Frog &       mistook)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!Clark &     brave)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!Park &    claire)", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// "!foo | bar" => not foo or bar
	//
	Describe("expression with 'or not' operator (A̅+B)", func() {
		Context("in line with A and B", func() {
			It("should be true", func() {
				Expect(SearchString("(!lazy | fox)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!the | fence)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!over | the)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!are | brevity)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!is | is)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!encouraged | regardless)", longTargetProse)).To(Equal(true))

				// with extra spaces
				Expect(SearchString("(!lazy |   fox)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!the |  fence)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!over |   the)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!are |  brevity)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!is |  is)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!encouraged |      regardless)", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("(!lazy | abc)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!the | fend)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!over | they)", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!are | bravado)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!is | si)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!encouraged | regarding)", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("(!laser | fox)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!than | fence)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!overtly | the)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!arsenal | brevity)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!island | is)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!encouraging | regardless)", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with not B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("(!laser | abc)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!than | fend)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!overtly | they)", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!arsenal | brevado)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!island | si)", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!encouraging | regarding)", longTargetProse)).To(Equal(true))
			})
		})
	})

	// QUOTES SECTION
	//
	// !(((A & !B))) | ((((A & B)))) | ((((!A & !B)))) | ((((!A & B))))
	//
	Describe("complex example with double quotes", func() {
		Context(" ", func() {
			It(" ", func() {
				Expect(SearchString("!(((\"lazy\" & !\"dog\"))) | ((((\"lazy\" & \"dog\")))) | !((((\"lazy\" | \"dog\")))) | ((((!\"lazy\" & \"dog\"))))", shortTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo) => (foo)
	//
	Describe("singular expression (A) with double quotes", func() {
		Context("in line with A", func() {
			It("should be true", func() {
				Expect(SearchString("\"The\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"jumped\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"fence\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"Programmers\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"a\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"variable\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line without A", func() {
			It("should be false", func() {
				Expect(SearchString("\"Thee\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"jumper\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"felt\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"Programmed\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"alt\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"variety\"", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo) => (foo)
	//
	Describe("singular expression (A̅) with double quotes", func() {
		Context("in line with A", func() {
			It("should be false", func() {
				Expect(SearchString("!\"The\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"jumped\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"Programmers\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("!\"a\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("!\"variable\"", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line without A", func() {
			It("should be true", func() {
				Expect(SearchString("!\"Thee\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"jumper\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"felt\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"Programmed\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("!\"alt\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("!\"variety\"", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// foo | bar => foo or bar
	//
	Describe("expression with or operator (A+B) with double quotes", func() {
		Context("in line with A and B", func() {
			It("should be true", func() {
				Expect(SearchString("\"lazy\" | \"fox\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"the\" | \"fence\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"over\" | \"the\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"are\" | \"brevity\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"is\" | \"is\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"encouraged\" | \"regardless\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("\"lazy\" | \"abc\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"the\" | \"fend\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"over\" | \"they\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"are\" | \"bravado\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"is\" | \"si\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"encouraged\" | \"regarding\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("\"laser\" | \"fox\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"than\" | \"fence\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"overtly\" | \"the\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"arsenal\" | \"brevity\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"island\" | \"is\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"encouraging\" | \"regardless\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not B and not A", func() {
			It("should be false", func() {
				Expect(SearchString("\"laser\" | \"abc\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"than\" | \"fend\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"overtly\" | \"they\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"arsenal\" | \"brevado\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"island\" | \"si\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"encouraging\" | \"regarding\"", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// foo & bar => foo and bar
	//
	Describe("expression with and operator (A.B) with double quotes", func() {
		Context("in line with A and B", func() {
			It("should be true", func() {
				Expect(SearchString("\"lazy\" & \"fox\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"The\" & \"over\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"jumped\" & \"fence\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"Programmers\" & \"mistake\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("∆√∫ & \"brevity\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"mistake\" & \"clarity\"", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("\"lazy\" & \"foxy\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"The\" & \"overt\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"jumped\" & \"fencing\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"Programmers\" & \"mistook\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("∆√∫ & \"brave\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"mistake\" & \"claire\"", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and B", func() {
			It("should be false", func() {
				Expect(SearchString("\"lazyish\" & \"fox\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"Then\" & \"over\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"jumpy\" & \"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"Programmered\" & \"mistake\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"through-fair\" & \"brevity\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"mistakes\" & \"clarity\"", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("\"ladle\" & \"foxy\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"Thorn\" & \"overt\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"jam\" & \"fencing\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"Frog\" & \"mistook\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"Clark\" & \"brave\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"Park\" & \"claire\"", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// foo|(bar|baz) => foo or (bar or baz)
	//
	Describe("expression with and operator A+(B+C) with double quotes", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("\"lazy\"|(\"the\"|\"jumped\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"fence\"|(\"fence\"|\"fence\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"fox\"|(\"fence\"|\"The\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"Programmers\"|(\"mistake\"|\"variable\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"clarity\"|(\"often\"|\"achieved\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"use\"|(\"long\"|\"are\")", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("\"lazy\"|(\"the\"|\"zumped\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"fence\"|(\"fence\"|\"zence\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"fox\"|(\"fence\"|\"zhe\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"Programmers\"|(\"mistake\"|\"zariable\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("∆√∫|(\"often\"|\"zchieved\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"use\"|(\"long\"|\"zre\")", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("\"lazy\"|(\"th¬kje\"|\"jumped\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"fence\"|(\"fkjence\"|\"fence\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"fox\"|(\"fenjce\"|\"the\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"Programmers\"|(\"mis˚take\"|\"variable\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"clarity\"|(\"ofte¬n\"|\"achieved\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"use\"|(\"loßng\"|\"are\")", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("\"ladπle\"|(\"the\"|\"jumped\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"deønse\"|(\"fence\"|\"fence\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"bo´´x\"|(\"fence\"|\"the\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"Grammers\"|(\"mistake\"|\"variable\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"Cl®arity\"|(\"often\"|\"achieved\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"US†A\"|(\"long\"|\"are\")", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("\"lazys\"|(\"these\"|\"jumpeds\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"fencåes\"|(\"fe∂nces\"|\"fßences\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"foxs\"|(\"fences\"|\"thes\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"Proƒgrammering\"|(\"mßistakes\"|\"variabl∂es\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"claritys\"|(\"oftens\"|\"achi∫eveds\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"uses\"|(\"longs\"|\"ares\")", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo|bar)|baz => (foo or bar) or baz
	//
	Describe("expression with and operator (A+B)+C with double quotes", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"the\"|\"jumped\")|\"lazy\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fence\"|\"fence\")|\"fence\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fence\"|\"The\")|\"fox\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"mistake\"|\"variable\")|\"Programmers\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"often\"|\"achieved\")|\"clarity\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"long\"|\"are\")|\"use\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"the\"|\"zumped\")|\"lazy\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fence\"|\"zence\")|\"fence\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fence\"|\"zhe\")|\"fox\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"mistake\"|\"zariable\")|\"Programmers\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"often\"|\"zchieved\")|∆√∫", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"long\"|\"zre\")|\"use\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"th¬kje\"|\"jumped\")|\"lazy\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fkjence\"|\"fence\")|\"fence\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fenjce\"|\"the\")|\"fox\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"mis˚take\"|\"variable\")|\"Programmers\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"ofte¬n\"|\"achieved\")|\"clarity\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"loßng\"|\"are\")|\"use\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"the\"|\"jumped\")|\"ladπle\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fence\"|\"fence\")|\"deønse\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fence\"|\"the\")|\"bo´´x\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"mistake\"|\"variable\")|\"Grammers\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"often\"|\"achieved\")|\"Cl®arity\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"long\"|\"are\")|\"US†A\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"these\"|\"jumpeds\")|\"lazys\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fe∂nces\"|\"fßences\")|\"fencåes\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fences\"|\"thes\")|\"foxs\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"mßistakes\"|\"variabl∂es\")|\"Proƒgrammering\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"oftens\"|\"achi∫eveds\")|\"claritys\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"longs\"|\"ares\")|\"uses\"", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo&bar)&baz => (foo and bar) and baz
	//
	Describe("expression with and operator (A.B).C with double quotes", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"the\"&\"jumped\")&\"lazy\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fence\"&\"fence\")&\"fence\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fence\"&\"The\")&\"fox\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"mistake\"&\"variable\")&\"Programmers\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"often\"&\"achieved\")&\"clarity\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"long\"&\"are\")&\"use\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"the\"&\"zumped\")&\"lazy\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fence\"&\"zence\")&\"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fence\"&\"zhe\")&\"fox\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"mistake\"&\"zariable\")&\"Programmers\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"often\"&\"zchieved\")&∆√∫", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"long\"&\"zre\")&\"use\"", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"th¬kje\"&\"jumped\")&\"lazy\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fkjence\"&\"fence\")&\"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fenjce\"&\"the\")&\"fox\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"mis˚take\"&\"variable\")&\"Programmers\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"ofte¬n\"&\"achieved\")&\"clarity\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"loßng\"&\"are\")&\"use\"", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"the\"&\"jumped\")&\"ladπle\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fence\"&\"fence\")&\"deønse\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fence\"&\"the\")&\"bo´´x\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"mistake\"&\"variable\")&\"Grammers\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"often\"&\"achieved\")&\"Cl®arity\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"long\"&\"are\")&\"US†A\"", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"these\"&\"jumpeds\")&\"lazys\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fe∂nces\"&\"fßences\")&\"fencåes\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fences\"&\"thes\")&\"foxs\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"mßistakes\"&\"variabl∂es\")&\"Proƒgrammering\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"oftens\"&\"achi∫eveds\")&\"claritys\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"longs\"&\"ares\")&\"uses\"", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// ((foo&bar)&baz) => ((foo and bar) and baz)
	//
	Describe("expression with and operator ((A.B).C)", func() {
		Context("in line with A and B and C)", func() {
			It("should be false)", func() {
				Expect(SearchString("((\"fence\"&\"fence\")&\"fence\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("((\"the\"&\"jumped\")&\"lazy\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("((\"fence\"&\"The\")&\"fox\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("((\"mistake\"&\"variable\")&\"Programmers\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("((\"often\"&\"achieved\")&\"clarity\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("((\"long\"&\"are\")&\"use\")", longTargetProse)).To(Equal(true))
			})
		})

		Context("(in line with A and B and not C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((\"the\"&\"zumped\")&\"lazy\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fence\"&\"zence\")&\"fence\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fence\"&\"zhe\")&\"fox\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"mistake\"&\"zariable\")&\"Programmers\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"often\"&\"zchieved\")&∆√∫)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"long\"&\"zre\")&\"use\")", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with A and not B and C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((\"th¬kje\"&\"jumped\")&\"lazy\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fkjence\"&\"fence\")&\"fence\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fenjce\"&\"the\")&\"fox\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"mis˚take\"&\"variable\")&\"Programmers\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"ofte¬n\"&\"achieved\")&\"clarity\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"loßng\"&\"are\")&\"use\")", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with not A and B and C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((\"the\"&\"jumped\")&\"ladπle\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fence\"&\"fence\")&\"deønse\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fence\"&\"the\")&\"bo´´x\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"mistake\"&\"variable\")&\"Grammers\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"often\"&\"achieved\")&\"Cl®arity\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"long\"&\"are\")&\"US†A\")", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with not A and not B and not C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((\"these\"&\"jumpeds\")&\"lazys\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fe∂nces\"&\"fßences\")&\"fencåes\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fences\"&\"thes\")&\"foxs\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"mßistakes\"&\"variabl∂es\")&\"Proƒgrammering\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"oftens\"&\"achi∫eveds\")&\"claritys\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"longs\"&\"ares\")&\"uses\")", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// foo|!(bar|!baz) => foo or not (bar or not baz)
	//
	Describe("expression with or-not operator A+(B̅+C̅) with double quotes", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("\"lazy\"|!(\"the\"|!\"jumped\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"fence\"|!(\"fence\"|!\"fence\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"fox\"|!(\"fence\"|!\"The\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"Programmers\"|!(\"mistake\"|!\"variable\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"clarity\"|!(\"often\"|!\"achieved\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"use\"|!(\"long\"|!\"are\")", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("\"lazy\"|!(\"the\"|!\"zumped\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"fence\"|!(\"fence\"|!\"zence\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"fox\"|!(\"fence\"|!\"zhe\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"Programmers\"|!(\"mistake\"|!\"zariable\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("∆√∫|!(\"often\"|!\"zchieved\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"use\"|!(\"long\"|!\"zre\")", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("\"lazy\"|!(\"th¬kje\"|!\"jumped\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"fence\"|!(\"fkjence\"|!\"fence\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"fox\"|!(\"fenjce\"|!\"the\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"Programmers\"|!(\"mis˚take\"|!\"variable\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"clarity\"|!(\"ofte¬n\"|!\"achieved\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"use\"|!(\"loßng\"|!\"are\")", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("\"ladπle\"|!(\"the\"|!\"jumped\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"deønse\"|!(\"fence\"|!\"fence\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"bo´´x\"|!(\"fence\"|!\"the\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"Grammers\"|!(\"mistake\"|!\"variable\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"Cl®arity\"|!(\"often\"|!\"achieved\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"US†A\"|!(\"long\"|!\"are\")", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("\"lazys\"|!(\"these\"|!\"jumpeds\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"fencåes\"|!(\"fe∂nces\"|!\"fßences\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"foxs\"|!(\"fences\"|!\"thes\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"Proƒgrammering\"|!(\"mßistakes\"|!\"variabl∂es\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"claritys\"|!(\"oftens\"|!\"achi∫eveds\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"uses\"|!(\"longs\"|!\"ares\")", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo|!bar)|!baz => (foo or not bar) or not baz
	//
	Describe("expression with or not operator (A+B̅)+C̅ with double quotes", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"the\"|!\"jumped\")|!\"lazy\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fence\"|!\"fence\")|!\"fence\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fence\"|!\"The\")|!\"fox\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"mistake\"|!\"variable\")|!\"Programmers\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"often\"|!\"achieved\")|!\"clarity\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"long\"|!\"are\")|!\"use\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"the\"|!\"jumped\")|!\"zazy\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fence\"|!\"fence\")|!\"zence\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fence\"|!\"the\")|!\"zox\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"mistake\"|!\"variable\")|!\"zrogrammers\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"often\"|!\"achieved\")|!\"zlap\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"long\"|!\"are\")|!\"zuse\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"the\"|!\"jomped\")|!\"lazy\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fence\"|!\"fonce\")|!\"fence\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fence\"|!\"tho\")|!\"fox\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"mistake\"|!\"voriable\")|!\"Programmers\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"often\"|!\"ochieved\")|!\"clarity\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"long\"|!\"ore\")|!\"use\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"thez\"|!\"jumped\")|!\"lazy\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fencez\"|!\"fence\")|!\"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fencez\"|!\"the\")|!\"fox\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"mistakez\"|!\"variable\")|!\"Programmers\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"oftenz\"|!\"achieved\")|!\"clarity\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"longz\"|!\"are\")|!\"use\"", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"these\"|!\"jumpeds\")|!\"lazys\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fe∂nces\"|!\"fßences\")|!\"fencåes\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"fences\"|!\"thes\")|!\"foxs\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(\"mßistakes\"|!\"variabl∂es\")|!\"Proƒgrammering\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"oftens\"|!\"achi∫eveds\")|!\"claritys\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("(\"longs\"|!\"ares\")|!\"uses\"", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// (foo&!bar)&!baz => (foo and not bar) and not baz
	//
	Describe("expression with and-not operator (A.B̅).C̅ with double quotes", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"the\"&!\"jumped\")&!\"lazy\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fence\"&!\"fence\")&!\"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fence\"&!\"The\")&!\"fox\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"mistake\"&!\"variable\")&!\"Programmers\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"often\"&!\"achieved\")&!\"clarity\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"long\"&!\"are\")&!\"use\"", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"the\"&!\"zumped\")&!\"lazy\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fence\"&!\"zence\")&!\"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fence\"&!\"zhe\")&!\"fox\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"mistake\"&!\"zariable\")&!\"Programmers\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"often\"&!\"zchieved\")&!∆√∫", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"long\"&!\"zre\")&!\"use\"", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"th¬kje\"&!\"jumped\")&!\"lazy\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fkjence\"&!\"fence\")&!\"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fenjce\"&!\"the\")&!\"fox\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"mis˚take\"&!\"variable\")&!\"Programmers\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"ofte¬n\"&!\"achieved\")&!\"clarity\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"loßng\"&!\"are\")&!\"use\"", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"the\"&!\"jumped\")&!\"ladπle\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fence\"&!\"fence\")&!\"deønse\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fence\"&!\"the\")&!\"bo´´x\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"mistake\"&!\"variable\")&!\"Grammers\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"often\"&!\"achieved\")&!\"Cl®arity\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"long\"&!\"are\")&!\"US†A\"", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("(\"these\"&!\"jumpeds\")&!\"lazys\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fe∂nces\"&!\"fßences\")&!\"fencåes\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"fences\"&!\"thes\")&!\"foxs\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(\"mßistakes\"&!\"variabl∂es\")&!\"Proƒgrammering\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"oftens\"&!\"achi∫eveds\")&!\"claritys\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("(\"longs\"&!\"ares\")&!\"uses\"", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// ((foo&!bar)&!baz) => ((foo and not bar) and not baz)
	//
	Describe("expression with and-not operator ((A.B̅).C̅) with double quotes", func() {
		Context("in line with A and B and C)", func() {
			It("should be false)", func() {
				Expect(SearchString("((\"fence\"&!\"fence\")&!\"fence\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"the\"&!\"jumped\")&!\"lazy\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fence\"&!\"The\")&!\"fox\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"mistake\"&!\"variable\")&!\"Programmers\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"often\"&!\"achieved\")&!\"clarity\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"long\"&!\"are\")&!\"use\")", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with A and B and not C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((\"the\"&!\"zumped\")&!\"lazy\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fence\"&!\"zence\")&!\"fence\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fence\"&!\"zhe\")&!\"fox\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"mistake\"&!\"zariable\")&!\"Programmers\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"often\"&!\"zchieved\")&!∆√∫)", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"long\"&!\"zre\")&!\"use\")", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with A and not B and C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((\"th¬kje\"&!\"jumped\")&!\"lazy\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fkjence\"&!\"fence\")&!\"fence\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fenjce\"&!\"the\")&!\"fox\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"mis˚take\"&!\"variable\")&!\"Programmers\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"ofte¬n\"&!\"achieved\")&!\"clarity\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"loßng\"&!\"are\")&!\"use\")", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with not A and B and C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((\"the\"&!\"jumped\")&!\"ladπle\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fence\"&!\"fence\")&!\"deønse\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fence\"&!\"the\")&!\"bo´´x\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"mistake\"&!\"variable\")&!\"Grammers\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"often\"&!\"achieved\")&!\"Cl®arity\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"long\"&!\"are\")&!\"US†A\")", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with not A and not B and not C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("((\"these\"&!\"jumpeds\")&!\"lazys\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fe∂nces\"&!\"fßences\")&!\"fencåes\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"fences\"&!\"thes\")&!\"foxs\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("((\"mßistakes\"&!\"variabl∂es\")&!\"Proƒgrammering\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"oftens\"&!\"achi∫eveds\")&!\"claritys\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("((\"longs\"&!\"ares\")&!\"uses\")", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// "foo &! bar" => foo and not bar
	//
	Describe("expression with and-not operator (A.B̅) with double quotes", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("\"lazy\" &! \"fox\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"The\" &! \"over\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"jumped\" &! \"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"Programmers\" &! \"mistake\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("∆√∫ &! \"brevity\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"mistake\" &! \"clarity\"", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("\"lazy\" &! \"foxy\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"The\" &! \"overt\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"jumped\" &! \"fencing\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"Programmers\" &! \"mistook\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("∆√∫ &! \"brave\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"mistake\" &! \"claire\"", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with not A and B", func() {
			It("should be false", func() {
				Expect(SearchString("\"lazyish\" &! \"fox\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"Then\" &! \"over\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"jumpy\" &! \"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"Prop\" &! \"mistake\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"through-fair\" &! \"brevity\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"mistakes\" &! \"clarity\"", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("\"ladle\" &    ! \"foxy\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"Thorn\" &   ! \"overt\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"jam\" & ! \"fencing\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"Frog\" &      ! \"mistook\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"Clark\" &    ! \"brave\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"Park\" &   ! \"claire\"", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// "foo |! bar" => foo or not bar
	//
	Describe("expression with 'or not' operator (A+B̅) with double quotes", func() {
		Context("in line with A and B", func() {
			It("should be true", func() {
				Expect(SearchString("\"lazy\" |! \"fox\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"the\" |! \"fence\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"over\" |! \"the\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"are\" |! \"brevity\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"is\" |! \"is\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"encouraged\" |! \"regardless\"", longTargetProse)).To(Equal(true))

				Expect(SearchString("\"lazy\" |  ! \"fox\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"the\" | ! \"fence\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"over\" |  ! \"the\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"are\" | ! \"brevity\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"is\" | ! \"is\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"encouraged\" |     ! \"regardless\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B with double quotes", func() {
			It("should be true", func() {
				Expect(SearchString("\"lazy\" |! \"abc\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"the\" |! \"fend\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"over\" |! \"they\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"are\" |! \"bravado\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"is\" |! \"si\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"encouraged\" | \"regarding\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with B and not A with double quotes", func() {
			It("should be false", func() {
				Expect(SearchString("\"laser\" |! \"fox\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"than\" |! \"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"overtly\" |! \"the\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("\"arsenal\" |! \"brevity\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"island\" |! \"is\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("\"encouraging\" |! \"regardless\"", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("\"laser\" |! \"abc\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"than\" |! \"fend\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"overtly\" |! \"they\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("\"arsenal\" |! \"brevado\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"island\" |! \"si\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("\"encouraging\" |! \"regarding\"", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo &! bar" => not foo and not bar
	//
	Describe("expression with and-not operator (A̅.B̅) with double quotes", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("!\"lazy\" &! \"fox\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"The\" &! \"over\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"jumped\" &! \"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"Programmers\" &! \"mistake\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("!\"∆√∫\" &! \"brevity\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("!\"mistake\" &! \"clarity\"", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("!\"lazy\" &! \"foxy\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"The\" &! \"overt\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"jumped\" &! \"fencing\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"Programmers\" &! \"mistook\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("!\"∆√∫\" &! \"brave\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("!\"mistake\" &! \"claire\"", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and B", func() {
			It("should be false", func() {
				Expect(SearchString("!\"lazyish\" &! \"fox\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"Then\" &! \"over\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"jumpy\" &! \"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"Prop\" &! \"mistake\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("!\"through-fair\" &! \"brevity\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("!\"mistakes\" &! \"clarity\"", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("!\"ladle\" &    ! \"foxy\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"Thorn\" &   ! \"overt\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"jam\" & ! \"fencing\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"Frog\" &      ! \"mistook\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("!\"Clark\" &    ! \"brave\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("!\"Park\" &   ! \"claire\"", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo |! bar" => not foo or not bar
	//
	Describe("expression with 'or not' operator (A̅+B̅) with double quotes", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("!\"lazy\" |! \"fox\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"the\" |! \"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"over\" |! \"the\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"are\" |! \"brevity\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("!\"is\" |! \"is\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("!\"encouraged\" |! \"regardless\"", longTargetProse)).To(Equal(false))

				// with extra spaces
				Expect(SearchString("!\"lazy\" |  ! \"fox\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"the\" | ! \"fence\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"over\" |  ! \"the\"", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!\"are\" | ! \"brevity\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("!\"is\" | ! \"is\"", longTargetProse)).To(Equal(false))
				Expect(SearchString("!\"encouraged\" |     ! \"regardless\"", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("!\"lazy\" |! \"abc\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"the\" |! \"fend\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"over\" |! \"they\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"are\" |! \"bravado\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("!\"is\" |! \"si\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("!\"encouraged\" |! \"regarding\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("!\"laser\" |! \"fox\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"than\" |! \"fence\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"overtly\" |! \"the\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"arsenal\" |! \"brevity\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("!\"island\" |! \"is\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("!\"encouraging\" |! \"regardless\"", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("!\"laser\" |! \"abc\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"than\" |! \"fend\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"overtly\" |! \"they\"", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!\"arsenal\" |! \"brevado\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("!\"island\" |! \"si\"", longTargetProse)).To(Equal(true))
				Expect(SearchString("!\"encouraging\" |! \"regarding\"", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo &! bar" => not foo and not bar
	//
	Describe("expression with and-not operator (A̅.B̅) with double quotes", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("(!\"lazy\" &! \"fox\"))", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"The\" &! \"over\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"jumped\" &! \"fence\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"Programmers\" &! \"mistake\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"∆√∫\" &! \"brevity\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"mistake\" &! \"clarity\")", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("(!\"lazy\" &! \"foxy\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"The\" &! \"overt\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"jumped\" &! \"fencing\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"Programmers\" &! \"mistook\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"∆√∫\" &! \"brave\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"mistake\" &! \"claire\")", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and B", func() {
			It("should be false", func() {
				Expect(SearchString("(!\"lazyish\" &! \"fox\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"Then\" &! \"over\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"jumpy\" &! \"fence\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"Prop\" &! \"mistake\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"through-fair\" &! \"brevity\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"mistakes\" &! \"clarity\")", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("(!\"ladle\" &    ! \"foxy\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"Thorn\" &   ! \"overt\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"jam\" & ! \"fencing\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"Frog\" &      ! \"mistook\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"Clark\" &    ! \"brave\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"Park\" &   ! \"claire\")", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo |! bar" => not foo or not bar
	//
	Describe("expression with 'or not' operator (A̅+B̅) with double quotes", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("(!\"lazy\" |! \"fox\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"the\" |! \"fence\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"over\" |! \"the\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"are\" |! \"brevity\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"is\" |! \"is\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"encouraged\" |! \"regardless\")", longTargetProse)).To(Equal(false))

				// with extra spaces
				Expect(SearchString("(!\"lazy\" |  ! \"fox\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"the\" | ! \"fence\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"over\" |  ! \"the\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"are\" | ! \"brevity\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"is\" | ! \"is\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"encouraged\" |     ! \"regardless\")", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("(!\"lazy\" |! \"abc\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"the\" |! \"fend\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"over\" |! \"they\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"are\" |! \"bravado\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"is\" |! \"si\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"encouraged\" |! \"regarding\")", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("(!\"laser\" |! \"fox\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"than\" |! \"fence\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"overtly\" |! \"the\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"arsenal\" |! \"brevity\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"island\" |! \"is\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"encouraging\" |! \"regardless\")", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("(!\"laser\" |! \"abc\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"than\" |! \"fend\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"overtly\" |! \"they\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"arsenal\" |! \"brevado\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"island\" |! \"si\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"encouraging\" |! \"regarding\")", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo &bar" => not foo and bar
	//
	Describe("expression with and-not operator (A̅.B) with double quotes", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("(!\"lazy\" &\"fox\"))", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"The\" &\"over\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"jumped\" &\"fence\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"Programmers\" &\"mistake\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"∆√∫\" &\"brevity\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"mistake\" &\"clarity\")", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("(!\"lazy\" &\"foxy\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"The\" &\"overt\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"jumped\" &\"fencing\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"Programmers\" &\"mistook\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"∆√∫\" &\"brave\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"mistake\" &\"claire\")", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and B", func() {
			It("should be true", func() {
				Expect(SearchString("(!\"lazyish\" &\"fox\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"Then\" &\"over\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"jumpy\" &\"fence\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"Prop\" &\"mistake\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"through-fair\" &\"brevity\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"mistakes\" &\"clarity\")", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("(!\"ladle\" &    \"foxy\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"Thorn\" &   \"overt\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"jam\" &  \"fencing\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"Frog\" &       \"mistook\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"Clark\" &     \"brave\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"Park\" &    \"claire\")", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// "!foo | bar" => not foo or bar
	//
	Describe("expression with 'or not' operator (A̅+B) with double quotes", func() {
		Context("in line with A and B", func() {
			It("should be true", func() {
				Expect(SearchString("(!\"lazy\" | \"fox\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"the\" | \"fence\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"over\" | \"the\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"are\" | \"brevity\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"is\" | \"is\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"encouraged\" | \"regardless\")", longTargetProse)).To(Equal(true))

				// with extra spaces
				Expect(SearchString("(!\"lazy\" |   \"fox\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"the\" |  \"fence\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"over\" |   \"the\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"are\" |  \"brevity\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"is\" |  \"is\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"encouraged\" |      \"regardless\")", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("(!\"lazy\" | \"abc\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"the\" | \"fend\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"over\" | \"they\")", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"are\" | \"bravado\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"is\" | \"si\")", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!\"encouraged\" | \"regarding\")", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("(!\"laser\" | \"fox\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"than\" | \"fence\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"overtly\" | \"the\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"arsenal\" | \"brevity\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"island\" | \"is\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"encouraging\" | \"regardless\")", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with not B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("(!\"laser\" | \"abc\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"than\" | \"fend\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"overtly\" | \"they\")", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"arsenal\" | \"brevado\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"island\" | \"si\")", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!\"encouraging\" | \"regarding\")", longTargetProse)).To(Equal(true))
			})
		})
	})

	Describe("complex example with single quotes", func() {
		Context(" ", func() {
			It(" ", func() {
				Expect(SearchString("!((('lazy' & !'dog'))) | (((('lazy' & 'dog')))) | !(((('lazy' | 'dog')))) | ((((!'lazy' & 'dog'))))", shortTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo) => (foo)
	//
	Describe("singular expression (A) with single quotes", func() {
		Context("in line with A", func() {
			It("should be true", func() {
				Expect(SearchString("'The'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'jumped'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'fence'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'Programmers'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'a'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'variable'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line without A", func() {
			It("should be false", func() {
				Expect(SearchString("'Thee'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'jumper'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'felt'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'Programmed'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'alt'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'variety'", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo) => (foo)
	//
	Describe("singular expression (A̅) with single quotes", func() {
		Context("in line with A", func() {
			It("should be false", func() {
				Expect(SearchString("!'The'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'jumped'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'Programmers'", longTargetProse)).To(Equal(false))
				Expect(SearchString("!'a'", longTargetProse)).To(Equal(false))
				Expect(SearchString("!'variable'", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line without A", func() {
			It("should be true", func() {
				Expect(SearchString("!'Thee'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'jumper'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'felt'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'Programmed'", longTargetProse)).To(Equal(true))
				Expect(SearchString("!'alt'", longTargetProse)).To(Equal(true))
				Expect(SearchString("!'variety'", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// foo | bar => foo or bar
	//
	Describe("expression with or operator (A+B) with single quotes", func() {
		Context("in line with A and B", func() {
			It("should be true", func() {
				Expect(SearchString("'lazy' | 'fox'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'the' | 'fence'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'over' | 'the'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'are' | 'brevity'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'is' | 'is'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'encouraged' | 'regardless'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("'lazy' | 'abc'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'the' | 'fend'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'over' | 'they'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'are' | 'bravado'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'is' | 'si'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'encouraged' | 'regarding'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("'laser' | 'fox'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'than' | 'fence'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'overtly' | 'the'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'arsenal' | 'brevity'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'island' | 'is'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'encouraging' | 'regardless'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not B and not A", func() {
			It("should be false", func() {
				Expect(SearchString("'laser' | 'abc'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'than' | 'fend'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'overtly' | 'they'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'arsenal' | 'brevado'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'island' | 'si'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'encouraging' | 'regarding'", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// foo & bar => foo and bar
	//
	Describe("expression with and operator (A.B) with single quotes", func() {
		Context("in line with A and B", func() {
			It("should be true", func() {
				Expect(SearchString("'lazy' & 'fox'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'The' & 'over'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'jumped' & 'fence'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'Programmers' & 'mistake'", longTargetProse)).To(Equal(true))
				Expect(SearchString("∆√∫ & 'brevity'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'mistake' & 'clarity'", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("'lazy' & 'foxy'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'The' & 'overt'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'jumped' & 'fencing'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'Programmers' & 'mistook'", longTargetProse)).To(Equal(false))
				Expect(SearchString("∆√∫ & 'brave'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'mistake' & 'claire'", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and B", func() {
			It("should be false", func() {
				Expect(SearchString("'lazyish' & 'fox'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'Then' & 'over'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'jumpy' & 'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'Programmered' & 'mistake'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'through-fair' & 'brevity'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'mistakes' & 'clarity'", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("'ladle' & 'foxy'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'Thorn' & 'overt'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'jam' & 'fencing'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'Frog' & 'mistook'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'Clark' & 'brave'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'Park' & 'claire'", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// foo|(bar|baz) => foo or (bar or baz)
	//
	Describe("expression with and operator A+(B+C) with single quotes", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("'lazy'|('the'|'jumped')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'fence'|('fence'|'fence')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'fox'|('fence'|'The')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'Programmers'|('mistake'|'variable')", longTargetProse)).To(Equal(true))
				Expect(SearchString("'clarity'|('often'|'achieved')", longTargetProse)).To(Equal(true))
				Expect(SearchString("'use'|('long'|'are')", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("'lazy'|('the'|'zumped')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'fence'|('fence'|'zence')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'fox'|('fence'|'zhe')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'Programmers'|('mistake'|'zariable')", longTargetProse)).To(Equal(true))
				Expect(SearchString("∆√∫|('often'|'zchieved')", longTargetProse)).To(Equal(true))
				Expect(SearchString("'use'|('long'|'zre')", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("'lazy'|('th¬kje'|'jumped')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'fence'|('fkjence'|'fence')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'fox'|('fenjce'|'the')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'Programmers'|('mis˚take'|'variable')", longTargetProse)).To(Equal(true))
				Expect(SearchString("'clarity'|('ofte¬n'|'achieved')", longTargetProse)).To(Equal(true))
				Expect(SearchString("'use'|('loßng'|'are')", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("'ladπle'|('the'|'jumped')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'deønse'|('fence'|'fence')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'bo´´x'|('fence'|'the')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'Grammers'|('mistake'|'variable')", longTargetProse)).To(Equal(true))
				Expect(SearchString("'Cl®arity'|('often'|'achieved')", longTargetProse)).To(Equal(true))
				Expect(SearchString("'US†A'|('long'|'are')", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("'lazys'|('these'|'jumpeds')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'fencåes'|('fe∂nces'|'fßences')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'foxs'|('fences'|'thes')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'Proƒgrammering'|('mßistakes'|'variabl∂es')", longTargetProse)).To(Equal(false))
				Expect(SearchString("'claritys'|('oftens'|'achi∫eveds')", longTargetProse)).To(Equal(false))
				Expect(SearchString("'uses'|('longs'|'ares')", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo|bar)|baz => (foo or bar) or baz
	//
	Describe("expression with and operator (A+B)+C with single quotes", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("('the'|'jumped')|'lazy'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fence'|'fence')|'fence'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fence'|'The')|'fox'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('mistake'|'variable')|'Programmers'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('often'|'achieved')|'clarity'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('long'|'are')|'use'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("('the'|'zumped')|'lazy'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fence'|'zence')|'fence'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fence'|'zhe')|'fox'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('mistake'|'zariable')|'Programmers'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('often'|'zchieved')|∆√∫", longTargetProse)).To(Equal(true))
				Expect(SearchString("('long'|'zre')|'use'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("('th¬kje'|'jumped')|'lazy'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fkjence'|'fence')|'fence'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fenjce'|'the')|'fox'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('mis˚take'|'variable')|'Programmers'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('ofte¬n'|'achieved')|'clarity'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('loßng'|'are')|'use'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("('the'|'jumped')|'ladπle'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fence'|'fence')|'deønse'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fence'|'the')|'bo´´x'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('mistake'|'variable')|'Grammers'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('often'|'achieved')|'Cl®arity'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('long'|'are')|'US†A'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("('these'|'jumpeds')|'lazys'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fe∂nces'|'fßences')|'fencåes'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fences'|'thes')|'foxs'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('mßistakes'|'variabl∂es')|'Proƒgrammering'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('oftens'|'achi∫eveds')|'claritys'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('longs'|'ares')|'uses'", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo&bar)&baz => (foo and bar) and baz
	//
	Describe("expression with and operator (A.B).C with single quotes", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("('the'&'jumped')&'lazy'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fence'&'fence')&'fence'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fence'&'The')&'fox'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('mistake'&'variable')&'Programmers'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('often'&'achieved')&'clarity'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('long'&'are')&'use'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("('the'&'zumped')&'lazy'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fence'&'zence')&'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fence'&'zhe')&'fox'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('mistake'&'zariable')&'Programmers'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('often'&'zchieved')&∆√∫", longTargetProse)).To(Equal(false))
				Expect(SearchString("('long'&'zre')&'use'", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("('th¬kje'&'jumped')&'lazy'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fkjence'&'fence')&'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fenjce'&'the')&'fox'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('mis˚take'&'variable')&'Programmers'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('ofte¬n'&'achieved')&'clarity'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('loßng'&'are')&'use'", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("('the'&'jumped')&'ladπle'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fence'&'fence')&'deønse'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fence'&'the')&'bo´´x'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('mistake'&'variable')&'Grammers'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('often'&'achieved')&'Cl®arity'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('long'&'are')&'US†A'", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("('these'&'jumpeds')&'lazys'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fe∂nces'&'fßences')&'fencåes'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fences'&'thes')&'foxs'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('mßistakes'&'variabl∂es')&'Proƒgrammering'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('oftens'&'achi∫eveds')&'claritys'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('longs'&'ares')&'uses'", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// ((foo&bar)&baz) => ((foo and bar) and baz)
	//
	Describe("expression with and operator ((A.B).C)", func() {
		Context("in line with A and B and C)", func() {
			It("should be false)", func() {
				Expect(SearchString("(('fence'&'fence')&'fence')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(('the'&'jumped')&'lazy')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(('fence'&'The')&'fox')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(('mistake'&'variable')&'Programmers')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(('often'&'achieved')&'clarity')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(('long'&'are')&'use')", longTargetProse)).To(Equal(true))
			})
		})

		Context("(in line with A and B and not C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("(('the'&'zumped')&'lazy')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fence'&'zence')&'fence')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fence'&'zhe')&'fox')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('mistake'&'zariable')&'Programmers')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('often'&'zchieved')&∆√∫)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('long'&'zre')&'use')", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with A and not B and C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("(('th¬kje'&'jumped')&'lazy')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fkjence'&'fence')&'fence')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fenjce'&'the')&'fox')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('mis˚take'&'variable')&'Programmers')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('ofte¬n'&'achieved')&'clarity')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('loßng'&'are')&'use')", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with not A and B and C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("(('the'&'jumped')&'ladπle')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fence'&'fence')&'deønse')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fence'&'the')&'bo´´x')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('mistake'&'variable')&'Grammers')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('often'&'achieved')&'Cl®arity')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('long'&'are')&'US†A')", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with not A and not B and not C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("(('these'&'jumpeds')&'lazys')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fe∂nces'&'fßences')&'fencåes')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fences'&'thes')&'foxs')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('mßistakes'&'variabl∂es')&'Proƒgrammering')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('oftens'&'achi∫eveds')&'claritys')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('longs'&'ares')&'uses')", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// foo|!(bar|!baz) => foo or not (bar or not baz)
	//
	Describe("expression with or-not operator A+(B̅+C̅) with single quotes", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("'lazy'|!('the'|!'jumped')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'fence'|!('fence'|!'fence')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'fox'|!('fence'|!'The')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'Programmers'|!('mistake'|!'variable')", longTargetProse)).To(Equal(true))
				Expect(SearchString("'clarity'|!('often'|!'achieved')", longTargetProse)).To(Equal(true))
				Expect(SearchString("'use'|!('long'|!'are')", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("'lazy'|!('the'|!'zumped')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'fence'|!('fence'|!'zence')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'fox'|!('fence'|!'zhe')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'Programmers'|!('mistake'|!'zariable')", longTargetProse)).To(Equal(true))
				Expect(SearchString("∆√∫|!('often'|!'zchieved')", longTargetProse)).To(Equal(true))
				Expect(SearchString("'use'|!('long'|!'zre')", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("'lazy'|!('th¬kje'|!'jumped')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'fence'|!('fkjence'|!'fence')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'fox'|!('fenjce'|!'the')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'Programmers'|!('mis˚take'|!'variable')", longTargetProse)).To(Equal(true))
				Expect(SearchString("'clarity'|!('ofte¬n'|!'achieved')", longTargetProse)).To(Equal(true))
				Expect(SearchString("'use'|!('loßng'|!'are')", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("'ladπle'|!('the'|!'jumped')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'deønse'|!('fence'|!'fence')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'bo´´x'|!('fence'|!'the')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'Grammers'|!('mistake'|!'variable')", longTargetProse)).To(Equal(false))
				Expect(SearchString("'Cl®arity'|!('often'|!'achieved')", longTargetProse)).To(Equal(false))
				Expect(SearchString("'US†A'|!('long'|!'are')", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("'lazys'|!('these'|!'jumpeds')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'fencåes'|!('fe∂nces'|!'fßences')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'foxs'|!('fences'|!'thes')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'Proƒgrammering'|!('mßistakes'|!'variabl∂es')", longTargetProse)).To(Equal(false))
				Expect(SearchString("'claritys'|!('oftens'|!'achi∫eveds')", longTargetProse)).To(Equal(false))
				Expect(SearchString("'uses'|!('longs'|!'ares')", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// (foo|!bar)|!baz => (foo or not bar) or not baz
	//
	Describe("expression with or not operator (A+B̅)+C̅ with single quotes", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("('the'|!'jumped')|!'lazy'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fence'|!'fence')|!'fence'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fence'|!'The')|!'fox'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('mistake'|!'variable')|!'Programmers'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('often'|!'achieved')|!'clarity'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('long'|!'are')|!'use'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("('the'|!'jumped')|!'zazy'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fence'|!'fence')|!'zence'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fence'|!'the')|!'zox'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('mistake'|!'variable')|!'zrogrammers'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('often'|!'achieved')|!'zlap'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('long'|!'are')|!'zuse'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("('the'|!'jomped')|!'lazy'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fence'|!'fonce')|!'fence'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fence'|!'tho')|!'fox'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('mistake'|!'voriable')|!'Programmers'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('often'|!'ochieved')|!'clarity'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('long'|!'ore')|!'use'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("('thez'|!'jumped')|!'lazy'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fencez'|!'fence')|!'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fencez'|!'the')|!'fox'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('mistakez'|!'variable')|!'Programmers'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('oftenz'|!'achieved')|!'clarity'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('longz'|!'are')|!'use'", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("('these'|!'jumpeds')|!'lazys'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fe∂nces'|!'fßences')|!'fencåes'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('fences'|!'thes')|!'foxs'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("('mßistakes'|!'variabl∂es')|!'Proƒgrammering'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('oftens'|!'achi∫eveds')|!'claritys'", longTargetProse)).To(Equal(true))
				Expect(SearchString("('longs'|!'ares')|!'uses'", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// (foo&!bar)&!baz => (foo and not bar) and not baz
	//
	Describe("expression with and-not operator (A.B̅).C̅ with single quotes", func() {
		Context("in line with A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("('the'&!'jumped')&!'lazy'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fence'&!'fence')&!'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fence'&!'The')&!'fox'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('mistake'&!'variable')&!'Programmers'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('often'&!'achieved')&!'clarity'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('long'&!'are')&!'use'", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("('the'&!'zumped')&!'lazy'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fence'&!'zence')&!'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fence'&!'zhe')&!'fox'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('mistake'&!'zariable')&!'Programmers'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('often'&!'zchieved')&!∆√∫", longTargetProse)).To(Equal(false))
				Expect(SearchString("('long'&!'zre')&!'use'", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and not B and C", func() {
			It("should be false", func() {
				Expect(SearchString("('th¬kje'&!'jumped')&!'lazy'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fkjence'&!'fence')&!'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fenjce'&!'the')&!'fox'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('mis˚take'&!'variable')&!'Programmers'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('ofte¬n'&!'achieved')&!'clarity'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('loßng'&!'are')&!'use'", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and B and C", func() {
			It("should be false", func() {
				Expect(SearchString("('the'&!'jumped')&!'ladπle'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fence'&!'fence')&!'deønse'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fence'&!'the')&!'bo´´x'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('mistake'&!'variable')&!'Grammers'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('often'&!'achieved')&!'Cl®arity'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('long'&!'are')&!'US†A'", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not A and not B and not C", func() {
			It("should be false", func() {
				Expect(SearchString("('these'&!'jumpeds')&!'lazys'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fe∂nces'&!'fßences')&!'fencåes'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('fences'&!'thes')&!'foxs'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("('mßistakes'&!'variabl∂es')&!'Proƒgrammering'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('oftens'&!'achi∫eveds')&!'claritys'", longTargetProse)).To(Equal(false))
				Expect(SearchString("('longs'&!'ares')&!'uses'", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// ((foo&!bar)&!baz) => ((foo and not bar) and not baz)
	//
	Describe("expression with and-not operator ((A.B̅).C̅) with single quotes", func() {
		Context("in line with A and B and C)", func() {
			It("should be false)", func() {
				Expect(SearchString("(('fence'&!'fence')&!'fence')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('the'&!'jumped')&!'lazy')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fence'&!'The')&!'fox')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('mistake'&!'variable')&!'Programmers')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('often'&!'achieved')&!'clarity')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('long'&!'are')&!'use')", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with A and B and not C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("(('the'&!'zumped')&!'lazy')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fence'&!'zence')&!'fence')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fence'&!'zhe')&!'fox')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('mistake'&!'zariable')&!'Programmers')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('often'&!'zchieved')&!∆√∫)", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('long'&!'zre')&!'use')", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with A and not B and C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("(('th¬kje'&!'jumped')&!'lazy')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fkjence'&!'fence')&!'fence')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fenjce'&!'the')&!'fox')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('mis˚take'&!'variable')&!'Programmers')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('ofte¬n'&!'achieved')&!'clarity')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('loßng'&!'are')&!'use')", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with not A and B and C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("(('the'&!'jumped')&!'ladπle')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fence'&!'fence')&!'deønse')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fence'&!'the')&!'bo´´x')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('mistake'&!'variable')&!'Grammers')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('often'&!'achieved')&!'Cl®arity')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('long'&!'are')&!'US†A')", longTargetProse)).To(Equal(false))
			})
		})

		Context("(in line with not A and not B and not C)", func() {
			It("(should be false)", func() {
				Expect(SearchString("(('these'&!'jumpeds')&!'lazys')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fe∂nces'&!'fßences')&!'fencåes')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('fences'&!'thes')&!'foxs')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(('mßistakes'&!'variabl∂es')&!'Proƒgrammering')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('oftens'&!'achi∫eveds')&!'claritys')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(('longs'&!'ares')&!'uses')", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// "foo &! bar" => foo and not bar
	//
	Describe("expression with and-not operator (A.B̅) with single quotes", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("'lazy' &! 'fox'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'The' &! 'over'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'jumped' &! 'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'Programmers' &! 'mistake'", longTargetProse)).To(Equal(false))
				Expect(SearchString("∆√∫ &! 'brevity'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'mistake' &! 'clarity'", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("'lazy' &! 'foxy'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'The' &! 'overt'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'jumped' &! 'fencing'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'Programmers' &! 'mistook'", longTargetProse)).To(Equal(true))
				Expect(SearchString("∆√∫ &! 'brave'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'mistake' &! 'claire'", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with not A and B", func() {
			It("should be false", func() {
				Expect(SearchString("'lazyish' &! 'fox'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'Then' &! 'over'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'jumpy' &! 'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'Prop' &! 'mistake'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'through-fair' &! 'brevity'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'mistakes' &! 'clarity'", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("'ladle' &    ! 'foxy'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'Thorn' &   ! 'overt'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'jam' & ! 'fencing'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'Frog' &      ! 'mistook'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'Clark' &    ! 'brave'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'Park' &   ! 'claire'", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// "foo |! bar" => foo or not bar
	//
	Describe("expression with 'or not' operator (A+B̅) with single quotes", func() {
		Context("in line with A and B", func() {
			It("should be true", func() {
				Expect(SearchString("'lazy' |! 'fox'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'the' |! 'fence'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'over' |! 'the'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'are' |! 'brevity'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'is' |! 'is'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'encouraged' |! 'regardless'", longTargetProse)).To(Equal(true))

				Expect(SearchString("'lazy' |  ! 'fox'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'the' | ! 'fence'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'over' |  ! 'the'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'are' | ! 'brevity'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'is' | ! 'is'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'encouraged' |     ! 'regardless'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with A and not B with single quotes", func() {
			It("should be true", func() {
				Expect(SearchString("'lazy' |! 'abc'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'the' |! 'fend'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'over' |! 'they'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'are' |! 'bravado'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'is' |! 'si'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'encouraged' | 'regarding'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with B and not A with single quotes", func() {
			It("should be false", func() {
				Expect(SearchString("'laser' |! 'fox'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'than' |! 'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'overtly' |! 'the'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("'arsenal' |! 'brevity'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'island' |! 'is'", longTargetProse)).To(Equal(false))
				Expect(SearchString("'encouraging' |! 'regardless'", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with not B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("'laser' |! 'abc'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'than' |! 'fend'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'overtly' |! 'they'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("'arsenal' |! 'brevado'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'island' |! 'si'", longTargetProse)).To(Equal(true))
				Expect(SearchString("'encouraging' |! 'regarding'", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo &! bar" => not foo and not bar
	//
	Describe("expression with and-not operator (A̅.B̅) with single quotes", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("!'lazy' &! 'fox'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'The' &! 'over'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'jumped' &! 'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'Programmers' &! 'mistake'", longTargetProse)).To(Equal(false))
				Expect(SearchString("!'∆√∫' &! 'brevity'", longTargetProse)).To(Equal(false))
				Expect(SearchString("!'mistake' &! 'clarity'", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("!'lazy' &! 'foxy'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'The' &! 'overt'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'jumped' &! 'fencing'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'Programmers' &! 'mistook'", longTargetProse)).To(Equal(false))
				Expect(SearchString("!'∆√∫' &! 'brave'", longTargetProse)).To(Equal(false))
				Expect(SearchString("!'mistake' &! 'claire'", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and B", func() {
			It("should be false", func() {
				Expect(SearchString("!'lazyish' &! 'fox'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'Then' &! 'over'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'jumpy' &! 'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'Prop' &! 'mistake'", longTargetProse)).To(Equal(false))
				Expect(SearchString("!'through-fair' &! 'brevity'", longTargetProse)).To(Equal(false))
				Expect(SearchString("!'mistakes' &! 'clarity'", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("!'ladle' &    ! 'foxy'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'Thorn' &   ! 'overt'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'jam' & ! 'fencing'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'Frog' &      ! 'mistook'", longTargetProse)).To(Equal(true))
				Expect(SearchString("!'Clark' &    ! 'brave'", longTargetProse)).To(Equal(true))
				Expect(SearchString("!'Park' &   ! 'claire'", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo |! bar" => not foo or not bar
	//
	Describe("expression with 'or not' operator (A̅+B̅) with single quotes", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("!'lazy' |! 'fox'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'the' |! 'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'over' |! 'the'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'are' |! 'brevity'", longTargetProse)).To(Equal(false))
				Expect(SearchString("!'is' |! 'is'", longTargetProse)).To(Equal(false))
				Expect(SearchString("!'encouraged' |! 'regardless'", longTargetProse)).To(Equal(false))

				// with extra spaces
				Expect(SearchString("!'lazy' |  ! 'fox'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'the' | ! 'fence'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'over' |  ! 'the'", shortTargetProse)).To(Equal(false))
				Expect(SearchString("!'are' | ! 'brevity'", longTargetProse)).To(Equal(false))
				Expect(SearchString("!'is' | ! 'is'", longTargetProse)).To(Equal(false))
				Expect(SearchString("!'encouraged' |     ! 'regardless'", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("!'lazy' |! 'abc'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'the' |! 'fend'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'over' |! 'they'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'are' |! 'bravado'", longTargetProse)).To(Equal(true))
				Expect(SearchString("!'is' |! 'si'", longTargetProse)).To(Equal(true))
				Expect(SearchString("!'encouraged' |! 'regarding'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("!'laser' |! 'fox'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'than' |! 'fence'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'overtly' |! 'the'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'arsenal' |! 'brevity'", longTargetProse)).To(Equal(true))
				Expect(SearchString("!'island' |! 'is'", longTargetProse)).To(Equal(true))
				Expect(SearchString("!'encouraging' |! 'regardless'", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("!'laser' |! 'abc'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'than' |! 'fend'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'overtly' |! 'they'", shortTargetProse)).To(Equal(true))
				Expect(SearchString("!'arsenal' |! 'brevado'", longTargetProse)).To(Equal(true))
				Expect(SearchString("!'island' |! 'si'", longTargetProse)).To(Equal(true))
				Expect(SearchString("!'encouraging' |! 'regarding'", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo &! bar" => not foo and not bar
	//
	Describe("expression with and-not operator (A̅.B̅) with single quotes", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("(!'lazy' &! 'fox'))", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'The' &! 'over')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'jumped' &! 'fence')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'Programmers' &! 'mistake')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'∆√∫' &! 'brevity')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'mistake' &! 'clarity')", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("(!'lazy' &! 'foxy')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'The' &! 'overt')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'jumped' &! 'fencing')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'Programmers' &! 'mistook')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'∆√∫' &! 'brave')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'mistake' &! 'claire')", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and B", func() {
			It("should be false", func() {
				Expect(SearchString("(!'lazyish' &! 'fox')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'Then' &! 'over')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'jumpy' &! 'fence')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'Prop' &! 'mistake')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'through-fair' &! 'brevity')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'mistakes' &! 'clarity')", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("(!'ladle' &    ! 'foxy')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'Thorn' &   ! 'overt')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'jam' & ! 'fencing')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'Frog' &      ! 'mistook')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'Clark' &    ! 'brave')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'Park' &   ! 'claire')", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo |! bar" => not foo or not bar
	//
	Describe("expression with 'or not' operator (A̅+B̅) with single quotes", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("(!'lazy' |! 'fox')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'the' |! 'fence')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'over' |! 'the')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'are' |! 'brevity')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'is' |! 'is')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'encouraged' |! 'regardless')", longTargetProse)).To(Equal(false))

				// with extra spaces
				Expect(SearchString("(!'lazy' |  ! 'fox')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'the' | ! 'fence')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'over' |  ! 'the')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'are' | ! 'brevity')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'is' | ! 'is')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'encouraged' |     ! 'regardless')", longTargetProse)).To(Equal(false))
			})
		})

		Context("in line with A and not B", func() {
			It("should be true", func() {
				Expect(SearchString("(!'lazy' |! 'abc')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'the' |! 'fend')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'over' |! 'they')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'are' |! 'bravado')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'is' |! 'si')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'encouraged' |! 'regarding')", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("(!'laser' |! 'fox')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'than' |! 'fence')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'overtly' |! 'the')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'arsenal' |! 'brevity')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'island' |! 'is')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'encouraging' |! 'regardless')", longTargetProse)).To(Equal(true))
			})
		})

		Context("in line with not B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("(!'laser' |! 'abc')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'than' |! 'fend')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'overtly' |! 'they')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'arsenal' |! 'brevado')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'island' |! 'si')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'encouraging' |! 'regarding')", longTargetProse)).To(Equal(true))
			})
		})
	})

	//
	// "!foo &bar" => not foo and bar
	//
	Describe("expression with and-not operator (A̅.B) with single quotes", func() {
		Context("in line with A and B", func() {
			It("should be false", func() {
				Expect(SearchString("(!'lazy' &'fox'))", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'The' &'over')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'jumped' &'fence')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'Programmers' &'mistake')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'∆√∫' &'brevity')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'mistake' &'clarity')", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("(!'lazy' &'foxy')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'The' &'overt')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'jumped' &'fencing')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'Programmers' &'mistook')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'∆√∫' &'brave')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'mistake' &'claire')", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with not A and B", func() {
			It("should be true", func() {
				Expect(SearchString("(!'lazyish' &'fox')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'Then' &'over')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'jumpy' &'fence')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'Prop' &'mistake')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'through-fair' &'brevity')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'mistakes' &'clarity')", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with not A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("(!'ladle' &    'foxy')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'Thorn' &   'overt')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'jam' &  'fencing')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'Frog' &       'mistook')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'Clark' &     'brave')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'Park' &    'claire')", longTargetProse)).To(Equal(false))
			})
		})
	})

	//
	// "!foo | bar" => not foo or bar
	//
	Describe("expression with 'or not' operator (A̅+B) with single quotes", func() {
		Context("in line with A and B", func() {
			It("should be true", func() {
				Expect(SearchString("(!'lazy' | 'fox')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'the' | 'fence')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'over' | 'the')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'are' | 'brevity')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'is' | 'is')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'encouraged' | 'regardless')", longTargetProse)).To(Equal(true))

				// with extra spaces
				Expect(SearchString("(!'lazy' |   'fox')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'the' |  'fence')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'over' |   'the')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'are' |  'brevity')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'is' |  'is')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'encouraged' |      'regardless')", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with A and not B", func() {
			It("should be false", func() {
				Expect(SearchString("(!'lazy' | 'abc')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'the' | 'fend')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'over' | 'they')", shortTargetProse)).To(Equal(false))
				Expect(SearchString("(!'are' | 'bravado')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'is' | 'si')", longTargetProse)).To(Equal(false))
				Expect(SearchString("(!'encouraged' | 'regarding')", longTargetProse)).To(Equal(false))
			})
		})
		Context("in line with B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("(!'laser' | 'fox')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'than' | 'fence')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'overtly' | 'the')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'arsenal' | 'brevity')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'island' | 'is')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'encouraging' | 'regardless')", longTargetProse)).To(Equal(true))
			})
		})
		Context("in line with not B and not A", func() {
			It("should be true", func() {
				Expect(SearchString("(!'laser' | 'abc')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'than' | 'fend')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'overtly' | 'they')", shortTargetProse)).To(Equal(true))
				Expect(SearchString("(!'arsenal' | 'brevado')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'island' | 'si')", longTargetProse)).To(Equal(true))
				Expect(SearchString("(!'encouraging' | 'regarding')", longTargetProse)).To(Equal(true))
			})
		})
	})
	//
	// nothing => nothing
	//
	Describe("empty expression", func() {
		Context("compared to line", func() {
			It("should be false", func() {
				Expect(SearchString("", shortTargetProse)).To(Equal(false))
				Expect(SearchString("", longTargetProse)).To(Equal(false))
			})
		})
	})
})
