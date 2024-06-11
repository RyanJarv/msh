package gluedatacatalogdatabasestatechange

type GlueDataCatalogDatabaseStateChange struct {
    DatabaseName string `json:"databaseName"`
    TypeOfChange string `json:"typeOfChange"`
    ChangedTables []string `json:"changedTables"`
}

func (g *GlueDataCatalogDatabaseStateChange) SetDatabaseName(databaseName string) {
    g.DatabaseName = databaseName
}

func (g *GlueDataCatalogDatabaseStateChange) SetTypeOfChange(typeOfChange string) {
    g.TypeOfChange = typeOfChange
}

func (g *GlueDataCatalogDatabaseStateChange) SetChangedTables(changedTables []string) {
    g.ChangedTables = changedTables
}
