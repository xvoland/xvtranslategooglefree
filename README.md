# Google Translate FREE in golang

![Go](https://github.com/xvoland/xvtranslategooglefree/actions/workflows/go.yml/badge.svg)


Package xvtranslategooglefree provides functionality for translating text without using the Google Translate API

## Description:
This package allows you to translate text from one language to another without using the Google Translate API.
It supports various language pairs and includes error handling for common issues such as invalid language codes.

# How To Use

```
import (
	translate "v1/libs/xvtranslategooglefree"
)

    translatedText, err := translate.Translate("Hello", "en", "es")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(translatedText) // Output: Hola
```
