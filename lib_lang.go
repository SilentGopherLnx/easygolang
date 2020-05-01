package easygolang

const DEFAULT_LANG = "en"

func InitLang(fname string) *LangArr {
	var langs *LangArr
	txt, ok := FileTextRead(fname)
	if !ok {
		Prln("ERROR READING FILE WITH TRANSLATION!")
		langs = LangArrNew("")
	} else {
		//Prln("LANG FILE LOADED!")
		langs = LangArrNew(txt)
	}
	langs.SetLang(DEFAULT_LANG)
	//langs.SetLang("ru")
	return langs
}

type langStrs struct {
	title string
	strs  map[string]string
}

func (a *langStrs) GetLangString(str_code string) (string, bool) {
	v, ok := a.strs[StringTrim(str_code)]
	if ok {
		return v, true
	} else {
		return "[?]" + str_code, false
	}
}

type LangArr struct {
	lang_current string
	lang_codes   []string
	lang_strs    map[string]*langStrs
}

func (a *LangArr) SetLang(lang_code string) {
	if StringLength(lang_code) > 0 {
		a.lang_current = lang_code
	} else {
		a.lang_current = DEFAULT_LANG
	}
}

func (a *LangArr) GetStr(str_code string) string {
	l, ok := a.lang_strs[StringTrim(a.lang_current)]
	if ok {
		v, yes := l.GetLangString(str_code)
		if yes {
			return v
		}
	}
	l2, ok2 := a.lang_strs[DEFAULT_LANG]
	if ok2 {
		v, yes := l2.GetLangString(str_code)
		if yes {
			return v
		}
	}
	return "[" + a.lang_current + "]" + str_code

}

func (a *LangArr) GetLangsCodes() []string {
	return a.lang_codes
}

func (a *LangArr) GetLangsTitles(codes []string) []string {
	arr := []string{}
	for j := 0; j < len(codes); j++ {
		arr = append(arr)
	}
	return arr
}

func LangArrNew(txt string) *LangArr {
	larr := LangArr{lang_strs: make(map[string]*langStrs), lang_codes: []string{}, lang_current: "en"}
	arr := StringSplit("\n"+txt, "\n!")
	for j := 0; j < len(arr); j++ {
		lang1 := StringSplitLines(arr[j])
		if len(lang1) > 0 {
			title := StringSplit(lang1[0], "/")
			if len(title) == 2 {
				larr.lang_codes = append(larr.lang_codes, title[0])
				lang := langStrs{title: title[1], strs: make(map[string]string)}
				larr.lang_strs[title[0]] = &lang
				for k := 1; k < len(lang1); k++ {
					ab := StringSplit(lang1[k], "=")
					if len(ab) >= 2 {
						lang.strs[ab[0]] = ab[1]
						//Prln(title[0] + " - " + ab[0] + " - " + ab[1])
					}
				}
			}
		}
	}
	return &larr
}
