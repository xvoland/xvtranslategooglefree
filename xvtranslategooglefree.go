/*
Package xvtranslategooglefree provides functionality for translating text without using the Google Translate API.

Author: Vitalii Tereshchuk
URL:    https://dotoca.net
Created: 04.01.2015
License: MIT

Description:
This package allows you to translate text from one language to another without using the Google Translate API.
It supports various language pairs and includes error handling for common issues such as invalid language codes.

Usage:
Call the Translate function with the source text, source language, and target language.

Example:

import (
	translate "github.com/xvoland/xvtranslategooglefree"
)

    translatedText, err := translate.Translate("Glory to Ukraine", "en", "es")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(translatedText) // Output: Слава Україні

*/

package xvtranslategooglefree

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var supportedLanguages = map[string]bool{
	"ab": true, "ace": true, "ach": true, "af": true, "sq": true, "alz": true, "am": true, "ar": true, "hy": true, "as": true,
	"awa": true, "ay": true, "az": true, "ban": true, "bm": true, "ba": true, "eu": true, "btx": true, "bts": true, "bbc": true,
	"be": true, "bem": true, "bn": true, "bew": true, "bho": true, "bik": true, "bs": true, "br": true, "bg": true, "bua": true,
	"yue": true, "ca": true, "ceb": true, "ny": true, "zh-CN": true, "zh-TW": true, "cv": true, "co": true, "crh": true, "hr": true,
	"cs": true, "da": true, "din": true, "dv": true, "doi": true, "dov": true, "nl": true, "dz": true, "en": true, "eo": true,
	"et": true, "ee": true, "fj": true, "fil": true, "fi": true, "fr": true, "fr-FR": true, "fr-CA": true, "fy": true, "ff": true,
	"gaa": true, "gl": true, "lg": true, "ka": true, "de": true, "el": true, "gn": true, "gu": true, "ht": true, "cnh": true,
	"ha": true, "haw": true, "iw": true, "hil": true, "hi": true, "hmn": true, "hu": true, "hrx": true, "is": true, "ig": true,
	"ilo": true, "id": true, "ga": true, "it": true, "ja": true, "jw": true, "kn": true, "pam": true, "kk": true, "km": true,
	"cgg": true, "rw": true, "ktu": true, "gom": true, "ko": true, "kri": true, "ku": true, "ckb": true, "ky": true, "lo": true,
	"ltg": true, "la": true, "lv": true, "lij": true, "li": true, "ln": true, "lt": true, "lmo": true, "luo": true, "lb": true,
	"mk": true, "mai": true, "mak": true, "mg": true, "ms": true, "ms-Arab": true, "ml": true, "mt": true, "mi": true, "mr": true,
	"chm": true, "mni-Mtei": true, "min": true, "lus": true, "mn": true, "my": true, "nr": true, "new": true, "ne": true, "nso": true,
	"no": true, "nus": true, "oc": true, "or": true, "om": true, "pag": true, "pap": true, "ps": true, "fa": true, "pl": true,
	"pt": true, "pt-PT": true, "pt-BR": true, "pa": true, "pa-Arab": true, "qu": true, "rom": true, "ro": true, "rn": true, "ru": true,
	"sm": true, "sg": true, "sa": true, "gd": true, "sr": true, "st": true, "crs": true, "shn": true, "sn": true, "scn": true,
	"szl": true, "sd": true, "si": true, "sk": true, "sl": true, "so": true, "es": true, "su": true, "sw": true, "ss": true,
	"sv": true, "tg": true, "ta": true, "tt": true, "te": true, "tet": true, "th": true, "ti": true, "ts": true, "tn": true,
	"tr": true, "tk": true, "ak": true, "uk": true, "ur": true, "ug": true, "uz": true, "vi": true, "cy": true, "xh": true,
	"yi": true, "yo": true, "yua": true, "zu": true,
}

var httpClient = &http.Client{
	Timeout: 30 * time.Second, // Set a timeout for the request
}

// EncodeURI encodes the given string for use in a URL
func encodeURI(s string) string {
	return url.QueryEscape(s)
}

// Translate translates a source text from one language to another using Google Translate
func Translate(source, sourceLang, targetLang string) (string, error) {

	// Check for empty source text
	if source == "" {
		return "", fmt.Errorf("source text is empty")
	}

	// Check for valid source and target languages
	if !isValidLanguage(sourceLang) {
		return "", fmt.Errorf("invalid source language: %s", sourceLang)
	}
	if !isValidLanguage(targetLang) {
		return "", fmt.Errorf("invalid target language: %s", targetLang)
	}

	encodedSource := encodeURI(source)
	// client=dict-chrome-ex
	url := "https://translate.googleapis.com/translate_a/single?client=gtx" +
		"&soc-app=1&soc-platform=1&soc-device=1&ie=UTF-8&oe=UTF-8" +
		"&sl=" + sourceLang + "&tl=" + targetLang + "&dt=t&q=" + encodedSource

	// fmt.Println(url)

	// Create the GET request with proper headers
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check for error status codes
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	// Read the response body
	var result []interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling response body: %w", err)
	}

	// Parse the translated text
	return extractTranslation(result)
}

// extractTranslation parses the translation from the response result
func extractTranslation(result []interface{}) (string, error) {
	if len(result) == 0 {
		return "", errors.New("no translation found in the response")
	}

	var translatedText []string
	inner := result[0]
	for _, slice := range inner.([]interface{}) {
		for _, translated := range slice.([]interface{}) {
			translatedText = append(translatedText, fmt.Sprintf("%v", translated))
			break
		}
	}
	if len(translatedText) == 0 {
		return "", errors.New("no valid translated text found")
	}

	return strings.Join(translatedText, ""), nil
}

// Helper function to check if a language is valid
func isValidLanguage(language string) bool {
	_, exists := supportedLanguages[language]
	return exists
}
