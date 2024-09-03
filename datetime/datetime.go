package datetime

import (
	"strconv"
	"strings"
	"time"

	"github.com/KyawBo/common-library/constants"
)

const (
	DATETIME_FORMAT_PRICE_DUE_IN_DB     = "2006-01-02 00:00:00"
	DATETIME_FORMAT_PRICE_DUE_IN_REDIS  = "02012006"
	DATETIME_FORMAT_DATE_IN_RESPONSE_EN = "2 January 2006"
	DATETIME_FORMAT_DATE_IN_SLIP        = "2 January 2006 15:04:05"
	DATETIME_FORMAT_YYYYMMDD            = "20060102"
	DATETIME_FORMAT_DD_MM_YYYY          = "02-01-2006"
)

func ToThaiMonth(in string) string {
	out := in
	out = strings.Replace(out, "January", "ม.ค.", 1)
	out = strings.Replace(out, "February", "ก.พ.", 1)
	out = strings.Replace(out, "March", "มี.ค.", 1)
	out = strings.Replace(out, "April", "เม.ย.", 1)
	out = strings.Replace(out, "May", "พ.ค.", 1)
	out = strings.Replace(out, "June", "มิ.ย.", 1)
	out = strings.Replace(out, "July", "ก.ค.", 1)
	out = strings.Replace(out, "August", "ส.ค.", 1)
	out = strings.Replace(out, "September", "ก.ย.", 1)
	out = strings.Replace(out, "October", "ต.ค.", 1)
	out = strings.Replace(out, "November", "พ.ย.", 1)
	out = strings.Replace(out, "December", "ธ.ค.", 1)

	out = strings.Replace(out, "01 ", "1 ", 1)
	out = strings.Replace(out, "02 ", "2 ", 1)
	out = strings.Replace(out, "03 ", "3 ", 1)
	out = strings.Replace(out, "04 ", "4 ", 1)
	out = strings.Replace(out, "05 ", "5 ", 1)
	out = strings.Replace(out, "06 ", "6 ", 1)
	out = strings.Replace(out, "07 ", "7 ", 1)
	out = strings.Replace(out, "08 ", "8 ", 1)
	out = strings.Replace(out, "09 ", "9 ", 1)
	return out
}

func ToThaiMonthFull(in string) string {
	out := in
	out = strings.Replace(out, "January", "มกราคม", 1)
	out = strings.Replace(out, "February", "กุมภาพันธ์", 1)
	out = strings.Replace(out, "March", "มีนาคม", 1)
	out = strings.Replace(out, "April", "เมษายน", 1)
	out = strings.Replace(out, "May", "พฤษภาคม", 1)
	out = strings.Replace(out, "June", "มิถุนายน", 1)
	out = strings.Replace(out, "July", "กรกฎาคม", 1)
	out = strings.Replace(out, "August", "สิงหาคม", 1)
	out = strings.Replace(out, "September", "กันยายน", 1)
	out = strings.Replace(out, "October", "ตุลาคม", 1)
	out = strings.Replace(out, "November", "พฤศจิกายน", 1)
	out = strings.Replace(out, "December", "ธันวาคม", 1)

	return out
}

func FromThaiMonth(in string) string {
	out := in
	out = strings.Replace(out, "ม.ค.", "January", 1)
	out = strings.Replace(out, "ก.พ.", "February", 1)
	out = strings.Replace(out, "มี.ค.", "March", 1)
	out = strings.Replace(out, "เม.ย.", "April", 1)
	out = strings.Replace(out, "พ.ค.", "May", 1)
	out = strings.Replace(out, "มิ.ย.", "June", 1)
	out = strings.Replace(out, "ก.ค.", "July", 1)
	out = strings.Replace(out, "ส.ค.", "August", 1)
	out = strings.Replace(out, "ก.ย.", "September", 1)
	out = strings.Replace(out, "ต.ค.", "October", 1)
	out = strings.Replace(out, "พ.ย.", "November", 1)
	out = strings.Replace(out, "ธ.ค.", "December", 1)

	out = strings.Replace(out, "มกราคม", "January", 1)
	out = strings.Replace(out, "กุมภาพันธ์", "February", 1)
	out = strings.Replace(out, "มีนาคม", "March", 1)
	out = strings.Replace(out, "เมษายน", "April", 1)
	out = strings.Replace(out, "พฤษภาคม", "May", 1)
	out = strings.Replace(out, "มิถุนายน", "June", 1)
	out = strings.Replace(out, "กรกฎาคม", "July", 1)
	out = strings.Replace(out, "สิงหาคม", "August", 1)
	out = strings.Replace(out, "กันยายน", "September", 1)
	out = strings.Replace(out, "ตุลาคม", "October", 1)
	out = strings.Replace(out, "พฤศจิกายน", "November", 1)
	out = strings.Replace(out, "ธันวาคม", "December", 1)

	return out
}

