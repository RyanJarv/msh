package guarddutyfinding

import (
    "time"
)


type Service struct {
    Action Action `json:"action,omitempty"`
    AdditionalInfo AdditionalInfo `json:"additionalInfo"`
    AwsApiCallAction AwsApiCallAction `json:"awsApiCallAction,omitempty"`
    EbsVolumeScanDetails EbsVolumeScanDetails `json:"ebsVolumeScanDetails,omitempty"`
    Evidence Evidence `json:"evidence,omitempty"`
    Archived bool `json:"archived"`
    Count float64 `json:"count"`
    DetectorId string `json:"detectorId"`
    EventFirstSeen time.Time `json:"eventFirstSeen"`
    EventLastSeen time.Time `json:"eventLastSeen"`
    FeatureName string `json:"featureName,omitempty"`
    ResourceRole string `json:"resourceRole,omitempty"`
    ServiceName string `json:"serviceName"`
}

func (s *Service) SetAction(action Action) {
    s.Action = action
}

func (s *Service) SetAdditionalInfo(additionalInfo AdditionalInfo) {
    s.AdditionalInfo = additionalInfo
}

func (s *Service) SetAwsApiCallAction(awsApiCallAction AwsApiCallAction) {
    s.AwsApiCallAction = awsApiCallAction
}

func (s *Service) SetEbsVolumeScanDetails(ebsVolumeScanDetails EbsVolumeScanDetails) {
    s.EbsVolumeScanDetails = ebsVolumeScanDetails
}

func (s *Service) SetEvidence(evidence Evidence) {
    s.Evidence = evidence
}

func (s *Service) SetArchived(archived bool) {
    s.Archived = archived
}

func (s *Service) SetCount(count float64) {
    s.Count = count
}

func (s *Service) SetDetectorId(detectorId string) {
    s.DetectorId = detectorId
}

func (s *Service) SetEventFirstSeen(eventFirstSeen time.Time) {
    s.EventFirstSeen = eventFirstSeen
}

func (s *Service) SetEventLastSeen(eventLastSeen time.Time) {
    s.EventLastSeen = eventLastSeen
}

func (s *Service) SetFeatureName(featureName string) {
    s.FeatureName = featureName
}

func (s *Service) SetResourceRole(resourceRole string) {
    s.ResourceRole = resourceRole
}

func (s *Service) SetServiceName(serviceName string) {
    s.ServiceName = serviceName
}
