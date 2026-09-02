# Changelog

## [0.5.0](https://github.com/TeamWarp/warp-sdk-go/compare/v0.5.1...v0.5.0) (2026-09-02)


### ⚠ BREAKING CHANGES

* **api:** 156 breaking changes to the SDK surface.
    - Response of `offers.void` changed from `objects_5` to `objects_6`.
    - Response of `offers.extendDeadline` changed from `objects_5` to `objects_6`.
    - Response of `offers.resend` changed from `objects_5` to `objects_6`.
    - Response of `timeOff.policies.get` changed from `objects_9` to `objects_10`.
    - Response of `workers.get` changed from `objects_10` to `objects_11`.
    - Property `public_text_worker_custom_field.display` type changed from `union_37` to `union_39`.
    - Property `public_number_worker_custom_field.display` type changed from `union_37` to `union_39`.
    - Property `public_date_worker_custom_field.display` type changed from `union_37` to `union_39`.
    - Property `public_boolean_worker_custom_field.display` type changed from `union_37` to `union_39`.
    - Property `public_currency_worker_custom_field.display` type changed from `union_37` to `union_39`.
    - Property `public_percentage_worker_custom_field.display` type changed from `union_37` to `union_39`.
    - Property `public_select_worker_custom_field.display` type changed from `union_37` to `union_39`.
    - Property `public_multi_select_worker_custom_field.display` type changed from `union_37` to `union_39`.
    - Property `objects_5.id` type changed from `string` to `string`.
    - Added required property `objects_5.code`.
    - Added required property `objects_5.name`.
    - Added required property `objects_5.track`.
    - Removed required property `objects_5.status`.
    - Removed required property `objects_5.workerType`.
    - Removed required property `objects_5.candidate`.
    - Removed required property `objects_5.position`.
    - Removed required property `objects_5.department`.
    - Removed required property `objects_5.workplace`.
    - Removed required property `objects_5.manager`.
    - Removed required property `objects_5.sentBy`.
    - Removed required property `objects_5.compensation`.
    - Removed required property `objects_5.offerUrl`.
    - Removed required property `objects_5.expirationTime`.
    - Removed required property `objects_5.lastViewedAt`.
    - Removed required property `objects_5.createdAt`.
    - Added required property `objects_7.firstName`.
    - Added required property `objects_7.lastName`.
    - Added required property `objects_7.email`.
    - Added required property `objects_7.contractorDetails`.
    - Removed required property `objects_7.title`.
    - Removed required property `objects_7.startDate`.
    - Removed required property `objects_7.country`.
    - Removed required property `objects_7.scopeOfWork`.
    - Added required property `objects_8.title`.
    - Added required property `objects_8.startDate`.
    - Added required property `objects_8.country`.
    - Added required property `objects_8.scopeOfWork`.
    - Removed required property `objects_8.basePay`.
    - Removed required property `objects_8.signOnBonus`.
    - Removed required property `objects_8.relocationBonus`.
    - Removed required property `objects_8.stock`.
    - Schema `union_19` shape changed.
    - Schema `union_20` shape changed.
    - Added required property `objects_9.basePay`.
    - Added required property `objects_9.signOnBonus`.
    - Added required property `objects_9.relocationBonus`.
    - Added required property `objects_9.stock`.
    - Removed required property `objects_9.id`.
    - Removed required property `objects_9.timeOffTypeId`.
    - Removed required property `objects_9.timeOffTypeName`.
    - Removed required property `objects_9.paid`.
    - Removed required property `objects_9.isUnlimited`.
    - Removed required property `objects_9.schedule`.
    - Removed required property `objects_9.unit`.
    - Removed required property `objects_9.name`.
    - Removed required property `objects_9.description`.
    - Removed required property `objects_9.hoursWorkedPerChunk`.
    - Removed required property `objects_9.minutesPerChunk`.
    - Removed required property `objects_9.minutesPerPeriod`.
    - Added required property `objects_6.id`.
    - Added required property `objects_6.status`.
    - Added required property `objects_6.workerType`.
    - Added required property `objects_6.candidate`.
    - Added required property `objects_6.position`.
    - Added required property `objects_6.department`.
    - Added required property `objects_6.workplace`.
    - Added required property `objects_6.manager`.
    - Added required property `objects_6.sentBy`.
    - Added required property `objects_6.compensation`.
    - Added required property `objects_6.offerUrl`.
    - Added required property `objects_6.expirationTime`.
    - Added required property `objects_6.lastViewedAt`.
    - Added required property `objects_6.createdAt`.
    - Removed required property `objects_6.firstName`.
    - Removed required property `objects_6.lastName`.
    - Removed required property `objects_6.email`.
    - Removed required property `objects_6.contractorDetails`.
    - Schema `union_21` changed from `enum(us_w2 | us_1099 | global_contractor)` to `string | null`.
    - Schema `union_22` changed from `string | number | union_2` to `enum(us_w2 | us_1099 | global_contractor)`.
    - Property `public_paycheck_worker.workerType` type changed from `union_21 & string` to `union_22 & string`.
    - Property `objects_10.id` type changed from `string` to `string`.
    - Added required property `objects_10.timeOffTypeId`.
    - Added required property `objects_10.timeOffTypeName`.
    - Added required property `objects_10.paid`.
    - Added required property `objects_10.isUnlimited`.
    - Added required property `objects_10.schedule`.
    - Added required property `objects_10.unit`.
    - Added required property `objects_10.name`.
    - Added required property `objects_10.description`.
    - Added required property `objects_10.hoursWorkedPerChunk`.
    - Added required property `objects_10.minutesPerChunk`.
    - Added required property `objects_10.minutesPerPeriod`.
    - Removed required property `objects_10.position`.
    - Removed required property `objects_10.type`.
    - Removed required property `objects_10.status`.
    - Removed required property `objects_10.startDate`.
    - Removed required property `objects_10.endDate`.
    - Removed required property `objects_10.isBusiness`.
    - Removed required property `objects_10.businessName`.
    - Removed required property `objects_10.firstName`.
    - Removed required property `objects_10.lastName`.
    - Removed required property `objects_10.email`.
    - Removed required property `objects_10.workEmail`.
    - Removed required property `objects_10.preferredName`.
    - Removed required property `objects_10.displayName`.
    - Removed required property `objects_10.timeZone`.
    - Removed required property `objects_10.department`.
    - Removed required property `objects_10.compensation`.
    - Removed optional property `objects_10.customFields`.
    - Schema `union_23` shape changed.
    - Property `time_off_policy_not_found_encoded.id` type changed from `union_22` to `union_23`.
    - Schema `union_24` shape changed.
    - Schema `union_25` changed from `enum(pending | approved | denied)` to `string | null`.
    - Schema `union_26` shape changed.
    - Schema `union_27` shape changed.
    - Schema `union_28` changed from `string | null` to `enum(employee | contractor)`.
    - Schema `union_29` shape changed.
    - Schema `union_30` shape changed.
    - Schema `union_31` shape changed.
    - Schema `union_32` shape changed.
    - Schema `union_34` shape changed.
    - Schema `union_35` shape changed.
    - Schema `union_36` shape changed.
    - Schema `union_37` shape changed.
    - Schema `union_39` changed from `enum(remote | office)` to `string | null`.
    - Schema `union_38` shape changed.
    - Added required property `objects_11.id`.
    - Added required property `objects_11.position`.
    - Added required property `objects_11.type`.
    - Added required property `objects_11.status`.
    - Added required property `objects_11.startDate`.
    - Added required property `objects_11.endDate`.
    - Added required property `objects_11.isBusiness`.
    - Added required property `objects_11.businessName`.
    - Added required property `objects_11.firstName`.
    - Added required property `objects_11.lastName`.
    - Added required property `objects_11.email`.
    - Added required property `objects_11.workEmail`.
    - Added required property `objects_11.preferredName`.
    - Added required property `objects_11.displayName`.
    - Added required property `objects_11.timeZone`.
    - Added required property `objects_11.department`.
    - Added required property `objects_11.compensation`.
    - Removed required property `objects_11.line1`.
    - Removed optional property `objects_11.line2`.
    - Removed required property `objects_11.city`.
    - Removed required property `objects_11.postalCode`.
    - Removed required property `objects_11.state`.
    - Removed required property `objects_11.country`.
    - Schema `union_40` changed from `enum(active | archived)` to `string | null`.
    - Property `workplace_already_exists_encoded.id` type changed from `union_38` to `union_40`.
