package iotanalyticsdatasetlifecyclenotification

type IoTAnalyticsDataSetLifeCycleNotification struct {
    EventDetailVersion string `json:"event-detail-version"`
    VersionId string `json:"version-id"`
    State string `json:"state"`
    Message string `json:"message"`
    DatasetName string `json:"dataset-name"`
    ContentDeliveryRuleIndex float64 `json:"content-delivery-rule-index"`
}

func (i *IoTAnalyticsDataSetLifeCycleNotification) SetEventDetailVersion(eventDetailVersion string) {
    i.EventDetailVersion = eventDetailVersion
}

func (i *IoTAnalyticsDataSetLifeCycleNotification) SetVersionId(versionId string) {
    i.VersionId = versionId
}

func (i *IoTAnalyticsDataSetLifeCycleNotification) SetState(state string) {
    i.State = state
}

func (i *IoTAnalyticsDataSetLifeCycleNotification) SetMessage(message string) {
    i.Message = message
}

func (i *IoTAnalyticsDataSetLifeCycleNotification) SetDatasetName(datasetName string) {
    i.DatasetName = datasetName
}

func (i *IoTAnalyticsDataSetLifeCycleNotification) SetContentDeliveryRuleIndex(contentDeliveryRuleIndex float64) {
    i.ContentDeliveryRuleIndex = contentDeliveryRuleIndex
}
