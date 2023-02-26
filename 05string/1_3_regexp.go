package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	regexp.MatchString("123+", "1234")
	//true <nil>

	r, err := regexp.Compile(`1a`)
	if err != nil {
		panic(err)
	}
	regexp.MustCompile(`123\q`)
	//срабатывает паника

	r.ReplaceAllString("er1aewt1a", "333")
	// er333ewt333
	r.FindAllString("er1aewt1a", -1)
	//[1a 1a]
	r.FindAllString("er1aewt1a", 1)
	//[1a]

	r.FindString("er1aewt1a")
	// "a1"
	r.FindString("er1ewt1b")
	// ""

	r.MatchString("1a")
	// true
	r.MatchString("11")
	// false
	r.MatchString("1a1")
	// true

	// \A начало строки
	// \z конец строки
	r2 := regexp.MustCompile(`\A1a\z`)
	r2.MatchString("1a")
	// true
	r2.MatchString("1a1")
	// false

	// Классы символов
	// .  любой символ
	// \d 0-9
	// \D не 0-9
	// \s пробельный символ \n\t\r' '
	// \S не \s
	// \w 0-9A-Za-z
	// \W не \w
	r3 := regexp.MustCompile(`\A\d\d.\d\d.\d\d\d\d\z`)

	r3.MatchString("01.01.2012")
	// true
	r3.FindString("01-0122012")
	// "01-0122012"
	r3.FindString("01.01.201a")
	// ""

	// специальные символы, нужно экранировать, чтобы рассматривались как обычные
	// ^ $ * + ? { } [ ] \ | ( )
	// .
	r4 := regexp.MustCompile(`\A\d\d\.\d\d\.\d\d\d\d\z`)
	r4.FindString("01-0122012")
	// ""

	// повторы символов
	// * 0 и более
	// + 1 и более
	// {n} n раз
	// {min, max} от min до max раз
	r5 := regexp.MustCompile(`\A\d+\.\d*\.\d{4}\z`)
	r5.MatchString("..2012")
	// false
	r5.MatchString("111..2012")
	// true

	r6 := regexp.MustCompile(`\A\d{2}\.\d{2}\.\d{4}\z`)
	r6.MatchString("1..2012")
	// false
	r6.MatchString("01.01.2012")
	// true

	// группировка
	r7 := regexp.MustCompile(`(\d{2}\.){2}\d{4}`)
	r7.MatchString("1.01.2012")
	// false
	r7.MatchString("01.01.2012")
	// true
	r7.FindAllStringSubmatch("01.01.201202.02.2013", -1)
	// [[01.01.2012 01.] [02.02.2013 02.]]

	// или (объединение нескольких выражений)
	r8 := regexp.MustCompile(`(\A(\d{2}\.){2}\d{4}\z)|(\d{2}:{\d{2})`)
	r8.MatchString("18:30")
	// true
	r8.MatchString("01.01.2012")
	// true
	r8.MatchString("ase18:30asdf")
	// true

	// описание класса символов
	r9 := regexp.MustCompile(`(\A(\d{2}[.-/]){2}\d{4}\z)`)
	r9.MatchString("12-12-1223")
	// false
	r9.MatchString("01/01/2012")
	// true
	r9.MatchString("01.01.2012")
	// true

	//правильное экранирование дефиса
	r10 := regexp.MustCompile(`(\A(\d{2}[.\-/]){2}\d{4}\z)`)
	r10.MatchString("12-12-1223")
	// true
	r10.MatchString("01/01/2012")
	// true
	r10.FindStringSubmatch("12-24-1223")
	// [12-24-1223 12-24-1223 24-]

	// отрицание
	r9_2 := regexp.MustCompile(`(\A(\d{2}[^.-/]){2}\d{4}\z)`)
	r9_2.MatchString("12-12-1223")
	// true
	r9_2.MatchString("01/01/2012")
	// false
	r9_2.MatchString("01.01.2012")
	// false

	//именованные группы через угловые скобки
	r11 := regexp.MustCompile(`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})T(?P<Hour>\d{2}):(?P<Minute>\d{2}):(?P<Second>\d{2})`)
	matches := r11.FindStringSubmatch("2014-12-22T12:23:29")
	year, _ := strconv.Atoi(matches[r11.SubexpIndex("Year")])
	month, _ := strconv.Atoi(matches[r11.SubexpIndex("Month")])
	day, _ := strconv.Atoi(matches[r11.SubexpIndex("Day")])
	hour, _ := strconv.Atoi(matches[r11.SubexpIndex("Hour")])
	minute, _ := strconv.Atoi(matches[r11.SubexpIndex("Minute")])
	second, _ := strconv.Atoi(matches[r11.SubexpIndex("Second")])
	fmt.Printf("Год %v, месяц %v, день %v, часов %v, минут %v, секунд %v\n", year, month, day, hour, minute, second)
	//Год 2014, месяц 12, день 22, часов 12, минут 23, секунд 29

	//поиск всех вхождений
	r12 := regexp.MustCompile(`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})T(?P<Hour>\d{2}):(?P<Minute>\d{2}):(?P<Second>\d{2})`)
	matches2 := r12.FindAllStringSubmatch(strings.Repeat("2014-12-22T12:23:20", 5), -1)
	for _, val := range matches2 {
		year, _ := strconv.Atoi(val[r12.SubexpIndex("Year")])
		month, _ := strconv.Atoi(val[r12.SubexpIndex("Month")])
		day, _ := strconv.Atoi(val[r12.SubexpIndex("Day")])
		hour, _ := strconv.Atoi(val[r12.SubexpIndex("Hour")])
		minute, _ := strconv.Atoi(val[r12.SubexpIndex("Minute")])
		second, _ := strconv.Atoi(val[r12.SubexpIndex("Second")])
		fmt.Printf("Год %v, месяц %v, день %v, часов %v, минут %v, секунд %v\n", year, month, day, hour, minute, second)
		//Год 2014, месяц 12, день 22, часов 12, минут 23, секунд 20 5 раз
	}

}