* **api:** Renamed SDK from `WarpApi` to `Warp`.
* **api:** Property `missing_required_company_permissions_encoded.requiredPermissions` type changed from `Array<object>` to `Array<object>`.
* **api:** URL of environment `production` changed from `https://api.joinwarp.com` to `https://api.joinwarp.com/public`.
* **api:** 51 breaking changes to the SDK surface.
    - Property `objects_4.value` type changed from `union_12` to `public_custom_field_value_output`.
    - Schema `union_12` shape changed.
    - Property `department_already_exists_encoded.id` type changed from `union_13` to `union_12`.
    - Schema `union_13` changed from `string | null` to `enum(draft | sent | accepted | …)`.
    - Schema `union_14` shape changed.
    - Schema `union_15` shape changed.
    - Property `objects_7.country` type changed from `union_16` to `union_15`.
    - Schema `union_16` changed from `enum(AD | AE | AF | …)` to `object | null`.
    - Schema `union_17` shape changed.
    - Schema `union_18` shape changed.
    - Schema `union_19` shape changed.
    - Property `objects_8.basePay` type changed from `object` to `object`.
    - Property `objects_8.signOnBonus` type changed from `union_20` to `union_19`.
    - Property `objects_8.relocationBonus` type changed from `union_20` to `union_19`.
    - Property `objects_5.status` type changed from `union_14` to `union_13`.
    - Property `objects_5.workerType` type changed from `union_15` to `union_14`.
    - Property `objects_5.department` type changed from `union_17` to `union_16`.
    - Property `objects_5.workplace` type changed from `union_18` to `union_17`.
    - Property `objects_5.manager` type changed from `union_19` to `union_18`.
    - Property `invalid_offer_status_error_encoded.status` type changed from `union_14` to `union_13`.
    - Schema `union_20` shape changed.
    - Schema `union_21` changed from `string | null` to `enum(us_w2 | us_1099 | global_contractor)`.
    - Property `public_paycheck_worker.workerType` type changed from `union_22 & string` to `union_21 & string`.
    - Schema `union_22` changed from `enum(us_w2 | us_1099 | global_contractor)` to `string | number | union_2`.
    - Property `time_off_policy_not_found_encoded.id` type changed from `union_23` to `union_22`.
    - Schema `union_23` shape changed.
    - Schema `union_24` shape changed.
    - Schema `union_25` changed from `string | null` to `enum(pending | approved | denied)`.
    - Schema `union_26` shape changed.
    - Schema `union_27` shape changed.
    - Schema `union_28` changed from `enum(employee | contractor)` to `string | null`.
    - Schema `union_29` shape changed.
    - Schema `union_30` shape changed.
    - Schema `union_31` shape changed.
    - Schema `union_32` shape changed.
    - Schema `union_34` shape changed.
    - Schema `union_35` shape changed.
    - Schema `union_37` shape changed.
    - Schema `union_36` shape changed.
    - Property `objects_10.type` type changed from `union_28` to `union_27`.
    - Property `objects_10.status` type changed from `union_27` to `union_26`.
    - Property `objects_10.endDate` type changed from `union_29` to `union_28`.
    - Property `objects_10.isBusiness` type changed from `union_30` to `union_29`.
    - Property `objects_10.businessName` type changed from `union_31` to `union_30`.
    - Property `objects_10.workEmail` type changed from `union_32` to `union_31`.
    - Property `objects_10.preferredName` type changed from `union_33` to `union_32`.
    - Property `objects_10.timeZone` type changed from `union_34` to `union_33`.
    - Property `objects_10.department` type changed from `union_35` to `union_34`.
    - Property `objects_10.compensation` type changed from `union_36` to `union_35`.
    - Property `objects_10.customFields` type changed from `union_37` to `union_36`.
    - Property `address_invalid_encoded.suggestedAlternative` type changed from `object | null` to `object | null`.
