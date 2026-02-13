package db

import _ "embed"

//go:embed schema.sql
var DDL string

//go:embed seeds/seeds.sql
var Seeds string
