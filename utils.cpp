#ifdef CUPTI_ENABLED

#include "utils.hpp"
#include <iostream>
#include <vector>

#include <cupti_target.h>
#include <cupti_profiler_target.h>
#include <cupti_callbacks.h>
#include <cupti_driver_cbid.h>
#include <nvperf_host.h>
#include <cuda.h>
#include <cuda_runtime.h>
#include <string>
#include <stdio.h>
#include <stdlib.h>
#include "Metric.h"
#include "Eval.h"

#define NVPW_API_CALL(apiFuncCall)                                         \
do {                                                                       \
  NVPA_Status _status = apiFuncCall;                                       \
  if (_status != NVPA_STATUS_SUCCESS) {                                    \
    fprintf(stderr, "%s:%d: error: function %s failed with error %d.\n",   \
            __FILE__, __LINE__, #apiFuncCall, _status);                    \
    exit(-1);                                                              \
  }                                                                        \
} while (0)

#define CUPTI_API_CALL(apiFuncCall)                                        \
do {                                                                       \
  CUptiResult _status = apiFuncCall;                                       \
  if (_status != CUPTI_SUCCESS) {                                          \
    const char *errstr;                                                    \
    cuptiGetResultString(_status, &errstr);                                \
    fprintf(stderr, "%s:%d: error: function %s failed with error %s.\n",   \
            __FILE__, __LINE__, #apiFuncCall, errstr);                     \
    exit(-1);                                                              \
  }                                                                        \
} while (0)

#define DRIVER_API_CALL(apiFuncCall)                                       \
do {                                                                       \
  CUresult _status = apiFuncCall;                                          \
  if (_status != CUDA_SUCCESS) {                                           \
    fprintf(stderr, "%s:%d: error: function %s failed with error %d.\n",   \
            __FILE__, __LINE__, #apiFuncCall, _status);                    \
    exit(-1);                                                              \
  }                                                                        \
} while (0)

#define RUNTIME_API_CALL(apiFuncCall)                                      \
do {                                                                       \
  cudaError_t _status = apiFuncCall;                                       \
  if (_status != cudaSuccess) {                                            \
    fprintf(stderr, "%s:%d: error: function %s failed with error %s.\n",   \
            __FILE__, __LINE__, #apiFuncCall, cudaGetErrorString(_status));\
    exit(-1);                                                              \
  }                                                                        \
} while (0)

struct ProfilingData_t {
  int numRanges = 100000;
  std::string chipName;
  std::vector<std::string> metricNames;
  CUpti_ProfilerRange profilerRange = CUPTI_AutoRange;
  CUpti_ProfilerReplayMode profilerReplayMode = CUPTI_KernelReplay;
  std::vector<uint8_t> counterDataImagePrefix;
  std::vector<uint8_t> configImage;
  std::vector<uint8_t> counterDataImage;
  std::vector<uint8_t> counterDataScratchBuffer;
  std::vector<double> metricData;
  int numMetrics = 0;
}cur;

void enableProfiling() {
  CUpti_Profiler_EnableProfiling_Params enableProfilingParams = { CUpti_Profiler_EnableProfiling_Params_STRUCT_SIZE };
  CUPTI_API_CALL(cuptiProfilerEnableProfiling(&enableProfilingParams));
}

void disableProfiling() {
  CUpti_Profiler_DisableProfiling_Params disableProfilingParams = { CUpti_Profiler_DisableProfiling_Params_STRUCT_SIZE };
  CUPTI_API_CALL(cuptiProfilerDisableProfiling(&disableProfilingParams));
}

void beginSession() {
  CUpti_Profiler_BeginSession_Params beginSessionParams = { CUpti_Profiler_BeginSession_Params_STRUCT_SIZE };
  beginSessionParams.ctx = NULL;
  beginSessionParams.counterDataImageSize = cur.counterDataImage.size();
  beginSessionParams.pCounterDataImage = &cur.counterDataImage[0];
  beginSessionParams.counterDataScratchBufferSize = cur.counterDataScratchBuffer.size();
  beginSessionParams.pCounterDataScratchBuffer = &cur.counterDataScratchBuffer[0];
  beginSessionParams.range = cur.profilerRange;
  beginSessionParams.replayMode = cur.profilerReplayMode;
  beginSessionParams.maxRangesPerPass = cur.numRanges;
  beginSessionParams.maxLaunchesPerPass = cur.numRanges;
  CUPTI_API_CALL(cuptiProfilerBeginSession(&beginSessionParams));
}

void setConfig() {
  CUpti_Profiler_SetConfig_Params setConfigParams = { CUpti_Profiler_SetConfig_Params_STRUCT_SIZE };
  setConfigParams.pConfig = &cur.configImage[0];
  setConfigParams.configSize = cur.configImage.size();
  setConfigParams.passIndex = 0;
  CUPTI_API_CALL(cuptiProfilerSetConfig(&setConfigParams));
}

void createCounterDataImage(int numRanges,
  std::vector<uint8_t>& counterDataImagePrefix,
  std::vector<uint8_t>& counterDataScratchBuffer,
  std::vector<uint8_t>& counterDataImage
) {
  CUpti_Profiler_CounterDataImageOptions counterDataImageOptions;
  counterDataImageOptions.pCounterDataPrefix = &counterDataImagePrefix[0];
  counterDataImageOptions.counterDataPrefixSize = counterDataImagePrefix.size();
  counterDataImageOptions.maxNumRanges = numRanges;
  counterDataImageOptions.maxNumRangeTreeNodes = numRanges;
  counterDataImageOptions.maxRangeNameLength = 64;

  CUpti_Profiler_CounterDataImage_CalculateSize_Params calculateSizeParams = { CUpti_Profiler_CounterDataImage_CalculateSize_Params_STRUCT_SIZE };
  calculateSizeParams.pOptions = &counterDataImageOptions;
  calculateSizeParams.sizeofCounterDataImageOptions = CUpti_Profiler_CounterDataImageOptions_STRUCT_SIZE;
  CUPTI_API_CALL(cuptiProfilerCounterDataImageCalculateSize(&calculateSizeParams));

  CUpti_Profiler_CounterDataImage_Initialize_Params initializeParams = { CUpti_Profiler_CounterDataImage_Initialize_Params_STRUCT_SIZE };
  initializeParams.sizeofCounterDataImageOptions = CUpti_Profiler_CounterDataImageOptions_STRUCT_SIZE;
  initializeParams.pOptions = &counterDataImageOptions;
  initializeParams.counterDataImageSize = calculateSizeParams.counterDataImageSize;
  counterDataImage.resize(calculateSizeParams.counterDataImageSize);
  initializeParams.pCounterDataImage = &counterDataImage[0];
  CUPTI_API_CALL(cuptiProfilerCounterDataImageInitialize(&initializeParams));

  CUpti_Profiler_CounterDataImage_CalculateScratchBufferSize_Params scratchBufferSizeParams = { CUpti_Profiler_CounterDataImage_CalculateScratchBufferSize_Params_STRUCT_SIZE };
  scratchBufferSizeParams.counterDataImageSize = calculateSizeParams.counterDataImageSize;
  scratchBufferSizeParams.pCounterDataImage = initializeParams.pCounterDataImage;
  CUPTI_API_CALL(cuptiProfilerCounterDataImageCalculateScratchBufferSize(&scratchBufferSizeParams));
  counterDataScratchBuffer.resize(scratchBufferSizeParams.counterDataScratchBufferSize);

  CUpti_Profiler_CounterDataImage_InitializeScratchBuffer_Params initScratchBufferParams = { CUpti_Profiler_CounterDataImage_InitializeScratchBuffer_Params_STRUCT_SIZE };
  initScratchBufferParams.counterDataImageSize = calculateSizeParams.counterDataImageSize;
  initScratchBufferParams.pCounterDataImage = initializeParams.pCounterDataImage;
  initScratchBufferParams.counterDataScratchBufferSize = scratchBufferSizeParams.counterDataScratchBufferSize;
  initScratchBufferParams.pCounterDataScratchBuffer = &counterDataScratchBuffer[0];
  CUPTI_API_CALL(cuptiProfilerCounterDataImageInitializeScratchBuffer(&initScratchBufferParams));
}

void setupProfiling() {
  NVPW_InitializeHost_Params initializeHostParams = { NVPW_InitializeHost_Params_STRUCT_SIZE };
  NVPW_API_CALL(NVPW_InitializeHost(&initializeHostParams));

  if (cur.metricNames.size()) {
      if (!NV::Metric::Config::GetConfigImage(cur.chipName, cur.metricNames, cur.configImage))
      {
          std::cout << "Failed to create configImage" << std::endl;
          exit(-1);
      }
      if (!NV::Metric::Config::GetCounterDataPrefixImage(cur.chipName, cur.metricNames, cur.counterDataImagePrefix))
      {
          std::cout << "Failed to create counterDataImagePrefix" << std::endl;
          exit(-1);
      }
  }

  createCounterDataImage(cur.numRanges, cur.counterDataImagePrefix,
                          cur.counterDataScratchBuffer, cur.counterDataImage);

  beginSession();
  setConfig();
}

void stopProfiling() {
  CUpti_Profiler_UnsetConfig_Params unsetConfigParams = { CUpti_Profiler_UnsetConfig_Params_STRUCT_SIZE };
  CUPTI_API_CALL(cuptiProfilerUnsetConfig(&unsetConfigParams));

  // CUpti_Profiler_DeInitialize_Params profilerDeInitializeParams = {CUpti_Profiler_DeInitialize_Params_STRUCT_SIZE};

  CUpti_Profiler_EndSession_Params endSessionParams = { CUpti_Profiler_EndSession_Params_STRUCT_SIZE };
  CUPTI_API_CALL(cuptiProfilerEndSession(&endSessionParams));
  // cuptiProfilerDeInitialize(&profilerDeInitializeParams);
}

void onCallback(int start) {
  if(start) {
  }
  else {
  }
}

void startProfiling(char *goMetrics) {
  if (goMetrics == NULL) {
    return;
  }

  cur = ProfilingData_t();

  char* metricName = strtok(goMetrics, ",");
  while (metricName != NULL) {
    cur.metricNames.push_back(metricName);
    cur.numMetrics++;
    metricName = strtok(NULL, ",");
  }

  CUpti_Profiler_Initialize_Params profilerInitializeParams = { CUpti_Profiler_Initialize_Params_STRUCT_SIZE };
  CUPTI_API_CALL(cuptiProfilerInitialize(&profilerInitializeParams));

  CUpti_Device_GetChipName_Params getChipNameParams = { CUpti_Device_GetChipName_Params_STRUCT_SIZE };
  getChipNameParams.deviceIndex = 0;

  CUPTI_API_CALL(cuptiDeviceGetChipName(&getChipNameParams));
  cur.chipName = getChipNameParams.pChipName;

  setupProfiling();
  enableProfiling();
}

double* endProfiling(uint64_t* len) {
  if(cur.numMetrics == 0) {
    *len = 0;
    return NULL;
  }
  disableProfiling();
  stopProfiling();
  cur.metricData = NV::Metric::Eval::GetMetricValues(cur.chipName, cur.counterDataImage, cur.metricNames);
  *len = cur.metricData.size();
  return cur.metricData.data();
}

bool getProfilingAPI(void) {
  int ccMajor = 0;
  CUdevice cuDevice;
  DRIVER_API_CALL(cuDeviceGet(&cuDevice, 0));
  DRIVER_API_CALL(cuDeviceGetAttribute(&ccMajor, CU_DEVICE_ATTRIBUTE_COMPUTE_CAPABILITY_MAJOR, cuDevice));
  return ccMajor >= 7;
}

#endif // CUPTI_ENABLED
