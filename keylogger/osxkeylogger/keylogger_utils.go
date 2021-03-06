package osxkeylogger

import "github.com/rikonor/keysig/keylogger/keyboard"

var keyCodeConversionTable = map[int]keyboard.Key{
	0x00: keyboard.A,
	0x01: keyboard.S,
	0x02: keyboard.D,
	0x03: keyboard.F,
	0x04: keyboard.H,
	0x05: keyboard.G,
	0x06: keyboard.Z,
	0x07: keyboard.X,
	0x08: keyboard.C,
	0x09: keyboard.V,
	0x0B: keyboard.B,
	0x0C: keyboard.Q,
	0x0D: keyboard.W,
	0x0E: keyboard.E,
	0x0F: keyboard.R,
	0x10: keyboard.Y,
	0x11: keyboard.T,
	0x12: keyboard.One,
	0x13: keyboard.Two,
	0x14: keyboard.Three,
	0x15: keyboard.Four,
	0x16: keyboard.Six,
	0x17: keyboard.Five,
	0x18: keyboard.Equals,
	0x19: keyboard.Nine,
	0x1A: keyboard.Seven,
	0x1B: keyboard.Dash,
	0x1C: keyboard.Eight,
	0x1D: keyboard.Zero,
	0x1E: keyboard.RightBracket,
	0x1F: keyboard.O,
	0x20: keyboard.U,
	0x21: keyboard.LeftBracket,
	0x22: keyboard.I,
	0x23: keyboard.P,
	0x25: keyboard.L,
	0x26: keyboard.J,
	0x27: keyboard.Apostrophe,
	0x28: keyboard.K,
	0x29: keyboard.Semicolon,
	0x2A: keyboard.BackSlash,
	0x2B: keyboard.Comma,
	0x2C: keyboard.ForwardSlash,
	0x2D: keyboard.N,
	0x2E: keyboard.M,
	0x2F: keyboard.Period,
	0x32: keyboard.Tilde,
	0x24: keyboard.Enter,
	0x30: keyboard.Tab,
	0x31: keyboard.Space,
	0x33: keyboard.Delete,
	0x35: keyboard.Escape,
	0x36: keyboard.RightSuper,
	0x37: keyboard.LeftSuper,
	0x38: keyboard.LeftShift,
	0x39: keyboard.CapsLock,
	0x3A: keyboard.LeftAlt,
	0x3B: keyboard.LeftCtrl,
	0x3C: keyboard.RightShift,
	0x3D: keyboard.RightAlt,
	0x3E: keyboard.RightCtrl,
	0x3F: keyboard.Invalid, // Function
	0x40: keyboard.F17,
	0x4F: keyboard.F18,
	0x50: keyboard.F19,
	0x5A: keyboard.F20,
	0x60: keyboard.F5,
	0x61: keyboard.F6,
	0x62: keyboard.F7,
	0x63: keyboard.F3,
	0x64: keyboard.F8,
	0x65: keyboard.F9,
	0x67: keyboard.F11,
	0x69: keyboard.F13,
	0x6A: keyboard.F16,
	0x6B: keyboard.F14,
	0x6D: keyboard.F10,
	0x6F: keyboard.F12,
	0x71: keyboard.F15,
	0x72: keyboard.Help,
	0x73: keyboard.Home,
	0x74: keyboard.PageUp,
	0x75: keyboard.Delete,
	0x76: keyboard.F4,
	0x77: keyboard.End,
	0x78: keyboard.F2,
	0x79: keyboard.PageDown,
	0x7A: keyboard.F1,
	0x7B: keyboard.ArrowLeft,
	0x7C: keyboard.ArrowRight,
	0x7D: keyboard.ArrowDown,
	0x7E: keyboard.ArrowUp,
}

func convertKeyCode(keyCode int) keyboard.Key {
	// Search for keyCode in the conversion table
	k, ok := keyCodeConversionTable[keyCode]
	if !ok {
		return keyboard.Invalid
	}
	return k
}

var stateCodeConversionTable = map[int]keyboard.State{
	0: keyboard.InvalidState,
	1: keyboard.Down,
	2: keyboard.Up,
}

func convertStateCode(stateCode int) keyboard.State {
	// Search for stateCode in the conversion table
	s, ok := stateCodeConversionTable[stateCode]
	if !ok {
		return keyboard.InvalidState
	}
	return s
}
