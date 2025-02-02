package sdl

type BlendOperation uint32

const (
	BlendOperationAdd BlendOperation = iota + 1
	BlendOperationSubtract
	BlendOperationRevSubtract
	BlendOperationMinimum
	BlendOperationMaximum
)

type BlendFactor uint32

const (
	BlendFactorZero BlendFactor = iota + 1
	BlendFactorOne
	BlendFactorSrcColor
	BlendFactorOneMinusSrcColor
	BlendFactorSrcAlpha
	BlendFactorOneMinusSrcAlpha
	BlendFactorDstColor
	BlendFactorOneMinusDstColor
	BlendFactorDstAlpha
	BlendFactorOneMinusDstAlpha
)

// func ComposeCustomBlendMode(srcColorFactor BlendFactor, dstColorFactor BlendFactor, colorOperation BlendOperation, srcAlphaFactor BlendFactor, dstAlphaFactor BlendFactor, alphaOperation BlendOperation) BlendMode {
//	return sdlComposeCustomBlendMode(srcColorFactor, dstColorFactor, colorOperation, srcAlphaFactor, dstAlphaFactor, alphaOperation)
// }
