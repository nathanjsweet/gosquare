package gosquare

import "fmt"

// RetrieveBusinessBatchRequest returns a BatchRequest object for RetrieveBusiness,
// along with a unique request id.
func RetrieveBusinessBatchRequest(token string) (*BatchRequest, string) {
	return newBatchRequest("GET", "/v1/me", token, nil)
}

// ListLocationsBatchRequest returns a BatchRequest object for ListLocations,
// along with a unique request id.
func ListLocationsBatchRequest(token string) (*BatchRequest, string) {
	return newBatchRequest("GET", "/v1/me/locations", token, nil)
}

// CreateEmployeeBatchRequest returns a BatchRequest object for CreateEmployee,
// along with a unique request id.
func CreateEmployeeBatchRequest(token string, reqObj *CreateEmployeeReqObject) (*BatchRequest, string) {
	return newBatchRequest("POST", "/v1/me/employees", token, reqObj)
}

// ListEmployeesBatchRequest returns a BatchRequest object for ListEmployees,
// along with a unique request id.
func ListEmployeesBatchRequest(token string, order, beginUpdatedAt, endUpdatedAt, beginCreatedAt, endCreatedAt, status, externalID string, limit int) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/me/employees?order=%s&begin_updated_at=%s&end_updated_at=%s&begin_created_at=%s&end_created_at=%s&status=%s&external_id=%s&limit=%d", order, beginUpdatedAt, endUpdatedAt, beginCreatedAt, endCreatedAt, status, externalID, limit), token, nil)
}

// RetrieveEmployeeBatchRequest returns a BatchRequest object for RetrieveEmployee,
// along with a unique request id.
func RetrieveEmployeeBatchRequest(token, employeeID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/me/employees/%s", employeeID), token, nil)
}

// UpdateEmployeeBatchRequest returns a BatchRequest object for UpdateEmployee,
// along with a unique request id.
func UpdateEmployeeBatchRequest(token, employeeID string, reqObj *UpdateEmployeeReqObject) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/me/employees/%s", employeeID), token, reqObj)
}

// CreateRoleBatchRequest returns a BatchRequest object for CreateRole,
// along with a unique request id.
func CreateRoleBatchRequest(token string, reqObj *CreateRoleReqObject) (*BatchRequest, string) {
	return newBatchRequest("POST", "/v1/me/roles", token, reqObj)
}

// ListRolesBatchRequest returns a BatchRequest object for ListRoles,
// along with a unique request id.
func ListRolesBatchRequest(token, order string, limit int) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/me/roles?order=%s&limit=%d", order, limit), token, nil)
}

// RetrieveRoleBatchRequest returns a BatchRequest object for RetrieveRole,
// along with a unique request id.
func RetrieveRoleBatchRequest(token, roleID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/me/roles/%s", roleID), token, nil)
}

// UpdateRoleBatchRequest returns a BatchRequest object for UpdateRole,
// along with a unique request id.
func UpdateRoleBatchRequest(token, roleID string, reqObj *UpdateRoleReqObject) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/me/roles/%s", roleID), token, reqObj)
}

// CreateTimecardBatchRequest returns a BatchRequest object for CreateTimecard,
// along with a unique request id.
func CreateTimecardBatchRequest(token string, reqObj *CreateTimecardReqObject) (*BatchRequest, string) {
	return newBatchRequest("POST", "/v1/me/timecards", token, reqObj)
}

// ListTimecardsBatchRequest returns a BatchRequest object for ListTimecards,
// along with a unique request id.
func ListTimecardsBatchRequest(token, order, employeeID, beginClockinTime, endClockinTime, beginClockoutTime, endClockoutTime, beginUpdatedAt, endUpdatedAt string, deleted bool, limit int) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/me/timecards?order=%s&employee_id=%s&begin_clockin_time=%s&end_clockin_time=%s&begin_clockout_time=%s&end_clockout_time=%s&begin_updated_at=%s&end_updated_at=%s&deleted=%t&limit=%d", order, employeeID, beginClockinTime, endClockinTime, beginClockoutTime, endClockoutTime, beginUpdatedAt, endUpdatedAt, deleted, limit), token, nil)
}

// RetrieveTimecardBatchRequest returns a BatchRequest object for RetrieveTimecard,
// along with a unique request id.
func RetrieveTimecardBatchRequest(token, timecardID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/me/timecards/%s", timecardID), token, nil)
}

