package commands

import (
    "encoding/json"
    "fmt"
    "os"
)

type Note struct {
    Text string `json:"text"`
}

func SaveNote(text string) {
    var notes []Note

    file, _ := os.ReadFile("data.json")
    json.Unmarshal(file, &notes)

    notes = append(notes, Note{Text: text})

    updated, _ := json.Marshal(notes)
    os.WriteFile("data.json", updated, 0644)

    fmt.Println("Note saved!")
}

func ShowNotes() {
    var notes []Note

    file, _ := os.ReadFile("data.json")
    json.Unmarshal(file, &notes)

    for i, n := range notes {
        fmt.Println(i+1, "-", n.Text)
    }
}