func FromThaiMonthSemi(in string) string {
	out := in

	out = strings.Replace(out, "มกราคม", "ม.ค.", 1)
	out = strings.Replace(out, "กุมภาพันธ์", "ก.พ.", 1)
	out = strings.Replace(out, "มีนาคม", "มี.ค.", 1)
	out = strings.Replace(out, "เมษายน", "เม.ย.", 1)
	out = strings.Replace(out, "พฤษภาคม", "พ.ค.", 1)
	out = strings.Replace(out, "มิถุนายน", "มิ.ย.", 1)
	out = strings.Replace(out, "กรกฎาคม", "ก.ค.", 1)
	out = strings.Replace(out, "สิงหาคม", "ส.ค.", 1)
	out = strings.Replace(out, "กันยายน", "ก.ย.", 1)
	out = strings.Replace(out, "ตุลาคม", "ต.ค.", 1)
	out = strings.Replace(out, "พฤศจิกายน", "พ.ย.", 1)
	out = strings.Replace(out, "ธันวาคม", "ธ.ค.", 1)

	return out
}

func ToThaiMonthNum(in string) string {
	out := in
	out = strings.Replace(out, "01", "ม.ค.", 1)
	out = strings.Replace(out, "02", "ก.พ.", 1)
	out = strings.Replace(out, "03", "มี.ค.", 1)
	out = strings.Replace(out, "04", "เม.ย.", 1)
	out = strings.Replace(out, "05", "พ.ค.", 1)
	out = strings.Replace(out, "06", "มิ.ย.", 1)
	out = strings.Replace(out, "07", "ก.ค.", 1)
	out = strings.Replace(out, "08", "ส.ค.", 1)
	out = strings.Replace(out, "09", "ก.ย.", 1)
	out = strings.Replace(out, "10", "ต.ค.", 1)
	out = strings.Replace(out, "11", "พ.ย.", 1)
	out = strings.Replace(out, "12", "ธ.ค.", 1)

	return out
}

func ToLottoRewardDate(round int, year int) string {
	outDay := ""
	outMonth := ""

	intYear := 543 + year
	outYear := strconv.Itoa(intYear)

	if round == 1 || round == 2 {
		outDay = "30 "
		outMonth = "ธ.ค. "
	} else if round == 3 || round == 4 {
		outDay = "17 "
		outMonth = "ม.ค. "
	} else if round == 5 || round == 6 {
		outDay = "1 "
		outMonth = "ก.พ. "
	} else if round == 7 || round == 8 {
		outDay = "17 "
		outMonth = "ก.พ. "
	} else if round == 9 || round == 10 {
		outDay = "1 "
		outMonth = "มี.ค. "
	} else if round == 11 || round == 12 {
		outDay = "16 "
		outMonth = "มี.ค. "
	} else if round == 13 || round == 14 {
		outDay = "1 "
		outMonth = "เม.ย. "
	} else if round == 15 || round == 16 {
		outDay = "16 "
		outMonth = "เม.ย. "
	} else if round == 17 || round == 18 {
		outDay = "2 "
		outMonth = "พ.ค. "
	} else if round == 19 || round == 20 {
		outDay = "16 "
		outMonth = "พ.ค. "
	} else if round == 21 || round == 22 {
		outDay = "1 "
		outMonth = "มิ.ย. "
	} else if round == 23 || round == 24 {
		outDay = "16 "
		outMonth = "มิ.ย. "
	} else if round == 25 || round == 26 {
		outDay = "1 "
		outMonth = "ก.ค. "
	} else if round == 27 || round == 28 {
		outDay = "16 "
		outMonth = "ก.ค. "
	} else if round == 29 || round == 30 {
		outDay = "1 "
		outMonth = "ส.ค. "
	} else if round == 31 || round == 32 {
		outDay = "16 "
		outMonth = "ส.ค. "
	} else if round == 33 || round == 34 {
		outDay = "1 "
		outMonth = "ก.ย. "
	} else if round == 35 || round == 36 {
		outDay = "16 "
		outMonth = "ก.ย. "
	} else if round == 37 || round == 38 {
		outDay = "1 "
		outMonth = "ต.ค. "
	} else if round == 39 || round == 40 {
		outDay = "16 "
		outMonth = "ต.ค. "
	} else if round == 41 || round == 42 {
		outDay = "1 "
		outMonth = "พ.ย. "
	} else if round == 43 || round == 44 {
		outDay = "16 "
		outMonth = "พ.ย. "
	} else if round == 45 || round == 46 {
		outDay = "1 "
		outMonth = "ธ.ค. "
	} else if round == 47 || round == 48 {
		outDay = "16 "
		outMonth = "ธ.ค. "
	}
	strOut := outDay + outMonth + outYear

	return strOut
}

