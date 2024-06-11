package devopsgurunewanomalyassociation

type SourceDetail struct {
    DataIdentifiers DataIdentifiers `json:"dataIdentifiers,omitempty"`
    DataSource string `json:"dataSource,omitempty"`
    DataSourceDetail string `json:"dataSourceDetail,omitempty"`
}

func (s *SourceDetail) SetDataIdentifiers(dataIdentifiers DataIdentifiers) {
    s.DataIdentifiers = dataIdentifiers
}

func (s *SourceDetail) SetDataSource(dataSource string) {
    s.DataSource = dataSource
}

func (s *SourceDetail) SetDataSourceDetail(dataSourceDetail string) {
    s.DataSourceDetail = dataSourceDetail
}
