package main

import (
    "github.com/seldonsmule/simplezip"
      "log"
      "fmt"
      "flag"
//      "os"
)

func help(){

  fmt.Println("zipunzip - Example uses of the simplezip package")
  fmt.Println()
  flag.PrintDefaults()
  fmt.Println()
  fmt.Println("cmds:")
  fmt.Println("dzip - Zips a directory")
  fmt.Println("dunzip - UnZips a directory")

}


func main() {

    cmdPtr := flag.String("cmd", "help", "Command to run")
    targetPtr := flag.String("target", "tmp.zip", "Command to run")
    sourcePtr := flag.String("source", "./", "Command to run")

    flag.Parse()

    sz := simplezip.NewSimpleZip()

    switch *cmdPtr {

      case "dzip":
        fmt.Printf("Directory Zip source[%s] target[%s]\n", 
                   *sourcePtr, *targetPtr)
        if err := sz.ZipDir(*sourcePtr, *targetPtr); err != nil {
            log.Fatal(err)
        }

      case "dunzip":
        fmt.Printf("Directory UnZip source[%s] target[%s]\n", 
                   *sourcePtr, *targetPtr)

        files, err := sz.UnZipDir(*sourcePtr, *targetPtr)

        if err != nil {
            log.Fatal(err)
        }

        j := 0

        for range files {

          fmt.Printf("Index[%d] is name[%s]\n", j, files[j])

          j++

        }


      case "help":
        help()

      default:
        fmt.Printf("Unknown command: %s\n", *cmdPtr)

    }


/*
    if err := sz.ZipDir("./", "../testFolder.zip"); err != nil {
        log.Fatal(err)
    }
*/

}


