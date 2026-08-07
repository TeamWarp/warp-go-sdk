# Changelog

## [0.3.0](https://github.com/TeamWarp/warp-sdk-go/compare/v0.2.0...v0.3.0) (2026-08-07)


### ⚠ BREAKING CHANGES

* **api:** 6 breaking changes to the SDK surface.
    - Removed operation `benefits.healthPlans.benefitsList` (`GET /v1/benefits/health_plans`).
    - Removed operation `benefits.healthPlans.benefitsGet` (`GET /v1/benefits/health_plans/{id}`).
    - Removed operation `benefits.retirementPlans.benefitsList` (`GET /v1/benefits/retirement_plans`).
    - Removed operation `benefits.retirementPlans.benefitsGet` (`GET /v1/benefits/retirement_plans/{id}`).
    - Removed operation `benefits.deductions.benefitsList` (`GET /v1/benefits/deductions`).
    - Removed operation `benefits.deductions.benefitsGet` (`GET /v1/benefits/deductions/{id}`).
* **api:** Renamed SDK from `WarpApi` to `Warp`.
* **api:** 20 breaking changes to the SDK surface.
    - Removed operation `customWorkerFields.list` (`GET /v1/custom-worker-fields`).
    - Removed operation `customWorkerFields.create` (`POST /v1/custom-worker-fields`).
    - Removed operation `customWorkerFields.retrieve` (`GET /v1/custom-worker-fields/{id}`).
    - Removed operation `customWorkerFields.update` (`PATCH /v1/custom-worker-fields/{id}`).
    - Removed operation `customWorkerFields.archive` (`POST /v1/custom-worker-fields/{id}/archive`).
    - Removed operation `customWorkerFields.createOption` (`POST /v1/custom-worker-fields/{id}/options`).
    - Removed operation `customWorkerFields.updateOption` (`PATCH /v1/custom-worker-field-options/{id}`).
    - Removed operation `customWorkerFields.deleteOption` (`DELETE /v1/custom-worker-field-options/{id}`).
    - Removed operation `customWorkerFields.archiveOption` (`POST /v1/custom-worker-field-options/{id}/archive`).
    - Removed operation `customWorkerFields.listValues` (`GET /v1/worker-custom-field-values`).
    - Removed operation `customWorkerFields.upsertValue` (`PUT /v1/worker-custom-field-values`).
    - Removed operation `customWorkerFields.clearValue` (`DELETE /v1/worker-custom-field-values`).
    - Removed schema `invalid_custom_worker_field_operation`.
    - Removed schema `custom_worker_field_already_exists`.
    - Removed schema `custom_worker_field_option_already_exists`.
    - Removed schema `custom_worker_field_not_found`.
    - Removed schema `custom_worker_field_option_not_found`.
    - Removed schema `custom_worker_field_option_in_use`.
    - Removed schema `invalid_custom_worker_field_value`.
    - Removed schema `custom_worker_field_worker_not_found`.