* **api:** 52 breaking changes to the SDK surface.
    - Schema `union_12` shape changed.
    - Property `objects_4.value` type changed from `object | object | object | object | object | object | object | object` to `union_12`.
    - Schema `union_13` changed from `enum(draft | sent | accepted | …)` to `string | null`.
    - Property `department_already_exists_encoded.id` type changed from `union_12` to `union_13`.
    - Schema `union_14` shape changed.
    - Schema `union_15` shape changed.
    - Schema `union_16` changed from `object | null` to `enum(AD | AE | AF | …)`.
    - Property `objects_7.country` type changed from `union_15` to `union_16`.
    - Schema `union_17` shape changed.
    - Schema `union_18` shape changed.
    - Schema `union_19` shape changed.
    - Schema `union_20` shape changed.
    - Property `objects_8.basePay` type changed from `object` to `object`.
    - Property `objects_8.signOnBonus` type changed from `union_19` to `union_20`.
    - Property `objects_8.relocationBonus` type changed from `union_19` to `union_20`.
    - Property `objects_5.status` type changed from `union_13` to `union_14`.
    - Property `objects_5.workerType` type changed from `union_14` to `union_15`.
    - Property `objects_5.department` type changed from `union_16` to `union_17`.
    - Property `objects_5.workplace` type changed from `union_17` to `union_18`.
    - Property `objects_5.manager` type changed from `union_18` to `union_19`.
    - Property `invalid_offer_status_error_encoded.status` type changed from `union_13` to `union_14`.
    - Schema `union_21` changed from `enum(us_w2 | us_1099 | global_contractor)` to `string | null`.
    - Schema `union_22` changed from `string | number | union_2` to `enum(us_w2 | us_1099 | global_contractor)`.
    - Property `public_paycheck_worker.workerType` type changed from `union_21 & string` to `union_22 & string`.
    - Schema `union_23` shape changed.
    - Property `time_off_policy_not_found_encoded.id` type changed from `union_22` to `union_23`.
    - Schema `union_24` shape changed.
    - Schema `union_25` changed from `enum(pending | approved | denied)` to `string | null`.
    - Schema `union_26` shape changed.
    - Schema `union_27` shape changed.
    - Schema `union_28` changed from `string | null` to `enum(employee | contractor)`.
    - Schema `union_29` shape changed.
    - Schema `union_30` shape changed.
    - Schema `union_31` shape changed.
    - Schema `union_32` shape changed.
    - Schema `union_34` shape changed.
    - Schema `union_35` shape changed.
    - Schema `union_36` shape changed.
    - Schema `union_37` changed from `enum(remote | office)` to `Array<object> | null`.
    - Property `objects_10.type` type changed from `union_27` to `union_28`.
    - Property `objects_10.status` type changed from `union_26` to `union_27`.
    - Property `objects_10.endDate` type changed from `union_28` to `union_29`.
    - Property `objects_10.isBusiness` type changed from `union_29` to `union_30`.
    - Property `objects_10.businessName` type changed from `union_30` to `union_31`.
    - Property `objects_10.workEmail` type changed from `union_31` to `union_32`.
    - Property `objects_10.preferredName` type changed from `union_32` to `union_33`.
    - Property `objects_10.timeZone` type changed from `union_33` to `union_34`.
    - Property `objects_10.department` type changed from `union_34` to `union_35`.
    - Property `objects_10.compensation` type changed from `union_35` to `union_36`.
    - Schema `union_38` changed from `enum(active | archived)` to `string | null`.
    - Property `workplace_already_exists_encoded.id` type changed from `union_36` to `union_38`.
    - Property `address_invalid_encoded.suggestedAlternative` type changed from `object | null` to `object | null`.
* **api:** 2 breaking changes to the SDK surface.
    - Removed operation `payroll.retrievePaycheck` (`GET /v1/paychecks/{id}`).
    - Removed operation `payroll.retrieve` (`GET /v1/payrolls/{id}`).
