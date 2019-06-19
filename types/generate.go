//go:generate go get github.com/alvaroloes/enumer
//go:generate enumer -type=CUpti_ActivityFlag -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityKind -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityObjectKind -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityOverheadKind -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityComputeApiKind -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityPCSamplingStallReason -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityPCSamplingPeriod -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityMemcpyKind -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityMemoryKind -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityPreemptionKind -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityEnvironmentKind -json -text -yaml -sql
//go:generate enumer -type=CUpti_EnvironmentClocksThrottleReason -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityUnifiedMemoryCounterScope -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityUnifiedMemoryCounterKind -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityUnifiedMemoryAccessType -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityUnifiedMemoryMigrationCause -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityUnifiedMemoryRemoteMapCause -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityInstructionClass -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityPartitionedGlobalCacheConfig -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivitySynchronizationType -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityStreamFlag -json -text -yaml -sql
//go:generate enumer -type=CUpti_LinkFlag -json -text -yaml -sql
//go:generate enumer -type=CUpti_DevType -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityAttribute -json -text -yaml -sql
//go:generate enumer -type=CUpti_ActivityThreadIdType -json -text -yaml -sql
//go:generate enumer -type=CUPTI_RUNTIME_TRACE_CBID -json -text -yaml -sql
//go:generate enumer -type=CUdevice_attribute -json -text -yaml -sql
//go:generate enumer -type=CUresult -json -text -yaml -sql
//go:generate enumer -type=CUDAMemcpyKind -json -text -yaml -sql
//go:generate enumer -type=CUDAError -json -text -yaml -sql
//go:generate enumer -type=CUpti_ApiCallbackSite -json -text -yaml -sql
//go:generate enumer -type=CUpti_CallbackDomain -json -text -yaml -sql
//go:generate enumer -type=CUpti_CallbackIdResource -json -text -yaml -sql
//go:generate enumer -type=CUpti_CallbackIdSync -json -text -yaml -sql
//go:generate enumer -type=CUptiResult -json -text -yaml -sql
//go:generate enumer -type=CUpti_nvtx_api_trace_cbid -json -text -yaml -sql
//go:generate enumer -type=CUpti_MetricCategory -json -text -yaml -sql
//go:generate enumer -type=CUpti_MetricEvaluationMode -json -text -yaml -sql
//go:generate enumer -type=CUpti_MetricValueKind -json -text -yaml -sql
//go:generate enumer -type=CUpti_MetricValueUtilizationLevel -json -text -yaml -sql
//go:generate enumer -type=CUpti_MetricAttribute -json -text -yaml -sql
//go:generate enumer -type=CUpti_MetricPropertyDeviceClass -json -text -yaml -sql
//go:generate enumer -type=CUpti_MetricPropertyID -json -text -yaml -sql
//go:generate enumer -type=CUpti_driver_api_trace_cbid -json -text -yaml -sql
//go:generate enumer -type=CUpti_DeviceAttributeDeviceClass -json -text -yaml -sql
//go:generate enumer -type=CUpti_DeviceAttribute -json -text -yaml -sql
//go:generate enumer -type=CUpti_EventDomainAttribute -json -text -yaml -sql
//go:generate enumer -type=CUpti_EventCollectionMethod -json -text -yaml -sql
//go:generate enumer -type=CUpti_EventGroupAttribute -json -text -yaml -sql
//go:generate enumer -type=CUpti_EventAttribute -json -text -yaml -sql
//go:generate enumer -type=CUpti_EventCollectionMode -json -text -yaml -sql
//go:generate enumer -type=CUpti_EventCategory -json -text -yaml -sql
//go:generate enumer -type=CUpti_ReadEventFlags -json -text -yaml -sql

package types
