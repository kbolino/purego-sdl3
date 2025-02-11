package sdl

type GPUSwapchainComposition uint32

const (
	GPUSwapchainCompositionSdr GPUSwapchainComposition = iota
	GPUSwapchainCompositionSdrLinear
	GPUSwapchainCompositionHdrExtendedLinear
	GPUSwapchainCompositionHdr10St2084
)

type GPUPresentMode uint32

const (
	GPUPresentModeVsync GPUPresentMode = iota
	GPUPresentModeImmediate
	GPUPresentModeMailbox
)

type GPUSamplerAddressMode uint32

const (
	GPUSamplerAddressModeRepeat GPUSamplerAddressMode = iota
	GPUSamplerAddressModeMirroredRepeat
	GPUSamplerAddressModeClampToEdge
)

type GPUSamplerMipmapMode uint32

const (
	GPUSamplerMipmapModeNearest GPUSamplerMipmapMode = iota
	GPUSamplerMipmapModeLinear
)

type GPUFilter uint32

const (
	GPUFilterNearest GPUFilter = iota
	GPUFilterLinear
)

type GPUBlendFactor uint32

const (
	GPUBlendFactorInvalid GPUBlendFactor = iota
	GPUBlendFactorZero
	GPUBlendFactorOne
	GPUBlendFactorSrcColor
	GPUBlendFactorOneMinusSrcColor
	GPUBlendFactorDstColor
	GPUBlendFactorOneMinusDstColor
	GPUBlendFactorSrcAlpha
	GPUBlendFactorOneMinusSrcAlpha
	GPUBlendFactorDstAlpha
	GPUBlendFactorOneMinusDstAlpha
	GPUBlendFactorConstantColor
	GPUBlendFactorOneMinusConstantColor
	GPUBlendFactorSrcAlphaSaturate
)

type GPUBlendOp uint32

const (
	GPUBlendOpInvalid GPUBlendOp = iota
	GPUBlendOpAdd
	GPUBlendOpSubtract
	GPUBlendOpReverseSubtract
	GPUBlendOpMin
	GPUBlendOpMax
)

type GPUStencilOp uint32

const (
	GPUStencilOpInvalid GPUStencilOp = iota
	GPUStencilOpKeep
	GPUStencilOpZero
	GPUStencilOpReplace
	GPUStencilOpIncrementAndClamp
	GPUStencilOpDecrementAndClamp
	GPUStencilOpInvert
	GPUStencilOpIncrementAndWrap
	GPUStencilOpDecrementAndWrap
)

type GPUCompareOp uint32

const (
	GPUCompareOpInvalid GPUCompareOp = iota
	GPUCompareOpNever
	GPUCompareOpLess
	GPUCompareOpEqual
	GPUCompareOpLessOrEqual
	GPUCompareOpGreater
	GPUCompareOpNotEqual
	GPUCompareOpGreaterOrEqual
	GPUCompareOpAlways
)

type GPUFrontFace uint32

const (
	GPUFrontFaceCounterClockwise GPUFrontFace = iota
	GPUFrontFaceClockwise
)

type GPUCullMode uint32

const (
	GPUCullModeNone GPUCullMode = iota
	GPUCullModeFront
	GPUCullModeBack
)

type GPUFillMode uint32

const (
	GPUFillModeFill GPUFillMode = iota
	GPUFillModeLine
)

type GPUVertexInputRate uint32

const (
	GPUVertexInputRateVertex GPUVertexInputRate = iota
	GPUVertexInputRateInstance
)

type GPUVertexElementFormat uint32

