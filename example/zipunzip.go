package main

import (
    "simplezip"
      "log"
)


func main() {

    sz := NewSimpleZip()

    if err := zipDirectory("./", "../testFolder.zip"); err != nil {
        log.Fatal(err)
    }

}


