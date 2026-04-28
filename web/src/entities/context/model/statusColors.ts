/**
 * Maps ERP order status strings to MUI Chip color variants.
 * Status values correspond to `OrderStatus` in erp-service/internal/domain/order.go.
 */
export const ORDER_STATUS_COLOR: Record<
  string,
  "default" | "warning" | "info" | "success"
> = {
  not_initialized: "default",
  registered: "default",
  printed: "default",
  started: "info",
  finished: "success",
  post_calculated: "success",
  delivered: "success",
  historical: "default",
};

/**
 * Maps ERP operation status strings to MUI Chip color variants.
 * Status values correspond to `OperationStatus` in erp-service/internal/domain/order_operation.go.
 */
export const OPERATION_STATUS_COLOR: Record<
  string,
  "default" | "warning" | "info" | "success"
> = {
  none: "default",
  started: "info",
  partially_shipped: "warning",
  fully_shipped: "success",
  partially_reported: "warning",
  finished: "success",
};
