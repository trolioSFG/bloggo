module github.com/trolioSFG/bloggo

go 1.23.2

replace github.com/trolioSFG/blogconfig => ./internal/config

replace github.com/trolioSFG/database => ./internal/database

require (
	github.com/lib/pq v1.10.9
	github.com/trolioSFG/blogconfig v0.0.0-00010101000000-000000000000
	github.com/trolioSFG/database v0.0.0-00010101000000-000000000000
)

require github.com/google/uuid v1.6.0 // indirect
