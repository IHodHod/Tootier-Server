package global

const (
	Fa       = "fa"
	En       = "en"
	Language = Fa
)

type Lang struct {
	MSG_OK                     string
	MSG_ERR                    string
	MSG_ERR_USERNAME_EXISTS    string
	MSG_ERR_EMAIL_EXISTS       string
	MSG_ERR_PHONENUMBER_EXISTS string
	MSG_ERR_CAN_NOT_CERATE_USER string
	MSG_SUCCESS_SIGNUP string
}

var FaLang = Lang{
	MSG_OK:                     "درخواست شما موفقیت آمیز بود",
	MSG_ERR:                    "درخواست شما با اشکال روبرو شد",
	MSG_ERR_USERNAME_EXISTS:    "این نام کاربری قبلا ایجاد شده است",
	MSG_ERR_EMAIL_EXISTS:       "این ایمیل قبلا استفاده شده است",
	MSG_ERR_PHONENUMBER_EXISTS: "این شماره موبایل قبلا استفاده شده است",
	MSG_ERR_CAN_NOT_CERATE_USER: "خطا در ثبت نام کاربر",
	MSG_SUCCESS_SIGNUP: "به توتیر خوش آمدید" ,
}

var EnLang = Lang{
	MSG_OK:                     "Your request has been succeed",
	MSG_ERR:                    "Your request has been failed",
	MSG_ERR_USERNAME_EXISTS:    "Your Username already created",
	MSG_ERR_EMAIL_EXISTS:       "Your Email already created",
	MSG_ERR_PHONENUMBER_EXISTS: "Your Phone Number already created",
	MSG_ERR_CAN_NOT_CERATE_USER: "Signup user failed",
	MSG_SUCCESS_SIGNUP: "Welcome To Tootier" ,
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
