package easygolang

const FILE_COMMENT = "#"

const OPTIONS_TYPE_STRING = "String"
const OPTIONS_TYPE_ARRAY = "Array"

type OptionsStorageItem struct {
	value_storaged string
	value_default  string
	value_type     string //string,num,bool,arr
	values_allowed []string
	value_min      float64
	value_max      float64
	value_step     float64
	key_comment    string
	file_order     int
}

type OptionsStorage struct {
	arr map[string]*OptionsStorageItem
}

func NewOptionsStorage() *OptionsStorage {
	return &OptionsStorage{arr: make(map[string]*OptionsStorageItem)}
}

func (s *OptionsStorage) AddRecord_String(order int, key string, value_default string, key_comment string) {
	s.arr[key] = &OptionsStorageItem{value_type: OPTIONS_TYPE_STRING, value_storaged: value_default, value_default: value_default, key_comment: key_comment, file_order: order}
}

func (s *OptionsStorage) AddRecord_Array(order int, key string, value_default string, values_allowed []string, key_comment string) {
	v := values_allowed
	if len(v) == 0 {
		v = []string{""}
	}
	s.arr[key] = &OptionsStorageItem{value_type: OPTIONS_TYPE_ARRAY, value_storaged: value_default, value_default: value_default, values_allowed: v, key_comment: key_comment, file_order: order}
}

func (s *OptionsStorage) GetRecordType(key string) string {
	item := s.arr[key]
	if item != nil {
		return item.value_type
	}
	return ""
}

func (s *OptionsStorage) GetRecordValuesArray(key string) []string {
	item := s.arr[key]
	if item != nil {
		return item.values_allowed
	}
	return []string{}
}

func (s *OptionsStorage) GetRecordComment(key string) string {
	item := s.arr[key]
	if item != nil {
		return item.key_comment
	}
	return ""
}

func (s *OptionsStorage) RecordsValues_Load(fname string) {
	text, ok := FileTextRead(fname)
	if !ok {
		return
	}
	arr := StringSplitLines(text)
	for j := 0; j < len(arr); j++ {
		str := arr[j]
		if len(str) > 0 && StringPart(str, 1, len(FILE_COMMENT)) != FILE_COMMENT {
			ind := StringFind(str, "=")
			if ind > 0 {
				key := StringPart(str, 1, ind-1)
				value := StringPart(str, ind+1, 0)
				item := s.arr[key]
				if item != nil {
					item.value_storaged = StringReplace(StringReplace(value, "\\n", "\n"), "\\\\", "\\")
					switch item.value_type {
					case OPTIONS_TYPE_STRING:

					case OPTIONS_TYPE_ARRAY:
						s.SetArrayIndex(key, s.GetArrayIndex(key))
					}
				}
			}
		}
	}
}

func (s *OptionsStorage) RecordsValues_Save(fname string) {
	strs := []string{}
	keys := s.GetRecordsKeys()
	for j := 0; j < len(keys); j++ {
		v := s.arr[keys[j]]
		strs = append(strs, FILE_COMMENT+" "+v.key_comment)
		strs = append(strs, keys[j]+"="+StringReplace(StringReplace(v.value_storaged, "\\", "\\\\"), "\n", "\\n"))
		strs = append(strs, "")
	}
	FileTextWrite(fname, StringJoin(strs, "\n"))
}

func (s *OptionsStorage) GetRecordsKeys() []string {
	keys := make([]string, 0, len(s.arr))
	for k := range s.arr {
		keys = append(keys, k)
	}
	SortArray(keys, func(i, j int) bool {
		return s.arr[keys[i]].file_order < s.arr[keys[j]].file_order
	})
	return keys
}

func (s *OptionsStorage) SetString(key string, value string) {
	item := s.arr[key]
	if item != nil {
		item.value_storaged = value
	}
}

func (s *OptionsStorage) GetString(key string) string {
	item := s.arr[key]
	if item != nil {
		return item.value_storaged
	}
	return ""
}

func (s *OptionsStorage) SetArrayIndex(key string, value int) {
	item := s.arr[key]
	if item != nil {
		if value > -1 && value < len(item.values_allowed) {
			item.value_storaged = item.values_allowed[value]
		} else {
			ind_def := StringInArray(item.value_default, item.values_allowed)
			if ind_def > -1 {
				item.value_storaged = item.value_default
			} else {
				item.value_storaged = item.values_allowed[0]
			}
		}
	}
}

func (s *OptionsStorage) GetArrayIndex(key string) int {
	item := s.arr[key]
	if item != nil {
		ind := StringInArray(item.value_storaged, item.values_allowed)
		if ind > -1 {
			return ind
		} else {
			ind_def := StringInArray(item.value_default, item.values_allowed)
			if ind_def > -1 {
				return ind_def
			}
		}
	}
	return 0
}
