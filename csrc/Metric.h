#pragma once

#include <Parser.h>
#include <Utils.h>
#include <nvperf_host.h>
#include <nvperf_cuda_host.h>
#include <iostream>
#include <ScopeExit.h>

namespace NV {
    namespace Metric {
        namespace Config {

            bool GetRawMetricRequests(NVPA_MetricsContext* pMetricsContext,
                                      std::vector<std::string> metricNames,
                                      std::vector<NVPA_RawMetricRequest>& rawMetricRequests,
                                      std::vector<std::string>& temp) {
                std::string reqName;
                bool isolated = true;
                bool keepInstances = true;

                for (auto& metricName : metricNames)
                {
                    NV::Metric::Parser::ParseMetricNameString(metricName, &reqName, &isolated, &keepInstances);
                    /* Bug in collection with collection of metrics without instances, keep it to true*/
                    keepInstances = true;
                    NVPW_MetricsContext_GetMetricProperties_Begin_Params getMetricPropertiesBeginParams = { NVPW_MetricsContext_GetMetricProperties_Begin_Params_STRUCT_SIZE };
                    getMetricPropertiesBeginParams.pMetricsContext = pMetricsContext;
                    getMetricPropertiesBeginParams.pMetricName = reqName.c_str();

                    RETURN_IF_NVPW_ERROR(false, NVPW_MetricsContext_GetMetricProperties_Begin(&getMetricPropertiesBeginParams));

                    for (const char** ppMetricDependencies = getMetricPropertiesBeginParams.ppRawMetricDependencies; *ppMetricDependencies; ++ppMetricDependencies)
                    {
                        temp.push_back(*ppMetricDependencies);
                    }
                    NVPW_MetricsContext_GetMetricProperties_End_Params getMetricPropertiesEndParams = { NVPW_MetricsContext_GetMetricProperties_End_Params_STRUCT_SIZE };
                    getMetricPropertiesEndParams.pMetricsContext = pMetricsContext;
                    RETURN_IF_NVPW_ERROR(false, NVPW_MetricsContext_GetMetricProperties_End(&getMetricPropertiesEndParams));
                }

                for (auto& rawMetricName : temp)
                {
                    NVPA_RawMetricRequest metricRequest = { NVPA_RAW_METRIC_REQUEST_STRUCT_SIZE };
                    metricRequest.pMetricName = rawMetricName.c_str();
                    metricRequest.isolated = isolated;
                    metricRequest.keepInstances = keepInstances;
                    rawMetricRequests.push_back(metricRequest);
                }

                return true;
            }

