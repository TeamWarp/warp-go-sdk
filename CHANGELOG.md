# Changelog

## v0.1.4 (2026-07-15)

# Breaking changes

- **Package import path changed**: Update imports from `github.com/TeamWarp/warp-go-sdk` to `github.com/marclave/warp-go-sdk`, and the package name from `warphr` to `warpgosdk`.
- **NewClient now returns a pointer**: `NewClient()` returns `*Client` instead of `Client`. This is a breaking change if you were storing the result as a value type.
- **Service field access changed**: Service fields on the client (e.g., `client.TimeOff`) are now pointers (`*TimeOffService`), and service `Options` field is now public (`Options` instead of `options`).
- **DepartmentService.List signature changed**: Now returns `*DepartmentListResponse` instead of a paginated cursor page. Update code that was iterating through paginated results.
- **Response types restructured**: Department, Worker, and Workplace response types have changed structure—they now include `Issues`, `Message`, and `Tag` fields. Check your code if you were accessing specific fields from these responses.
- **Method signatures simplified**: All Execute/Get/Post/Patch/Delete methods now use `interface{}` instead of `any` for params and response types (functionally equivalent but note if you have type constraints).

# Changed

- Helper function `Opt[T]()` is now `F[T]()` for creating field values.
- Helper functions like `StringPtr()`, `IntPtr()`, etc. have been removed; use `F()` or `Null()` directly.
- `File()` helper is now `FileParam()`.
- Query parameter types changed from `param.Opt[T]` to `param.Field[T]`.
- Request body parameter types changed from `param.Opt[T]` to `param.Field[T]`.
- `URLQuery()` method on params no longer returns an error; signature is now `URLQuery() url.Values`.
- Path escaping removed in some methods (e.g., `url.PathEscape()` calls removed).

# Added

- New `VERSIONING.md` documenting the manual versioning policy.
- New `defaultHTTPClient()` function for better HTTP connection handling.
- Extended environment variable support: `WARP_CUSTOM_HEADERS` can now be set to define custom headers via newline-separated `key: value` pairs.

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