const (
	GPUVertexElementFormatInvalid GPUVertexElementFormat = iota
	GPUVertexElementFormatInt
	GPUVertexElementFormatInt2
	GPUVertexElementFormatInt3
	GPUVertexElementFormatInt4
	GPUVertexElementFormatUint
	GPUVertexElementFormatUint2
	GPUVertexElementFormatUint3
	GPUVertexElementFormatUint4
	GPUVertexElementFormatFloat
	GPUVertexElementFormatFloat2
	GPUVertexElementFormatFloat3
	GPUVertexElementFormatFloat4
	GPUVertexElementFormatByte2
	GPUVertexElementFormatByte4
	GPUVertexElementFormatUbyte2
	GPUVertexElementFormatUbyte4
	GPUVertexElementFormatByte2Norm
	GPUVertexElementFormatByte4Norm
	GPUVertexElementFormatUbyte2Norm
	GPUVertexElementFormatUbyte4Norm
	GPUVertexElementFormatShort2
	GPUVertexElementFormatShort4
	GPUVertexElementFormatUshort2
	GPUVertexElementFormatUshort4
	GPUVertexElementFormatShort2Norm
	GPUVertexElementFormatShort4Norm
	GPUVertexElementFormatUshort2Norm
	GPUVertexElementFormatUshort4Norm
	GPUVertexElementFormatHalf2
	GPUVertexElementFormatHalf4
)

type GPUShaderStage uint32

const (
	GPUShaderStageVertex GPUShaderStage = iota
	GPUShaderStageFragment
)

type GPUTransferBufferUsage uint32

const (
	GPUTransferBufferUsageUpload GPUTransferBufferUsage = iota
	GPUTransferBufferUsageDownload
)

type GPUCubeMapFace uint32

const (
	GPUCubeMapFacePositiveX GPUCubeMapFace = iota
	GPUCubeMapFaceNegativeX
	GPUCubeMapFacePositiveY
	GPUCubeMapFaceNegativeY
	GPUCubeMapFacePositiveZ
	GPUCubeMapFaceNegativeZ
)

type GPUSampleCount uint32

const (
	GPUSampleCount1 GPUSampleCount = iota
	GPUSampleCount2
	GPUSampleCount4
	GPUSampleCount8
)

type GPUTextureType uint32

const (
	GPUTextureType2D GPUTextureType = iota
	GPUTextureType2DArray
	GPUTextureType3D
	GPUTextureTypeCube
	GPUTextureTypeCubeArray
)

type GPUTextureFormat uint32

