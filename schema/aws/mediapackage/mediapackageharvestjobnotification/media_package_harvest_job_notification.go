package mediapackageharvestjobnotification

type MediaPackageHarvestJobNotification struct {
    HarvestJob Harvest_job `json:"harvest_job"`
}

func (m *MediaPackageHarvestJobNotification) SetHarvestJob(harvestJob Harvest_job) {
    m.HarvestJob = harvestJob
}
