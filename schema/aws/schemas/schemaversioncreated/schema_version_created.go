package schemaversioncreated

import (
    "time"
)


type SchemaVersionCreated struct {
    CreationDate time.Time `json:"CreationDate"`
    RegistryName string `json:"RegistryName"`
    SchemaName string `json:"SchemaName"`
    SchemaType string `json:"SchemaType"`
    Version string `json:"Version"`
}

func (s *SchemaVersionCreated) SetCreationDate(creationDate time.Time) {
    s.CreationDate = creationDate
}

func (s *SchemaVersionCreated) SetRegistryName(registryName string) {
    s.RegistryName = registryName
}

func (s *SchemaVersionCreated) SetSchemaName(schemaName string) {
    s.SchemaName = schemaName
}

func (s *SchemaVersionCreated) SetSchemaType(schemaType string) {
    s.SchemaType = schemaType
}

func (s *SchemaVersionCreated) SetVersion(version string) {
    s.Version = version
}
