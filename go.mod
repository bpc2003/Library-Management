module backend

go 1.21.6

require (
	backend/models v0.0.0-00010101000000-000000000000
	backend/routers v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.9
)

require backend/controllers v0.0.0-00010101000000-000000000000 // indirect

replace backend/models => ./models

replace backend/controllers => ./controllers

replace backend/routers => ./routers
