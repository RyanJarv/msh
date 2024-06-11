package guarddutyfinding

type AffectedResources_1 struct {
    AWSCloudTrailTrail string `json:"AWS-CloudTrail-Trail,omitempty"`
    AWSEC2Instance string `json:"AWS-EC2-Instance,omitempty"`
    AWSS3Bucket string `json:"AWS-S3-Bucket,omitempty"`
}

func (a *AffectedResources_1) SetAWSCloudTrailTrail(aWSCloudTrailTrail string) {
    a.AWSCloudTrailTrail = aWSCloudTrailTrail
}

func (a *AffectedResources_1) SetAWSEC2Instance(aWSEC2Instance string) {
    a.AWSEC2Instance = aWSEC2Instance
}

func (a *AffectedResources_1) SetAWSS3Bucket(aWSS3Bucket string) {
    a.AWSS3Bucket = aWSS3Bucket
}