* **api:** 86 breaking changes to the SDK surface.
    - Renamed SDK from `Warp` to `WarpApi`.
    - `400` error response of `departments.list` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `departments.list` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `429` error response of `departments.list` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `departments.create` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `departments.create` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `409` error response of `departments.create` changed from `DepartmentAlreadyExists` to `department_already_exists`.
    - `429` error response of `departments.create` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `departments.update` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `departments.update` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `404` error response of `departments.update` changed from `DepartmentNotFound` to `department_not_found`.
    - `409` error response of `departments.update` changed from `DepartmentAlreadyExists` to `department_already_exists`.
    - `429` error response of `departments.update` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `timeOff.listAssignments` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `timeOff.listAssignments` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `429` error response of `timeOff.listAssignments` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `timeOff.listBalances` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `timeOff.listBalances` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `404` error response of `timeOff.listBalances` changed from `TimeOffPolicyNotFound` to `time_off_policy_not_found`.
    - `429` error response of `timeOff.listBalances` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `timeOff.listRequests` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `timeOff.listRequests` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `404` error response of `timeOff.listRequests` changed from `TimeOffRequestNotFoundError` to `time_off_request_not_found_error`.
    - `429` error response of `timeOff.listRequests` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `workers.list` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `workers.list` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `429` error response of `workers.list` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `workers.retrieve` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `workers.retrieve` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `404` error response of `workers.retrieve` changed from `WorkerNotFoundError` to `worker_not_found_error`.
    - `429` error response of `workers.retrieve` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `workers.delete` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `workers.delete` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `404` error response of `workers.delete` changed from `WorkerNotFoundError` to `worker_not_found_error`.
    - `409` error response of `workers.delete` changed from `CannotDeleteWorker` to `cannot_delete_worker`.
    - `429` error response of `workers.delete` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `401` error response of `workers.createEmployee` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `429` error response of `workers.createEmployee` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `workers.createContractor` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `workers.createContractor` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `429` error response of `workers.createContractor` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `workers.invite` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `workers.invite` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `404` error response of `workers.invite` changed from `WorkerNotFoundError` to `worker_not_found_error`.
    - `429` error response of `workers.invite` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `workplaces.list` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `workplaces.list` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `429` error response of `workplaces.list` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `401` error response of `workplaces.create` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `409` error response of `workplaces.create` changed from `WorkplaceAlreadyExists` to `workplace_already_exists`.
    - `429` error response of `workplaces.create` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `workplaces.update` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `workplaces.update` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `404` error response of `workplaces.update` changed from `WorkplaceNotFound` to `workplace_not_found`.
    - `409` error response of `workplaces.update` changed from `WorkplaceAlreadyExists` to `workplace_already_exists`.
    - `429` error response of `workplaces.update` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - Removed operation `timeOff.policies.list` (`GET /v1/time_off/policies`).
    - Removed operation `timeOff.policies.retrieve` (`GET /v1/time_off/policies/{id}`).
    - Removed schema `Date`.
    - Removed schema `HttpApiDecodeError`.
    - Removed schema `Issue`.
    - Removed schema `PropertyKey`.
    - Removed schema `InternalServerError`.
    - Removed schema `ApiKeyUnauthorized`.
    - Removed schema `RateLimitExceeded`.
    - Removed schema `DateTimeUtc`.
    - Removed schema `MissingRequiredCompanyPermissions`.
    - Removed schema `ApiNotEnabled`.
    - Removed schema `Trimmed`.
    - Removed schema `DepartmentAlreadyExists`.
    - Removed schema `DepartmentNotFound`.
    - Removed schema `ManagerNotFoundError`.
    - Removed schema `TimeOffPolicyNotFound`.
    - Removed schema `TimeOffRequestNotFoundError`.
    - Removed schema `WorkerNotFoundError`.
    - Removed schema `OfficeWorkLocation`.
    - Removed schema `RemoteWorkLocation`.
    - Removed schema `StateRegistrationRequired`.
    - Removed schema `WorkplaceNotFound`.
    - Removed schema `PayScheduleNotConfigured`.
    - Removed schema `SubscriptionLimitError`.
    - Removed schema `InvalidWorkerStatusError`.
    - Removed schema `WorkerAlreadyExistsError`.
    - Removed schema `CannotDeleteWorker`.
    - Removed schema `AddressInvalid`.
    - Removed schema `WorkplaceAlreadyExists`.
* **api:** Property `MissingRequiredCompanyPermissions.requiredPermissions` type changed from `Array<object>` to `Array<object>`.

### Features

