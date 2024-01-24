module backend/routers

go 1.21.6

replace backend/controllers => ../controllers

require (
	backend/controllers v0.0.0-00010101000000-000000000000
	backend/middleware v0.0.0-00010101000000-000000000000
)

replace backend/middleware => ../middleware