// UpdateTimecardBatchRequest returns a BatchRequest object for UpdateTimecard,
// along with a unique request id.
func UpdateTimecardBatchRequest(token, timecardID string, reqObj *UpdateTimecardReqObject) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/me/timecards/%s", timecardID), token, reqObj)
}

// DeleteTimecardBatchRequest returns a BatchRequest object for DeleteTimecard,
// along with a unique request id.
func DeleteTimecardBatchRequest(token, timecardID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/me/timecards/%s", timecardID), token, nil)
}

// ListTimecardEventsBatchRequest returns a BatchRequest object for ListTimecardEvents,
// along with a unique request id.
func ListTimecardEventsBatchRequest(token, timecardID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/me/timecards/%s/events", timecardID), token, nil)
}

// ListCashDrawerShiftsBatchRequest returns a BatchRequest object for ListCashDrawerShifts,
// along with a unique request id.
func ListCashDrawerShiftsBatchRequest(token, locationID, beginTime, endTime, order string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/cash-drawer-shifts?begin_time=%s&end_time=%s&order=%s", locationID, beginTime, endTime, order), token, nil)
}

// RetrieveCashDrawerShiftBatchRequest returns a BatchRequest object for RetrieveCashDrawerShift,
// along with a unique request id.
func RetrieveCashDrawerShiftBatchRequest(token, locationID, shiftID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/cash-drawer-shifts/%s", locationID, shiftID), token, nil)
}

// ListPaymentsBatchRequest returns a BatchRequest object for ListPayments,
// along with a unique request id.
func ListPaymentsBatchRequest(token, locationID, beginTime, endTime, order string, limit int) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/payments?begin_time=%s&end_time=%s&order=%s&limit=%d", locationID, beginTime, endTime, order, limit), token, nil)
}

// RetrievePaymentBatchRequest returns a BatchRequest object for RetrievePayment,
// along with a unique request id.
func RetrievePaymentBatchRequest(token, locationID, paymentID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/payments/%s", locationID, paymentID), token, nil)
}

// ListSettlementsBatchRequest returns a BatchRequest object for ListSettlements,
// along with a unique request id.
func ListSettlementsBatchRequest(token, locationID, beginTime, endTime, order string, limit int, status string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/settlements?begin_time=%s&end_time=%s&order=%s&limit=%d&status=%s", locationID, beginTime, endTime, order, limit, status), token, nil)
}

// RetrieveSettlementBatchRequest returns a BatchRequest object for RetrieveSettlement,
// along with a unique request id.
func RetrieveSettlementBatchRequest(token, locationID, settlementID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/settlements/%s", locationID, settlementID), token, nil)
}

// CreateRefundBatchRequest returns a BatchRequest object for CreateRefund,
// along with a unique request id.
func CreateRefundBatchRequest(token, locationID string, reqObj *CreateRefundReqObject) (*BatchRequest, string) {
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/refunds", locationID), token, reqObj)
}

// ListRefundsBatchRequest returns a BatchRequest object for ListRefunds,
// along with a unique request id.
func ListRefundsBatchRequest(token, locationID, beginTime, endTime, order string, limit int) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/refunds?begin_time=%s&end_time=%s&order=%s&limit=%d", locationID, beginTime, endTime, order, limit), token, nil)
}

// ListOrdersBatchRequest returns a BatchRequest object for ListOrders,
// along with a unique request id.
func ListOrdersBatchRequest(token, locationID string, limit int, order string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/orders?limit=%d&order=%s", locationID, limit, order), token, nil)
}

// RetrieveOrderBatchRequest returns a BatchRequest object for RetrieveOrder,
// along with a unique request id.
func RetrieveOrderBatchRequest(token, locationID, orderID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/orders/%s", locationID, orderID), token, nil)
}

// UpdateOrderBatchRequest returns a BatchRequest object for UpdateOrder,
// along with a unique request id.
func UpdateOrderBatchRequest(token, locationID, orderID string, reqObj *UpdateOrderReqObject) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/orders/%s", locationID, orderID), token, reqObj)
}

// ListBankAccountsBatchRequest returns a BatchRequest object for ListBankAccounts,
// along with a unique request id.
func ListBankAccountsBatchRequest(token, locationID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/bank-accounts", locationID), token, nil)
}

