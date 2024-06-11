package awsapicallviacloudtrail

type Settings struct {
    Timeseries Timeseries `json:"timeseries,omitempty"`
}

func (s *Settings) SetTimeseries(timeseries Timeseries) {
    s.Timeseries = timeseries
}
