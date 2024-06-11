package awsapicallviacloudtrail

type Badge struct {
    BadgeEnabled bool `json:"badgeEnabled,omitempty"`
    BadgeRequestUrl string `json:"badgeRequestUrl,omitempty"`
}

func (b *Badge) SetBadgeEnabled(badgeEnabled bool) {
    b.BadgeEnabled = badgeEnabled
}

func (b *Badge) SetBadgeRequestUrl(badgeRequestUrl string) {
    b.BadgeRequestUrl = badgeRequestUrl
}
