package sdl

type PenAxis uint32

const (
	PenAxisPressure PenAxis = iota
	PenAxisXTilt
	PenAxisYTilt
	PenAxisDistance
	PenAxisRotation
	PenAxisSlider
	PenAxisTangentialPressure
	PenAxisCount
)

type PenID uint32

type PenInputFlags uint32

const (
	PenInputDown PenInputFlags = 1 << iota
	PenInputButton1
	PenInputButton2
	PenInputButton3
	PenInputButton4
	PenInputButton5
	PenInputEraserTip PenInputFlags = 1 << 30
)
