# Changelog

## v0.1.4 (2026-07-16)

Looking at this diff, I can identify the following user-facing changes:

## Breaking changes

- **Package name changed**: `warphr` → `warpgosdk`. Update your imports from `"github.com/TeamWarp/warp-go-sdk"` (which still resolves to package `warphr`) to use package `warpgosdk` instead.

- **Client initialization now returns a pointer**: `NewClient()` now returns `*Client` instead of `Client`. You don't need to change call sites (Go handles this automatically), but be aware if you were taking the address of the result.

- **Service types now use pointers**: All service types (`TimeOffService`, `WorkerService`, `DepartmentService`, `WorkplaceService`) are now returned as pointers from their constructors. This is internal refactoring that shouldn't affect normal usage via the client.

- **Parameter field API changed**: The SDK now uses `param.Field[T]` instead of `param.Opt[T]`. Use the new `F[T]()` helper function to wrap optional parameters:
  - Old: `DepartmentNewParams{Name: "foo"}` 
  - New: `DepartmentNewParams{Name: sdk.F[string]("foo")}`

- **Removed pagination wrapper**: `List()` methods now return the response type directly (e.g., `*DepartmentListResponse`) instead of `*pagination.CursorPage[T]`. The `ListAutoPaging()` methods have been removed.

- **Removed `respjson` metadata**: Response types no longer include the `.JSON` field with `respjson.Field` metadata. Access raw JSON via `RawJSON()` method if needed.

## Changed

- Department methods reordered: `List()` is now the first method (was last), followed by `New()` and `Update()`.
- URL path encoding simplified: `url.PathEscape()` calls removed (paths are now escaped automatically).

## Added

- New `default_http_client.go` provides automatic response header timeout configuration for better handling of stuck connections.
- New `VERSIONING.md` documents the SDK's manual versioning policy.

Minor internal updates.

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