* **api:** 131 breaking changes to the SDK surface.
    - Schema `union` shape changed.
    - Schema `union_1` shape changed.
    - Property `public_worker_compensation.payRateId` type changed from `string` to `string`.
    - Property `public_worker_compensation.amount` type changed from `string` to `integer`.
    - Property `public_worker_compensation.currency` type changed from `union` to `union_1`.
    - Property `public_money_amount.amount` type changed from `string` to `integer`.
    - Property `public_money_amount.currency` type changed from `union` to `union_1`.
    - Property `public_health_plan_carrier.id` type changed from `string` to `string`.
    - Property `public_health_plan.id` type changed from `string` to `string`.
    - Property `public_health_plan.effectiveStartDate` type changed from `string` to `string`.
    - Property `public_health_plan.effectiveEndDate` type changed from `string | null` to `string | null`.
    - Property `health_plan_not_found_error_encoded.id` type changed from `string` to `string`.
    - Property `public_retirement_plan.id` type changed from `string` to `string`.
    - Property `public_retirement_plan.effectiveStartDate` type changed from `string` to `string`.
    - Property `public_retirement_plan.effectiveEndDate` type changed from `string | null` to `string | null`.
    - Property `retirement_plan_not_found_error_encoded.id` type changed from `string` to `string`.
    - Property `public_worker_reference.id` type changed from `string` to `string`.
    - Property `health_plan_reference.id` type changed from `string` to `string`.
    - Property `retirement_plan_reference.id` type changed from `string` to `string`.
    - Schema `union_2` changed from `string | null` to `enum(Infinity | -Infinity | NaN)`.
    - Property `percentage_contribution.percentage` type changed from `string | union_1` to `number | union_2`.
    - Property `public_benefit_deduction.id` type changed from `string` to `string`.
    - Property `public_benefit_deduction.effectiveStartDate` type changed from `string` to `string`.
    - Property `public_benefit_deduction.effectiveEndDate` type changed from `string | null` to `string | null`.
    - Property `benefit_deduction_not_found_error_encoded.id` type changed from `string` to `string`.
    - Schema `union_3` changed from `enum(text | number | date | …)` to `string | null`.
    - Schema `union_4` shape changed.
    - Schema `union_5` shape changed.
    - Schema `union_6` shape changed.
    - Schema `union_7` shape changed.
    - Schema `union_8` changed from `boolean | null` to `enum(admin | worker)`.
    - Schema `union_9` shape changed.
    - Property `objects.id` type changed from `string` to `string`.
    - Property `objects.description` type changed from `union_2` to `union_3`.
    - Property `objects.type` type changed from `union_3` to `union_4`.
    - Property `objects.status` type changed from `union_4` to `union_5`.
    - Property `objects.category` type changed from `union_5` to `union_6`.
    - Property `objects.accessLevel` type changed from `union_6` to `union_7`.
    - Property `objects.inputBy` type changed from `union_7` to `union_8`.
    - Property `objects.required` type changed from `union_8` to `union_9`.
    - Schema `union_10` shape changed.
    - Property `objects_2.label` type changed from `string & string` to `string`.
    - Property `objects_2.value` type changed from `string & string` to `string`.
    - Property `objects_2.sortOrder` type changed from `number | union_1 | null` to `number | union_2 | null`.
    - Schema `union_11` shape changed.
    - Property `objects_3.id` type changed from `string` to `string`.
    - Property `objects_3.sortOrder` type changed from `union_10` to `union_11`.
    - Property `custom_field_not_found_error_encoded.id` type changed from `string` to `string`.
    - Property `custom_field_option_not_found_error_encoded.id` type changed from `string` to `string`.
    - Property `custom_field_option_in_use_error_encoded.id` type changed from `string` to `string`.
    - Property `objects_4.id` type changed from `string` to `string`.
    - Property `objects_4.workerId` type changed from `string` to `string`.
    - Property `objects_4.fieldId` type changed from `string` to `string`.
    - Property `objects_4.value` type changed from `object | object | object | object | object | object | object | object` to `object | object | object | object | object | object | object | object`.
    - Property `custom_field_worker_not_found_error_encoded.id` type changed from `string` to `string`.
    - Schema `union_12` changed from `enum(draft | sent | accepted | …)` to `string | null`.
    - Property `department_already_exists_encoded.id` type changed from `union_11` to `union_12`.
    - Property `department_not_found_encoded.id` type changed from `string` to `string`.
    - Schema `union_13` shape changed.
    - Schema `union_14` shape changed.
    - Property `objects_6.email` type changed from `string<email>` to `string<email>`.
    - Schema `union_15` changed from `object | null` to `enum(AD | AE | AF | …)`.
    - Property `objects_7.startDate` type changed from `string` to `string`.
    - Property `objects_7.country` type changed from `union_14` to `union_15`.
    - Schema `union_16` shape changed.
    - Schema `union_17` shape changed.
    - Schema `union_18` shape changed.
    - Schema `union_19` shape changed.
    - Property `objects_8.basePay` type changed from `object` to `object`.
    - Property `objects_8.signOnBonus` type changed from `union_18` to `union_19`.
    - Property `objects_8.relocationBonus` type changed from `union_18` to `union_19`.
    - Property `objects_8.stock` type changed from `object | null` to `object | null`.
    - Property `objects_5.id` type changed from `string` to `string`.
    - Property `objects_5.status` type changed from `union_12` to `union_13`.
    - Property `objects_5.workerType` type changed from `union_13` to `union_14`.
    - Property `objects_5.department` type changed from `union_15` to `union_16`.
    - Property `objects_5.workplace` type changed from `union_16` to `union_17`.
    - Property `objects_5.manager` type changed from `union_17` to `union_18`.
    - Property `workplace_not_found_encoded.id` type changed from `string` to `string`.
    - Property `manager_not_found_error_encoded.id` type changed from `string` to `string`.
    - Property `offer_not_found_error_encoded.id` type changed from `string` to `string`.
    - Property `invalid_offer_status_error_encoded.id` type changed from `string` to `string`.
    - Property `invalid_offer_status_error_encoded.status` type changed from `union_12` to `union_13`.
    - Property `public_pay_rate.id` type changed from `string` to `string`.
    - Property `public_pay_rate.amount` type changed from `string` to `integer`.
    - Property `public_pay_rate.currency` type changed from `union` to `union_1`.
    - Property `public_pay_rate.effectiveStartDate` type changed from `string | null` to `string | null`.
    - Property `public_pay_rate.effectiveEndDate` type changed from `string | null` to `string | null`.
    - Property `pay_rate_not_found_error_encoded.id` type changed from `string` to `string`.
    - Schema `union_20` shape changed.
    - Schema `union_21` changed from `string | null` to `enum(us_w2 | us_1099 | global_contractor)`.
    - Property `objects_9.id` type changed from `string` to `string`.
    - Property `objects_9.timeOffTypeId` type changed from `string` to `string`.
    - Property `objects_9.hoursWorkedPerChunk` type changed from `union_10 | null` to `union_11 | null`.
    - Property `objects_9.minutesPerChunk` type changed from `union_10 | null` to `union_11 | null`.
    - Property `objects_9.minutesPerPeriod` type changed from `union_10 | null` to `union_11 | null`.
    - Schema `union_22` changed from `enum(pending | approved | denied)` to `string | number | union_2`.
    - Property `time_off_policy_not_found_encoded.id` type changed from `union_19` to `union_22`.
    - Schema `union_23` changed from `enum(draft | invited | onboarding | …)` to `string | null`.
    - Schema `union_24` changed from `enum(employee | contractor)` to `string | null`.
    - Schema `union_25` changed from `string | null` to `enum(pending | approved | denied)`.
    - Property `time_off_request_not_found_error_encoded.id` type changed from `string | number | union_1` to `string | number | union_2`.
    - Schema `union_26` changed from `boolean | null` to `enum(draft | invited | onboarding | …)`.
    - Schema `union_27` changed from `string | null` to `enum(employee | contractor)`.
    - Schema `union_28` shape changed.
    - Schema `union_29` shape changed.
    - Schema `union_31` shape changed.
    - Schema `union_32` shape changed.
    - Schema `union_33` shape changed.
    - Schema `union_34` changed from `enum(remote | office)` to `object | null`.
    - Schema `union_35` changed from `enum(active | archived)` to `public_worker_compensation | null`.
    - Property `objects_10.id` type changed from `string` to `string`.
    - Property `objects_10.type` type changed from `union_24` to `union_27`.
    - Property `objects_10.status` type changed from `union_23` to `union_26`.
    - Property `objects_10.startDate` type changed from `string` to `string`.
    - Property `objects_10.endDate` type changed from `union_25` to `union_28`.
    - Property `objects_10.isBusiness` type changed from `union_26` to `union_29`.
    - Property `objects_10.businessName` type changed from `union_27` to `union_30`.
    - Property `objects_10.email` type changed from `string<email>` to `string<email>`.
    - Property `objects_10.workEmail` type changed from `union_28` to `union_31`.
    - Property `objects_10.preferredName` type changed from `union_29` to `union_32`.
    - Property `objects_10.timeZone` type changed from `union_30` to `union_33`.
    - Property `objects_10.department` type changed from `union_31` to `union_34`.
    - Property `objects_10.compensation` type changed from `union_32` to `union_35`.
    - Property `worker_not_found_error_encoded.id` type changed from `string` to `string`.
    - Property `office_work_location.workplaceId` type changed from `string` to `string`.
    - Property `invalid_worker_status_error_encoded.id` type changed from `string` to `string`.
    - Property `cannot_delete_worker_encoded.id` type changed from `string` to `string`.
    - Property `objects_11.line1` type changed from `string` to `string`.
    - Property `workplace_already_exists_encoded.id` type changed from `union_33` to `union_36`.
    - Property `address_invalid_encoded.suggestedAlternative` type changed from `object | null` to `object | null`.
* **api:** 8 breaking changes to the SDK surface.
    - Property `missing_required_company_permissions_encoded.requiredPermissions` type changed from `Array<object>` to `Array<object>`.
    - Property `objects_4.value` type changed from `object | object | object | object | object | object | object | object` to `object | object | object | object | object | object | object | object`.
    - Property `objects_6.contractorDetails` type changed from `object | null` to `object | null`.
    - Schema `union_15` shape changed.
    - Schema `union_16` shape changed.
    - Schema `union_17` shape changed.
    - Property `objects_8.basePay` type changed from `object` to `object`.
    - Property `address_invalid_encoded.suggestedAlternative` type changed from `object | null` to `object | null`.
* **api:** `422` error response of `offers.create` changed from `invalid_expiration_time_error_encoded` to `application/json`.
* **api:** 2 breaking changes to the SDK surface.
    - Added required body field `voidReason` to `offers.void`.
    - Added required request body to `offers.void`.
* **api:** 8 breaking changes to the SDK surface.
    - Property `missing_required_company_permissions_encoded.requiredPermissions` type changed from `Array<object>` to `Array<object>`.
    - Property `objects_4.value` type changed from `object | object | object | object | object | object | object | object` to `object | object | object | object | object | object | object | object`.
    - Property `objects_6.contractorDetails` type changed from `object | null` to `object | null`.
    - Schema `union_15` shape changed.
    - Schema `union_16` shape changed.
    - Schema `union_17` shape changed.
    - Property `objects_8.basePay` type changed from `object` to `object`.
    - Property `address_invalid_encoded.suggestedAlternative` type changed from `object | null` to `object | null`.
