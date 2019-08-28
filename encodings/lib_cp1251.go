package encodings

import "bytes"

var TABLE_CP1251 map[int]string

func init() {
	TABLE_CP1251 = make(map[int]string)

	TABLE_CP1251[185] = "№"

	TABLE_CP1251[192] = "А"
	TABLE_CP1251[193] = "Б"
	TABLE_CP1251[194] = "В"
	TABLE_CP1251[195] = "Г"
	TABLE_CP1251[196] = "Д"
	TABLE_CP1251[197] = "Е"
	TABLE_CP1251[168] = "Ё"
	TABLE_CP1251[198] = "Ж"
	TABLE_CP1251[199] = "З"
	TABLE_CP1251[200] = "И"
	TABLE_CP1251[201] = "Й"
	TABLE_CP1251[202] = "К"
	TABLE_CP1251[203] = "Л"
	TABLE_CP1251[204] = "М"
	TABLE_CP1251[205] = "Н"
	TABLE_CP1251[206] = "О"
	TABLE_CP1251[207] = "П"
	TABLE_CP1251[208] = "Р"
	TABLE_CP1251[209] = "С"
	TABLE_CP1251[210] = "Т"
	TABLE_CP1251[211] = "У"
	TABLE_CP1251[212] = "Ф"
	TABLE_CP1251[213] = "Х"
	TABLE_CP1251[214] = "Ц"
	TABLE_CP1251[215] = "Ч"
	TABLE_CP1251[216] = "Ш"
	TABLE_CP1251[217] = "Щ"
	TABLE_CP1251[218] = "Ъ"
	TABLE_CP1251[219] = "Ы"
	TABLE_CP1251[220] = "Ь"
	TABLE_CP1251[221] = "Э"
	TABLE_CP1251[222] = "Ю"
	TABLE_CP1251[223] = "Я"

	TABLE_CP1251[224] = "а"
	TABLE_CP1251[225] = "б"
	TABLE_CP1251[226] = "в"
	TABLE_CP1251[227] = "г"
	TABLE_CP1251[228] = "д"
	TABLE_CP1251[229] = "е"
	TABLE_CP1251[184] = "ё"
	TABLE_CP1251[230] = "ж"
	TABLE_CP1251[231] = "з"
	TABLE_CP1251[232] = "и"
	TABLE_CP1251[233] = "й"
	TABLE_CP1251[234] = "к"
	TABLE_CP1251[235] = "л"
	TABLE_CP1251[236] = "м"
	TABLE_CP1251[237] = "н"
	TABLE_CP1251[238] = "о"
	TABLE_CP1251[239] = "п"
	TABLE_CP1251[240] = "р"
	TABLE_CP1251[241] = "с"
	TABLE_CP1251[242] = "т"
	TABLE_CP1251[243] = "у"
	TABLE_CP1251[244] = "ф"
	TABLE_CP1251[245] = "х"
	TABLE_CP1251[246] = "ц"
	TABLE_CP1251[247] = "ч"
	TABLE_CP1251[248] = "ш"
	TABLE_CP1251[249] = "щ"
	TABLE_CP1251[250] = "ъ"
	TABLE_CP1251[251] = "ы"
	TABLE_CP1251[252] = "ь"
	TABLE_CP1251[253] = "э"
	TABLE_CP1251[254] = "ю"
	TABLE_CP1251[255] = "я"
}

func FromCP1251(str string) string {
	sz := len(str)
	//runes := []rune(str)
	//out:=""
	var buffer bytes.Buffer
	//var b strings.Builder
	for j := 0; j < sz; j++ {
		char := int(str[j])
		newstr, ok := TABLE_CP1251[char]
		if !ok {
			// newstr := TABLE_CP1251[char]
			// if newstr == "" {
			//out += string(runes[j : j+1])
			//buffer.WriteString(string(runes[j : j+1]))
			buffer.Write([]byte{str[j]})
			//b.WriteString(string(runes[j : j+1]))
		} else {
			//out += newstr
			buffer.WriteString(newstr)
			//b.WriteString(newstr)
		}
	}
	//return b.String()
	return buffer.String() // fastest method
	//return out
}

//https://github.com/JLarky/nikmed-stats/blob/master/app/cp1251_utf8/cp1251_utf8.go
//http://unicode.org/Public/MAPPINGS/VENDORS/MICSFT/WINDOWS/CP1251.TXT