const (
	GPUTextureFormatInvalid GPUTextureFormat = iota
	GPUTextureFormatA8Unorm
	GPUTextureFormatR8Unorm
	GPUTextureFormatR8G8Unorm
	GPUTextureFormatR8G8B8A8Unorm
	GPUTextureFormatR16Unorm
	GPUTextureFormatR16G16Unorm
	GPUTextureFormatR16G16B16A16Unorm
	GPUTextureFormatR10G10B10A2Unorm
	GPUTextureFormatB5G6R5Unorm
	GPUTextureFormatB5G5R5A1Unorm
	GPUTextureFormatB4G4R4A4Unorm
	GPUTextureFormatB8G8R8A8Unorm
	GPUTextureFormatBc1RgbaUnorm
	GPUTextureFormatBc2RgbaUnorm
	GPUTextureFormatBc3RgbaUnorm
	GPUTextureFormatBc4RUnorm
	GPUTextureFormatBc5RgUnorm
	GPUTextureFormatBc7RgbaUnorm
	GPUTextureFormatBc6HRgbFloat
	GPUTextureFormatBc6HRgbUfloat
	GPUTextureFormatR8Snorm
	GPUTextureFormatR8G8Snorm
	GPUTextureFormatR8G8B8A8Snorm
	GPUTextureFormatR16Snorm
	GPUTextureFormatR16G16Snorm
	GPUTextureFormatR16G16B16A16Snorm
	GPUTextureFormatR16Float
	GPUTextureFormatR16G16Float
	GPUTextureFormatR16G16B16A16Float
	GPUTextureFormatR32Float
	GPUTextureFormatR32G32Float
	GPUTextureFormatR32G32B32A32Float
	GPUTextureFormatR11G11B10Ufloat
	GPUTextureFormatR8Uint
	GPUTextureFormatR8G8Uint
	GPUTextureFormatR8G8B8A8Uint
	GPUTextureFormatR16Uint
	GPUTextureFormatR16G16Uint
	GPUTextureFormatR16G16B16A16Uint
	GPUTextureFormatR32Uint
	GPUTextureFormatR32G32Uint
	GPUTextureFormatR32G32B32A32Uint
	GPUTextureFormatR8Int
	GPUTextureFormatR8G8Int
	GPUTextureFormatR8G8B8A8Int
	GPUTextureFormatR16Int
	GPUTextureFormatR16G16Int
	GPUTextureFormatR16G16B16A16Int
	GPUTextureFormatR32Int
	GPUTextureFormatR32G32Int
	GPUTextureFormatR32G32B32A32Int
	GPUTextureFormatR8G8B8A8UnormSrgb
	GPUTextureFormatB8G8R8A8UnormSrgb
	GPUTextureFormatBc1RgbaUnormSrgb
	GPUTextureFormatBc2RgbaUnormSrgb
	GPUTextureFormatBc3RgbaUnormSrgb
	GPUTextureFormatBc7RgbaUnormSrgb
	GPUTextureFormatD16Unorm
	GPUTextureFormatD24Unorm
	GPUTextureFormatD32Float
	GPUTextureFormatD24UnormS8Uint
	GPUTextureFormatD32FloatS8Uint
	GPUTextureFormatAstc4x4Unorm
	GPUTextureFormatAstc5x4Unorm
	GPUTextureFormatAstc5x5Unorm
	GPUTextureFormatAstc6x5Unorm
	GPUTextureFormatAstc6x6Unorm
	GPUTextureFormatAstc8x5Unorm
	GPUTextureFormatAstc8x6Unorm
	GPUTextureFormatAstc8x8Unorm
	GPUTextureFormatAstc10x5Unorm
	GPUTextureFormatAstc10x6Unorm
	GPUTextureFormatAstc10x8Unorm
	GPUTextureFormatAstc10x10Unorm
	GPUTextureFormatAstc12x10Unorm
	GPUTextureFormatAstc12x12Unorm
	GPUTextureFormatAstc4x4UnormSrgb
	GPUTextureFormatAstc5x4UnormSrgb
	GPUTextureFormatAstc5x5UnormSrgb
	GPUTextureFormatAstc6x5UnormSrgb
	GPUTextureFormatAstc6x6UnormSrgb
	GPUTextureFormatAstc8x5UnormSrgb
	GPUTextureFormatAstc8x6UnormSrgb
	GPUTextureFormatAstc8x8UnormSrgb
	GPUTextureFormatAstc10x5UnormSrgb
	GPUTextureFormatAstc10x6UnormSrgb
	GPUTextureFormatAstc10x8UnormSrgb
	GPUTextureFormatAstc10x10UnormSrgb
	GPUTextureFormatAstc12x10UnormSrgb
	GPUTextureFormatAstc12x12UnormSrgb
	GPUTextureFormatAstc4x4Float
	GPUTextureFormatAstc5x4Float
	GPUTextureFormatAstc5x5Float
	GPUTextureFormatAstc6x5Float
	GPUTextureFormatAstc6x6Float
	GPUTextureFormatAstc8x5Float
	GPUTextureFormatAstc8x6Float
	GPUTextureFormatAstc8x8Float
	GPUTextureFormatAstc10x5Float
	GPUTextureFormatAstc10x6Float
	GPUTextureFormatAstc10x8Float
	GPUTextureFormatAstc10x10Float
	GPUTextureFormatAstc12x10Float
	GPUTextureFormatAstc12x12Float
)

type GPUIndexElementSize uint32

const (
	GPUIndexElementSize16Bit GPUIndexElementSize = iota
	GPUIndexElementSize32Bit
)

type GPUStoreOp uint32

const (
	GPUStoreOpStore GPUStoreOp = iota
	GPUStoreOpDontCare
	GPUStoreOpResolve
	GPUStoreOpResolveAndStore
)

type GPULoadOp uint32

const (
	GPULoadOpLoad GPULoadOp = iota
	GPULoadOpClear
	GPULoadOpDontCare
)

type GPUPrimitiveType uint32

