package mediapackageharvestjobnotification

import (
    "time"
)


type Harvest_job struct {
    S3Destination S3_destination `json:"s3_destination"`
    Arn string `json:"arn"`
    CreatedAt time.Time `json:"created_at"`
    EndTime time.Time `json:"end_time"`
    Id string `json:"id"`
    OriginEndpointId string `json:"origin_endpoint_id"`
    StartTime time.Time `json:"start_time"`
    Status string `json:"status"`
}

func (h *Harvest_job) SetS3Destination(s3Destination S3_destination) {
    h.S3Destination = s3Destination
}

func (h *Harvest_job) SetArn(arn string) {
    h.Arn = arn
}

func (h *Harvest_job) SetCreatedAt(createdAt time.Time) {
    h.CreatedAt = createdAt
}

func (h *Harvest_job) SetEndTime(endTime time.Time) {
    h.EndTime = endTime
}

func (h *Harvest_job) SetId(id string) {
    h.Id = id
}

func (h *Harvest_job) SetOriginEndpointId(originEndpointId string) {
    h.OriginEndpointId = originEndpointId
}

func (h *Harvest_job) SetStartTime(startTime time.Time) {
    h.StartTime = startTime
}

func (h *Harvest_job) SetStatus(status string) {
    h.Status = status
}