// RetrieveBankAccountBatchRequest returns a BatchRequest object for RetrieveBankAccount,
// along with a unique request id.
func RetrieveBankAccountBatchRequest(token, locationID, bankAccountID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/bank-accounts/%s", locationID, bankAccountID), token, nil)
}

// CreateItemBatchRequest returns a BatchRequest object for CreateItem,
// along with a unique request id.
func CreateItemBatchRequest(token, locationID string, reqObj *CreateItemReqObject) (*BatchRequest, string) {
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/items", locationID), token, reqObj)
}

// ListItemsBatchRequest returns a BatchRequest object for ListItems,
// along with a unique request id.
func ListItemsBatchRequest(token, locationID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/items", locationID), token, nil)
}

// RetrieveItemBatchRequest returns a BatchRequest object for RetrieveItem,
// along with a unique request id.
func RetrieveItemBatchRequest(token, locationID, itemID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/items/%s", locationID, itemID), token, nil)
}

// UpdateItemBatchRequest returns a BatchRequest object for UpdateItem,
// along with a unique request id.
func UpdateItemBatchRequest(token, locationID, itemID string, reqObj *UpdateItemReqObject) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/items/%s", locationID, itemID), token, reqObj)
}

// DeleteItemBatchRequest returns a BatchRequest object for DeleteItem,
// along with a unique request id.
func DeleteItemBatchRequest(token, locationID, itemID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/items/%s", locationID, itemID), token, nil)
}

// CreateVariationBatchRequest returns a BatchRequest object for CreateVariation,
// along with a unique request id.
func CreateVariationBatchRequest(token, locationID, itemID string, reqObj *CreateVariationReqObject) (*BatchRequest, string) {
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/items/%s/variations", locationID, itemID), token, reqObj)
}

// UpdateVariationBatchRequest returns a BatchRequest object for UpdateVariation,
// along with a unique request id.
func UpdateVariationBatchRequest(token, locationID, itemID, variationID string, reqObj *UpdateVariationReqObject) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/items/%s/variations/%s", locationID, itemID, variationID), token, reqObj)
}

// DeleteVariationBatchRequest returns a BatchRequest object for DeleteVariation,
// along with a unique request id.
func DeleteVariationBatchRequest(token, locationID, itemID, variationID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/items/%s/variations/%s", locationID, itemID, variationID), token, nil)
}

// ListInventoryBatchRequest returns a BatchRequest object for ListInventory,
// along with a unique request id.
func ListInventoryBatchRequest(token, locationID string, limit int) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/inventory?limit=%d", locationID, limit), token, nil)
}

// AdjustInventoryBatchRequest returns a BatchRequest object for AdjustInventory,
// along with a unique request id.
func AdjustInventoryBatchRequest(token, locationID, variationID string, reqObj *AdjustInventoryReqObject) (*BatchRequest, string) {
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/inventory/%s", locationID, variationID), token, reqObj)
}

// CreateModifierListBatchRequest returns a BatchRequest object for CreateModifierList,
// along with a unique request id.
func CreateModifierListBatchRequest(token, locationID string, reqObj *CreateModifierListReqObject) (*BatchRequest, string) {
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/modifier-lists", locationID), token, reqObj)
}

// ListModifierListsBatchRequest returns a BatchRequest object for ListModifierLists,
// along with a unique request id.
func ListModifierListsBatchRequest(token, locationID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/modifier-lists", locationID), token, nil)
}

// RetrieveModifierListBatchRequest returns a BatchRequest object for RetrieveModifierList,
// along with a unique request id.
func RetrieveModifierListBatchRequest(token, locationID, modifierListID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/modifier-lists/%s", locationID, modifierListID), token, nil)
}

// UpdateModifierListBatchRequest returns a BatchRequest object for UpdateModifierList,
// along with a unique request id.
func UpdateModifierListBatchRequest(token, locationID, modifierListID string, reqObj *UpdateModifierListReqObject) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/modifier-lists/%s", locationID, modifierListID), token, reqObj)
}

// DeleteModifierListBatchRequest returns a BatchRequest object for DeleteModifierList,
// along with a unique request id.
func DeleteModifierListBatchRequest(token, locationID, modifierListID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/modifier-lists/%s", locationID, modifierListID), token, nil)
}