const (
	GPUPrimitiveTypeTrianglelist GPUPrimitiveType = iota
	GPUPrimitiveTypeTrianglestrip
	GPUPrimitiveTypeLinelist
	GPUPrimitiveTypeLinestrip
	GPUPrimitiveTypePointlist
)

type GPUDevice struct{}

// func AcquireGPUCommandBuffer(device *GPUDevice) *GPUCommandBuffer {
//	return sdlAcquireGPUCommandBuffer(device)
// }

// func AcquireGPUSwapchainTexture(command_buffer *GPUCommandBuffer, window *Window, swapchain_texture **GPUTexture, swapchain_texture_width *uint32, swapchain_texture_height *uint32) bool {
//	return sdlAcquireGPUSwapchainTexture(command_buffer, window, swapchain_texture, swapchain_texture_width, swapchain_texture_height)
// }

// func BeginGPUComputePass(command_buffer *GPUCommandBuffer, storage_texture_bindings *GPUStorageTextureReadWriteBinding, num_storage_texture_bindings uint32, storage_buffer_bindings *GPUStorageBufferReadWriteBinding, num_storage_buffer_bindings uint32) *GPUComputePass {
//	return sdlBeginGPUComputePass(command_buffer, storage_texture_bindings, num_storage_texture_bindings, storage_buffer_bindings, num_storage_buffer_bindings)
// }

// func BeginGPUCopyPass(command_buffer *GPUCommandBuffer) *GPUCopyPass {
//	return sdlBeginGPUCopyPass(command_buffer)
// }

// func BeginGPURenderPass(command_buffer *GPUCommandBuffer, color_target_infos *GPUColorTargetInfo, num_color_targets uint32, depth_stencil_target_info *GPUDepthStencilTargetInfo) *GPURenderPass {
//	return sdlBeginGPURenderPass(command_buffer, color_target_infos, num_color_targets, depth_stencil_target_info)
// }

// func BindGPUComputePipeline(compute_pass *GPUComputePass, compute_pipeline *GPUComputePipeline)  {
//	sdlBindGPUComputePipeline(compute_pass, compute_pipeline)
// }

// func BindGPUComputeSamplers(compute_pass *GPUComputePass, first_slot uint32, texture_sampler_bindings *GPUTextureSamplerBinding, num_bindings uint32)  {
//	sdlBindGPUComputeSamplers(compute_pass, first_slot, texture_sampler_bindings, num_bindings)
// }

// func BindGPUComputeStorageBuffers(compute_pass *GPUComputePass, first_slot uint32, storage_buffers **GPUBuffer, num_bindings uint32)  {
//	sdlBindGPUComputeStorageBuffers(compute_pass, first_slot, storage_buffers, num_bindings)
// }

// func BindGPUComputeStorageTextures(compute_pass *GPUComputePass, first_slot uint32, storage_textures **GPUTexture, num_bindings uint32)  {
//	sdlBindGPUComputeStorageTextures(compute_pass, first_slot, storage_textures, num_bindings)
// }

// func BindGPUFragmentSamplers(render_pass *GPURenderPass, first_slot uint32, texture_sampler_bindings *GPUTextureSamplerBinding, num_bindings uint32)  {
//	sdlBindGPUFragmentSamplers(render_pass, first_slot, texture_sampler_bindings, num_bindings)
// }

// func BindGPUFragmentStorageBuffers(render_pass *GPURenderPass, first_slot uint32, storage_buffers **GPUBuffer, num_bindings uint32)  {
//	sdlBindGPUFragmentStorageBuffers(render_pass, first_slot, storage_buffers, num_bindings)
// }

// func BindGPUFragmentStorageTextures(render_pass *GPURenderPass, first_slot uint32, storage_textures **GPUTexture, num_bindings uint32)  {
//	sdlBindGPUFragmentStorageTextures(render_pass, first_slot, storage_textures, num_bindings)
// }

// func BindGPUGraphicsPipeline(render_pass *GPURenderPass, graphics_pipeline *GPUGraphicsPipeline)  {
//	sdlBindGPUGraphicsPipeline(render_pass, graphics_pipeline)
// }

