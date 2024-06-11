package schemacreated

import (
    "time"
)


type SchemaCreated struct {
    CreationDate time.Time `json:"CreationDate"`
    RegistryName string `json:"RegistryName"`
    SchemaName string `json:"SchemaName"`
    SchemaType string `json:"SchemaType"`
    Version string `json:"Version"`
}

func (s *SchemaCreated) SetCreationDate(creationDate time.Time) {
    s.CreationDate = creationDate
}

func (s *SchemaCreated) SetRegistryName(registryName string) {
    s.RegistryName = registryName
}

func (s *SchemaCreated) SetSchemaName(schemaName string) {
    s.SchemaName = schemaName
}

func (s *SchemaCreated) SetSchemaType(schemaType string) {
    s.SchemaType = schemaType
}

func (s *SchemaCreated) SetVersion(version string) {
    s.Version = version
}