// ApplyModifierListBatchRequest returns a BatchRequest object for ApplyModifierList,
// along with a unique request id.
func ApplyModifierListBatchRequest(token, locationID, itemID, modifierListID string) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/items/%s/modifier-lists/%s", locationID, itemID, modifierListID), token, nil)
}

// RemoveModifierListBatchRequest returns a BatchRequest object for RemoveModifierList,
// along with a unique request id.
func RemoveModifierListBatchRequest(token, locationID, itemID, modifierListID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/items/%s/modifier-lists/%s", locationID, itemID, modifierListID), token, nil)
}

// CreateModifierOptionBatchRequest returns a BatchRequest object for CreateModifierOption,
// along with a unique request id.
func CreateModifierOptionBatchRequest(token, locationID, modifierListID string, reqObj *CreateModifierOptionReqObject) (*BatchRequest, string) {
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/modifier-lists/%s/modifier-options", locationID, modifierListID), token, reqObj)
}

// UpdateModifierOptionBatchRequest returns a BatchRequest object for UpdateModifierOption,
// along with a unique request id.
func UpdateModifierOptionBatchRequest(token, locationID, modifierListID, modifierOptionID string, reqObj *UpdateModifierOptionReqObject) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/modifier-lists/%s/modifier-options/%s", locationID, modifierListID, modifierOptionID), token, reqObj)
}

// DeleteModifierOptionBatchRequest returns a BatchRequest object for DeleteModifierOption,
// along with a unique request id.
func DeleteModifierOptionBatchRequest(token, locationID, modifierListID, modifierOptionID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/modifier-lists/%s/modifier-options/%s", locationID, modifierListID, modifierOptionID), token, nil)
}

// CreateCategoryBatchRequest returns a BatchRequest object for CreateCategory,
// along with a unique request id.
func CreateCategoryBatchRequest(token, locationID string, reqObj *CreateCategoryReqObject) (*BatchRequest, string) {
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/categories", locationID), token, reqObj)
}

// ListCategoriesBatchRequest returns a BatchRequest object for ListCategories,
// along with a unique request id.
func ListCategoriesBatchRequest(token, locationID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/categories", locationID), token, nil)
}

// UpdateCategoryBatchRequest returns a BatchRequest object for UpdateCategory,
// along with a unique request id.
func UpdateCategoryBatchRequest(token, locationID, categoryID string, reqObj *UpdateCategoryReqObject) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/categories/%s", locationID, categoryID), token, reqObj)
}

// DeleteCategoryBatchRequest returns a BatchRequest object for DeleteCategory,
// along with a unique request id.
func DeleteCategoryBatchRequest(token, locationID, categoryID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/categories/%s", locationID, categoryID), token, nil)
}

// CreateDiscountBatchRequest returns a BatchRequest object for CreateDiscount,
// along with a unique request id.
func CreateDiscountBatchRequest(token, locationID string, reqObj *CreateDiscountReqObject) (*BatchRequest, string) {
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/discounts", locationID), token, reqObj)
}

// ListDiscountsBatchRequest returns a BatchRequest object for ListDiscounts,
// along with a unique request id.
func ListDiscountsBatchRequest(token, locationID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/discounts", locationID), token, nil)
}

// UpdateDiscountBatchRequest returns a BatchRequest object for UpdateDiscount,
// along with a unique request id.
func UpdateDiscountBatchRequest(token, locationID, discountID string, reqObj *UpdateDiscountReqObject) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/discounts/%s", locationID, discountID), token, reqObj)
}

// DeleteDiscountBatchRequest returns a BatchRequest object for DeleteDiscount,
// along with a unique request id.
func DeleteDiscountBatchRequest(token, locationID, discountID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/discounts/%s", locationID, discountID), token, nil)
}

// CreateFeeBatchRequest returns a BatchRequest object for CreateFee,
// along with a unique request id.
func CreateFeeBatchRequest(token, locationID string, reqObj *CreateFeeReqObject) (*BatchRequest, string) {
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/fees", locationID), token, reqObj)
}

// ListFeesBatchRequest returns a BatchRequest object for ListFees,
// along with a unique request id.
func ListFeesBatchRequest(token, locationID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/fees", locationID), token, nil)
}

