package main

import "testing"

func TestGreet(t *testing.T) {
	type testCase struct {
		lang             language
		expectedGreeting string
	}

	var tests = map[string]testCase{
		"English": {
			lang:             "en",
			expectedGreeting: "Hello, World!",
		},
		"French": {
			lang:             "fr",
			expectedGreeting: "Bonjour, le monde!",
		},
		"Akkadian, not supported": {
			lang:             "akk",
			expectedGreeting: `unsupported language: "akk"`,
		},
		"Greek": {
			lang:             "el",
			expectedGreeting: "Χαίρετε Κόσμε!",
		},
		"Hebrew": {
			lang:             "he",
			expectedGreeting: "שלום עולם!",
		},
		"Urdu": {
			lang:             "ur",
			expectedGreeting: "ہیلو دنیا",
		},
		"Vietnamese": {
			lang:             "vi",
			expectedGreeting: "Xin chào Thế Giới",
		},
		"Chinese": {
			lang:             "zh",
			expectedGreeting: "你好，世界！",
		},
		"Spanish": {
			lang:             "es",
			expectedGreeting: "Hola, Mundo!",
		},
		"Italian": {
			lang:             "it",
			expectedGreeting: "Ciao, mondo!",
		},
		"Japanese": {
			lang:             "ja",
			expectedGreeting: "こんにちは世界!",
		},
		"Empty": {
			lang:             "",
			expectedGreeting: `unsupported language: ""`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			greeting := greet(tc.lang)

			if greeting != tc.expectedGreeting {
				t.Errorf("expected: %q, got: %q", tc.expectedGreeting, greeting)
			}
		})
	}
}
