package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

var groundhogDay time.Time = time.Date(2020, time.February, 2, 0, 0, 0, 0, time.UTC)
var anotherDay time.Time = time.Date(2020, time.February, 3, 0, 0, 0, 0, time.UTC)

const testFilename string = ".testgoroundhog"

func TestResultExistsFalse(t *testing.T) {
	filepath := filepath.Join(getHomeDirectory(), testFilename)

	if resultExists(filepath) != false {
		t.Error("unexpected result")
	}
}

func TestResultExistsTrue(t *testing.T) {
	filepath := filepath.Join(getHomeDirectory(), testFilename)

	file, err := os.Create(filepath)
	if err != nil {
		t.Fatal(err.Error())
	}

	if resultExists(filepath) != true {
		t.Error("unexpected result")
	}

	file.Close()

	if err = os.Remove(filepath); err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetThisYearResultsAlreadyRanThisYear(t *testing.T) {
	previousResults := "20201"
	got := getThisYearResults(previousResults, groundhogDay)
	if got.alreadyHappened != true {
		t.Error("got <false> want <true>")
	}
}

func TestGetThisYearResultsDidNotRunThisYear(t *testing.T) {
	previousResults := "20191"
	got := getThisYearResults(previousResults, groundhogDay)
	if got.alreadyHappened != false {
		t.Error("got <true> want <false>")
	}
}

func TestGetThisYearResultSawShadow(t *testing.T) {
	previousResults := "20201"
	got := getThisYearResults(previousResults, groundhogDay)
	if got.sawShadow != true {
		t.Error("got <false> want <true>")
	}
}

func TestGetThisYearResultDidntSeeShadow(t *testing.T) {
	previousResults := "20200"
	got := getThisYearResults(previousResults, groundhogDay)
	if got.sawShadow != false {
		t.Error("got <true> want <false>")
	}
}

func TestCelebrateOnTheDay(t *testing.T) {
	if err := celebrateGroundhogDay(groundhogDay); err != nil {
		t.Error(err.Error())
	}
}

func TestCelebrateOnAnotherDay(t *testing.T) {
	if err := celebrateGroundhogDay(anotherDay); err == nil {
		t.Error("should have thrown error")
	}
}
