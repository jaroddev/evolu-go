package record

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Row struct {
	Generation  int
	BestFitness float64
}

type Recorder struct {
	Rows []Row
}

func NewRecorder() Recorder {
	return Recorder{
		Rows: make([]Row, 0),
	}
}

func (recorder *Recorder) Record(Generation int, BestFitness float64) {
	newRow := &Row{Generation, BestFitness}
	recorder.Rows = append(recorder.Rows, *newRow)
}

func (r *Recorder) Save() {
	records := [][]string{
		{"generation", "fitness"},
	}

	w := csv.NewWriter(os.Stdout)

	for _, row := range r.Rows {
		records = append(records, []string{
			fmt.Sprint(row.Generation), fmt.Sprint(row.BestFitness),
		})
	}

	w.WriteAll(records) // calls Flush internally
	if err := w.Error(); err != nil {
		fmt.Println("whouah, got an error man")
	}
}
