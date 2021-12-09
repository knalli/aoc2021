package day08

func buildDigitSpecs() []DigitType {
	result := make([]DigitType, 0)
	result = append(result, DigitType{
		Value: 0,
		Signals: map[int32]bool{
			'a': true,
			'c': true,
			'f': true,
			'g': true,
			'e': true,
			'b': true,
		},
	})
	result = append(result, DigitType{
		Value: 1,
		Signals: map[int32]bool{
			'c': true,
			'f': true,
		},
	})
	result = append(result, DigitType{
		Value: 2,
		Signals: map[int32]bool{
			'a': true,
			'c': true,
			'g': true,
			'e': true,
			'd': true,
		},
	})
	result = append(result, DigitType{
		Value: 3,
		Signals: map[int32]bool{
			'a': true,
			'c': true,
			'f': true,
			'g': true,
			'd': true,
		},
	})
	result = append(result, DigitType{
		Value: 4,
		Signals: map[int32]bool{
			'c': true,
			'f': true,
			'd': true,
			'b': true,
		},
	})
	result = append(result, DigitType{
		Value: 5,
		Signals: map[int32]bool{
			'a': true,
			'f': true,
			'g': true,
			'd': true,
			'b': true,
		},
	})
	result = append(result, DigitType{
		Value: 6,
		Signals: map[int32]bool{
			'a': true,
			'f': true,
			'g': true,
			'e': true,
			'd': true,
			'b': true,
		},
	})
	result = append(result, DigitType{
		Value: 7,
		Signals: map[int32]bool{
			'a': true,
			'c': true,
			'f': true,
		},
	})
	result = append(result, DigitType{
		Value: 8,
		Signals: map[int32]bool{
			'a': true,
			'c': true,
			'f': true,
			'g': true,
			'e': true,
			'd': true,
			'b': true,
		},
	})
	result = append(result, DigitType{
		Value: 9,
		Signals: map[int32]bool{
			'a': true,
			'c': true,
			'f': true,
			'g': true,
			'd': true,
			'b': true,
		},
	})
	return result
}