            bool GetConfigImage(std::string chipName, std::vector<std::string> metricNames, std::vector<uint8_t>& configImage, const uint8_t* pCounterAvailabilityImage = NULL) 
            {
                NVPW_CUDA_MetricsContext_Create_Params metricsContextCreateParams = { NVPW_CUDA_MetricsContext_Create_Params_STRUCT_SIZE };
                metricsContextCreateParams.pChipName = chipName.c_str();
                RETURN_IF_NVPW_ERROR(false, NVPW_CUDA_MetricsContext_Create(&metricsContextCreateParams));

                NVPW_MetricsContext_Destroy_Params metricsContextDestroyParams = { NVPW_MetricsContext_Destroy_Params_STRUCT_SIZE };
                metricsContextDestroyParams.pMetricsContext = metricsContextCreateParams.pMetricsContext;
                SCOPE_EXIT([&]() { NVPW_MetricsContext_Destroy((NVPW_MetricsContext_Destroy_Params *)&metricsContextDestroyParams); });
                
                std::vector<NVPA_RawMetricRequest> rawMetricRequests;
                std::vector<std::string> temp;
                GetRawMetricRequests(metricsContextCreateParams.pMetricsContext, metricNames, rawMetricRequests, temp);

                NVPW_CUDA_RawMetricsConfig_Create_Params nvpw_metricsConfigCreateParams {};
                nvpw_metricsConfigCreateParams.structSize = NVPW_CUDA_RawMetricsConfig_Create_Params_STRUCT_SIZE;
                nvpw_metricsConfigCreateParams.activityKind = NVPA_ACTIVITY_KIND_PROFILER;
                nvpw_metricsConfigCreateParams.pChipName = chipName.c_str();
                RETURN_IF_NVPW_ERROR(false, NVPW_CUDA_RawMetricsConfig_Create(&nvpw_metricsConfigCreateParams));

                if(pCounterAvailabilityImage)
                {
                    NVPW_RawMetricsConfig_SetCounterAvailability_Params setCounterAvailabilityParams = {NVPW_RawMetricsConfig_SetCounterAvailability_Params_STRUCT_SIZE};
                    setCounterAvailabilityParams.pRawMetricsConfig = nvpw_metricsConfigCreateParams.pRawMetricsConfig;
                    setCounterAvailabilityParams.pCounterAvailabilityImage = pCounterAvailabilityImage;
                    RETURN_IF_NVPW_ERROR(false, NVPW_RawMetricsConfig_SetCounterAvailability(&setCounterAvailabilityParams));
                }

                NVPW_RawMetricsConfig_Destroy_Params rawMetricsConfigDestroyParams = { NVPW_RawMetricsConfig_Destroy_Params_STRUCT_SIZE };
                rawMetricsConfigDestroyParams.pRawMetricsConfig = nvpw_metricsConfigCreateParams.pRawMetricsConfig;
                SCOPE_EXIT([&]() { NVPW_RawMetricsConfig_Destroy((NVPW_RawMetricsConfig_Destroy_Params *)&rawMetricsConfigDestroyParams); });

                NVPW_RawMetricsConfig_BeginPassGroup_Params beginPassGroupParams = { NVPW_RawMetricsConfig_BeginPassGroup_Params_STRUCT_SIZE };
                beginPassGroupParams.pRawMetricsConfig = nvpw_metricsConfigCreateParams.pRawMetricsConfig;
                RETURN_IF_NVPW_ERROR(false, NVPW_RawMetricsConfig_BeginPassGroup(&beginPassGroupParams));

                NVPW_RawMetricsConfig_AddMetrics_Params addMetricsParams = { NVPW_RawMetricsConfig_AddMetrics_Params_STRUCT_SIZE };
                addMetricsParams.pRawMetricsConfig = nvpw_metricsConfigCreateParams.pRawMetricsConfig;
                addMetricsParams.pRawMetricRequests = &rawMetricRequests[0];
                addMetricsParams.numMetricRequests = rawMetricRequests.size();
                RETURN_IF_NVPW_ERROR(false, NVPW_RawMetricsConfig_AddMetrics(&addMetricsParams));

                NVPW_RawMetricsConfig_EndPassGroup_Params endPassGroupParams = { NVPW_RawMetricsConfig_EndPassGroup_Params_STRUCT_SIZE };
                endPassGroupParams.pRawMetricsConfig = nvpw_metricsConfigCreateParams.pRawMetricsConfig;
                RETURN_IF_NVPW_ERROR(false, NVPW_RawMetricsConfig_EndPassGroup(&endPassGroupParams));

                NVPW_RawMetricsConfig_GenerateConfigImage_Params generateConfigImageParams = { NVPW_RawMetricsConfig_GenerateConfigImage_Params_STRUCT_SIZE };
                generateConfigImageParams.pRawMetricsConfig = nvpw_metricsConfigCreateParams.pRawMetricsConfig;
                RETURN_IF_NVPW_ERROR(false, NVPW_RawMetricsConfig_GenerateConfigImage(&generateConfigImageParams));

                NVPW_RawMetricsConfig_GetConfigImage_Params getConfigImageParams = { NVPW_RawMetricsConfig_GetConfigImage_Params_STRUCT_SIZE };
                getConfigImageParams.pRawMetricsConfig = nvpw_metricsConfigCreateParams.pRawMetricsConfig;
                getConfigImageParams.bytesAllocated = 0;
                getConfigImageParams.pBuffer = NULL;
                RETURN_IF_NVPW_ERROR(false, NVPW_RawMetricsConfig_GetConfigImage(&getConfigImageParams));

                configImage.resize(getConfigImageParams.bytesCopied);

                getConfigImageParams.bytesAllocated = configImage.size();
                getConfigImageParams.pBuffer = &configImage[0];
                RETURN_IF_NVPW_ERROR(false, NVPW_RawMetricsConfig_GetConfigImage(&getConfigImageParams));

                return true;
            }

