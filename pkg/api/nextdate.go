package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func ruleOfDay(now, date time.Time, repeat string) (string, error) {

	d, err := strconv.Atoi(repeat)
	if err != nil || d > 400 || d <= 0 {
		return "", errors.New("invalid interval")
	}
	for {
		date = date.AddDate(0, 0, d)
		if date.After(now) {
			break
		}
	}
	return date.Format(DateFormat), nil
}

func ruleOfYear(now, date time.Time) (string, error) {

	for {
		date = date.AddDate(1, 0, 0)
		if date.After(now) {
			break
		}
	}
	return date.Format(DateFormat), nil
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if len(repeat) == 0 {
		return "", errors.New("repeat is empty")
	}
	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", errors.New("date format error")
	}

	repeatSlice := strings.Split(repeat, " ")
	rule := repeatSlice[0]

	switch {

	case rule == "d" && len(repeatSlice) > 1:
		return ruleOfDay(now, date, repeatSlice[1])
	case rule == "y" && len(repeatSlice) == 1:
		return ruleOfYear(now, date)
	default:
		return "", errors.New("rule format error")
	}
}
func nextDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "wrong method", http.StatusMethodNotAllowed)
	}
	err := r.ParseForm()
	if err != nil {
		return
	}
	now, err := time.Parse(DateFormat, r.Form.Get("now"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	dstart := r.Form.Get("date")
	repeat := r.Form.Get("repeat")

	result, err := NextDate(now, dstart, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "string")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(result))
	if err != nil {
		panic(err)
	}
}
