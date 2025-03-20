# Google Translate FREE in golang

![Go](https://github.com/xvoland/xvtranslategooglefree/actions/workflows/go.yml/badge.svg)


Package xvtranslategooglefree provides functionality for translating text without using the Google Translate API

## Description:
This package allows you to translate text from one language to another without using the Google Translate API.
It supports various language pairs and includes error handling for common issues such as invalid language codes.

# How To Use

```golang
import (
	translate "github.com/xvoland/xvtranslategooglefree"
)

    translatedText, err := translate.Translate("Glory to Ukraine", "en", "uk")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(translatedText) // Output: Слава Україні
```

<br /><br />
### Donation

No matter what, I’ll keep working on improving the app because I love seeing people benefit from it and achieve their goals. But even a $1 donation can make a big impact. It helps cover essentials like hosting costs and the time I dedicate to coding. Your support would be incredibly appreciated and would mean so much to me. Thank you!

<p align="center">
  <a href="https://paypal.me/xvoland" target="blank"><img align="center" src="https://raw.githubusercontent.com/xvoland/xvoland/main/images/paypal.png" alt="PayPal" width="250" /></a>
</p>