package spacexstats

import (
    "encoding/json"
    "fmt"
    "net/http"
    "bufio"
    "strings"
    "time"
)

type NextMission struct{
    Title string
    FullTitle string
    Description string

    Name string
    Slug string
    Contractor string
    Summary string
    Status string

    LaunchDateTime time.Time
    TimeKnown bool

    Created time.Time
    Updated time.Time
}

func Fetch() (*NextMission, error) {
    resp, err := http.Get("https://spacexstats.com/")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    reader := bufio.NewReader(resp.Body)
    line, err := reader.ReadString('\n')
    if err != nil {
        return nil, err
    }

    rawJson := strings.TrimSuffix(strings.TrimPrefix(line, "<script>window.laravel = window.laravel || {};laravel.statistics = "), ";</script><!DOCTYPE html>\n")

    jsonObj := [][]map[string]interface{}{}
    err = json.Unmarshal([]byte(rawJson), &jsonObj)
    if err != nil {
        return nil, err
    }

    m := jsonObj[0][0]
    mr := m["result"].(map[string]interface{})

    loc, err := time.LoadLocation("America/New_York")
    if err != nil {
        return nil, err
    }

    launchTime, err := time.ParseInLocation("2006-01-02 15:04:05", mr["launch_date_time"].(string), loc)
    if err != nil {
        return nil, err
    }
    launchTime = launchTime.UTC()

    created, err := time.ParseInLocation("2006-01-02 15:04:05", mr["created_at"].(string), loc)
    if err != nil {
        return nil, err
    }
    created = created.UTC()

    updated, err := time.ParseInLocation("2006-01-02 15:04:05", mr["updated_at"].(string), loc)
    if err != nil {
        return nil, err
    }
    updated = updated.UTC()

    return &NextMission{
        Title: m["name"].(string),
        FullTitle: m["full_title"].(string),
        Description: m["description"].(string),

        Name: mr["name"].(string),
        Slug: mr["slug"].(string),
        Contractor: mr["contractor"].(string),
        Summary: mr["summary"].(string),
        Status: mr["status"].(string),

        LaunchDateTime: launchTime,
        TimeKnown: int(mr["launch_specificity"].(float64)) != 6,

        Created: created,
        Updated: updated,
    }, nil
}

func (self *NextMission) PrintSummary() {
    launchTimeFormat := "2006-01-02 15:04"
    if !self.TimeKnown {
        launchTimeFormat = strings.TrimSuffix("2006-01-02 15:04", " 15:04")
    }
    launchTime := self.LaunchDateTime.Format(launchTimeFormat) + " UTC"
    fmt.Printf("# %s\n%s\n\nStatus: %s\nLaunch time: %s\nCreated: %s\nUpdated: %s\n\n",
        self.FullTitle,
        self.Description,
        self.Status,
        launchTime,
        self.Created.Format("2006-01-02 15:04") + " UTC",
        self.Updated.Format("2006-01-02 15:04") + " UTC",
    )
}

func (self *NextMission) StartCountdown() {
    for {
        select {
        case <-time.After(1 * time.Second):
            launch := self.formatCountdownToLaunch()
            fmt.Printf("\rCountdown until Launch: %s", launch)
        }
    }
}

func (self *NextMission) formatCountdownToLaunch() string {
    ld := self.LaunchDateTime
    format := "%d days, %d hours, %d minutes, %d seconds"

    if !ld.After(time.Now().UTC()) {
        return "NOW / Already launched"
    }

    if !self.TimeKnown {
        ld = ld.Round(24 * time.Hour)
        format = format + " (specific time unknown)"
    }

    d := ld.Sub(time.Now().UTC())
    hours := int(d.Hours())
    minutes := int(d.Minutes()) - hours * 60
    seconds := int(d.Seconds()) - hours * 60 * 60 - minutes * 60

    days := hours / 24
    hours = hours - days * 24

    return fmt.Sprintf(format, days, hours, minutes, seconds)
}
