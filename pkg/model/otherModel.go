package model

type Examination struct {
	Id         int
	Cost       float64
	CheckCount int
	Remainder  float64
	CardType   float64
}

type CheckProject struct {
	Id    int
	Name  string
	Money float64
}

var CheckProjectMap = map[int]*CheckProject{
	1: {
		Id:    1,
		Name:  "内科",
		Money: 30,
	},
	2: {
		Id:    2,
		Name:  "外科",
		Money: 30,
	},
	3: {
		Id:    3,
		Name:  "耳鼻喉科",
		Money: 30,
	},
	4: {
		Id:    4,
		Name:  "肝功",
		Money: 199,
	},
	5: {
		Id:    5,
		Name:  "血糖",
		Money: 30,
	},
	6: {
		Id:    6,
		Name:  "血脂",
		Money: 30,
	},
	7: {
		Id:    7,
		Name:  "肾功",
		Money: 299,
	},
}

type BookingPay struct {
	Remainder  float64 `json:"remainder"`
	CheckCount int     `json:"check_count"`
}

var ExaminationMap = map[int]*Examination{
	1: {
		Id:         1,
		Cost:       300,
		CheckCount: 0,
		Remainder:  300,
		CardType:   9,
	},
	2: {
		Id:         2,
		Cost:       500,
		CheckCount: 0,
		Remainder:  500,
		CardType:   8,
	},
	3: {
		Id:         3,
		Cost:       1000,
		CheckCount: 0,
		Remainder:  1000,
		CardType:   7,
	},
	4: {
		Id:         4,
		Cost:       2000,
		CheckCount: 0,
		Remainder:  2000,
		CardType:   6,
	},
	5: {
		Id:         5,
		Cost:       569,
		CheckCount: 1,
		Remainder:  0,
		CardType:   0,
	},
	6: {
		Id:         6,
		Cost:       2699,
		CheckCount: 5,
		Remainder:  0,
		CardType:   0,
	},
	7: {
		Id:         7,
		Cost:       4999,
		CheckCount: 10,
		Remainder:  0,
		CardType:   0,
	},
}
