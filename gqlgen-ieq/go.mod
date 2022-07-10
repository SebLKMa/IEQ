module github.com/seblkma/ieq/gqlgen-ieq

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/lib/pq v1.9.0 // indirect
	github.com/seblkma/ieq/db/postgres v0.0.0-00010101000000-000000000000
	github.com/seblkma/ieq/models v0.0.0-00010101000000-000000000000 // indirect
	github.com/vektah/gqlparser/v2 v2.1.0
)

replace (
	github.com/seblkma/ieq/db/postgres => ../db/postgres
	github.com/seblkma/ieq/models => ../models
)