// UpdateFeeBatchRequest returns a BatchRequest object for UpdateFee,
// along with a unique request id.
func UpdateFeeBatchRequest(token, locationID, feeID string, reqObj *UpdateFeeReqObject) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/fees/%s", locationID, feeID), token, reqObj)
}

// DeleteFeeBatchRequest returns a BatchRequest object for DeleteFee,
// along with a unique request id.
func DeleteFeeBatchRequest(token, locationID, feeID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/fees/%s", locationID, feeID), token, nil)
}

// ApplyFeeBatchRequest returns a BatchRequest object for ApplyFee,
// along with a unique request id.
func ApplyFeeBatchRequest(token, locationID, itemID, feeID string) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/items/%s/fees/%s", locationID, itemID, feeID), token, nil)
}

// RemoveFeeBatchRequest returns a BatchRequest object for RemoveFee,
// along with a unique request id.
func RemoveFeeBatchRequest(token, locationID, itemID, feeID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/items/%s/fees/%s", locationID, itemID, feeID), token, nil)
}

// CreatePageBatchRequest returns a BatchRequest object for CreatePage,
// along with a unique request id.
func CreatePageBatchRequest(token, locationID string, reqObj *CreatePageReqObject) (*BatchRequest, string) {
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/pages", locationID), token, reqObj)
}

// ListPagesBatchRequest returns a BatchRequest object for ListPages,
// along with a unique request id.
func ListPagesBatchRequest(token, locationID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/pages", locationID), token, nil)
}

// UpdatePageBatchRequest returns a BatchRequest object for UpdatePage,
// along with a unique request id.
func UpdatePageBatchRequest(token, locationID, pageID string, reqObj *UpdatePageReqObject) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/pages/%s", locationID, pageID), token, reqObj)
}

// DeletePageBatchRequest returns a BatchRequest object for DeletePage,
// along with a unique request id.
func DeletePageBatchRequest(token, locationID, pageID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/pages/%s", locationID, pageID), token, nil)
}

// UpdateCellBatchRequest returns a BatchRequest object for UpdateCell,
// along with a unique request id.
func UpdateCellBatchRequest(token, locationID, pageID string, reqObj *UpdateCellReqObject) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/pages/%s/cells", locationID, pageID), token, reqObj)
}

// DeleteCellBatchRequest returns a BatchRequest object for DeleteCell,
// along with a unique request id.
func DeleteCellBatchRequest(token, locationID, pageID string, row, column int) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/pages/%s/cells?row=%d&column=%d", locationID, pageID, row, column), token, nil)
}

// ListWebhooksBatchRequest returns a BatchRequest object for ListWebhooks,
// along with a unique request id.
func ListWebhooksBatchRequest(token, locationID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/webhooks", locationID), token, nil)
}

// UpdateWebhooksBatchRequest returns a BatchRequest object for UpdateWebhooks,
// along with a unique request id.
func UpdateWebhooksBatchRequest(token, locationID string) (*BatchRequest, string) {
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/webhooks", locationID), token, nil)
}

// ListSubscriptionsBatchRequest returns a BatchRequest object for ListSubscriptions,
// along with a unique request id.
func ListSubscriptionsBatchRequest(token, clientID, merchantID string, limit int) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/oauth2/clients/%s/subscriptions?merchant_id=%s&limit=%d", clientID, merchantID, limit), token, nil)
}

// RetrieveSubscriptionBatchRequest returns a BatchRequest object for RetrieveSubscription,
// along with a unique request id.
func RetrieveSubscriptionBatchRequest(token, clientID, subscriptionID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/oauth2/clients/%s/subscriptions/%s", clientID, subscriptionID), token, nil)
}

// ListSubscriptionPlansBatchRequest returns a BatchRequest object for ListSubscriptionPlans,
// along with a unique request id.
func ListSubscriptionPlansBatchRequest(token, clientID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/oauth2/clients/%s/plans", clientID), token, nil)
}

// RetrieveSubscriptionPlanBatchRequest returns a BatchRequest object for RetrieveSubscriptionPlan,
// along with a unique request id.
func RetrieveSubscriptionPlanBatchRequest(token, clientID, planID string) (*BatchRequest, string) {
	return newBatchRequest("GET", fmt.Sprintf("/oauth2/clients/%s/plans/%s", clientID, planID), token, nil)
}
