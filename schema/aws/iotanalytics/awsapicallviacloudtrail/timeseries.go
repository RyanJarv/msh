package awsapicallviacloudtrail

type Timeseries struct {
    DimensionAttributes []TimeseriesItem `json:"dimensionAttributes,omitempty"`
    MeasureAttributes []TimeseriesItem `json:"measureAttributes,omitempty"`
    TimestampAttribute string `json:"timestampAttribute,omitempty"`
}

func (t *Timeseries) SetDimensionAttributes(dimensionAttributes []TimeseriesItem) {
    t.DimensionAttributes = dimensionAttributes
}

func (t *Timeseries) SetMeasureAttributes(measureAttributes []TimeseriesItem) {
    t.MeasureAttributes = measureAttributes
}

func (t *Timeseries) SetTimestampAttribute(timestampAttribute string) {
    t.TimestampAttribute = timestampAttribute
}
