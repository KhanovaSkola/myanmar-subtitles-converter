package main

import (
    "fmt"
    "regexp"
    "os"
    "bufio"
    "io/ioutil"
    "github.com/google/myanmar-tools/clients/go"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func read_file(f string) string {
    dat, err := ioutil.ReadFile(f)
    check(err)
    return string(dat)
}

func main() {
    zgDetector := myanmartools.NewZawgyiDetector()
    var subs string
    re := regexp.MustCompile("\n00:[^\n]+\n")
    const ytidsFilename = "myanmar_ytids.dat"
    const subFormat = "vtt"

    fytids, err := os.Open(ytidsFilename)
    check(err)
    fileScanner := bufio.NewScanner(fytids)
    fileScanner.Split(bufio.ScanLines)
    var ytids []string

    for fileScanner.Scan() {
      ytids = append(ytids, fileScanner.Text())
    }
    fytids.Close()

    var fname string
    for _, ytid := range ytids {
      fname = "subs_original/" + ytid + ".my." + subFormat
      data, err := ioutil.ReadFile(fname)
      if err != nil {
        fmt.Printf("%s\tNOT FOUND\n", ytid)
        continue
      }
      subs = string(data)
      //Remove subtitle timestamps to speed up detection
      subs = re.ReplaceAllString(subs, "")
      score := zgDetector.GetZawgyiProbability(subs)
      fmt.Printf("%s\t%f\n",ytid, score)
    }
}