* **api:** Schema `union_20` shape changed.
* **api:** 5 breaking changes to the SDK surface.
    - Added required property `public_worker_compensation.per`.
    - Removed required property `public_worker_compensation.basis`.
    - Added required property `public_pay_rate.per`.
    - Removed required property `public_pay_rate.basis`.
    - Removed schema `public_pay_rate_basis`.
* **api:** 2 breaking changes to the SDK surface.
    - Renamed SDK from `Warp` to `WarpApi`.
    - Removed operation `benefits.deductions.retrieve` (`GET /v1/benefits/deductions/{id}`).
* **api:** Removed operation `benefits.deductions.get` (`GET /v1/benefits/deductions/{id}`).
* **api:** 2 breaking changes to the SDK surface.
    - Schema `public_pay_rate_basis` shape changed.
    - Added required property `public_worker_compensation.basis`.
* **api:** 4 breaking changes to the SDK surface.
    - Removed operation `customFields.retrieve` (`GET /v1/custom_fields/{id}`).
    - Removed operation `timeOff.policies.timeOffGet` (`GET /v1/time_off/policies`).
    - Removed operation `timeOff.policies.timeOffGet2` (`GET /v1/time_off/policies/{id}`).
    - Removed operation `workers.retrieve` (`GET /v1/workers/{id}`).
* **api:** 2 breaking changes to the SDK surface.
    - Added required property `public_pay_rate.worker`.
    - Removed required property `public_pay_rate.workerId`.
* **api:** 8 breaking changes to the SDK surface.
    - Removed operation `payRates.list` (`GET /v1/pay_rates`).
    - Removed operation `payRates.retrieve` (`GET /v1/pay_rates/{id}`).
    - Removed operation `timeOff.policies.list` (`GET /v1/time_off/policies`).
    - Removed operation `timeOff.policies.get` (`GET /v1/time_off/policies/{id}`).
    - Removed schema `public_pay_rate_type`.
    - Removed schema `public_pay_rate_basis`.
    - Removed schema `public_pay_rate`.
    - Removed schema `pay_rate_not_found_error_encoded`.
* **api:** 7 breaking changes to the SDK surface.
    - Removed operation `timeOff.policies.timeOffGet` (`GET /v1/time_off/policies`).
    - Removed operation `timeOff.policies.timeOffGet2` (`GET /v1/time_off/policies/{id}`).
    - Schema `union_32` shape changed.
    - Added required property `objects_10.compensation`.
    - Schema `union_33` changed from `enum(remote | office)` to `string | null`.
    - Schema `union_34` shape changed.
    - Property `workplace_already_exists_encoded.id` type changed from `union_32` to `union_33`.
* **api:** 3 breaking changes to the SDK surface.
    - Property `objects_6.email` type changed from `string` to `string<email>`.
    - Schema `union_28` shape changed.
    - Property `objects_10.email` type changed from `string` to `string<email>`.
