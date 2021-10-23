module launcher-tables

go 1.16

require (
	github.com/kolide/kit v0.0.0-20191023141830-6312ecc11c23
	github.com/kolide/launcher v0.11.22
	github.com/osquery/osquery-go v0.0.0-20210622151333-99b4efa62ec5
)

replace github.com/knightsc/system_policy => github.com/zwass/system_policy v1.1.1-0.20211023013151-0e75d1be2802

replace github.com/kolide/launcher => github.com/zwass/launcher v0.0.0-20211023013720-3738413857ac
