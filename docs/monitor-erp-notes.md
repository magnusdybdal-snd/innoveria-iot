# Monitor ERP

### Manufacturing orders
what is happening right now
- https://api.monitor.se/api/Monitor.API.Manufacturing.ManufacturingOrderOperationReporting.html

The most useful monitor erp data fields:
- What order is running? `ManufacturingOrderNumber`
- What operation? `OperationNumber`
- Which machine? `WorkCenterId`
- When? `ReportingTimestamp`
- What happened? `Type` + `NodeStatus`


### Machines
- https://api.monitor.se/api/Monitor.API.Manufacturing.WorkCenter.html

- returns what a machine is.
- used for mapping id to name
