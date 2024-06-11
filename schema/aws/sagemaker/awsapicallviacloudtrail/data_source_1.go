package awsapicallviacloudtrail

type DataSource_1 struct {
    S3DataSource S3DataSource_1 `json:"s3DataSource"`
}

func (d *DataSource_1) SetS3DataSource(s3DataSource S3DataSource_1) {
    d.S3DataSource = s3DataSource
}
