// Code generated by "enumer -type=CUPTI_RUNTIME_TRACE_CBID -json -text -yaml -sql"; DO NOT EDIT.

//
package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const _CUPTI_RUNTIME_TRACE_CBIDName = "CUPTI_RUNTIME_TRACE_CBID_INVALIDCUPTI_RUNTIME_TRACE_CBID_cudaDriverGetVersion_v3020CUPTI_RUNTIME_TRACE_CBID_cudaRuntimeGetVersion_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGetDeviceCount_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGetDeviceProperties_v3020CUPTI_RUNTIME_TRACE_CBID_cudaChooseDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGetChannelDesc_v3020CUPTI_RUNTIME_TRACE_CBID_cudaCreateChannelDesc_v3020CUPTI_RUNTIME_TRACE_CBID_cudaConfigureCall_v3020CUPTI_RUNTIME_TRACE_CBID_cudaSetupArgument_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGetLastError_v3020CUPTI_RUNTIME_TRACE_CBID_cudaPeekAtLastError_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGetErrorString_v3020CUPTI_RUNTIME_TRACE_CBID_cudaLaunch_v3020CUPTI_RUNTIME_TRACE_CBID_cudaFuncSetCacheConfig_v3020CUPTI_RUNTIME_TRACE_CBID_cudaFuncGetAttributes_v3020CUPTI_RUNTIME_TRACE_CBID_cudaSetDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGetDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaSetValidDevices_v3020CUPTI_RUNTIME_TRACE_CBID_cudaSetDeviceFlags_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMalloc_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMallocPitch_v3020CUPTI_RUNTIME_TRACE_CBID_cudaFree_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMallocArray_v3020CUPTI_RUNTIME_TRACE_CBID_cudaFreeArray_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMallocHost_v3020CUPTI_RUNTIME_TRACE_CBID_cudaFreeHost_v3020CUPTI_RUNTIME_TRACE_CBID_cudaHostAlloc_v3020CUPTI_RUNTIME_TRACE_CBID_cudaHostGetDevicePointer_v3020CUPTI_RUNTIME_TRACE_CBID_cudaHostGetFlags_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemGetInfo_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy2D_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyToArray_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy2DToArray_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyFromArray_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy2DFromArray_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyArrayToArray_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy2DArrayToArray_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyToSymbol_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyFromSymbol_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyAsync_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyToArrayAsync_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyFromArrayAsync_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy2DAsync_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy2DToArrayAsync_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy2DFromArrayAsync_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyToSymbolAsync_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyFromSymbolAsync_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemset_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemset2D_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemsetAsync_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemset2DAsync_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGetSymbolAddress_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGetSymbolSize_v3020CUPTI_RUNTIME_TRACE_CBID_cudaBindTexture_v3020CUPTI_RUNTIME_TRACE_CBID_cudaBindTexture2D_v3020CUPTI_RUNTIME_TRACE_CBID_cudaBindTextureToArray_v3020CUPTI_RUNTIME_TRACE_CBID_cudaUnbindTexture_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGetTextureAlignmentOffset_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGetTextureReference_v3020CUPTI_RUNTIME_TRACE_CBID_cudaBindSurfaceToArray_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGetSurfaceReference_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGLSetGLDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGLRegisterBufferObject_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGLMapBufferObject_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGLUnmapBufferObject_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGLUnregisterBufferObject_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGLSetBufferObjectMapFlags_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGLMapBufferObjectAsync_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGLUnmapBufferObjectAsync_v3020CUPTI_RUNTIME_TRACE_CBID_cudaWGLGetDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsGLRegisterImage_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsGLRegisterBuffer_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsUnregisterResource_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsResourceSetMapFlags_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsMapResources_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsUnmapResources_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsResourceGetMappedPointer_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsSubResourceGetMappedArray_v3020CUPTI_RUNTIME_TRACE_CBID_cudaVDPAUGetDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaVDPAUSetVDPAUDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsVDPAURegisterVideoSurface_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsVDPAURegisterOutputSurface_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D11GetDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D11GetDevices_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D11SetDirect3DDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsD3D11RegisterResource_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D10GetDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D10GetDevices_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D10SetDirect3DDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsD3D10RegisterResource_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D10RegisterResource_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D10UnregisterResource_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D10MapResources_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D10UnmapResources_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D10ResourceSetMapFlags_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D10ResourceGetSurfaceDimensions_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D10ResourceGetMappedArray_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D10ResourceGetMappedPointer_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D10ResourceGetMappedSize_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D10ResourceGetMappedPitch_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9GetDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9GetDevices_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9SetDirect3DDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9GetDirect3DDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsD3D9RegisterResource_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9RegisterResource_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9UnregisterResource_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9MapResources_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9UnmapResources_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9ResourceSetMapFlags_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9ResourceGetSurfaceDimensions_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9ResourceGetMappedArray_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9ResourceGetMappedPointer_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9ResourceGetMappedSize_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9ResourceGetMappedPitch_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9Begin_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9End_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9RegisterVertexBuffer_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9UnregisterVertexBuffer_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9MapVertexBuffer_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D9UnmapVertexBuffer_v3020CUPTI_RUNTIME_TRACE_CBID_cudaThreadExit_v3020CUPTI_RUNTIME_TRACE_CBID_cudaSetDoubleForDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaSetDoubleForHost_v3020CUPTI_RUNTIME_TRACE_CBID_cudaThreadSynchronize_v3020CUPTI_RUNTIME_TRACE_CBID_cudaThreadGetLimit_v3020CUPTI_RUNTIME_TRACE_CBID_cudaThreadSetLimit_v3020CUPTI_RUNTIME_TRACE_CBID_cudaStreamCreate_v3020CUPTI_RUNTIME_TRACE_CBID_cudaStreamDestroy_v3020CUPTI_RUNTIME_TRACE_CBID_cudaStreamSynchronize_v3020CUPTI_RUNTIME_TRACE_CBID_cudaStreamQuery_v3020CUPTI_RUNTIME_TRACE_CBID_cudaEventCreate_v3020CUPTI_RUNTIME_TRACE_CBID_cudaEventCreateWithFlags_v3020CUPTI_RUNTIME_TRACE_CBID_cudaEventRecord_v3020CUPTI_RUNTIME_TRACE_CBID_cudaEventDestroy_v3020CUPTI_RUNTIME_TRACE_CBID_cudaEventSynchronize_v3020CUPTI_RUNTIME_TRACE_CBID_cudaEventQuery_v3020CUPTI_RUNTIME_TRACE_CBID_cudaEventElapsedTime_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMalloc3D_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMalloc3DArray_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemset3D_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemset3DAsync_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy3D_v3020CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy3DAsync_v3020CUPTI_RUNTIME_TRACE_CBID_cudaThreadSetCacheConfig_v3020CUPTI_RUNTIME_TRACE_CBID_cudaStreamWaitEvent_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D11GetDirect3DDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaD3D10GetDirect3DDevice_v3020CUPTI_RUNTIME_TRACE_CBID_cudaThreadGetCacheConfig_v3020CUPTI_RUNTIME_TRACE_CBID_cudaPointerGetAttributes_v4000CUPTI_RUNTIME_TRACE_CBID_cudaHostRegister_v4000CUPTI_RUNTIME_TRACE_CBID_cudaHostUnregister_v4000CUPTI_RUNTIME_TRACE_CBID_cudaDeviceCanAccessPeer_v4000CUPTI_RUNTIME_TRACE_CBID_cudaDeviceEnablePeerAccess_v4000CUPTI_RUNTIME_TRACE_CBID_cudaDeviceDisablePeerAccess_v4000CUPTI_RUNTIME_TRACE_CBID_cudaPeerRegister_v4000CUPTI_RUNTIME_TRACE_CBID_cudaPeerUnregister_v4000CUPTI_RUNTIME_TRACE_CBID_cudaPeerGetDevicePointer_v4000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyPeer_v4000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyPeerAsync_v4000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy3DPeer_v4000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy3DPeerAsync_v4000CUPTI_RUNTIME_TRACE_CBID_cudaDeviceReset_v3020CUPTI_RUNTIME_TRACE_CBID_cudaDeviceSynchronize_v3020CUPTI_RUNTIME_TRACE_CBID_cudaDeviceGetLimit_v3020CUPTI_RUNTIME_TRACE_CBID_cudaDeviceSetLimit_v3020CUPTI_RUNTIME_TRACE_CBID_cudaDeviceGetCacheConfig_v3020CUPTI_RUNTIME_TRACE_CBID_cudaDeviceSetCacheConfig_v3020CUPTI_RUNTIME_TRACE_CBID_cudaProfilerInitialize_v4000CUPTI_RUNTIME_TRACE_CBID_cudaProfilerStart_v4000CUPTI_RUNTIME_TRACE_CBID_cudaProfilerStop_v4000CUPTI_RUNTIME_TRACE_CBID_cudaDeviceGetByPCIBusId_v4010CUPTI_RUNTIME_TRACE_CBID_cudaDeviceGetPCIBusId_v4010CUPTI_RUNTIME_TRACE_CBID_cudaGLGetDevices_v4010CUPTI_RUNTIME_TRACE_CBID_cudaIpcGetEventHandle_v4010CUPTI_RUNTIME_TRACE_CBID_cudaIpcOpenEventHandle_v4010CUPTI_RUNTIME_TRACE_CBID_cudaIpcGetMemHandle_v4010CUPTI_RUNTIME_TRACE_CBID_cudaIpcOpenMemHandle_v4010CUPTI_RUNTIME_TRACE_CBID_cudaIpcCloseMemHandle_v4010CUPTI_RUNTIME_TRACE_CBID_cudaArrayGetInfo_v4010CUPTI_RUNTIME_TRACE_CBID_cudaFuncSetSharedMemConfig_v4020CUPTI_RUNTIME_TRACE_CBID_cudaDeviceGetSharedMemConfig_v4020CUPTI_RUNTIME_TRACE_CBID_cudaDeviceSetSharedMemConfig_v4020CUPTI_RUNTIME_TRACE_CBID_cudaCreateTextureObject_v5000CUPTI_RUNTIME_TRACE_CBID_cudaDestroyTextureObject_v5000CUPTI_RUNTIME_TRACE_CBID_cudaGetTextureObjectResourceDesc_v5000CUPTI_RUNTIME_TRACE_CBID_cudaGetTextureObjectTextureDesc_v5000CUPTI_RUNTIME_TRACE_CBID_cudaCreateSurfaceObject_v5000CUPTI_RUNTIME_TRACE_CBID_cudaDestroySurfaceObject_v5000CUPTI_RUNTIME_TRACE_CBID_cudaGetSurfaceObjectResourceDesc_v5000CUPTI_RUNTIME_TRACE_CBID_cudaMallocMipmappedArray_v5000CUPTI_RUNTIME_TRACE_CBID_cudaGetMipmappedArrayLevel_v5000CUPTI_RUNTIME_TRACE_CBID_cudaFreeMipmappedArray_v5000CUPTI_RUNTIME_TRACE_CBID_cudaBindTextureToMipmappedArray_v5000CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsResourceGetMappedMipmappedArray_v5000CUPTI_RUNTIME_TRACE_CBID_cudaStreamAddCallback_v5000CUPTI_RUNTIME_TRACE_CBID_cudaStreamCreateWithFlags_v5000CUPTI_RUNTIME_TRACE_CBID_cudaGetTextureObjectResourceViewDesc_v5000CUPTI_RUNTIME_TRACE_CBID_cudaDeviceGetAttribute_v5000CUPTI_RUNTIME_TRACE_CBID_cudaStreamDestroy_v5050CUPTI_RUNTIME_TRACE_CBID_cudaStreamCreateWithPriority_v5050CUPTI_RUNTIME_TRACE_CBID_cudaStreamGetPriority_v5050CUPTI_RUNTIME_TRACE_CBID_cudaStreamGetFlags_v5050CUPTI_RUNTIME_TRACE_CBID_cudaDeviceGetStreamPriorityRange_v5050CUPTI_RUNTIME_TRACE_CBID_cudaMallocManaged_v6000CUPTI_RUNTIME_TRACE_CBID_cudaOccupancyMaxActiveBlocksPerMultiprocessor_v6000CUPTI_RUNTIME_TRACE_CBID_cudaStreamAttachMemAsync_v6000CUPTI_RUNTIME_TRACE_CBID_cudaGetErrorName_v6050CUPTI_RUNTIME_TRACE_CBID_cudaOccupancyMaxActiveBlocksPerMultiprocessor_v6050CUPTI_RUNTIME_TRACE_CBID_cudaLaunchKernel_v7000CUPTI_RUNTIME_TRACE_CBID_cudaGetDeviceFlags_v7000CUPTI_RUNTIME_TRACE_CBID_cudaLaunch_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaLaunchKernel_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy2D_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyToArray_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy2DToArray_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyFromArray_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy2DFromArray_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyArrayToArray_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy2DArrayToArray_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyToSymbol_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyFromSymbol_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyAsync_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyToArrayAsync_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyFromArrayAsync_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy2DAsync_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy2DToArrayAsync_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy2DFromArrayAsync_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyToSymbolAsync_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpyFromSymbolAsync_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemset_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemset2D_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemsetAsync_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemset2DAsync_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaStreamGetPriority_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaStreamGetFlags_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaStreamSynchronize_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaStreamQuery_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaStreamAttachMemAsync_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaEventRecord_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemset3D_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemset3DAsync_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy3D_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy3DAsync_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaStreamWaitEvent_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaStreamAddCallback_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy3DPeer_ptds_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemcpy3DPeerAsync_ptsz_v7000CUPTI_RUNTIME_TRACE_CBID_cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemPrefetchAsync_v8000CUPTI_RUNTIME_TRACE_CBID_cudaMemPrefetchAsync_ptsz_v8000CUPTI_RUNTIME_TRACE_CBID_cudaMemAdvise_v8000CUPTI_RUNTIME_TRACE_CBID_cudaDeviceGetP2PAttribute_v8000CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsEGLRegisterImage_v7000CUPTI_RUNTIME_TRACE_CBID_cudaEGLStreamConsumerConnect_v7000CUPTI_RUNTIME_TRACE_CBID_cudaEGLStreamConsumerDisconnect_v7000CUPTI_RUNTIME_TRACE_CBID_cudaEGLStreamConsumerAcquireFrame_v7000CUPTI_RUNTIME_TRACE_CBID_cudaEGLStreamConsumerReleaseFrame_v7000CUPTI_RUNTIME_TRACE_CBID_cudaEGLStreamProducerConnect_v7000CUPTI_RUNTIME_TRACE_CBID_cudaEGLStreamProducerDisconnect_v7000CUPTI_RUNTIME_TRACE_CBID_cudaEGLStreamProducerPresentFrame_v7000CUPTI_RUNTIME_TRACE_CBID_cudaEGLStreamProducerReturnFrame_v7000CUPTI_RUNTIME_TRACE_CBID_cudaGraphicsResourceGetMappedEglFrame_v7000CUPTI_RUNTIME_TRACE_CBID_cudaMemRangeGetAttribute_v8000CUPTI_RUNTIME_TRACE_CBID_cudaMemRangeGetAttributes_v8000CUPTI_RUNTIME_TRACE_CBID_cudaEGLStreamConsumerConnectWithFlags_v7000CUPTI_RUNTIME_TRACE_CBID_cudaLaunchCooperativeKernel_v9000CUPTI_RUNTIME_TRACE_CBID_cudaLaunchCooperativeKernel_ptsz_v9000CUPTI_RUNTIME_TRACE_CBID_cudaEventCreateFromEGLSync_v9000CUPTI_RUNTIME_TRACE_CBID_cudaLaunchCooperativeKernelMultiDevice_v9000CUPTI_RUNTIME_TRACE_CBID_cudaFuncSetAttribute_v9000CUPTI_RUNTIME_TRACE_CBID_SIZE"

