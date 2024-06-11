package gluedatacatalogtablestatechange

type GlueDataCatalogTableStateChange struct {
    DatabaseName string `json:"databaseName"`
    ChangedPartitions []string `json:"changedPartitions"`
    TypeOfChange string `json:"typeOfChange"`
    TableName string `json:"tableName"`
}

func (g *GlueDataCatalogTableStateChange) SetDatabaseName(databaseName string) {
    g.DatabaseName = databaseName
}

func (g *GlueDataCatalogTableStateChange) SetChangedPartitions(changedPartitions []string) {
    g.ChangedPartitions = changedPartitions
}

func (g *GlueDataCatalogTableStateChange) SetTypeOfChange(typeOfChange string) {
    g.TypeOfChange = typeOfChange
}

func (g *GlueDataCatalogTableStateChange) SetTableName(tableName string) {
    g.TableName = tableName
}
