package sagemakertransformjobstatechange

type DataSource struct {
    S3DataSource S3DataSource `json:"S3DataSource"`
}

func (d *DataSource) SetS3DataSource(s3DataSource S3DataSource) {
    d.S3DataSource = s3DataSource
}