var _CUPTI_RUNTIME_TRACE_CBIDIndex = [...]uint16{0, 32, 83, 135, 184, 238, 285, 334, 386, 434, 482, 529, 579, 628, 669, 722, 774, 818, 862, 912, 961, 1002, 1048, 1087, 1133, 1177, 1222, 1265, 1309, 1364, 1411, 1456, 1497, 1540, 1588, 1638, 1688, 1740, 1793, 1848, 1897, 1948, 1994, 2047, 2102, 2150, 2205, 2262, 2316, 2372, 2413, 2456, 2502, 2550, 2601, 2649, 2695, 2743, 2796, 2844, 2904, 2958, 3011, 3065, 3113, 3170, 3222, 3276, 3335, 3395, 3452, 3511, 3558, 3616, 3675, 3736, 3798, 3853, 3910, 3977, 4045, 4094, 4148, 4216, 4285, 4334, 4384, 4441, 4505, 4554, 4604, 4661, 4725, 4781, 4839, 4891, 4945, 5004, 5072, 5134, 5198, 5259, 5321, 5369, 5418, 5474, 5530, 5593, 5648, 5705, 5756, 5809, 5867, 5934, 5995, 6058, 6118, 6179, 6223, 6265, 6324, 6385, 6439, 6495, 6540, 6593, 6644, 6696, 6745, 6794, 6841, 6889, 6941, 6987, 7033, 7088, 7134, 7181, 7232, 7277, 7328, 7371, 7419, 7462, 7510, 7553, 7601, 7656, 7706, 7763, 7820, 7875, 7930, 7977, 8026, 8080, 8137, 8195, 8242, 8291, 8346, 8391, 8441, 8488, 8540, 8586, 8638, 8687, 8736, 8791, 8846, 8899, 8947, 8994, 9048, 9100, 9147, 9199, 9252, 9302, 9353, 9405, 9452, 9509, 9568, 9627, 9681, 9736, 9799, 9861, 9915, 9970, 10033, 10088, 10145, 10198, 10260, 10334, 10386, 10442, 10509, 10562, 10610, 10669, 10721, 10770, 10833, 10881, 10957, 11012, 11059, 11135, 11182, 11231, 11277, 11329, 11375, 11423, 11476, 11531, 11586, 11643, 11701, 11761, 11815, 11871, 11922, 11980, 12040, 12093, 12153, 12215, 12274, 12335, 12381, 12429, 12480, 12533, 12590, 12644, 12701, 12752, 12812, 12863, 12911, 12964, 13012, 13065, 13120, 13177, 13229, 13286, 13371, 13422, 13478, 13522, 13578, 13637, 13696, 13758, 13822, 13886, 13945, 14007, 14071, 14134, 14202, 14257, 14313, 14381, 14439, 14502, 14559, 14628, 14679, 14708}

