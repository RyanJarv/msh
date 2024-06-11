package maciealert

import (
    "time"
)


type SummaryItem struct {
    Count float64 `json:"count"`
    Start time.Time `json:"start"`
    End time.Time `json:"end"`
}

func (s *SummaryItem) SetCount(count float64) {
    s.Count = count
}

func (s *SummaryItem) SetStart(start time.Time) {
    s.Start = start
}

func (s *SummaryItem) SetEnd(end time.Time) {
    s.End = end
}
