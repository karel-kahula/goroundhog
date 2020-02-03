package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	goroundhogFilename string = ".goroundhog"
	groundhog          string = "           __________________\n          /                  \\\n         |    ___      ___    |\n       __|   /   \\    /   \\   |__\n      /      | o |    | o |       \\\n     |  C    \\___/    \\___/    D  |\n      \\__                       __/\n         |        ====         |\n         |         ][          |\n         |                     |"
	shadow             string = "         \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n          \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n         \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n           \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\n             \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\"
)

type groundhogDayResult struct {
	alreadyHappened bool
	sawShadow       bool
}

func main() {
	today := time.Now()
	if err := celebrateGroundhogDay(today); err != nil {
		panic(fmt.Sprintf("ground hog day is cancelled:\n%v", err))
	}
}

func celebrateGroundhogDay(today time.Time) error {
	if isGroundhogDay(today) {
		thisGroundhogDay, err := getPreviousRunResults(today)
		if err != nil {
			return fmt.Errorf("problem getting today's results:\n%w", err)
		}

		if !thisGroundhogDay.alreadyHappened {
			thisGroundhogDay = consultWithGroundhog(today)

			if err := recordResult(getFilepath(), thisGroundhogDay, today); err != nil {
				return fmt.Errorf("failed to record results:\n%w", err)
			}
		}

		printGroundhogDayResult(thisGroundhogDay)
	} else {
		return fmt.Errorf("It's not Groundhog Day!")
	}

	return nil
}

func isGroundhogDay(today time.Time) bool {
	return today.Month() == time.February && today.Day() == 2
}

func getPreviousRunResults(today time.Time) (*groundhogDayResult, error) {

	thisGroundhogDay := &groundhogDayResult{}
	filepath := getFilepath()
	if resultExists(filepath) {
		previousResults, err := readFile(filepath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file <%s>:\n%w", filepath, err)
		}

		thisGroundhogDay = getThisYearResults(previousResults, today)
	}

	return thisGroundhogDay, nil
}

func getFilepath() string {
	return filepath.Join(getHomeDirectory(), goroundhogFilename)
}

func getHomeDirectory() string {
	currentUser, _ := user.Current()
	return currentUser.HomeDir
}

func resultExists(filepath string) bool {
	fileInfo, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !fileInfo.IsDir()
}

func readFile(filepath string) (string, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to read file <%s>:\n%w", filepath, err)
	}
	return string(bytes), nil
}

func getThisYearResults(previousResults string, today time.Time) *groundhogDayResult {
	return &groundhogDayResult{
		alreadyHappened: strings.HasPrefix(previousResults, strconv.Itoa(today.Year())),
		sawShadow:       strings.HasSuffix(previousResults, "1"),
	}
}

func consultWithGroundhog(today time.Time) *groundhogDayResult {
	rand.Seed(today.UTC().UnixNano())

	return &groundhogDayResult{
		alreadyHappened: true,
		sawShadow:       rand.Intn(2) == 1,
	}
}

func printGroundhogDayResult(result *groundhogDayResult) {
	fmt.Println(groundhog)
	if result.sawShadow {
		fmt.Println(shadow)
		fmt.Println("6 More weeks of winter")
	} else {
		fmt.Println("It looks like an early spring")
	}
}

func recordResult(filepath string, result *groundhogDayResult, today time.Time) error {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file <%s>:\n%w", filepath, err)
	}
	defer file.Close()

	if _, err := file.WriteString(resultAsFileContents(today, result)); err != nil {
		return fmt.Errorf("failed to write to file <%s>:\n%w", filepath, err)
	}

	return nil
}

func resultAsFileContents(today time.Time, result *groundhogDayResult) string {
	sawShadowString := "0"
	if result.sawShadow {
		sawShadowString = "1"
	}

	return strconv.Itoa(today.Year()) + sawShadowString
}
