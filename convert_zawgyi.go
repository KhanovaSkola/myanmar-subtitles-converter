package main

import (
    "fmt"
    "os"
    "bufio"
    "io/ioutil"
    "github.com/Rabbit-Converter/Rabbit-Go"
)

func main() {
    var subs string
    const ytidsFilename = "zawgyi_ytids.dat"
    const subFormat = "vtt"
    var ytids []string

    fytids, err := os.Open(ytidsFilename)
    if err != nil {
      panic(err)
    }
    defer fytids.Close()

    fileScanner := bufio.NewScanner(fytids)
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
      ytids = append(ytids, fileScanner.Text())
    }

    for _, ytid := range ytids {
      fname_in := "subs_original/" + ytid + ".my." + subFormat
      fname_out := "subs_converted/" + ytid + ".my." + subFormat
      data, err := ioutil.ReadFile(fname_in)
      if err != nil {
        fmt.Printf("%s\tNOT FOUND\n", ytid)
        continue
      }
      subs = string(data)
      subs_uni := rabbit.Zg2uni(subs)
      data = []byte(subs_uni)
      err = ioutil.WriteFile(fname_out, data, 0644)
      if err != nil {
        fmt.Printf("Could not write to file %s\n", fname_out)
        return
      }
    }
}
