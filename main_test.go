package main

// import "testing"

// func TestHexToDec(t *testing.T) {
// 	testCases := map[string]struct {
// 		hex      string
// 		expected string
// 	}{
// 		"valid_hex": {
// 			hex:      "1E",
// 			expected: "30",
// 		},
// 		"invalid_hex": {
// 			hex:      "XYZ",
// 			expected: "",
// 		},
// 	}

// 	for name, tc := range testCases {
// 		t.Run(name, func(t *testing.T) {
// 			actual := hexToDec(tc.hex)
// 			if actual != tc.expected {
// 				t.Errorf("Expected decimal value %s for hex %s, but got %s", tc.expected, tc.hex, actual)
// 			}
// 		})
// 	}
// }

// func TestBinToDec(t *testing.T) {
// 	testCases := map[string]struct {
// 		bin      string
// 		expected string
// 	}{
// 		"valid_bin": {
// 			bin:      "1010",
// 			expected: "10",
// 		},
// 		"invalid_bin": {
// 			bin:      "XYZ",
// 			expected: "",
// 		},
// 	}

// 	for name, tc := range testCases {
// 		t.Run(name, func(t *testing.T) {
// 			actual := binToDec(tc.bin)
// 			if actual != tc.expected {
// 				t.Errorf("Expected decimal value %s for binary %s, but got %s", tc.expected, tc.bin, actual)
// 			}
// 		})
// 	}
// }

// func TestModifyByNumber(t *testing.T) {
// 	testCases := map[string]struct {
// 		word           string
// 		action         string
// 		number         string
// 		expectedOutput string
// 	}{
// 		"lowercase": {
// 			word:           "This is a TEST",
// 			action:         "(low)",
// 			number:         "1",
// 			expectedOutput: "this is a test",
// 		},
// 		"uppercase": {
// 			word:           "This is a test",
// 			action:         "(up)",
// 			number:         "2",
// 			expectedOutput: "This is A TEST",
// 		},
// 		"capitalize": {
// 			word:           "this is a test",
// 			action:         "(cap)",
// 			number:         "3",
// 			expectedOutput: "this Is A Test",
// 		},
// 	}

// 	for name, tc := range testCases {
// 		t.Run(name, func(t *testing.T) {
// 			actualOutput := modifyByNumber(tc.word, tc.action, tc.number)
// 			if actualOutput != tc.expectedOutput {
// 				t.Errorf("Expected modified word '%s', but got '%s'", tc.expectedOutput, actualOutput)
// 			}
// 		})
// 	}
// }

// func TestModifyText(t *testing.T) {
// 	inputText := `1E (hex) files were added
// It has been 10 (bin) years
// Ready, set, go (up) !
// I should stop SHOUTING (low)
// Welcome to the Brooklyn bridge (cap)
// This is so exciting (up, 2)
// I was sitting over there ,and then BAMM !!
// I was thinking ... You were right
// As Elton John said: ' I am the most well-known homosexual in the world '
// There it was. A amazing rock!`

// 	expectedOutput := `30 files were added
// It has been 2 years
// Ready, set, GO!
// I should stop shouting
// Welcome to the Brooklyn Bridge
// This is SO EXCITING!
// I was sitting over there, and then BAMM!
// I was thinking... You were right
// As Elton John said: 'I am the most well-known homosexual in the world'
// There it was. An amazing rock!`

// 	actualOutput := modifyText(inputText)

// 	if actualOutput != expectedOutput {
// 		t.Errorf("Expected modified text:\n%s\n\nBut got:\n%s", expectedOutput, actualOutput)
// 	}
// }