            bool GetCounterDataPrefixImage(std::string chipName, std::vector<std::string> metricNames, std::vector<uint8_t>& counterDataImagePrefix) 
            {
                NVPW_CUDA_MetricsContext_Create_Params metricsContextCreateParams = { NVPW_CUDA_MetricsContext_Create_Params_STRUCT_SIZE };
                metricsContextCreateParams.pChipName = chipName.c_str();
                RETURN_IF_NVPW_ERROR(false, NVPW_CUDA_MetricsContext_Create(&metricsContextCreateParams));

                NVPW_MetricsContext_Destroy_Params metricsContextDestroyParams = { NVPW_MetricsContext_Destroy_Params_STRUCT_SIZE };
                metricsContextDestroyParams.pMetricsContext = metricsContextCreateParams.pMetricsContext;
                SCOPE_EXIT([&]() { NVPW_MetricsContext_Destroy((NVPW_MetricsContext_Destroy_Params *)&metricsContextDestroyParams); });

                std::vector<NVPA_RawMetricRequest> rawMetricRequests;
                std::vector<std::string> temp;
                GetRawMetricRequests(metricsContextCreateParams.pMetricsContext, metricNames, rawMetricRequests, temp);

                NVPW_CounterDataBuilder_Create_Params counterDataBuilderCreateParams = { NVPW_CounterDataBuilder_Create_Params_STRUCT_SIZE };
                counterDataBuilderCreateParams.pChipName = chipName.c_str();
                RETURN_IF_NVPW_ERROR(false, NVPW_CounterDataBuilder_Create(&counterDataBuilderCreateParams));

                NVPW_CounterDataBuilder_Destroy_Params counterDataBuilderDestroyParams = { NVPW_CounterDataBuilder_Destroy_Params_STRUCT_SIZE };
                counterDataBuilderDestroyParams.pCounterDataBuilder = counterDataBuilderCreateParams.pCounterDataBuilder;
                SCOPE_EXIT([&]() { NVPW_CounterDataBuilder_Destroy((NVPW_CounterDataBuilder_Destroy_Params *)&counterDataBuilderDestroyParams); });

                NVPW_CounterDataBuilder_AddMetrics_Params addMetricsParams = { NVPW_CounterDataBuilder_AddMetrics_Params_STRUCT_SIZE };
                addMetricsParams.pCounterDataBuilder = counterDataBuilderCreateParams.pCounterDataBuilder;
                addMetricsParams.pRawMetricRequests = &rawMetricRequests[0];
                addMetricsParams.numMetricRequests = rawMetricRequests.size();
                RETURN_IF_NVPW_ERROR(false, NVPW_CounterDataBuilder_AddMetrics(&addMetricsParams));

                size_t counterDataPrefixSize = 0;
                NVPW_CounterDataBuilder_GetCounterDataPrefix_Params getCounterDataPrefixParams = { NVPW_CounterDataBuilder_GetCounterDataPrefix_Params_STRUCT_SIZE };
                getCounterDataPrefixParams.pCounterDataBuilder = counterDataBuilderCreateParams.pCounterDataBuilder;
                getCounterDataPrefixParams.bytesAllocated = 0;
                getCounterDataPrefixParams.pBuffer = NULL;
                RETURN_IF_NVPW_ERROR(false, NVPW_CounterDataBuilder_GetCounterDataPrefix(&getCounterDataPrefixParams));

                counterDataImagePrefix.resize(getCounterDataPrefixParams.bytesCopied);

                getCounterDataPrefixParams.bytesAllocated = counterDataImagePrefix.size();
                getCounterDataPrefixParams.pBuffer = &counterDataImagePrefix[0];
                RETURN_IF_NVPW_ERROR(false, NVPW_CounterDataBuilder_GetCounterDataPrefix(&getCounterDataPrefixParams));

                return true;
            }
        }
    }
}