func (i CUPTI_RUNTIME_TRACE_CBID) String() string {
	if i < 0 || i >= CUPTI_RUNTIME_TRACE_CBID(len(_CUPTI_RUNTIME_TRACE_CBIDIndex)-1) {
		return fmt.Sprintf("CUPTI_RUNTIME_TRACE_CBID(%d)", i)
	}
	return _CUPTI_RUNTIME_TRACE_CBIDName[_CUPTI_RUNTIME_TRACE_CBIDIndex[i]:_CUPTI_RUNTIME_TRACE_CBIDIndex[i+1]]
}

var _CUPTI_RUNTIME_TRACE_CBIDValues = []CUPTI_RUNTIME_TRACE_CBID{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202, 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, 215, 216, 217, 218, 219, 220, 221, 222, 223, 224, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 235, 236, 237, 238, 239, 240, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 251, 252, 253, 254, 255, 256, 257, 258, 259, 260, 261, 262, 263, 264, 265, 266, 267, 268, 269, 270, 271, 272, 273, 274}

var _CUPTI_RUNTIME_TRACE_CBIDNameToValueMap = map[string]CUPTI_RUNTIME_TRACE_CBID{
	_CUPTI_RUNTIME_TRACE_CBIDName[0:32]:        0,
	_CUPTI_RUNTIME_TRACE_CBIDName[32:83]:       1,
	_CUPTI_RUNTIME_TRACE_CBIDName[83:135]:      2,
	_CUPTI_RUNTIME_TRACE_CBIDName[135:184]:     3,
	_CUPTI_RUNTIME_TRACE_CBIDName[184:238]:     4,
	_CUPTI_RUNTIME_TRACE_CBIDName[238:285]:     5,
	_CUPTI_RUNTIME_TRACE_CBIDName[285:334]:     6,
	_CUPTI_RUNTIME_TRACE_CBIDName[334:386]:     7,
	_CUPTI_RUNTIME_TRACE_CBIDName[386:434]:     8,
	_CUPTI_RUNTIME_TRACE_CBIDName[434:482]:     9,
	_CUPTI_RUNTIME_TRACE_CBIDName[482:529]:     10,
	_CUPTI_RUNTIME_TRACE_CBIDName[529:579]:     11,
	_CUPTI_RUNTIME_TRACE_CBIDName[579:628]:     12,
	_CUPTI_RUNTIME_TRACE_CBIDName[628:669]:     13,
	_CUPTI_RUNTIME_TRACE_CBIDName[669:722]:     14,
	_CUPTI_RUNTIME_TRACE_CBIDName[722:774]:     15,
	_CUPTI_RUNTIME_TRACE_CBIDName[774:818]:     16,
	_CUPTI_RUNTIME_TRACE_CBIDName[818:862]:     17,
	_CUPTI_RUNTIME_TRACE_CBIDName[862:912]:     18,
	_CUPTI_RUNTIME_TRACE_CBIDName[912:961]:     19,
	_CUPTI_RUNTIME_TRACE_CBIDName[961:1002]:    20,
	_CUPTI_RUNTIME_TRACE_CBIDName[1002:1048]:   21,
	_CUPTI_RUNTIME_TRACE_CBIDName[1048:1087]:   22,
	_CUPTI_RUNTIME_TRACE_CBIDName[1087:1133]:   23,
	_CUPTI_RUNTIME_TRACE_CBIDName[1133:1177]:   24,
	_CUPTI_RUNTIME_TRACE_CBIDName[1177:1222]:   25,
	_CUPTI_RUNTIME_TRACE_CBIDName[1222:1265]:   26,
	_CUPTI_RUNTIME_TRACE_CBIDName[1265:1309]:   27,
	_CUPTI_RUNTIME_TRACE_CBIDName[1309:1364]:   28,
	_CUPTI_RUNTIME_TRACE_CBIDName[1364:1411]:   29,
	_CUPTI_RUNTIME_TRACE_CBIDName[1411:1456]:   30,
	_CUPTI_RUNTIME_TRACE_CBIDName[1456:1497]:   31,
	_CUPTI_RUNTIME_TRACE_CBIDName[1497:1540]:   32,
	_CUPTI_RUNTIME_TRACE_CBIDName[1540:1588]:   33,
	_CUPTI_RUNTIME_TRACE_CBIDName[1588:1638]:   34,
	_CUPTI_RUNTIME_TRACE_CBIDName[1638:1688]:   35,
	_CUPTI_RUNTIME_TRACE_CBIDName[1688:1740]:   36,
	_CUPTI_RUNTIME_TRACE_CBIDName[1740:1793]:   37,
	_CUPTI_RUNTIME_TRACE_CBIDName[1793:1848]:   38,
	_CUPTI_RUNTIME_TRACE_CBIDName[1848:1897]:   39,
	_CUPTI_RUNTIME_TRACE_CBIDName[1897:1948]:   40,
	_CUPTI_RUNTIME_TRACE_CBIDName[1948:1994]:   41,
	_CUPTI_RUNTIME_TRACE_CBIDName[1994:2047]:   42,
	_CUPTI_RUNTIME_TRACE_CBIDName[2047:2102]:   43,
	_CUPTI_RUNTIME_TRACE_CBIDName[2102:2150]:   44,
	_CUPTI_RUNTIME_TRACE_CBIDName[2150:2205]:   45,
	_CUPTI_RUNTIME_TRACE_CBIDName[2205:2262]:   46,
	_CUPTI_RUNTIME_TRACE_CBIDName[2262:2316]:   47,
	_CUPTI_RUNTIME_TRACE_CBIDName[2316:2372]:   48,
	_CUPTI_RUNTIME_TRACE_CBIDName[2372:2413]:   49,
	_CUPTI_RUNTIME_TRACE_CBIDName[2413:2456]:   50,
	_CUPTI_RUNTIME_TRACE_CBIDName[2456:2502]:   51,
	_CUPTI_RUNTIME_TRACE_CBIDName[2502:2550]:   52,
	_CUPTI_RUNTIME_TRACE_CBIDName[2550:2601]:   53,
	_CUPTI_RUNTIME_TRACE_CBIDName[2601:2649]:   54,
	_CUPTI_RUNTIME_TRACE_CBIDName[2649:2695]:   55,
	_CUPTI_RUNTIME_TRACE_CBIDName[2695:2743]:   56,
	_CUPTI_RUNTIME_TRACE_CBIDName[2743:2796]:   57,
	_CUPTI_RUNTIME_TRACE_CBIDName[2796:2844]:   58,
	_CUPTI_RUNTIME_TRACE_CBIDName[2844:2904]:   59,
	_CUPTI_RUNTIME_TRACE_CBIDName[2904:2958]:   60,
	_CUPTI_RUNTIME_TRACE_CBIDName[2958:3011]:   61,
	_CUPTI_RUNTIME_TRACE_CBIDName[3011:3065]:   62,
	_CUPTI_RUNTIME_TRACE_CBIDName[3065:3113]:   63,
	_CUPTI_RUNTIME_TRACE_CBIDName[3113:3170]:   64,
	_CUPTI_RUNTIME_TRACE_CBIDName[3170:3222]:   65,
	_CUPTI_RUNTIME_TRACE_CBIDName[3222:3276]:   66,
	_CUPTI_RUNTIME_TRACE_CBIDName[3276:3335]:   67,
	_CUPTI_RUNTIME_TRACE_CBIDName[3335:3395]:   68,
	_CUPTI_RUNTIME_TRACE_CBIDName[3395:3452]:   69,
	_CUPTI_RUNTIME_TRACE_CBIDName[3452:3511]:   70,
	_CUPTI_RUNTIME_TRACE_CBIDName[3511:3558]:   71,
	_CUPTI_RUNTIME_TRACE_CBIDName[3558:3616]:   72,
	_CUPTI_RUNTIME_TRACE_CBIDName[3616:3675]:   73,
	_CUPTI_RUNTIME_TRACE_CBIDName[3675:3736]:   74,
	_CUPTI_RUNTIME_TRACE_CBIDName[3736:3798]:   75,
	_CUPTI_RUNTIME_TRACE_CBIDName[3798:3853]:   76,
	_CUPTI_RUNTIME_TRACE_CBIDName[3853:3910]:   77,
	_CUPTI_RUNTIME_TRACE_CBIDName[3910:3977]:   78,
	_CUPTI_RUNTIME_TRACE_CBIDName[3977:4045]:   79,
	_CUPTI_RUNTIME_TRACE_CBIDName[4045:4094]:   80,
	_CUPTI_RUNTIME_TRACE_CBIDName[4094:4148]:   81,
	_CUPTI_RUNTIME_TRACE_CBIDName[4148:4216]:   82,
	_CUPTI_RUNTIME_TRACE_CBIDName[4216:4285]:   83,
	_CUPTI_RUNTIME_TRACE_CBIDName[4285:4334]:   84,
	_CUPTI_RUNTIME_TRACE_CBIDName[4334:4384]:   85,
	_CUPTI_RUNTIME_TRACE_CBIDName[4384:4441]:   86,
	_CUPTI_RUNTIME_TRACE_CBIDName[4441:4505]:   87,
	_CUPTI_RUNTIME_TRACE_CBIDName[4505:4554]:   88,
	_CUPTI_RUNTIME_TRACE_CBIDName[4554:4604]:   89,
	_CUPTI_RUNTIME_TRACE_CBIDName[4604:4661]:   90,
	_CUPTI_RUNTIME_TRACE_CBIDName[4661:4725]:   91,
	_CUPTI_RUNTIME_TRACE_CBIDName[4725:4781]:   92,
	_CUPTI_RUNTIME_TRACE_CBIDName[4781:4839]:   93,
	_CUPTI_RUNTIME_TRACE_CBIDName[4839:4891]:   94,
	_CUPTI_RUNTIME_TRACE_CBIDName[4891:4945]:   95,
	_CUPTI_RUNTIME_TRACE_CBIDName[4945:5004]:   96,
	_CUPTI_RUNTIME_TRACE_CBIDName[5004:5072]:   97,
	_CUPTI_RUNTIME_TRACE_CBIDName[5072:5134]:   98,
	_CUPTI_RUNTIME_TRACE_CBIDName[5134:5198]:   99,
	_CUPTI_RUNTIME_TRACE_CBIDName[5198:5259]:   100,
	_CUPTI_RUNTIME_TRACE_CBIDName[5259:5321]:   101,
	_CUPTI_RUNTIME_TRACE_CBIDName[5321:5369]:   102,
	_CUPTI_RUNTIME_TRACE_CBIDName[5369:5418]:   103,
	_CUPTI_RUNTIME_TRACE_CBIDName[5418:5474]:   104,
	_CUPTI_RUNTIME_TRACE_CBIDName[5474:5530]:   105,
	_CUPTI_RUNTIME_TRACE_CBIDName[5530:5593]:   106,
	_CUPTI_RUNTIME_TRACE_CBIDName[5593:5648]:   107,
	_CUPTI_RUNTIME_TRACE_CBIDName[5648:5705]:   108,
	_CUPTI_RUNTIME_TRACE_CBIDName[5705:5756]:   109,
	_CUPTI_RUNTIME_TRACE_CBIDName[5756:5809]:   110,
	_CUPTI_RUNTIME_TRACE_CBIDName[5809:5867]:   111,
	_CUPTI_RUNTIME_TRACE_CBIDName[5867:5934]:   112,
	_CUPTI_RUNTIME_TRACE_CBIDName[5934:5995]:   113,
	_CUPTI_RUNTIME_TRACE_CBIDName[5995:6058]:   114,
	_CUPTI_RUNTIME_TRACE_CBIDName[6058:6118]:   115,
	_CUPTI_RUNTIME_TRACE_CBIDName[6118:6179]:   116,
	_CUPTI_RUNTIME_TRACE_CBIDName[6179:6223]:   117,
	_CUPTI_RUNTIME_TRACE_CBIDName[6223:6265]:   118,
	_CUPTI_RUNTIME_TRACE_CBIDName[6265:6324]:   119,
	_CUPTI_RUNTIME_TRACE_CBIDName[6324:6385]:   120,
	_CUPTI_RUNTIME_TRACE_CBIDName[6385:6439]:   121,
	_CUPTI_RUNTIME_TRACE_CBIDName[6439:6495]:   122,
	_CUPTI_RUNTIME_TRACE_CBIDName[6495:6540]:   123,
	_CUPTI_RUNTIME_TRACE_CBIDName[6540:6593]:   124,
	_CUPTI_RUNTIME_TRACE_CBIDName[6593:6644]:   125,
	_CUPTI_RUNTIME_TRACE_CBIDName[6644:6696]:   126,
	_CUPTI_RUNTIME_TRACE_CBIDName[6696:6745]:   127,
	_CUPTI_RUNTIME_TRACE_CBIDName[6745:6794]:   128,
	_CUPTI_RUNTIME_TRACE_CBIDName[6794:6841]:   129,
	_CUPTI_RUNTIME_TRACE_CBIDName[6841:6889]:   130,
	_CUPTI_RUNTIME_TRACE_CBIDName[6889:6941]:   131,
	_CUPTI_RUNTIME_TRACE_CBIDName[6941:6987]:   132,
	_CUPTI_RUNTIME_TRACE_CBIDName[6987:7033]:   133,
	_CUPTI_RUNTIME_TRACE_CBIDName[7033:7088]:   134,
	_CUPTI_RUNTIME_TRACE_CBIDName[7088:7134]:   135,
	_CUPTI_RUNTIME_TRACE_CBIDName[7134:7181]:   136,
	_CUPTI_RUNTIME_TRACE_CBIDName[7181:7232]:   137,
	_CUPTI_RUNTIME_TRACE_CBIDName[7232:7277]:   138,
	_CUPTI_RUNTIME_TRACE_CBIDName[7277:7328]:   139,
	_CUPTI_RUNTIME_TRACE_CBIDName[7328:7371]:   140,
	_CUPTI_RUNTIME_TRACE_CBIDName[7371:7419]:   141,
	_CUPTI_RUNTIME_TRACE_CBIDName[7419:7462]:   142,
	_CUPTI_RUNTIME_TRACE_CBIDName[7462:7510]:   143,
	_CUPTI_RUNTIME_TRACE_CBIDName[7510:7553]:   144,
	_CUPTI_RUNTIME_TRACE_CBIDName[7553:7601]:   145,
	_CUPTI_RUNTIME_TRACE_CBIDName[7601:7656]:   146,
	_CUPTI_RUNTIME_TRACE_CBIDName[7656:7706]:   147,
	_CUPTI_RUNTIME_TRACE_CBIDName[7706:7763]:   148,
	_CUPTI_RUNTIME_TRACE_CBIDName[7763:7820]:   149,
	_CUPTI_RUNTIME_TRACE_CBIDName[7820:7875]:   150,
	_CUPTI_RUNTIME_TRACE_CBIDName[7875:7930]:   151,
	_CUPTI_RUNTIME_TRACE_CBIDName[7930:7977]:   152,
	_CUPTI_RUNTIME_TRACE_CBIDName[7977:8026]:   153,
	_CUPTI_RUNTIME_TRACE_CBIDName[8026:8080]:   154,
	_CUPTI_RUNTIME_TRACE_CBIDName[8080:8137]:   155,
	_CUPTI_RUNTIME_TRACE_CBIDName[8137:8195]:   156,
	_CUPTI_RUNTIME_TRACE_CBIDName[8195:8242]:   157,
	_CUPTI_RUNTIME_TRACE_CBIDName[8242:8291]:   158,
	_CUPTI_RUNTIME_TRACE_CBIDName[8291:8346]:   159,
	_CUPTI_RUNTIME_TRACE_CBIDName[8346:8391]:   160,
	_CUPTI_RUNTIME_TRACE_CBIDName[8391:8441]:   161,
	_CUPTI_RUNTIME_TRACE_CBIDName[8441:8488]:   162,
	_CUPTI_RUNTIME_TRACE_CBIDName[8488:8540]:   163,
	_CUPTI_RUNTIME_TRACE_CBIDName[8540:8586]:   164,
	_CUPTI_RUNTIME_TRACE_CBIDName[8586:8638]:   165,
	_CUPTI_RUNTIME_TRACE_CBIDName[8638:8687]:   166,
	_CUPTI_RUNTIME_TRACE_CBIDName[8687:8736]:   167,
	_CUPTI_RUNTIME_TRACE_CBIDName[8736:8791]:   168,
	_CUPTI_RUNTIME_TRACE_CBIDName[8791:8846]:   169,
	_CUPTI_RUNTIME_TRACE_CBIDName[8846:8899]:   170,
	_CUPTI_RUNTIME_TRACE_CBIDName[8899:8947]:   171,
	_CUPTI_RUNTIME_TRACE_CBIDName[8947:8994]:   172,
	_CUPTI_RUNTIME_TRACE_CBIDName[8994:9048]:   173,
	_CUPTI_RUNTIME_TRACE_CBIDName[9048:9100]:   174,
	_CUPTI_RUNTIME_TRACE_CBIDName[9100:9147]:   175,
	_CUPTI_RUNTIME_TRACE_CBIDName[9147:9199]:   176,
	_CUPTI_RUNTIME_TRACE_CBIDName[9199:9252]:   177,
	_CUPTI_RUNTIME_TRACE_CBIDName[9252:9302]:   178,
	_CUPTI_RUNTIME_TRACE_CBIDName[9302:9353]:   179,
	_CUPTI_RUNTIME_TRACE_CBIDName[9353:9405]:   180,
	_CUPTI_RUNTIME_TRACE_CBIDName[9405:9452]:   181,
	_CUPTI_RUNTIME_TRACE_CBIDName[9452:9509]:   182,
	_CUPTI_RUNTIME_TRACE_CBIDName[9509:9568]:   183,
	_CUPTI_RUNTIME_TRACE_CBIDName[9568:9627]:   184,
	_CUPTI_RUNTIME_TRACE_CBIDName[9627:9681]:   185,
	_CUPTI_RUNTIME_TRACE_CBIDName[9681:9736]:   186,
	_CUPTI_RUNTIME_TRACE_CBIDName[9736:9799]:   187,
	_CUPTI_RUNTIME_TRACE_CBIDName[9799:9861]:   188,
	_CUPTI_RUNTIME_TRACE_CBIDName[9861:9915]:   189,
	_CUPTI_RUNTIME_TRACE_CBIDName[9915:9970]:   190,
	_CUPTI_RUNTIME_TRACE_CBIDName[9970:10033]:  191,
	_CUPTI_RUNTIME_TRACE_CBIDName[10033:10088]: 192,
	_CUPTI_RUNTIME_TRACE_CBIDName[10088:10145]: 193,
	_CUPTI_RUNTIME_TRACE_CBIDName[10145:10198]: 194,
	_CUPTI_RUNTIME_TRACE_CBIDName[10198:10260]: 195,
	_CUPTI_RUNTIME_TRACE_CBIDName[10260:10334]: 196,
	_CUPTI_RUNTIME_TRACE_CBIDName[10334:10386]: 197,
	_CUPTI_RUNTIME_TRACE_CBIDName[10386:10442]: 198,
	_CUPTI_RUNTIME_TRACE_CBIDName[10442:10509]: 199,
	_CUPTI_RUNTIME_TRACE_CBIDName[10509:10562]: 200,
	_CUPTI_RUNTIME_TRACE_CBIDName[10562:10610]: 201,
	_CUPTI_RUNTIME_TRACE_CBIDName[10610:10669]: 202,
	_CUPTI_RUNTIME_TRACE_CBIDName[10669:10721]: 203,
	_CUPTI_RUNTIME_TRACE_CBIDName[10721:10770]: 204,
	_CUPTI_RUNTIME_TRACE_CBIDName[10770:10833]: 205,
	_CUPTI_RUNTIME_TRACE_CBIDName[10833:10881]: 206,
	_CUPTI_RUNTIME_TRACE_CBIDName[10881:10957]: 207,
	_CUPTI_RUNTIME_TRACE_CBIDName[10957:11012]: 208,
	_CUPTI_RUNTIME_TRACE_CBIDName[11012:11059]: 209,
	_CUPTI_RUNTIME_TRACE_CBIDName[11059:11135]: 210,
	_CUPTI_RUNTIME_TRACE_CBIDName[11135:11182]: 211,
	_CUPTI_RUNTIME_TRACE_CBIDName[11182:11231]: 212,
	_CUPTI_RUNTIME_TRACE_CBIDName[11231:11277]: 213,
	_CUPTI_RUNTIME_TRACE_CBIDName[11277:11329]: 214,
	_CUPTI_RUNTIME_TRACE_CBIDName[11329:11375]: 215,
	_CUPTI_RUNTIME_TRACE_CBIDName[11375:11423]: 216,
	_CUPTI_RUNTIME_TRACE_CBIDName[11423:11476]: 217,
	_CUPTI_RUNTIME_TRACE_CBIDName[11476:11531]: 218,
	_CUPTI_RUNTIME_TRACE_CBIDName[11531:11586]: 219,
	_CUPTI_RUNTIME_TRACE_CBIDName[11586:11643]: 220,
	_CUPTI_RUNTIME_TRACE_CBIDName[11643:11701]: 221,
	_CUPTI_RUNTIME_TRACE_CBIDName[11701:11761]: 222,
	_CUPTI_RUNTIME_TRACE_CBIDName[11761:11815]: 223,
	_CUPTI_RUNTIME_TRACE_CBIDName[11815:11871]: 224,
	_CUPTI_RUNTIME_TRACE_CBIDName[11871:11922]: 225,
	_CUPTI_RUNTIME_TRACE_CBIDName[11922:11980]: 226,
	_CUPTI_RUNTIME_TRACE_CBIDName[11980:12040]: 227,
	_CUPTI_RUNTIME_TRACE_CBIDName[12040:12093]: 228,
	_CUPTI_RUNTIME_TRACE_CBIDName[12093:12153]: 229,
	_CUPTI_RUNTIME_TRACE_CBIDName[12153:12215]: 230,
	_CUPTI_RUNTIME_TRACE_CBIDName[12215:12274]: 231,
	_CUPTI_RUNTIME_TRACE_CBIDName[12274:12335]: 232,
	_CUPTI_RUNTIME_TRACE_CBIDName[12335:12381]: 233,
	_CUPTI_RUNTIME_TRACE_CBIDName[12381:12429]: 234,
	_CUPTI_RUNTIME_TRACE_CBIDName[12429:12480]: 235,
	_CUPTI_RUNTIME_TRACE_CBIDName[12480:12533]: 236,
	_CUPTI_RUNTIME_TRACE_CBIDName[12533:12590]: 237,
	_CUPTI_RUNTIME_TRACE_CBIDName[12590:12644]: 238,
	_CUPTI_RUNTIME_TRACE_CBIDName[12644:12701]: 239,
	_CUPTI_RUNTIME_TRACE_CBIDName[12701:12752]: 240,
	_CUPTI_RUNTIME_TRACE_CBIDName[12752:12812]: 241,
	_CUPTI_RUNTIME_TRACE_CBIDName[12812:12863]: 242,
	_CUPTI_RUNTIME_TRACE_CBIDName[12863:12911]: 243,
	_CUPTI_RUNTIME_TRACE_CBIDName[12911:12964]: 244,
	_CUPTI_RUNTIME_TRACE_CBIDName[12964:13012]: 245,
	_CUPTI_RUNTIME_TRACE_CBIDName[13012:13065]: 246,
	_CUPTI_RUNTIME_TRACE_CBIDName[13065:13120]: 247,
	_CUPTI_RUNTIME_TRACE_CBIDName[13120:13177]: 248,
	_CUPTI_RUNTIME_TRACE_CBIDName[13177:13229]: 249,
	_CUPTI_RUNTIME_TRACE_CBIDName[13229:13286]: 250,
	_CUPTI_RUNTIME_TRACE_CBIDName[13286:13371]: 251,
	_CUPTI_RUNTIME_TRACE_CBIDName[13371:13422]: 252,
	_CUPTI_RUNTIME_TRACE_CBIDName[13422:13478]: 253,
	_CUPTI_RUNTIME_TRACE_CBIDName[13478:13522]: 254,
	_CUPTI_RUNTIME_TRACE_CBIDName[13522:13578]: 255,
	_CUPTI_RUNTIME_TRACE_CBIDName[13578:13637]: 256,
	_CUPTI_RUNTIME_TRACE_CBIDName[13637:13696]: 257,
	_CUPTI_RUNTIME_TRACE_CBIDName[13696:13758]: 258,
	_CUPTI_RUNTIME_TRACE_CBIDName[13758:13822]: 259,
	_CUPTI_RUNTIME_TRACE_CBIDName[13822:13886]: 260,
	_CUPTI_RUNTIME_TRACE_CBIDName[13886:13945]: 261,
	_CUPTI_RUNTIME_TRACE_CBIDName[13945:14007]: 262,
	_CUPTI_RUNTIME_TRACE_CBIDName[14007:14071]: 263,
	_CUPTI_RUNTIME_TRACE_CBIDName[14071:14134]: 264,
	_CUPTI_RUNTIME_TRACE_CBIDName[14134:14202]: 265,
	_CUPTI_RUNTIME_TRACE_CBIDName[14202:14257]: 266,
	_CUPTI_RUNTIME_TRACE_CBIDName[14257:14313]: 267,
	_CUPTI_RUNTIME_TRACE_CBIDName[14313:14381]: 268,
	_CUPTI_RUNTIME_TRACE_CBIDName[14381:14439]: 269,
	_CUPTI_RUNTIME_TRACE_CBIDName[14439:14502]: 270,
	_CUPTI_RUNTIME_TRACE_CBIDName[14502:14559]: 271,
	_CUPTI_RUNTIME_TRACE_CBIDName[14559:14628]: 272,
	_CUPTI_RUNTIME_TRACE_CBIDName[14628:14679]: 273,
	_CUPTI_RUNTIME_TRACE_CBIDName[14679:14708]: 274,
}