* **api:** add operation customFields.list (+20 more changes) ([9517c37](https://github.com/TeamWarp/warp-sdk-go/commit/9517c37fc8673889431c6688bfa34202ca1896e1))
* **api:** add schema public_money_amount (+1 more change) ([697c33d](https://github.com/TeamWarp/warp-sdk-go/commit/697c33dd6fc59aa298eeb249edca58aec2afae38))
* **api:** initial SDK generation ([8561f11](https://github.com/TeamWarp/warp-sdk-go/commit/8561f11f168401b3ced29c591fe7c59fb1bea64c))
* **api:** remove operation customWorkerFields.list (+19 more changes) ([46a5ca0](https://github.com/TeamWarp/warp-sdk-go/commit/46a5ca0b4e6a63cabed0d0f6806e0f58b28d6e55))
* **api:** update import names to warp ([87831bc](https://github.com/TeamWarp/warp-sdk-go/commit/87831bc83f412d1069ef9435c1ae7dd00823391f))
* **api:** update SDK name (+145 more changes) ([6ac93ee](https://github.com/TeamWarp/warp-sdk-go/commit/6ac93ee5b2ffb457a3c6a21bc75eaeeb2ca5f5fc))
* **api:** update SDK name (+27 more changes) ([1f0c2a5](https://github.com/TeamWarp/warp-sdk-go/commit/1f0c2a57eb0d071d907deef400eeddd839f3ae49))
* **api:** update SDK surface (14 changes) ([7e88858](https://github.com/TeamWarp/warp-sdk-go/commit/7e888580c1b460dd1349b53f8048dc642a243e68))
* **api:** update SDK surface (2 changes) ([ea40d6f](https://github.com/TeamWarp/warp-sdk-go/commit/ea40d6faf17e53ddcdb7b83d1d305602b82897d0))


### Chores

* **api:** regenerate SDK ([cd4d20f](https://github.com/TeamWarp/warp-sdk-go/commit/cd4d20f18cc55b7d99265ebee7b420af23a5d4e9))
* **api:** regenerate SDK ([5f693d8](https://github.com/TeamWarp/warp-sdk-go/commit/5f693d87cd0c64f107eb57356b8ddb22932f0d7a))
* **api:** regenerate SDK ([ac04211](https://github.com/TeamWarp/warp-sdk-go/commit/ac042115bb63da525c083edb0ceb0b6fd6acc527))
* **api:** update generated SDK content ([c30a172](https://github.com/TeamWarp/warp-sdk-go/commit/c30a172959b057054fa6e84320a98cc887889551))
* **api:** update generated SDK content ([e02c75c](https://github.com/TeamWarp/warp-sdk-go/commit/e02c75c417391e38cfbf130f1ee9ba2a6594d10a))

## [0.2.0](https://github.com/TeamWarp/warp-sdk-go/compare/v0.1.0...v0.2.0) (2026-07-29)


### ⚠ BREAKING CHANGES

* **api:** 86 breaking changes to the SDK surface.
    - Renamed SDK from `Warp` to `WarpApi`.
    - `400` error response of `departments.list` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `departments.list` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `429` error response of `departments.list` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `departments.create` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `departments.create` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `409` error response of `departments.create` changed from `DepartmentAlreadyExists` to `department_already_exists`.
    - `429` error response of `departments.create` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `departments.update` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `departments.update` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `404` error response of `departments.update` changed from `DepartmentNotFound` to `department_not_found`.
    - `409` error response of `departments.update` changed from `DepartmentAlreadyExists` to `department_already_exists`.
    - `429` error response of `departments.update` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `timeOff.listAssignments` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `timeOff.listAssignments` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `429` error response of `timeOff.listAssignments` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `timeOff.listBalances` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `timeOff.listBalances` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `404` error response of `timeOff.listBalances` changed from `TimeOffPolicyNotFound` to `time_off_policy_not_found`.
    - `429` error response of `timeOff.listBalances` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `timeOff.listRequests` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `timeOff.listRequests` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `404` error response of `timeOff.listRequests` changed from `TimeOffRequestNotFoundError` to `time_off_request_not_found_error`.
    - `429` error response of `timeOff.listRequests` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `workers.list` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `workers.list` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `429` error response of `workers.list` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `workers.retrieve` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `workers.retrieve` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `404` error response of `workers.retrieve` changed from `WorkerNotFoundError` to `worker_not_found_error`.
    - `429` error response of `workers.retrieve` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `workers.delete` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `workers.delete` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `404` error response of `workers.delete` changed from `WorkerNotFoundError` to `worker_not_found_error`.
    - `409` error response of `workers.delete` changed from `CannotDeleteWorker` to `cannot_delete_worker`.
    - `429` error response of `workers.delete` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `401` error response of `workers.createEmployee` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `429` error response of `workers.createEmployee` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `workers.createContractor` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `workers.createContractor` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `429` error response of `workers.createContractor` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `workers.invite` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `workers.invite` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `404` error response of `workers.invite` changed from `WorkerNotFoundError` to `worker_not_found_error`.
    - `429` error response of `workers.invite` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `workplaces.list` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `workplaces.list` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `429` error response of `workplaces.list` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `401` error response of `workplaces.create` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `409` error response of `workplaces.create` changed from `WorkplaceAlreadyExists` to `workplace_already_exists`.
    - `429` error response of `workplaces.create` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - `400` error response of `workplaces.update` changed from `HttpApiDecodeError` to `http_api_decode_error`.
    - `401` error response of `workplaces.update` changed from `ApiKeyUnauthorized` to `api_key_unauthorized`.
    - `404` error response of `workplaces.update` changed from `WorkplaceNotFound` to `workplace_not_found`.
    - `409` error response of `workplaces.update` changed from `WorkplaceAlreadyExists` to `workplace_already_exists`.
    - `429` error response of `workplaces.update` changed from `RateLimitExceeded` to `rate_limit_exceeded`.
    - Removed operation `timeOff.policies.list` (`GET /v1/time_off/policies`).
    - Removed operation `timeOff.policies.retrieve` (`GET /v1/time_off/policies/{id}`).
    - Removed schema `Date`.
    - Removed schema `HttpApiDecodeError`.
    - Removed schema `Issue`.
    - Removed schema `PropertyKey`.
    - Removed schema `InternalServerError`.
    - Removed schema `ApiKeyUnauthorized`.
    - Removed schema `RateLimitExceeded`.
    - Removed schema `DateTimeUtc`.
    - Removed schema `MissingRequiredCompanyPermissions`.
    - Removed schema `ApiNotEnabled`.
    - Removed schema `Trimmed`.
    - Removed schema `DepartmentAlreadyExists`.
    - Removed schema `DepartmentNotFound`.
    - Removed schema `ManagerNotFoundError`.
    - Removed schema `TimeOffPolicyNotFound`.
    - Removed schema `TimeOffRequestNotFoundError`.
    - Removed schema `WorkerNotFoundError`.
    - Removed schema `OfficeWorkLocation`.
    - Removed schema `RemoteWorkLocation`.
    - Removed schema `StateRegistrationRequired`.
    - Removed schema `WorkplaceNotFound`.
    - Removed schema `PayScheduleNotConfigured`.
    - Removed schema `SubscriptionLimitError`.
    - Removed schema `InvalidWorkerStatusError`.
    - Removed schema `WorkerAlreadyExistsError`.
    - Removed schema `CannotDeleteWorker`.
    - Removed schema `AddressInvalid`.
    - Removed schema `WorkplaceAlreadyExists`.
* **api:** Property `MissingRequiredCompanyPermissions.requiredPermissions` type changed from `Array<object>` to `Array<object>`.

### Features

* **api:** initial SDK generation ([8561f11](https://github.com/TeamWarp/warp-sdk-go/commit/8561f11f168401b3ced29c591fe7c59fb1bea64c))
* **api:** update contact email ([c51fcd0](https://github.com/TeamWarp/warp-sdk-go/commit/c51fcd07eeea2d307a884bdd9cf5c72e2b6cfbc6))
* **api:** update import names to warp ([87831bc](https://github.com/TeamWarp/warp-sdk-go/commit/87831bc83f412d1069ef9435c1ae7dd00823391f))
* **api:** update SDK name (+145 more changes) ([6ac93ee](https://github.com/TeamWarp/warp-sdk-go/commit/6ac93ee5b2ffb457a3c6a21bc75eaeeb2ca5f5fc))
* **api:** update SDK surface (2 changes) ([ea40d6f](https://github.com/TeamWarp/warp-sdk-go/commit/ea40d6faf17e53ddcdb7b83d1d305602b82897d0))


### Chores

* **api:** regenerate SDK ([ac04211](https://github.com/TeamWarp/warp-sdk-go/commit/ac042115bb63da525c083edb0ceb0b6fd6acc527))
* **api:** update generated SDK content ([e02c75c](https://github.com/TeamWarp/warp-sdk-go/commit/e02c75c417391e38cfbf130f1ee9ba2a6594d10a))
