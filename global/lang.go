package global

const (
	Fa       = "fa"
	En       = "en"
	Language = Fa
)

type Lang struct {
	MSG_OK string
	MSG_ERR string
}

var FaLang = Lang{
	MSG_OK: "درخواست شما موفقیت آمیز بود",
	MSG_ERR: "درخواست شما با اشکال روبرو شد",
}

var EnLang = Lang{
	MSG_OK: "Your request has been succeed",
	MSG_ERR: "Your request has been failed" ,
}



func GetLang() Lang {
	switch Language {
	case Fa:
		return FaLang
	case En:
		return EnLang
	default:
		return FaLang
	}
}