// CUPTI_RUNTIME_TRACE_CBIDString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CUPTI_RUNTIME_TRACE_CBIDString(s string) (CUPTI_RUNTIME_TRACE_CBID, error) {
	if val, ok := _CUPTI_RUNTIME_TRACE_CBIDNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CUPTI_RUNTIME_TRACE_CBID values", s)
}

// CUPTI_RUNTIME_TRACE_CBIDValues returns all values of the enum
func CUPTI_RUNTIME_TRACE_CBIDValues() []CUPTI_RUNTIME_TRACE_CBID {
	return _CUPTI_RUNTIME_TRACE_CBIDValues
}

// IsACUPTI_RUNTIME_TRACE_CBID returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CUPTI_RUNTIME_TRACE_CBID) IsACUPTI_RUNTIME_TRACE_CBID() bool {
	for _, v := range _CUPTI_RUNTIME_TRACE_CBIDValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CUPTI_RUNTIME_TRACE_CBID
func (i CUPTI_RUNTIME_TRACE_CBID) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CUPTI_RUNTIME_TRACE_CBID
func (i *CUPTI_RUNTIME_TRACE_CBID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CUPTI_RUNTIME_TRACE_CBID should be a string, got %s", data)
	}

	var err error
	*i, err = CUPTI_RUNTIME_TRACE_CBIDString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for CUPTI_RUNTIME_TRACE_CBID
func (i CUPTI_RUNTIME_TRACE_CBID) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for CUPTI_RUNTIME_TRACE_CBID
func (i *CUPTI_RUNTIME_TRACE_CBID) UnmarshalText(text []byte) error {
	var err error
	*i, err = CUPTI_RUNTIME_TRACE_CBIDString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for CUPTI_RUNTIME_TRACE_CBID
func (i CUPTI_RUNTIME_TRACE_CBID) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CUPTI_RUNTIME_TRACE_CBID
func (i *CUPTI_RUNTIME_TRACE_CBID) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = CUPTI_RUNTIME_TRACE_CBIDString(s)
	return err
}

func (i CUPTI_RUNTIME_TRACE_CBID) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *CUPTI_RUNTIME_TRACE_CBID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := CUPTI_RUNTIME_TRACE_CBIDString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}