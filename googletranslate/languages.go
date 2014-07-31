package googletranslate

type Language struct {
	Name, Code string
}

var languages = []Language{
	Language{"Afrikaans", "af"},
	Language{"Albanian", "sq"},
	Language{"Arabic", "ar"},
	Language{"Armenian", "hy"},
	Language{"Azerbaijani", "az"},
	Language{"Basque", "eu"},
	Language{"Belarusian", "be"},
	Language{"Bengali", "bn"},
	Language{"Bosnian", "bs"},
	Language{"Bulgarian", "bg"},
	Language{"Catalan", "ca"},
	Language{"Chinese", "zh"},
	Language{"Croatian", "hr"},
	Language{"Czech", "cs"},
	Language{"Danish", "da"},
	Language{"Dutch", "nl"},
	Language{"English", "en"},
	Language{"Esperanto", "eo"},
	Language{"Estonian", "et"},
	Language{"Finnish", "fi"},
	Language{"French", "fr"},
	Language{"Galician", "gl"},
	Language{"Georgian", "ka"},
	Language{"German", "de"},
	Language{"Greek", "el"},
	Language{"Gujarati", "gu"},
	Language{"Haitian", "ht"},
	Language{"Hausa", "ha"},
	Language{"Hebrew", "he"},
	Language{"Hindi", "hi"},
	Language{"Hungarian", "hu"},
	Language{"Icelandic", "is"},
	Language{"Igbo", "ig"},
	Language{"Indonesian", "id"},
	Language{"Irish", "ga"},
	Language{"Italian", "it"},
	Language{"Japanese", "ja"},
	Language{"Javanese", "jv"},
	Language{"Kannada", "kn"},
	Language{"Khmer", "km"},
	Language{"Korean", "ko"},
	Language{"Lao", "lo"},
	Language{"Latin", "la"},
	Language{"Latvian", "lv"},
	Language{"Lithuanian", "lt"},
	Language{"Macedonian", "mk"},
	Language{"Malay", "ms"},
	Language{"Maltese", "mt"},
	Language{"Maori", "mi"},
	Language{"Marathi", "mr"},
	Language{"Mongolian", "mn"},
	Language{"Nepali", "ne"},
	Language{"Norwegian", "no"},
	Language{"Persian", "fa"},
	Language{"Polish", "pl"},
	Language{"Portuguese", "pt"},
	Language{"Punjabi", "pa"},
	Language{"Romanian", "ro"},
	Language{"Russian", "ru"},
	Language{"Serbian", "sr"},
	Language{"Slovak", "sk"},
	Language{"Slovenian", "sl"},
	Language{"Somali", "so"},
	Language{"Spanish", "es"},
	Language{"Swahili", "sw"},
	Language{"Swedish", "sv"},
	Language{"Tamil", "ta"},
	Language{"Telugu", "te"},
	Language{"Thai", "th"},
	Language{"Turkish", "tr"},
	Language{"Ukrainian", "uk"},
	Language{"Urdu", "ur"},
	Language{"Vietnamese", "vi"},
	Language{"Welsh", "cy"},
	Language{"Yiddish", "yi"},
	Language{"Yoruba", "yo"},
	Language{"Zulu", "zu"},
}

func ListLanguages() string {
	langs := ""
	for _, lang := range languages {
		langs += lang.Code + " - " + lang.Name + "\n"
	}
	return langs
}

func inLangList(code string) bool {
	for _, lang := range languages {
		if lang.Code == code {
			return true
		}
	}

	return false
}