// func BindGPUIndexBuffer(render_pass *GPURenderPass, binding *GPUBufferBinding, index_element_size GPUIndexElementSize)  {
//	sdlBindGPUIndexBuffer(render_pass, binding, index_element_size)
// }

// func BindGPUVertexBuffers(render_pass *GPURenderPass, first_slot uint32, bindings *GPUBufferBinding, num_bindings uint32)  {
//	sdlBindGPUVertexBuffers(render_pass, first_slot, bindings, num_bindings)
// }

// func BindGPUVertexSamplers(render_pass *GPURenderPass, first_slot uint32, texture_sampler_bindings *GPUTextureSamplerBinding, num_bindings uint32)  {
//	sdlBindGPUVertexSamplers(render_pass, first_slot, texture_sampler_bindings, num_bindings)
// }

// func BindGPUVertexStorageBuffers(render_pass *GPURenderPass, first_slot uint32, storage_buffers **GPUBuffer, num_bindings uint32)  {
//	sdlBindGPUVertexStorageBuffers(render_pass, first_slot, storage_buffers, num_bindings)
// }

// func BindGPUVertexStorageTextures(render_pass *GPURenderPass, first_slot uint32, storage_textures **GPUTexture, num_bindings uint32)  {
//	sdlBindGPUVertexStorageTextures(render_pass, first_slot, storage_textures, num_bindings)
// }

// func BlitGPUTexture(command_buffer *GPUCommandBuffer, info *GPUBlitInfo)  {
//	sdlBlitGPUTexture(command_buffer, info)
// }

// func CalculateGPUTextureFormatSize(format GPUTextureFormat, width uint32, height uint32, depth_or_layer_count uint32) uint32 {
//	return sdlCalculateGPUTextureFormatSize(format, width, height, depth_or_layer_count)
// }

// func CancelGPUCommandBuffer(command_buffer *GPUCommandBuffer) bool {
//	return sdlCancelGPUCommandBuffer(command_buffer)
// }

// func ClaimWindowForGPUDevice(device *GPUDevice, window *Window) bool {
//	return sdlClaimWindowForGPUDevice(device, window)
// }

// func CopyGPUBufferToBuffer(copy_pass *GPUCopyPass, source *GPUBufferLocation, destination *GPUBufferLocation, size uint32, cycle bool)  {
//	sdlCopyGPUBufferToBuffer(copy_pass, source, destination, size, cycle)
// }

// func CopyGPUTextureToTexture(copy_pass *GPUCopyPass, source *GPUTextureLocation, destination *GPUTextureLocation, w uint32, h uint32, d uint32, cycle bool)  {
//	sdlCopyGPUTextureToTexture(copy_pass, source, destination, w, h, d, cycle)
// }

// func CreateGPUBuffer(device *GPUDevice, createinfo *GPUBufferCreateInfo) *GPUBuffer {
//	return sdlCreateGPUBuffer(device, createinfo)
// }

// func CreateGPUComputePipeline(device *GPUDevice, createinfo *GPUComputePipelineCreateInfo) *GPUComputePipeline {
//	return sdlCreateGPUComputePipeline(device, createinfo)
// }

// func CreateGPUDevice(format_flags GPUShaderFormat, debug_mode bool, name string) *GPUDevice {
//	return sdlCreateGPUDevice(format_flags, debug_mode, name)
// }

// func CreateGPUDeviceWithProperties(props PropertiesID) *GPUDevice {
//	return sdlCreateGPUDeviceWithProperties(props)
// }

// func CreateGPUGraphicsPipeline(device *GPUDevice, createinfo *GPUGraphicsPipelineCreateInfo) *GPUGraphicsPipeline {
//	return sdlCreateGPUGraphicsPipeline(device, createinfo)
// }

// func CreateGPUSampler(device *GPUDevice, createinfo *GPUSamplerCreateInfo) *GPUSampler {
//	return sdlCreateGPUSampler(device, createinfo)
// }

