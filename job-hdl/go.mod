module github.com/SENERGY-Platform/go-service-base/job-hdl

go 1.21

require (
	github.com/SENERGY-Platform/go-cc-job-handler v0.1.1
	github.com/SENERGY-Platform/go-service-base/job-hdl/lib v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.3.1
)

replace github.com/SENERGY-Platform/go-service-base/job-hdl/lib => ./lib
