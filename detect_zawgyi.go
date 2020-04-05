package main

import (
	"bufio"
	"fmt"
	"github.com/google/myanmar-tools/clients/go"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	const ytidsFilename = "myanmar_ytids.dat"
	const subFormat = "vtt"
	var subs string
	var ytids []string
	re := regexp.MustCompile("\n00:[^\n]+\n")

	zgDetector := myanmartools.NewZawgyiDetector()

	fytids, err := os.Open(ytidsFilename)
	if err != nil {
		fmt.Println("Could not open file %s", ytidsFilename)
		return
	}
	defer fytids.Close()

	fileScanner := bufio.NewScanner(fytids)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		ytids = append(ytids, fileScanner.Text())
	}

	var fname string
	for _, ytid := range ytids {
		fname = "subs_original/" + ytid + ".my." + subFormat
		data, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Printf("%s\tNOT FOUND\n", ytid)
			continue
		}
		subs = string(data)
		// Remove subtitle timestamps to speed up detection
		// Not quite sure whether the time savings are actually worth it
		subs = re.ReplaceAllString(subs, "")

		score := zgDetector.GetZawgyiProbability(subs)
		fmt.Printf("%s\t%f\n", ytid, score)
	}
}
