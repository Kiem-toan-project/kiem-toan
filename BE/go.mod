module github.com/Kiem-toan-project/kiem-toan/tree/main/BE

go 1.16

replace backend => ./

require (
	backend v0.0.0-00010101000000-000000000000
	github.com/go-resty/resty/v2 v2.6.0
	gopkg.in/yaml.v2 v2.4.0
)