* **api:** 214 breaking changes to the SDK surface.
    - query param `limit` on `benefits.healthPlans.list` is now required.
    - query param `statuses` on `benefits.healthPlans.list` is now required.
    - Serialization or defaults of query param `statuses` on `benefits.healthPlans.list` changed.
    - `401` error response of `benefits.healthPlans.list` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `benefits.healthPlans.list` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `benefits.healthPlans.get` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `benefits.healthPlans.get` changed from `health_plan_not_found_error` to `health_plan_not_found_error_encoded`.
    - `429` error response of `benefits.healthPlans.get` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - query param `limit` on `benefits.retirementPlans.list` is now required.
    - query param `statuses` on `benefits.retirementPlans.list` is now required.
    - Serialization or defaults of query param `statuses` on `benefits.retirementPlans.list` changed.
    - `401` error response of `benefits.retirementPlans.list` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `benefits.retirementPlans.list` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `benefits.retirementPlans.get` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `benefits.retirementPlans.get` changed from `retirement_plan_not_found_error` to `retirement_plan_not_found_error_encoded`.
    - `429` error response of `benefits.retirementPlans.get` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - query param `limit` on `benefits.deductions.list` is now required.
    - query param `statuses` on `benefits.deductions.list` is now required.
    - Serialization or defaults of query param `statuses` on `benefits.deductions.list` changed.
    - `401` error response of `benefits.deductions.list` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `benefits.deductions.list` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `benefits.deductions.get` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `benefits.deductions.get` changed from `benefit_deduction_not_found_error` to `benefit_deduction_not_found_error_encoded`.
    - `429` error response of `benefits.deductions.get` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `customFields.list` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `customFields.list` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `400` error response of `customFields.create` changed from `application/json` to `invalid_custom_field_operation_error_encoded`.
    - `401` error response of `customFields.create` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `customFields.create` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `customFields.retrieve` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `customFields.retrieve` changed from `custom_field_not_found_error` to `custom_field_not_found_error_encoded`.
    - `429` error response of `customFields.retrieve` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - Response of `customFields.update` changed from `none` to `objects`.
    - `400` error response of `customFields.update` changed from `application/json` to `invalid_custom_field_operation_error_encoded`.
    - `401` error response of `customFields.update` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `customFields.update` changed from `custom_field_not_found_error` to `custom_field_not_found_error_encoded`.
    - `409` error response of `customFields.update` changed from `custom_field_already_exists_error` to `custom_field_already_exists_error_encoded`.
    - `429` error response of `customFields.update` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - Response of `customFields.archive` changed from `none` to `objects`.
    - `401` error response of `customFields.archive` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `customFields.archive` changed from `custom_field_not_found_error` to `custom_field_not_found_error_encoded`.
    - `429` error response of `customFields.archive` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `400` error response of `customFields.createOption` changed from `application/json` to `invalid_custom_field_operation_error_encoded`.
    - `401` error response of `customFields.createOption` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `customFields.createOption` changed from `custom_field_not_found_error` to `custom_field_not_found_error_encoded`.
    - `409` error response of `customFields.createOption` changed from `custom_field_option_already_exists_error` to `custom_field_option_already_exists_error_encoded`.
    - `429` error response of `customFields.createOption` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - Response of `customFields.updateOption` changed from `none` to `objects_3`.
    - `400` error response of `customFields.updateOption` changed from `application/json` to `invalid_custom_field_operation_error_encoded`.
    - `401` error response of `customFields.updateOption` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `customFields.updateOption` changed from `custom_field_option_not_found_error` to `custom_field_option_not_found_error_encoded`.
    - `409` error response of `customFields.updateOption` changed from `custom_field_option_already_exists_error` to `custom_field_option_already_exists_error_encoded`.
    - `429` error response of `customFields.updateOption` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `customFields.deleteOption` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `customFields.deleteOption` changed from `custom_field_option_not_found_error` to `custom_field_option_not_found_error_encoded`.
    - `409` error response of `customFields.deleteOption` changed from `custom_field_option_in_use_error` to `custom_field_option_in_use_error_encoded`.
    - `429` error response of `customFields.deleteOption` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - Response of `customFields.archiveOption` changed from `none` to `objects_3`.
    - `401` error response of `customFields.archiveOption` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `customFields.archiveOption` changed from `custom_field_option_not_found_error` to `custom_field_option_not_found_error_encoded`.
    - `429` error response of `customFields.archiveOption` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `customFields.listValues` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `customFields.listValues` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - Response of `customFields.upsertValue` changed from `none` to `objects_4`.
    - `401` error response of `customFields.upsertValue` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `customFields.upsertValue` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `customFields.clearValue` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `customFields.clearValue` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - query param `limit` on `departments.list` is now required.
    - `401` error response of `departments.list` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `departments.list` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `departments.create` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `409` error response of `departments.create` changed from `department_already_exists` to `department_already_exists_encoded`.
    - `429` error response of `departments.create` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `departments.update` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `departments.update` changed from `department_not_found` to `department_not_found_encoded`.
    - `409` error response of `departments.update` changed from `department_already_exists` to `department_already_exists_encoded`.
    - `429` error response of `departments.update` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - query param `limit` on `offers.list` is now required.
    - `401` error response of `offers.list` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `offers.list` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `offers.create` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `422` error response of `offers.create` changed from `invalid_expiration_time_error` to `invalid_expiration_time_error_encoded`.
    - `429` error response of `offers.create` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - Response of `offers.void` changed from `none` to `objects_5`.
    - `401` error response of `offers.void` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `offers.void` changed from `offer_not_found_error` to `offer_not_found_error_encoded`.
    - `409` error response of `offers.void` changed from `invalid_offer_status_error` to `invalid_offer_status_error_encoded`.
    - `429` error response of `offers.void` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - Response of `offers.extendDeadline` changed from `none` to `objects_5`.
    - `401` error response of `offers.extendDeadline` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `offers.extendDeadline` changed from `offer_not_found_error` to `offer_not_found_error_encoded`.
    - `409` error response of `offers.extendDeadline` changed from `invalid_offer_status_error` to `invalid_offer_status_error_encoded`.
    - `422` error response of `offers.extendDeadline` changed from `invalid_expiration_time_error` to `invalid_expiration_time_error_encoded`.
    - `429` error response of `offers.extendDeadline` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - Response of `offers.resend` changed from `none` to `objects_5`.
    - `401` error response of `offers.resend` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `offers.resend` changed from `offer_not_found_error` to `offer_not_found_error_encoded`.
    - `409` error response of `offers.resend` changed from `invalid_offer_status_error` to `invalid_offer_status_error_encoded`.
    - `429` error response of `offers.resend` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - query param `limit` on `timeOff.listAssignments` is now required.
    - `401` error response of `timeOff.listAssignments` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `timeOff.listAssignments` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - query param `limit` on `timeOff.listBalances` is now required.
    - `401` error response of `timeOff.listBalances` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `timeOff.listBalances` changed from `time_off_policy_not_found` to `time_off_policy_not_found_encoded`.
    - `429` error response of `timeOff.listBalances` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - query param `limit` on `timeOff.listRequests` is now required.
    - `401` error response of `timeOff.listRequests` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `timeOff.listRequests` changed from `time_off_request_not_found_error` to `time_off_request_not_found_error_encoded`.
    - `429` error response of `timeOff.listRequests` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - query param `limit` on `timeOff.policies.timeOffGet` is now required.
    - `401` error response of `timeOff.policies.timeOffGet` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `timeOff.policies.timeOffGet` changed from `time_off_policy_not_found` to `time_off_policy_not_found_encoded`.
    - `429` error response of `timeOff.policies.timeOffGet` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - Response of `timeOff.policies.timeOffGet2` changed from `none` to `objects_9`.
    - `401` error response of `timeOff.policies.timeOffGet2` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `timeOff.policies.timeOffGet2` changed from `time_off_policy_not_found` to `time_off_policy_not_found_encoded`.
    - `429` error response of `timeOff.policies.timeOffGet2` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - query param `limit` on `workers.list` is now required.
    - `401` error response of `workers.list` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `workers.list` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - Response of `workers.retrieve` changed from `none` to `objects_10`.
    - `401` error response of `workers.retrieve` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `workers.retrieve` changed from `worker_not_found_error` to `worker_not_found_error_encoded`.
    - `429` error response of `workers.retrieve` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `workers.delete` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `workers.delete` changed from `worker_not_found_error` to `worker_not_found_error_encoded`.
    - `409` error response of `workers.delete` changed from `cannot_delete_worker` to `cannot_delete_worker_encoded`.
    - `429` error response of `workers.delete` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `400` error response of `workers.createEmployee` changed from `application/json` to `state_registration_required_encoded`.
    - `401` error response of `workers.createEmployee` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `workers.createEmployee` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `workers.createContractor` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `workers.createContractor` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `workers.invite` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `workers.invite` changed from `worker_not_found_error` to `worker_not_found_error_encoded`.
    - `429` error response of `workers.invite` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - query param `limit` on `workplaces.list` is now required.
    - `401` error response of `workplaces.list` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `429` error response of `workplaces.list` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `400` error response of `workplaces.create` changed from `application/json` to `address_invalid_encoded`.
    - `401` error response of `workplaces.create` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `409` error response of `workplaces.create` changed from `workplace_already_exists` to `workplace_already_exists_encoded`.
    - `429` error response of `workplaces.create` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - `401` error response of `workplaces.update` changed from `api_key_unauthorized` to `api_key_unauthorized_encoded`.
    - `404` error response of `workplaces.update` changed from `workplace_not_found` to `workplace_not_found_encoded`.
    - `409` error response of `workplaces.update` changed from `workplace_already_exists` to `workplace_already_exists_encoded`.
    - `429` error response of `workplaces.update` changed from `rate_limit_exceeded` to `rate_limit_exceeded_encoded`.
    - Property `public_money_amount.amount` type changed from `integer` to `string`.
    - Property `public_money_amount.currency` type changed from `enum(USD | AUD | BGN | …)` to `union`.
    - Property `public_health_plan_carrier.id` type changed from `string` to `string`.
    - Property `public_health_plan.id` type changed from `string` to `string`.
    - Property `public_health_plan.effectiveStartDate` type changed from `string` to `string`.
    - Property `public_health_plan.effectiveEndDate` type changed from `string | null` to `string | null`.
    - Property `public_health_plan.createdAt` type changed from `date` to `string`.
    - Property `public_health_plan.updatedAt` type changed from `date` to `string`.
    - Property `public_retirement_plan.id` type changed from `string` to `string`.
    - Property `public_retirement_plan.effectiveStartDate` type changed from `string` to `string`.
    - Property `public_retirement_plan.effectiveEndDate` type changed from `string | null` to `string | null`.
    - Property `public_retirement_plan.createdAt` type changed from `date` to `string`.
    - Property `public_retirement_plan.updatedAt` type changed from `date` to `string`.
    - Property `public_worker_reference.id` type changed from `string` to `string`.
    - Property `health_plan_reference.id` type changed from `string` to `string`.
    - Property `retirement_plan_reference.id` type changed from `string` to `string`.
    - Property `percentage_contribution.percentage` type changed from `number` to `string | union_1`.
    - Property `public_benefit_deduction.id` type changed from `string` to `string`.
    - Property `public_benefit_deduction.effectiveStartDate` type changed from `string` to `string`.
    - Property `public_benefit_deduction.effectiveEndDate` type changed from `string | null` to `string | null`.
    - Property `public_benefit_deduction.createdAt` type changed from `date_from_string` to `string`.
    - Property `public_benefit_deduction.updatedAt` type changed from `date_from_string` to `string`.
    - Property `office_work_location.workplaceId` type changed from `string` to `string`.
    - Removed schema `date`.
    - Removed schema `http_api_decode_error`.
    - Removed schema `issue`.
    - Removed schema `property_key`.
    - Removed schema `internal_server_error`.
    - Removed schema `api_key_unauthorized`.
    - Removed schema `rate_limit_exceeded`.
    - Removed schema `date_time_utc`.
    - Removed schema `missing_required_company_permissions`.
    - Removed schema `api_not_enabled`.
    - Removed schema `health_plan_not_found_error`.
    - Removed schema `retirement_plan_not_found_error`.
    - Removed schema `date_from_string`.
    - Removed schema `benefit_deduction_not_found_error`.
    - Removed schema `trimmed`.
    - Removed schema `non_empty_trimmed_string`.
    - Removed schema `invalid_custom_field_operation_error`.
    - Removed schema `custom_field_already_exists_error`.
    - Removed schema `custom_field_option_already_exists_error`.
    - Removed schema `custom_field_not_found_error`.
    - Removed schema `custom_field_option_not_found_error`.
    - Removed schema `custom_field_option_in_use_error`.
    - Removed schema `invalid_custom_field_value_error`.
    - Removed schema `custom_field_worker_not_found_error`.
    - Removed schema `department_already_exists`.
    - Removed schema `department_not_found`.
    - Removed schema `invalid_expiration_time_error`.
    - Removed schema `workplace_not_found`.
    - Removed schema `manager_not_found_error`.
    - Removed schema `offer_not_found_error`.
    - Removed schema `invalid_offer_status_error`.
    - Removed schema `time_off_policy_not_found`.
    - Removed schema `time_off_request_not_found_error`.
    - Removed schema `worker_not_found_error`.
    - Removed schema `state_registration_required`.
    - Removed schema `pay_schedule_not_configured`.
    - Removed schema `subscription_limit_error`.
    - Removed schema `invalid_worker_status_error`.
    - Removed schema `worker_already_exists_error`.
    - Removed schema `cannot_delete_worker`.
    - Removed schema `address_invalid`.
    - Removed schema `workplace_already_exists`.
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

