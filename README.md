# ASCII Art Web

A browser interface for generating ASCII banners. Enter a word or several lines of text, choose **Standard**, **Shadow**, or **Thinkertoy**, and submit the form. A Go HTTP server renders the banner and returns it as HTML; the browser does not need JavaScript to draw the result.

## Run locally

Requires **Go 1.22.5 or later**. The fonts, templates, and CSS are included in the repository; no external service is needed.

**Build note:** The renderer import is commented out in [main.go](Ascii-art-web/main.go). Restore `"web/assci"` in the import block before running the server. The directory is named `assci`, but it declares `package ascci`, which is why the handler calls `ascci.Makeandreturn`. Without the import, the build stops with `undefined: ascci`.

From the repository root:

```sh
cd Ascii-art-web
go run .
```

Open [http://localhost:7090](http://localhost:7090). Try entering `Hello` and `Go` on separate lines, choose **Shadow**, and select **Send**. Use the Enter key for a new line rather than typing `\n`. To compare fonts, submit the same text again with **Thinkertoy** selected.

## How a request becomes a banner

1. `GET /` serves the [input form](Ascii-art-web/template/HELLO.html). **Standard** is selected by default.
2. The form sends the text (`name`) and selected font filename (`asciid`) to `POST /submit`. The handler returns `400 Bad Request` for empty text and `405 Method Not Allowed` for non-POST requests to `/submit`.
3. The [renderer](Ascii-art-web/assci/ascii.go) reads the selected font file and joins matching glyph rows across each line of text. The handler passes those rows to a Go HTML template, which displays them in a `<pre>` block so spacing is preserved.
