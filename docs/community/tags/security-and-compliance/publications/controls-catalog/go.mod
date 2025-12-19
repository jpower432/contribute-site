module github.com/cncf/contribute-site/controls-catalog

go 1.23.0

toolchain go1.24.10

require (
	github.com/ossf/gemara v0.17.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/goccy/go-yaml v1.19.0 // indirect
	github.com/kr/text v0.2.0 // indirect
)

replace github.com/ossf/gemara => /tmp/gemara