* **api:** add body field voidReason on offers.void (+2 more changes) ([490263c](https://github.com/TeamWarp/warp-sdk-go/commit/490263cad6475b1333bce176d02396c9d76bdc45))
* **api:** add operation customFields.list (+20 more changes) ([9517c37](https://github.com/TeamWarp/warp-sdk-go/commit/9517c37fc8673889431c6688bfa34202ca1896e1))
* **api:** add operation levels.list ([821a4f1](https://github.com/TeamWarp/warp-sdk-go/commit/821a4f17e1972be68edd1d7329cb655c2650fece))
* **api:** add operation payRates.list (+5 more changes) ([5371520](https://github.com/TeamWarp/warp-sdk-go/commit/537152043f2b64853a67ffbd4593f2a34633ee90))
* **api:** add operation payroll.listPaychecks (+46 more changes) ([cb2cbc1](https://github.com/TeamWarp/warp-sdk-go/commit/cb2cbc1bce4560bd8d687c615e96f6e0f7c6a120))
* **api:** add property public_pay_rate.worker (+1 more change) ([f761bf2](https://github.com/TeamWarp/warp-sdk-go/commit/f761bf2d77535b1f5c10cc7a761065b9014ed2d6))
* **api:** add schema public_money_amount (+1 more change) ([697c33d](https://github.com/TeamWarp/warp-sdk-go/commit/697c33dd6fc59aa298eeb249edca58aec2afae38))
* **api:** api update ([eec6538](https://github.com/TeamWarp/warp-sdk-go/commit/eec65386699c3916c2b08b016459b56ba9d85971))
* **api:** initial SDK generation ([8561f11](https://github.com/TeamWarp/warp-sdk-go/commit/8561f11f168401b3ced29c591fe7c59fb1bea64c))
* **api:** remove operation benefits.deductions.get (+1 more change) ([6f44e30](https://github.com/TeamWarp/warp-sdk-go/commit/6f44e307725d1ed3fd34e975aa2f1256d0fd0f57))
* **api:** remove operation customFields.retrieve (+7 more changes) ([18a4e17](https://github.com/TeamWarp/warp-sdk-go/commit/18a4e17ba9f9b9f86a564a14ff85760c9b885143))
* **api:** remove operation customWorkerFields.list (+19 more changes) ([46a5ca0](https://github.com/TeamWarp/warp-sdk-go/commit/46a5ca0b4e6a63cabed0d0f6806e0f58b28d6e55))
* **api:** remove operation payRates.list (+9 more changes) ([eb0f673](https://github.com/TeamWarp/warp-sdk-go/commit/eb0f67357ec7fd5ebde13d61bc7a0bf9a1982a18))
* **api:** remove operation payroll.retrievePaycheck (+3 more changes) ([9f4b4bf](https://github.com/TeamWarp/warp-sdk-go/commit/9f4b4bf813262c0ece1b0e545f3202d0362e0a45))
* **api:** update 422 error response on offers.create (+2 more changes) ([7a242ca](https://github.com/TeamWarp/warp-sdk-go/commit/7a242ca2efc04e6a477207337f147d012f451846))
* **api:** update contact email ([c51fcd0](https://github.com/TeamWarp/warp-sdk-go/commit/c51fcd07eeea2d307a884bdd9cf5c72e2b6cfbc6))
* **api:** update environment production ([3f5d1c6](https://github.com/TeamWarp/warp-sdk-go/commit/3f5d1c6216b99440c23753c360bf1fbccbb1e4f7))
* **api:** update import names to warp ([87831bc](https://github.com/TeamWarp/warp-sdk-go/commit/87831bc83f412d1069ef9435c1ae7dd00823391f))
* **api:** update property objects_4.value (+68 more changes) ([2cc4138](https://github.com/TeamWarp/warp-sdk-go/commit/2cc4138291e9eb1863adb578781626a226b28590))
* **api:** update property objects_6.email (+2 more changes) ([ae95e54](https://github.com/TeamWarp/warp-sdk-go/commit/ae95e54fb7f77e6612e17a3b15326a92baa0c575))
* **api:** update response of offers.void (+166 more changes) ([6b8353b](https://github.com/TeamWarp/warp-sdk-go/commit/6b8353b77ec3c3c99cdaff1e428805b53fcf9c06))
* **api:** update schema public_pay_rate_basis (+1 more change) ([2056b24](https://github.com/TeamWarp/warp-sdk-go/commit/2056b24b554e4896678ce15164b2863306da3853))
* **api:** update schema union (+134 more changes) ([bd3a0be](https://github.com/TeamWarp/warp-sdk-go/commit/bd3a0be1a7e2dd86143d4d5cff06006ec5d7c817))
* **api:** update schema union_12 (+55 more changes) ([511d7dd](https://github.com/TeamWarp/warp-sdk-go/commit/511d7dda1c45b47f526076305dc9652f3b57ded4))
* **api:** update schema union_20 ([4a9b32c](https://github.com/TeamWarp/warp-sdk-go/commit/4a9b32c9c652fde1cfda5a27c138c4a60b26ff02))
* **api:** update SDK name (+1 more change) ([b9dfdd2](https://github.com/TeamWarp/warp-sdk-go/commit/b9dfdd2c3f6334d99b246f3a8ec3ee850487e4b5))
* **api:** update SDK name (+145 more changes) ([6ac93ee](https://github.com/TeamWarp/warp-sdk-go/commit/6ac93ee5b2ffb457a3c6a21bc75eaeeb2ca5f5fc))
* **api:** update SDK name (+2 more changes) ([49119b0](https://github.com/TeamWarp/warp-sdk-go/commit/49119b0532dfe43a446f85a1ed9fce2bbb338f4b))
* **api:** update SDK name (+27 more changes) ([1f0c2a5](https://github.com/TeamWarp/warp-sdk-go/commit/1f0c2a57eb0d071d907deef400eeddd839f3ae49))
* **api:** update SDK surface (14 changes) ([7e88858](https://github.com/TeamWarp/warp-sdk-go/commit/7e888580c1b460dd1349b53f8048dc642a243e68))
* **api:** update SDK surface (18 changes) ([dbfbbd5](https://github.com/TeamWarp/warp-sdk-go/commit/dbfbbd58ccf8afd8659614cacaffbdc4e0137819))
* **api:** update SDK surface (2 changes) ([5ab55a4](https://github.com/TeamWarp/warp-sdk-go/commit/5ab55a47bb4b9170bb5c5f29aac195ceb9d4c323))
* **api:** update SDK surface (2 changes) ([ea40d6f](https://github.com/TeamWarp/warp-sdk-go/commit/ea40d6faf17e53ddcdb7b83d1d305602b82897d0))
* **api:** update SDK surface (329 changes) ([8b89839](https://github.com/TeamWarp/warp-sdk-go/commit/8b89839a66d56e675b90697e69910e9e64107cd6))
* **api:** update SDK surface (7 changes) ([2f055ff](https://github.com/TeamWarp/warp-sdk-go/commit/2f055ffea1f7e2c3a0b3f67d4ee9d7f05b801bc8))
* **api:** update SDK surface (8 changes) ([5f10017](https://github.com/TeamWarp/warp-sdk-go/commit/5f10017a8c465edc287f8389f249c548b5a111ec))
* **api:** update SDK surface (9 changes) ([503b137](https://github.com/TeamWarp/warp-sdk-go/commit/503b13797edbb7d9729b17eb80ae0eb4506e439c))


### Chores

* **api:** regenerate SDK ([cca3c6a](https://github.com/TeamWarp/warp-sdk-go/commit/cca3c6a300ca1aad84aff2f6f7c36247df173300))
* **api:** regenerate SDK ([c141183](https://github.com/TeamWarp/warp-sdk-go/commit/c141183c3f6dabe0accc7356f9b42044246ddd98))
* **api:** regenerate SDK ([6121744](https://github.com/TeamWarp/warp-sdk-go/commit/6121744eac7639401e86e1bda8b31801dc75b27a))
* **api:** regenerate SDK ([cd4d20f](https://github.com/TeamWarp/warp-sdk-go/commit/cd4d20f18cc55b7d99265ebee7b420af23a5d4e9))
* **api:** regenerate SDK ([5f693d8](https://github.com/TeamWarp/warp-sdk-go/commit/5f693d87cd0c64f107eb57356b8ddb22932f0d7a))
* **api:** regenerate SDK ([ac04211](https://github.com/TeamWarp/warp-sdk-go/commit/ac042115bb63da525c083edb0ceb0b6fd6acc527))
* **api:** update generated SDK content ([1b8d601](https://github.com/TeamWarp/warp-sdk-go/commit/1b8d601eb75c9dbe8dd50da06ab616b68da7494e))
* **api:** update generated SDK content ([46c0f55](https://github.com/TeamWarp/warp-sdk-go/commit/46c0f55209633b0013bc906ac8beb5f4e63a52b0))
* **api:** update generated SDK content ([325146a](https://github.com/TeamWarp/warp-sdk-go/commit/325146a80da6bab28adedecbd4991c9118ec3bbd))
* **api:** update generated SDK content ([baf2df6](https://github.com/TeamWarp/warp-sdk-go/commit/baf2df6fb9a00e6d79a8c4cc26dff76ef4f12002))
* **api:** update generated SDK content ([1fe6087](https://github.com/TeamWarp/warp-sdk-go/commit/1fe6087666755e4082752d924b4b90b977b8b458))
* **api:** update generated SDK content ([47a204f](https://github.com/TeamWarp/warp-sdk-go/commit/47a204f18018b078c4b5f9e1da53e57245f7b258))
* **api:** update generated SDK content ([c30a172](https://github.com/TeamWarp/warp-sdk-go/commit/c30a172959b057054fa6e84320a98cc887889551))
* **api:** update generated SDK content ([e02c75c](https://github.com/TeamWarp/warp-sdk-go/commit/e02c75c417391e38cfbf130f1ee9ba2a6594d10a))
* configure new SDK language ([8dc14ec](https://github.com/TeamWarp/warp-sdk-go/commit/8dc14ecbba6683d8f8aaf719d4af3bd78a6bb889))
* release 0.5.0 ([3469d0b](https://github.com/TeamWarp/warp-sdk-go/commit/3469d0bc025a5446700fa23c7e925b1a7f959d33))
* set the release manifest to 0.5.1 ([9a5ec38](https://github.com/TeamWarp/warp-sdk-go/commit/9a5ec38f331f6516757850ff20e5ab57bd2747e5))
* update SDK settings ([d217727](https://github.com/TeamWarp/warp-sdk-go/commit/d217727a6569fb3b7627ba2e5d5a674e6418af84))

## [0.5.0](https://github.com/TeamWarp/warp-sdk-go/compare/v0.2.0...v0.5.0) (2026-08-10)


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
* release 0.5.0 ([3469d0b](https://github.com/TeamWarp/warp-sdk-go/commit/3469d0bc025a5446700fa23c7e925b1a7f959d33))

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
