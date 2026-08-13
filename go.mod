module github.com/TeamWarp/warp-sdk-go

go 1.22

require (
	github.com/tidwall/gjson v1.18.0
	github.com/tidwall/sjson v1.2.5
)

require (
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
)

retract [v0.4.0, v0.4.2] // Published under an incorrect module path; use github.com/TeamWarp/warp-go-sdk instead.

retract [v0.4.0, v0.5.1] // Published under an incorrect module path; use github.com/TeamWarp/warp-go-sdk instead.
