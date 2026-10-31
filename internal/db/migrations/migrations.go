package migrations

import "embed"

// FS contains ordered SQL migrations. Add the next change as
// 00000N_description.up.sql and a matching .down.sql file.
//
//go:embed *.sql
var FS embed.FS