// func CreateGPUShader(device *GPUDevice, createinfo *GPUShaderCreateInfo) *GPUShader {
//	return sdlCreateGPUShader(device, createinfo)
// }

// func CreateGPUTexture(device *GPUDevice, createinfo *GPUTextureCreateInfo) *GPUTexture {
//	return sdlCreateGPUTexture(device, createinfo)
// }

// func CreateGPUTransferBuffer(device *GPUDevice, createinfo *GPUTransferBufferCreateInfo) *GPUTransferBuffer {
//	return sdlCreateGPUTransferBuffer(device, createinfo)
// }

// func DestroyGPUDevice(device *GPUDevice)  {
//	sdlDestroyGPUDevice(device)
// }

// func DispatchGPUCompute(compute_pass *GPUComputePass, groupcount_x uint32, groupcount_y uint32, groupcount_z uint32)  {
//	sdlDispatchGPUCompute(compute_pass, groupcount_x, groupcount_y, groupcount_z)
// }

// func DispatchGPUComputeIndirect(compute_pass *GPUComputePass, buffer *GPUBuffer, offset uint32)  {
//	sdlDispatchGPUComputeIndirect(compute_pass, buffer, offset)
// }

// func DownloadFromGPUBuffer(copy_pass *GPUCopyPass, source *GPUBufferRegion, destination *GPUTransferBufferLocation)  {
//	sdlDownloadFromGPUBuffer(copy_pass, source, destination)
// }

// func DownloadFromGPUTexture(copy_pass *GPUCopyPass, source *GPUTextureRegion, destination *GPUTextureTransferInfo)  {
//	sdlDownloadFromGPUTexture(copy_pass, source, destination)
// }

// func DrawGPUIndexedPrimitives(render_pass *GPURenderPass, num_indices uint32, num_instances uint32, first_index uint32, vertex_offset int32, first_instance uint32)  {
//	sdlDrawGPUIndexedPrimitives(render_pass, num_indices, num_instances, first_index, vertex_offset, first_instance)
// }

// func DrawGPUIndexedPrimitivesIndirect(render_pass *GPURenderPass, buffer *GPUBuffer, offset uint32, draw_count uint32)  {
//	sdlDrawGPUIndexedPrimitivesIndirect(render_pass, buffer, offset, draw_count)
// }

// func DrawGPUPrimitives(render_pass *GPURenderPass, num_vertices uint32, num_instances uint32, first_vertex uint32, first_instance uint32)  {
//	sdlDrawGPUPrimitives(render_pass, num_vertices, num_instances, first_vertex, first_instance)
// }

// func DrawGPUPrimitivesIndirect(render_pass *GPURenderPass, buffer *GPUBuffer, offset uint32, draw_count uint32)  {
//	sdlDrawGPUPrimitivesIndirect(render_pass, buffer, offset, draw_count)
// }

// func EndGPUComputePass(compute_pass *GPUComputePass)  {
//	sdlEndGPUComputePass(compute_pass)
// }

// func EndGPUCopyPass(copy_pass *GPUCopyPass)  {
//	sdlEndGPUCopyPass(copy_pass)
// }

// func EndGPURenderPass(render_pass *GPURenderPass)  {
//	sdlEndGPURenderPass(render_pass)
// }

// func GenerateMipmapsForGPUTexture(command_buffer *GPUCommandBuffer, texture *GPUTexture)  {
//	sdlGenerateMipmapsForGPUTexture(command_buffer, texture)
// }

// func GetGPUDeviceDriver(device *GPUDevice) string {
//	return sdlGetGPUDeviceDriver(device)
// }

// func GetGPUDriver(index int32) string {
//	return sdlGetGPUDriver(index)
// }

// func GetGPUShaderFormats(device *GPUDevice) GPUShaderFormat {
//	return sdlGetGPUShaderFormats(device)
// }

// func GetGPUSwapchainTextureFormat(device *GPUDevice, window *Window) GPUTextureFormat {
//	return sdlGetGPUSwapchainTextureFormat(device, window)
// }

// func GetNumGPUDrivers() int32 {
//	return sdlGetNumGPUDrivers()
// }

