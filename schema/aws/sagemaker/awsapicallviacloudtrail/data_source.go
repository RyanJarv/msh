package awsapicallviacloudtrail

type DataSource struct {
    S3DataSource S3DataSource `json:"s3DataSource,omitempty"`
}

func (d *DataSource) SetS3DataSource(s3DataSource S3DataSource) {
    d.S3DataSource = s3DataSource
}
