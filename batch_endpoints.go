package gosquare

import "fmt"

// RetrieveBusinessBatchRequest returns a BatchRequest object for RetrieveBusiness,
// along with a unique request id.
func RetrieveBusinessBatchRequest(token string) (*BatchRequest, string) {
	v := new(Merchant)
	return newBatchRequest("GET", "/v1/me", token, nil, v)
}

// ListLocationsBatchRequest returns a BatchRequest object for ListLocations,
// along with a unique request id.
func ListLocationsBatchRequest(token string) (*BatchRequest, string) {
	v := make([]*Merchant, 0)
	return newBatchRequest("GET", "/v1/me/locations", token, nil, &v)
}

// CreateEmployeeBatchRequest returns a BatchRequest object for CreateEmployee,
// along with a unique request id.
func CreateEmployeeBatchRequest(token string, reqObj *CreateEmployeeReqObject) (*BatchRequest, string) {
	v := new(Employee)
	return newBatchRequest("POST", "/v1/me/employees", token, reqObj, v)
}

// ListEmployeesBatchRequest returns a BatchRequest object for ListEmployees,
// along with a unique request id.
func ListEmployeesBatchRequest(token string, order, beginUpdatedAt, endUpdatedAt, beginCreatedAt, endCreatedAt, status, externalID string, limit int) (*BatchRequest, string) {
	v := make([]*Employee, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/me/employees?order=%s&begin_updated_at=%s&end_updated_at=%s&begin_created_at=%s&end_created_at=%s&status=%s&external_id=%s&limit=%d", order, beginUpdatedAt, endUpdatedAt, beginCreatedAt, endCreatedAt, status, externalID, limit), token, nil, &v)
}

// RetrieveEmployeeBatchRequest returns a BatchRequest object for RetrieveEmployee,
// along with a unique request id.
func RetrieveEmployeeBatchRequest(token, employeeID string) (*BatchRequest, string) {
	v := new(Employee)
	return newBatchRequest("GET", fmt.Sprintf("/v1/me/employees/%s", employeeID), token, nil, v)
}

// UpdateEmployeeBatchRequest returns a BatchRequest object for UpdateEmployee,
// along with a unique request id.
func UpdateEmployeeBatchRequest(token, employeeID string, reqObj *UpdateEmployeeReqObject) (*BatchRequest, string) {
	v := new(Employee)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/me/employees/%s", employeeID), token, reqObj, v)
}

// CreateRoleBatchRequest returns a BatchRequest object for CreateRole,
// along with a unique request id.
func CreateRoleBatchRequest(token string, reqObj *CreateRoleReqObject) (*BatchRequest, string) {
	v := new(EmployeeRole)
	return newBatchRequest("POST", "/v1/me/roles", token, reqObj, v)
}

// ListRolesBatchRequest returns a BatchRequest object for ListRoles,
// along with a unique request id.
func ListRolesBatchRequest(token, order string, limit int) (*BatchRequest, string) {
	v := make([]*EmployeeRole, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/me/roles?order=%s&limit=%d", order, limit), token, nil, &v)
}

// RetrieveRoleBatchRequest returns a BatchRequest object for RetrieveRole,
// along with a unique request id.
func RetrieveRoleBatchRequest(token, roleID string) (*BatchRequest, string) {
	v := new(EmployeeRole)
	return newBatchRequest("GET", fmt.Sprintf("/v1/me/roles/%s", roleID), token, nil, v)
}

// UpdateRoleBatchRequest returns a BatchRequest object for UpdateRole,
// along with a unique request id.
func UpdateRoleBatchRequest(token, roleID string, reqObj *UpdateRoleReqObject) (*BatchRequest, string) {
	v := new(EmployeeRole)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/me/roles/%s", roleID), token, reqObj, v)
}

// CreateTimecardBatchRequest returns a BatchRequest object for CreateTimecard,
// along with a unique request id.
func CreateTimecardBatchRequest(token string, reqObj *CreateTimecardReqObject) (*BatchRequest, string) {
	v := new(Timecard)
	return newBatchRequest("POST", "/v1/me/timecards", token, reqObj, v)
}

// ListTimecardsBatchRequest returns a BatchRequest object for ListTimecards,
// along with a unique request id.
func ListTimecardsBatchRequest(token, order, employeeID, beginClockinTime, endClockinTime, beginClockoutTime, endClockoutTime, beginUpdatedAt, endUpdatedAt string, deleted bool, limit int) (*BatchRequest, string) {
	v := make([]*Timecard, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/me/timecards?order=%s&employee_id=%s&begin_clockin_time=%s&end_clockin_time=%s&begin_clockout_time=%s&end_clockout_time=%s&begin_updated_at=%s&end_updated_at=%s&deleted=%t&limit=%d", order, employeeID, beginClockinTime, endClockinTime, beginClockoutTime, endClockoutTime, beginUpdatedAt, endUpdatedAt, deleted, limit), token, nil, &v)
}

// RetrieveTimecardBatchRequest returns a BatchRequest object for RetrieveTimecard,
// along with a unique request id.
func RetrieveTimecardBatchRequest(token, timecardID string) (*BatchRequest, string) {
	v := new(Timecard)
	return newBatchRequest("GET", fmt.Sprintf("/v1/me/timecards/%s", timecardID), token, nil, v)
}

// UpdateTimecardBatchRequest returns a BatchRequest object for UpdateTimecard,
// along with a unique request id.
func UpdateTimecardBatchRequest(token, timecardID string, reqObj *UpdateTimecardReqObject) (*BatchRequest, string) {
	v := new(Timecard)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/me/timecards/%s", timecardID), token, reqObj, v)
}

// DeleteTimecardBatchRequest returns a BatchRequest object for DeleteTimecard,
// along with a unique request id.
func DeleteTimecardBatchRequest(token, timecardID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/me/timecards/%s", timecardID), token, nil, nil)
}

// ListTimecardEventsBatchRequest returns a BatchRequest object for ListTimecardEvents,
// along with a unique request id.
func ListTimecardEventsBatchRequest(token, timecardID string) (*BatchRequest, string) {
	v := make([]*TimecardEvent, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/me/timecards/%s/events", timecardID), token, nil, &v)
}

// ListCashDrawerShiftsBatchRequest returns a BatchRequest object for ListCashDrawerShifts,
// along with a unique request id.
func ListCashDrawerShiftsBatchRequest(token, locationID, beginTime, endTime, order string) (*BatchRequest, string) {
	v := make([]*CashDrawerShift, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/cash-drawer-shifts?begin_time=%s&end_time=%s&order=%s", locationID, beginTime, endTime, order), token, nil, &v)
}

// RetrieveCashDrawerShiftBatchRequest returns a BatchRequest object for RetrieveCashDrawerShift,
// along with a unique request id.
func RetrieveCashDrawerShiftBatchRequest(token, locationID, shiftID string) (*BatchRequest, string) {
	v := new(CashDrawerShift)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/cash-drawer-shifts/%s", locationID, shiftID), token, nil, v)
}

// ListPaymentsBatchRequest returns a BatchRequest object for ListPayments,
// along with a unique request id.
func ListPaymentsBatchRequest(token, locationID, beginTime, endTime, order string, limit int) (*BatchRequest, string) {
	v := make([]*Payment, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/payments?begin_time=%s&end_time=%s&order=%s&limit=%d", locationID, beginTime, endTime, order, limit), token, nil, &v)
}

// RetrievePaymentBatchRequest returns a BatchRequest object for RetrievePayment,
// along with a unique request id.
func RetrievePaymentBatchRequest(token, locationID, paymentID string) (*BatchRequest, string) {
	v := new(Payment)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/payments/%s", locationID, paymentID), token, nil, v)
}

// ListSettlementsBatchRequest returns a BatchRequest object for ListSettlements,
// along with a unique request id.
func ListSettlementsBatchRequest(token, locationID, beginTime, endTime, order string, limit int, status string) (*BatchRequest, string) {
	v := make([]*Settlement, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/settlements?begin_time=%s&end_time=%s&order=%s&limit=%d&status=%s", locationID, beginTime, endTime, order, limit, status), token, nil, &v)
}

// RetrieveSettlementBatchRequest returns a BatchRequest object for RetrieveSettlement,
// along with a unique request id.
func RetrieveSettlementBatchRequest(token, locationID, settlementID string) (*BatchRequest, string) {
	v := new(Settlement)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/settlements/%s", locationID, settlementID), token, nil, v)
}

// CreateRefundBatchRequest returns a BatchRequest object for CreateRefund,
// along with a unique request id.
func CreateRefundBatchRequest(token, locationID string, reqObj *CreateRefundReqObject) (*BatchRequest, string) {
	v := new(Refund)
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/refunds", locationID), token, reqObj, v)
}

// ListRefundsBatchRequest returns a BatchRequest object for ListRefunds,
// along with a unique request id.
func ListRefundsBatchRequest(token, locationID, beginTime, endTime, order string, limit int) (*BatchRequest, string) {
	v := make([]*Refund, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/refunds?begin_time=%s&end_time=%s&order=%s&limit=%d", locationID, beginTime, endTime, order, limit), token, nil, &v)
}

// ListOrdersBatchRequest returns a BatchRequest object for ListOrders,
// along with a unique request id.
func ListOrdersBatchRequest(token, locationID string, limit int, order string) (*BatchRequest, string) {
	v := make([]*Order, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/orders?limit=%d&order=%s", locationID, limit, order), token, nil, &v)
}

// RetrieveOrderBatchRequest returns a BatchRequest object for RetrieveOrder,
// along with a unique request id.
func RetrieveOrderBatchRequest(token, locationID, orderID string) (*BatchRequest, string) {
	v := new(Order)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/orders/%s", locationID, orderID), token, nil, v)
}

// UpdateOrderBatchRequest returns a BatchRequest object for UpdateOrder,
// along with a unique request id.
func UpdateOrderBatchRequest(token, locationID, orderID string, reqObj *UpdateOrderReqObject) (*BatchRequest, string) {
	v := new(Order)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/orders/%s", locationID, orderID), token, reqObj, v)
}

// ListBankAccountsBatchRequest returns a BatchRequest object for ListBankAccounts,
// along with a unique request id.
func ListBankAccountsBatchRequest(token, locationID string) (*BatchRequest, string) {
	v := make([]*BankAccount, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/bank-accounts", locationID), token, nil, &v)
}

// RetrieveBankAccountBatchRequest returns a BatchRequest object for RetrieveBankAccount,
// along with a unique request id.
func RetrieveBankAccountBatchRequest(token, locationID, bankAccountID string) (*BatchRequest, string) {
	v := new(BankAccount)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/bank-accounts/%s", locationID, bankAccountID), token, nil, v)
}

// CreateItemBatchRequest returns a BatchRequest object for CreateItem,
// along with a unique request id.
func CreateItemBatchRequest(token, locationID string, reqObj *CreateItemReqObject) (*BatchRequest, string) {
	v := new(Item)
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/items", locationID), token, reqObj, v)
}

// ListItemsBatchRequest returns a BatchRequest object for ListItems,
// along with a unique request id.
func ListItemsBatchRequest(token, locationID string) (*BatchRequest, string) {
	v := make([]*Item, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/items", locationID), token, nil, &v)
}

// RetrieveItemBatchRequest returns a BatchRequest object for RetrieveItem,
// along with a unique request id.
func RetrieveItemBatchRequest(token, locationID, itemID string) (*BatchRequest, string) {
	v := new(Item)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/items/%s", locationID, itemID), token, nil, v)
}

// UpdateItemBatchRequest returns a BatchRequest object for UpdateItem,
// along with a unique request id.
func UpdateItemBatchRequest(token, locationID, itemID string, reqObj *UpdateItemReqObject) (*BatchRequest, string) {
	v := new(Item)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/items/%s", locationID, itemID), token, reqObj, v)
}

// DeleteItemBatchRequest returns a BatchRequest object for DeleteItem,
// along with a unique request id.
func DeleteItemBatchRequest(token, locationID, itemID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/items/%s", locationID, itemID), token, nil, nil)
}

// UpdateVariationBatchRequest returns a BatchRequest object for UpdateVariation,
// along with a unique request id.
func UpdateVariationBatchRequest(token, locationID, itemID, variationID string, reqObj *UpdateVariationReqObject) (*BatchRequest, string) {
	v := new(ItemVariation)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/items/%s/variations/%s", locationID, itemID, variationID), token, reqObj, v)
}

// DeleteVariationBatchRequest returns a BatchRequest object for DeleteVariation,
// along with a unique request id.
func DeleteVariationBatchRequest(token, locationID, itemID, variationID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/items/%s/variations/%s", locationID, itemID, variationID), token, nil, nil)
}

// ListInventoryBatchRequest returns a BatchRequest object for ListInventory,
// along with a unique request id.
func ListInventoryBatchRequest(token, locationID string, limit int) (*BatchRequest, string) {
	v := make([]*InventoryEntry, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/inventory?limit=%d", locationID, limit), token, nil, &v)
}

// AdjustInventoryBatchRequest returns a BatchRequest object for AdjustInventory,
// along with a unique request id.
func AdjustInventoryBatchRequest(token, locationID, variationID string, reqObj *AdjustInventoryReqObject) (*BatchRequest, string) {
	v := new(InventoryEntry)
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/inventory/%s", locationID, variationID), token, reqObj, v)
}

// CreateModifierListBatchRequest returns a BatchRequest object for CreateModifierList,
// along with a unique request id.
func CreateModifierListBatchRequest(token, locationID string, reqObj *CreateModifierListReqObject) (*BatchRequest, string) {
	v := new(ModifierList)
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/modifier-lists", locationID), token, reqObj, v)
}

// ListModifierListsBatchRequest returns a BatchRequest object for ListModifierLists,
// along with a unique request id.
func ListModifierListsBatchRequest(token, locationID string) (*BatchRequest, string) {
	v := make([]*ModifierList, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/modifier-lists", locationID), token, nil, &v)
}

// RetrieveModifierListBatchRequest returns a BatchRequest object for RetrieveModifierList,
// along with a unique request id.
func RetrieveModifierListBatchRequest(token, locationID, modifierListID string) (*BatchRequest, string) {
	v := new(ModifierList)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/modifier-lists/%s", locationID, modifierListID), token, nil, v)
}

// UpdateModifierListBatchRequest returns a BatchRequest object for UpdateModifierList,
// along with a unique request id.
func UpdateModifierListBatchRequest(token, locationID, modifierListID string, reqObj *UpdateModifierListReqObject) (*BatchRequest, string) {
	v := new(ModifierList)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/modifier-lists/%s", locationID, modifierListID), token, reqObj, v)
}

// DeleteModifierListBatchRequest returns a BatchRequest object for DeleteModifierList,
// along with a unique request id.
func DeleteModifierListBatchRequest(token, locationID, modifierListID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/modifier-lists/%s", locationID, modifierListID), token, nil, nil)
}

// ApplyModifierListBatchRequest returns a BatchRequest object for ApplyModifierList,
// along with a unique request id.
func ApplyModifierListBatchRequest(token, locationID, itemID, modifierListID string) (*BatchRequest, string) {
	v := new(Item)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/items/%s/modifier-lists/%s", locationID, itemID, modifierListID), token, nil, v)
}

// RemoveModifierListBatchRequest returns a BatchRequest object for RemoveModifierList,
// along with a unique request id.
func RemoveModifierListBatchRequest(token, locationID, itemID, modifierListID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/items/%s/modifier-lists/%s", locationID, itemID, modifierListID), token, nil, nil)
}

// CreateModifierOptionBatchRequest returns a BatchRequest object for CreateModifierOption,
// along with a unique request id.
func CreateModifierOptionBatchRequest(token, locationID, modifierListID string, reqObj *CreateModifierOptionReqObject) (*BatchRequest, string) {
	v := new(ModifierOption)
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/modifier-lists/%s/modifier-options", locationID, modifierListID), token, reqObj, v)
}

// UpdateModifierOptionBatchRequest returns a BatchRequest object for UpdateModifierOption,
// along with a unique request id.
func UpdateModifierOptionBatchRequest(token, locationID, modifierListID, modifierOptionID string, reqObj *UpdateModifierOptionReqObject) (*BatchRequest, string) {
	v := new(ModifierOption)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/modifier-lists/%s/modifier-options/%s", locationID, modifierListID, modifierOptionID), token, reqObj, v)
}

// DeleteModifierOptionBatchRequest returns a BatchRequest object for DeleteModifierOption,
// along with a unique request id.
func DeleteModifierOptionBatchRequest(token, locationID, modifierListID, modifierOptionID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/modifier-lists/%s/modifier-options/%s", locationID, modifierListID, modifierOptionID), token, nil, nil)
}

// CreateCategoryBatchRequest returns a BatchRequest object for CreateCategory,
// along with a unique request id.
func CreateCategoryBatchRequest(token, locationID string, reqObj *CreateCategoryReqObject) (*BatchRequest, string) {
	v := new(Category)
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/categories", locationID), token, reqObj, v)
}

// ListCategoriesBatchRequest returns a BatchRequest object for ListCategories,
// along with a unique request id.
func ListCategoriesBatchRequest(token, locationID string) (*BatchRequest, string) {
	v := make([]*Category, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/categories", locationID), token, nil, &v)
}

// UpdateCategoryBatchRequest returns a BatchRequest object for UpdateCategory,
// along with a unique request id.
func UpdateCategoryBatchRequest(token, locationID, categoryID string, reqObj *UpdateCategoryReqObject) (*BatchRequest, string) {
	v := new(Category)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/categories/%s", locationID, categoryID), token, reqObj, v)
}

// DeleteCategoryBatchRequest returns a BatchRequest object for DeleteCategory,
// along with a unique request id.
func DeleteCategoryBatchRequest(token, locationID, categoryID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/categories/%s", locationID, categoryID), token, nil, nil)
}

// CreateDiscountBatchRequest returns a BatchRequest object for CreateDiscount,
// along with a unique request id.
func CreateDiscountBatchRequest(token, locationID string, reqObj *CreateDiscountReqObject) (*BatchRequest, string) {
	v := new(Discount)
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/discounts", locationID), token, reqObj, v)
}

// ListDiscountsBatchRequest returns a BatchRequest object for ListDiscounts,
// along with a unique request id.
func ListDiscountsBatchRequest(token, locationID string) (*BatchRequest, string) {
	v := make([]*Discount, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/discounts", locationID), token, nil, &v)
}

// UpdateDiscountBatchRequest returns a BatchRequest object for UpdateDiscount,
// along with a unique request id.
func UpdateDiscountBatchRequest(token, locationID, discountID string, reqObj *UpdateDiscountReqObject) (*BatchRequest, string) {
	v := new(Discount)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/discounts/%s", locationID, discountID), token, reqObj, v)
}

// DeleteDiscountBatchRequest returns a BatchRequest object for DeleteDiscount,
// along with a unique request id.
func DeleteDiscountBatchRequest(token, locationID, discountID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/discounts/%s", locationID, discountID), token, nil, nil)
}

// CreateFeeBatchRequest returns a BatchRequest object for CreateFee,
// along with a unique request id.
func CreateFeeBatchRequest(token, locationID string, reqObj *CreateFeeReqObject) (*BatchRequest, string) {
	v := new(Fee)
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/fees", locationID), token, reqObj, v)
}

// ListFeesBatchRequest returns a BatchRequest object for ListFees,
// along with a unique request id.
func ListFeesBatchRequest(token, locationID string) (*BatchRequest, string) {
	v := make([]*Fee, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/fees", locationID), token, nil, &v)
}

// UpdateFeeBatchRequest returns a BatchRequest object for UpdateFee,
// along with a unique request id.
func UpdateFeeBatchRequest(token, locationID, feeID string, reqObj *UpdateFeeReqObject) (*BatchRequest, string) {
	v := new(Fee)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/fees/%s", locationID, feeID), token, reqObj, v)
}

// DeleteFeeBatchRequest returns a BatchRequest object for DeleteFee,
// along with a unique request id.
func DeleteFeeBatchRequest(token, locationID, feeID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/fees/%s", locationID, feeID), token, nil, nil)
}

// ApplyFeeBatchRequest returns a BatchRequest object for ApplyFee,
// along with a unique request id.
func ApplyFeeBatchRequest(token, locationID, itemID, feeID string) (*BatchRequest, string) {
	v := new(Item)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/items/%s/fees/%s", locationID, itemID, feeID), token, nil, v)
}

// RemoveFeeBatchRequest returns a BatchRequest object for RemoveFee,
// along with a unique request id.
func RemoveFeeBatchRequest(token, locationID, itemID, feeID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/items/%s/fees/%s", locationID, itemID, feeID), token, nil, nil)
}

// CreatePageBatchRequest returns a BatchRequest object for CreatePage,
// along with a unique request id.
func CreatePageBatchRequest(token, locationID string, reqObj *CreatePageReqObject) (*BatchRequest, string) {
	v := new(Page)
	return newBatchRequest("POST", fmt.Sprintf("/v1/%s/pages", locationID), token, reqObj, v)
}

// ListPagesBatchRequest returns a BatchRequest object for ListPages,
// along with a unique request id.
func ListPagesBatchRequest(token, locationID string) (*BatchRequest, string) {
	v := make([]*Page, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/pages", locationID), token, nil, &v)
}

// UpdatePageBatchRequest returns a BatchRequest object for UpdatePage,
// along with a unique request id.
func UpdatePageBatchRequest(token, locationID, pageID string, reqObj *UpdatePageReqObject) (*BatchRequest, string) {
	v := new(Page)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/pages/%s", locationID, pageID), token, reqObj, v)
}

// DeletePageBatchRequest returns a BatchRequest object for DeletePage,
// along with a unique request id.
func DeletePageBatchRequest(token, locationID, pageID string) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/pages/%s", locationID, pageID), token, nil, nil)
}

// UpdateCellBatchRequest returns a BatchRequest object for UpdateCell,
// along with a unique request id.
func UpdateCellBatchRequest(token, locationID, pageID string, reqObj *UpdateCellReqObject) (*BatchRequest, string) {
	v := new(PageCell)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/pages/%s/cells", locationID, pageID), token, reqObj, v)
}

// DeleteCellBatchRequest returns a BatchRequest object for DeleteCell,
// along with a unique request id.
func DeleteCellBatchRequest(token, locationID, pageID string, row, column int) (*BatchRequest, string) {
	return newBatchRequest("DELETE", fmt.Sprintf("/v1/%s/pages/%s/cells?row=%d&column=%d", locationID, pageID, row, column), token, nil, nil)
}

// ListWebhooksBatchRequest returns a BatchRequest object for ListWebhooks,
// along with a unique request id.
func ListWebhooksBatchRequest(token, locationID string) (*BatchRequest, string) {
	v := make([]string, 0)
	return newBatchRequest("GET", fmt.Sprintf("/v1/%s/webhooks", locationID), token, nil, &v)
}

// UpdateWebhooksBatchRequest returns a BatchRequest object for UpdateWebhooks,
// along with a unique request id.
func UpdateWebhooksBatchRequest(token, locationID string) (*BatchRequest, string) {
	v := make([]string, 0)
	return newBatchRequest("PUT", fmt.Sprintf("/v1/%s/webhooks", locationID), token, nil, &v)
}

// ListSubscriptionsBatchRequest returns a BatchRequest object for ListSubscriptions,
// along with a unique request id.
func ListSubscriptionsBatchRequest(token, clientID, merchantID string, limit int) (*BatchRequest, string) {
	v := make([]*Subscription, 0)
	return newBatchRequest("GET", fmt.Sprintf("/oauth2/clients/%s/subscriptions?merchant_id=%s&limit=%d", clientID, merchantID, limit), token, nil, &v)
}

// RetrieveSubscriptionBatchRequest returns a BatchRequest object for RetrieveSubscription,
// along with a unique request id.
func RetrieveSubscriptionBatchRequest(token, clientID, subscriptionID string) (*BatchRequest, string) {
	v := new(Subscription)
	return newBatchRequest("GET", fmt.Sprintf("/oauth2/clients/%s/subscriptions/%s", clientID, subscriptionID), token, nil, v)
}

// ListSubscriptionPlansBatchRequest returns a BatchRequest object for ListSubscriptionPlans,
// along with a unique request id.
func ListSubscriptionPlansBatchRequest(token, clientID string) (*BatchRequest, string) {
	v := make([]*SubscriptionPlan, 0)
	return newBatchRequest("GET", fmt.Sprintf("/oauth2/clients/%s/plans", clientID), token, nil, &v)
}

// RetrieveSubscriptionPlanBatchRequest returns a BatchRequest object for RetrieveSubscriptionPlan,
// along with a unique request id.
func RetrieveSubscriptionPlanBatchRequest(token, clientID, planID string) (*BatchRequest, string) {
	v := new(SubscriptionPlan)
	return newBatchRequest("GET", fmt.Sprintf("/oauth2/clients/%s/plans/%s", clientID, planID), token, nil, v)
}
