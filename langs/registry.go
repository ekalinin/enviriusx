package langs

type LangCreator func() LangBuilder

var Langs = map[string]LangCreator{}

func Add(name string, lc LangCreator) {
	Langs[name] = lc
}