// func GPUSupportsProperties(props PropertiesID) bool {
//	return sdlGPUSupportsProperties(props)
// }

// func GPUSupportsShaderFormats(format_flags GPUShaderFormat, name string) bool {
//	return sdlGPUSupportsShaderFormats(format_flags, name)
// }

// func GPUTextureFormatTexelBlockSize(format GPUTextureFormat) uint32 {
//	return sdlGPUTextureFormatTexelBlockSize(format)
// }

// func GPUTextureSupportsFormat(device *GPUDevice, format GPUTextureFormat, type GPUTextureType, usage GPUTextureUsageFlags) bool {
//	return sdlGPUTextureSupportsFormat(device, format, type, usage)
// }

// func GPUTextureSupportsSampleCount(device *GPUDevice, format GPUTextureFormat, sample_count GPUSampleCount) bool {
//	return sdlGPUTextureSupportsSampleCount(device, format, sample_count)
// }

// func InsertGPUDebugLabel(command_buffer *GPUCommandBuffer, text string)  {
//	sdlInsertGPUDebugLabel(command_buffer, text)
// }

// func MapGPUTransferBuffer(device *GPUDevice, transfer_buffer *GPUTransferBuffer, cycle bool) unsafe.Pointer {
//	return sdlMapGPUTransferBuffer(device, transfer_buffer, cycle)
// }

// func PopGPUDebugGroup(command_buffer *GPUCommandBuffer)  {
//	sdlPopGPUDebugGroup(command_buffer)
// }

// func PushGPUComputeUniformData(command_buffer *GPUCommandBuffer, slot_index uint32, data unsafe.Pointer, length uint32)  {
//	sdlPushGPUComputeUniformData(command_buffer, slot_index, data, length)
// }

// func PushGPUDebugGroup(command_buffer *GPUCommandBuffer, name string)  {
//	sdlPushGPUDebugGroup(command_buffer, name)
// }

// func PushGPUFragmentUniformData(command_buffer *GPUCommandBuffer, slot_index uint32, data unsafe.Pointer, length uint32)  {
//	sdlPushGPUFragmentUniformData(command_buffer, slot_index, data, length)
// }

// func PushGPUVertexUniformData(command_buffer *GPUCommandBuffer, slot_index uint32, data unsafe.Pointer, length uint32)  {
//	sdlPushGPUVertexUniformData(command_buffer, slot_index, data, length)
// }

// func QueryGPUFence(device *GPUDevice, fence *GPUFence) bool {
//	return sdlQueryGPUFence(device, fence)
// }

// func ReleaseGPUBuffer(device *GPUDevice, buffer *GPUBuffer)  {
//	sdlReleaseGPUBuffer(device, buffer)
// }

// func ReleaseGPUComputePipeline(device *GPUDevice, compute_pipeline *GPUComputePipeline)  {
//	sdlReleaseGPUComputePipeline(device, compute_pipeline)
// }

// func ReleaseGPUFence(device *GPUDevice, fence *GPUFence)  {
//	sdlReleaseGPUFence(device, fence)
// }

// func ReleaseGPUGraphicsPipeline(device *GPUDevice, graphics_pipeline *GPUGraphicsPipeline)  {
//	sdlReleaseGPUGraphicsPipeline(device, graphics_pipeline)
// }

// func ReleaseGPUSampler(device *GPUDevice, sampler *GPUSampler)  {
//	sdlReleaseGPUSampler(device, sampler)
// }

// func ReleaseGPUShader(device *GPUDevice, shader *GPUShader)  {
//	sdlReleaseGPUShader(device, shader)
// }

// func ReleaseGPUTexture(device *GPUDevice, texture *GPUTexture)  {
//	sdlReleaseGPUTexture(device, texture)
// }

// func ReleaseGPUTransferBuffer(device *GPUDevice, transfer_buffer *GPUTransferBuffer)  {
//	sdlReleaseGPUTransferBuffer(device, transfer_buffer)
// }

