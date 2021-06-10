package constant

const TimeLocation string = "Asia/Shanghai"
const TimeFormatString string = "2006-01-02 15:04:05"
const TimeFormatString2 string = "2006年01月02日 15:04:05"

const (
	Token     = "Misaki"
	AppID     = "wx33cc6387acefe650"
	Appsecret = "7143fad1a2653361437faed615abd086"
)
const (
	VerificationCodeRedisKey = "guardian_verification_code_"
)

const (
	UserGenderMale   = 1
	UserGenderFemale = 2
)

const (
	UserPowerNormal = 1
	UserPowerAdmin  = 2
)

const UserOpenIdContextKey = "openId"

const (
	ExpenseStatusCost   = 1
	ExpenseStatusRefund = 2
)

var CheckProjectMap = map[int]string{
	1: "内科",
	2: "外科",
	3: "耳鼻喉科",
	4: "肝功",
	5: "血糖",
	6: "血脂",
	7: "肾功",
}

const BookCheckMaxCount = 200

const (
	PayTypeRemainder  = 1
	PayTypeCheckCount = 2
)

const CheckEndStatus = 8

const (
	MaxExaminationRemainder  = 1000000
	MaxExaminationCheckCount = 100
)

const (
	TemplateIdCheckWillStart = "rVykD4brIHuiZuELLkK45qdqSb6MSDQYyT1pqWDA_Mc"
	TemplateIdCheckStart     = "vKglW8YhFsUoB-PSyzGtP-0okYCrPZlnIwByb6gpybg"
	TemplateIdCheckAllFinish = "bIthJ8Vvf2YZ_boacnmaVw5OlBg5veqxBO9j9IY8KBE"
	TemplateIdSubscribe      = "MyKpx941KseLOVkswCnDw0ky0iYCS6TDgE49tTdJNLg"
	TemplateIdBeAdmin        = "ZzrMuqMZ_MGY_k3kImXwR1FZAn-GiGkz6U5GmMKbYKo"
	TemplateIdAdmin          = "JmXePkevoOWitoBK0wMaRXwOuuydtbGHnEulKxyXoMw"
)