func ToLottoRewardDateFullM(round int, year int) string {
	outDay := ""
	outMonth := ""

	intYear := 543 + year
	outYear := strconv.Itoa(intYear)

	if round == 1 || round == 2 {
		outDay = "30 "
		outMonth = "ธันวาคม "
	} else if round == 3 || round == 4 {
		outDay = "17 "
		outMonth = "มกราคม "
	} else if round == 5 || round == 6 {
		outDay = "1 "
		outMonth = "กุมภาพันธ์ "
	} else if round == 7 || round == 8 {
		outDay = "17 "
		outMonth = "กุมภาพันธ์ "
	} else if round == 9 || round == 10 {
		outDay = "1 "
		outMonth = "มีนาคม "
	} else if round == 11 || round == 12 {
		outDay = "16 "
		outMonth = "มีนาคม "
	} else if round == 13 || round == 14 {
		outDay = "1 "
		outMonth = "เมษายน "
	} else if round == 15 || round == 16 {
		outDay = "16 "
		outMonth = "เม.ย. "
	} else if round == 17 || round == 18 {
		outDay = "2 "
		outMonth = "พฤษภาคม "
	} else if round == 19 || round == 20 {
		outDay = "16 "
		outMonth = "พฤษภาคม "
	} else if round == 21 || round == 22 {
		outDay = "1 "
		outMonth = "มิถุนายน "
	} else if round == 23 || round == 24 {
		outDay = "16 "
		outMonth = "มิถุนายน "
	} else if round == 25 || round == 26 {
		outDay = "1 "
		outMonth = "กรกฎาคม "
	} else if round == 27 || round == 28 {
		outDay = "16 "
		outMonth = "กรกฎาคม "
	} else if round == 29 || round == 30 {
		outDay = "1 "
		outMonth = "สิงหาคม "
	} else if round == 31 || round == 32 {
		outDay = "16 "
		outMonth = "สิงหาคม "
	} else if round == 33 || round == 34 {
		outDay = "1 "
		outMonth = "กันยายน "
	} else if round == 35 || round == 36 {
		outDay = "16 "
		outMonth = "กันยายน "
	} else if round == 37 || round == 38 {
		outDay = "1 "
		outMonth = "ตุลาคม "
	} else if round == 39 || round == 40 {
		outDay = "16 "
		outMonth = "ตุลาคม "
	} else if round == 41 || round == 42 {
		outDay = "1 "
		outMonth = "พฤศจิกายน "
	} else if round == 43 || round == 44 {
		outDay = "16 "
		outMonth = "พฤศจิกายน "
	} else if round == 45 || round == 46 {
		outDay = "1 "
		outMonth = "ธันวาคม "
	} else if round == 47 || round == 48 {
		outDay = "16 "
		outMonth = "ธันวาคม "
	}
	strOut := outDay + outMonth + outYear
	return strOut
}

func ToCurrentRound() (int, time.Month, int) {
	currentTime := time.Now()
	y, m, day := currentTime.Date()

	layout1_ := "2 January 2006"
	currentTime, _ = time.Parse(layout1_, "1 "+m.String()+" "+strconv.Itoa(y))
	pDay := 0
	if m.String() != "January" && m.String() != "December" && m.String() != "April" && m.String() != "February" && m.String() != "May" {

		pDay = 16
		if day > 16 {
			pDay = 1
			y, m, day = currentTime.AddDate(0, 1, 0).Date()
		}
		if day == 1 {
			pDay = 1
		}
	} else if m.String() == "January" {
		pDay = 17
		if day > 17 {
			pDay = 1
			y, m, day = currentTime.AddDate(0, 1, 0).Date()
		}
	} else if m.String() == "February" {
		pDay = 17
		if day > 17 {
			pDay = 1
			y, m, day = currentTime.AddDate(0, 1, 0).Date()
		}
	} else if m.String() == "May" {
		pDay = 16
		if day > 16 {
			pDay = 1
			y, m, day = currentTime.AddDate(0, 1, 0).Date()
		} else if day <= 2 {
			pDay = 2
		}
	} else if m.String() == "December" {

		pDay = 16
		if day > 16 {
			pDay = 30
		}
		if day == 1 {
			pDay = 1
		}
		if day == 30 {
			pDay = 30
		}
		if day > 30 {
			pDay = 17
			y, m, day = currentTime.AddDate(0, 1, 0).Date()
		}
	} else if m.String() == "April" {
		pDay = 16
		if day > 16 {
			pDay = 2
			y, m, day = currentTime.AddDate(0, 1, 0).Date()
		}
	}

	return pDay, m, y
}

