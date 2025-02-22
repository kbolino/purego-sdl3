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

type BlendMode uint32

const (
	BlendModeNone               = 0x00000000
	BlendModeBlend              = 0x00000001
	BlendModeBlendPremultiplied = 0x00000010
	BlendModeAdd                = 0x00000002
	BlendModeAddPremultiplied   = 0x00000020
	BlendModeMod                = 0x00000004
	BlendModeMul                = 0x00000008
	BlendModeInvalid            = 0x7FFFFFFF
)

// func ComposeCustomBlendMode(srcColorFactor BlendFactor, dstColorFactor BlendFactor, colorOperation BlendOperation, srcAlphaFactor BlendFactor, dstAlphaFactor BlendFactor, alphaOperation BlendOperation) BlendMode {
//	return sdlComposeCustomBlendMode(srcColorFactor, dstColorFactor, colorOperation, srcAlphaFactor, dstAlphaFactor, alphaOperation)
// }
