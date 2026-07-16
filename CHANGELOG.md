# Changelog

## v0.1.4 (2026-07-16)

- **Breaking change**: The package name has changed from `warphr` to `warpgosdk`. Update all imports from `"github.com/TeamWarp/warp-go-sdk"` to use `warpgosdk` instead of `warphr`.

- **Breaking change**: `NewClient()` now returns `*Client` instead of `Client`. Update any code that doesn't already dereference the return value.

- **Breaking change**: Service fields on `Client` are now pointers (e.g., `client.TimeOff` is `*TimeOffService`). This shouldn't affect normal usage but may impact code that directly accesses these fields.

- **Breaking change**: The parameter builder functions have changed. Replace `String(value)`, `Int(value)`, `Bool(value)`, `Float(value)` calls with `F[T](value)` where `T` is the type. For example: `String("name")` becomes `F[string]("name")`.

- **Breaking change**: Response pagination has changed. `DepartmentService.List()` and similar methods now return `*DepartmentListResponse` directly instead of `*pagination.CursorPage[DepartmentListResponse]`. The `ListAutoPaging()` methods have been removed.

- **Breaking change**: `DepartmentService.List()` has moved to be called before `New()` and `Update()` in the service.

- **Changed**: Parameter structs no longer embed `paramObj`. Use `param.Field[T]` directly with the `F[]` helper for optional fields.

- **Changed**: Error handling and JSON unmarshaling have been refactored internally. The API surface remains the same, but error types and response JSON field handling are slightly different.

## 0.3.0 (2026-03-27)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/TeamWarp/warp-go-sdk/compare/v0.2.0...v0.3.0)

### Features

* **api:** update import names to warp ([b5716fb](https://github.com/TeamWarp/warp-go-sdk/commit/b5716fb6c2b7ce0a01d4adf22d60fc6c570fb887))

## 0.2.0 (2026-03-27)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/TeamWarp/warp-go-sdk/compare/v0.1.0...v0.2.0)

### Features

* **api:** update contact email ([a0a0026](https://github.com/TeamWarp/warp-go-sdk/commit/a0a0026c2373e72554e46c6ea1b46300be005419))

## 0.1.0 (2026-03-27)

Full Changelog: [v0.0.2...v0.1.0](https://github.com/TeamWarp/warp-go-sdk/compare/v0.0.2...v0.1.0)

### Features

* **api:** api update ([eec6538](https://github.com/TeamWarp/warp-go-sdk/commit/eec65386699c3916c2b08b016459b56ba9d85971))

## 0.0.2 (2026-03-27)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/TeamWarp/warp-go-sdk/compare/v0.0.1...v0.0.2)

### Chores

* configure new SDK language ([8dc14ec](https://github.com/TeamWarp/warp-go-sdk/commit/8dc14ecbba6683d8f8aaf719d4af3bd78a6bb889))
* update SDK settings ([d217727](https://github.com/TeamWarp/warp-go-sdk/commit/d217727a6569fb3b7627ba2e5d5a674e6418af84))