func IsRewardDay() bool {
	currentTime := time.Now()
	_, m, day := currentTime.Date()
	return (day == 2 && m.String() == "May") || ((day == 1 && m.String() != "January") && (day == 1 && m.String() != "May")) || ((day == 16 && m.String() != "January") && (day == 16 && m.String() != "February")) || (day == 30 && m.String() == "December") || (day == 17 && m.String() == "January") || (day == 17 && m.String() == "February")

}

func ToCurrentRoundInNumber() []int {
	currentTime := time.Now()
	y, m, day := currentTime.Date()

	layout1_ := "2 January 2006"
	currentTime, _ = time.Parse(layout1_, "1 "+m.String()+" "+strconv.Itoa(y))

	var out []int

	if m.String() == "January" {

		if day <= 17 {
			out = append(out, 3)
			out = append(out, 4)
		} else {
			out = append(out, 5)
			out = append(out, 6)
		}
	} else if m.String() == "February" {

		if day <= 17 {
			out = append(out, 7)
			out = append(out, 8)
		} else {
			out = append(out, 9)
			out = append(out, 10)
		}
	} else if m.String() == "March" {

		if day <= 16 {
			out = append(out, 11)
			out = append(out, 12)
		} else {
			out = append(out, 13)
			out = append(out, 14)
		}
	} else if m.String() == "April" {

		if day <= 16 {
			out = append(out, 15)
			out = append(out, 16)
		} else {
			out = append(out, 17)
			out = append(out, 18)
		}
	} else if m.String() == "May" {

		if day <= 16 {
			out = append(out, 19)
			out = append(out, 20)
		} else {
			out = append(out, 21)
			out = append(out, 22)
		}
	} else if m.String() == "June" {

		if day <= 16 {
			out = append(out, 23)
			out = append(out, 24)
		} else {
			out = append(out, 25)
			out = append(out, 26)
		}
	} else if m.String() == "July" {

		if day <= 16 {
			out = append(out, 27)
			out = append(out, 28)
		} else {
			out = append(out, 29)
			out = append(out, 30)
		}
	} else if m.String() == "August" {

		if day <= 16 {
			out = append(out, 31)
			out = append(out, 32)
		} else {
			out = append(out, 33)
			out = append(out, 34)
		}
	} else if m.String() == "September" {

		if day <= 16 {
			out = append(out, 35)
			out = append(out, 36)
		} else {
			out = append(out, 37)
			out = append(out, 38)
		}
	} else if m.String() == "October" {

		if day <= 16 {
			out = append(out, 39)
			out = append(out, 40)
		} else {
			out = append(out, 41)
			out = append(out, 42)
		}
	} else if m.String() == "November" {

		if day <= 16 {
			out = append(out, 43)
			out = append(out, 44)
		} else {
			out = append(out, 45)
			out = append(out, 46)
		}
	} else if m.String() == "December" {

		if day <= 16 {
			out = append(out, 47)
			out = append(out, 48)
		} else {
			out = append(out, 1)
			out = append(out, 2)
		}
	}

	return out
}
func ToCurrentYear() string {
	currentTime := time.Now()
	y, _, _ := currentTime.Date()
	y += 543
	return strconv.Itoa(y)[2:4]
}

func CalculateEmsEffDate(txnTime time.Time) time.Time {
	hr, _, _ := txnTime.Clock()
	effDate := time.Time{}
	if hr >= 23 {
		// EffDate = T+1 if >= 23.00
		effDate = txnTime.AddDate(0, 0, 1)
	} else {
		// EffDate = T if < 23.00
		effDate = txnTime
	}

	return effDate
}

func ConvertChristianYearToThaiYear(date time.Time) time.Time {
	if IsLeapYear(date) {
		y := date.Year() + 543
		return time.Date(y, date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
	}
	return date.AddDate(543, 0, 0)
}

func IsLeapYear(date time.Time) bool {
	endOfYear := time.Date(date.Year(), time.December, 31, 0, 0, 0, 0, time.Local)
	day := endOfYear.YearDay()
	return day == 366
}

func GetFormattedCurrentTime() string {
	currentTime := time.Now().Local()
	formattedCurrentTime := currentTime.Format(constants.DATE_FORMAT)
	return formattedCurrentTime
}