// func ReleaseWindowFromGPUDevice(device *GPUDevice, window *Window)  {
//	sdlReleaseWindowFromGPUDevice(device, window)
// }

// func SetGPUAllowedFramesInFlight(device *GPUDevice, allowed_frames_in_flight uint32) bool {
//	return sdlSetGPUAllowedFramesInFlight(device, allowed_frames_in_flight)
// }

// func SetGPUBlendConstants(render_pass *GPURenderPass, blend_constants FColor)  {
//	sdlSetGPUBlendConstants(render_pass, blend_constants)
// }

// func SetGPUBufferName(device *GPUDevice, buffer *GPUBuffer, text string)  {
//	sdlSetGPUBufferName(device, buffer, text)
// }

// func SetGPUScissor(render_pass *GPURenderPass, scissor *Rect)  {
//	sdlSetGPUScissor(render_pass, scissor)
// }

// func SetGPUStencilReference(render_pass *GPURenderPass, reference uint8)  {
//	sdlSetGPUStencilReference(render_pass, reference)
// }

// func SetGPUSwapchainParameters(device *GPUDevice, window *Window, swapchain_composition GPUSwapchainComposition, present_mode GPUPresentMode) bool {
//	return sdlSetGPUSwapchainParameters(device, window, swapchain_composition, present_mode)
// }

// func SetGPUTextureName(device *GPUDevice, texture *GPUTexture, text string)  {
//	sdlSetGPUTextureName(device, texture, text)
// }

// func SetGPUViewport(render_pass *GPURenderPass, viewport *GPUViewport)  {
//	sdlSetGPUViewport(render_pass, viewport)
// }

// func SubmitGPUCommandBuffer(command_buffer *GPUCommandBuffer) bool {
//	return sdlSubmitGPUCommandBuffer(command_buffer)
// }

// func SubmitGPUCommandBufferAndAcquireFence(command_buffer *GPUCommandBuffer) *GPUFence {
//	return sdlSubmitGPUCommandBufferAndAcquireFence(command_buffer)
// }

// func UnmapGPUTransferBuffer(device *GPUDevice, transfer_buffer *GPUTransferBuffer)  {
//	sdlUnmapGPUTransferBuffer(device, transfer_buffer)
// }

// func UploadToGPUBuffer(copy_pass *GPUCopyPass, source *GPUTransferBufferLocation, destination *GPUBufferRegion, cycle bool)  {
//	sdlUploadToGPUBuffer(copy_pass, source, destination, cycle)
// }

// func UploadToGPUTexture(copy_pass *GPUCopyPass, source *GPUTextureTransferInfo, destination *GPUTextureRegion, cycle bool)  {
//	sdlUploadToGPUTexture(copy_pass, source, destination, cycle)
// }

// func WaitAndAcquireGPUSwapchainTexture(command_buffer *GPUCommandBuffer, window *Window, swapchain_texture **GPUTexture, swapchain_texture_width *uint32, swapchain_texture_height *uint32) bool {
//	return sdlWaitAndAcquireGPUSwapchainTexture(command_buffer, window, swapchain_texture, swapchain_texture_width, swapchain_texture_height)
// }

// func WaitForGPUFences(device *GPUDevice, wait_all bool, fences **GPUFence, num_fences uint32) bool {
//	return sdlWaitForGPUFences(device, wait_all, fences, num_fences)
// }

// func WaitForGPUIdle(device *GPUDevice) bool {
//	return sdlWaitForGPUIdle(device)
// }

// func WaitForGPUSwapchain(device *GPUDevice, window *Window) bool {
//	return sdlWaitForGPUSwapchain(device, window)
// }

// func WindowSupportsGPUPresentMode(device *GPUDevice, window *Window, present_mode GPUPresentMode) bool {
//	return sdlWindowSupportsGPUPresentMode(device, window, present_mode)
// }

// func WindowSupportsGPUSwapchainComposition(device *GPUDevice, window *Window, swapchain_composition GPUSwapchainComposition) bool {
//	return sdlWindowSupportsGPUSwapchainComposition(device, window, swapchain_composition)
// }
