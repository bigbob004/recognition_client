package database

import (
	sq "github.com/Masterminds/squirrel"
)

// Формат плэйсхолдоров для запросов в бд
var PSQL = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
