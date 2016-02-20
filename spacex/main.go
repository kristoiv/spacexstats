package main

import (
    "log"
    "os"

    "github.com/kristoiv/spacexstats"
)

func main() {
    logger := log.New(os.Stderr, "[spacexstats] ", log.LstdFlags)

    sx, err := spacexstats.Fetch()
    if err != nil {
        logger.Fatalln(err)
    }

    sx.PrintSummary()
    sx.StartCountdown()
}
