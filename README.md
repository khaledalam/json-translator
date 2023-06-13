## JsonTranslator

#### Simple Go module that helps with translating json files using Google Translator.

[![Go Report Card](https://goreportcard.com/badge/github.com/khaledalam/json-translator)](https://goreportcard.com/report/github.com/khaledalam/json-translator)
[![Go Reference](https://pkg.go.dev/badge/github.com/khaledalam/json-translator?status.svg)](https://pkg.go.dev/github.com/khaledalam/json-translator?status)
![Test](https://github.com/khaledalam/json-translator/workflows/Test/badge.svg)


### Install:

```shell
go get github.com/khaledalam/json-translator
```

### Args:

- [x] `[source json file path]`
- [x] `[interval in millisecond]` 
- [x] `[Language_from]`
- [x] `[Language_to 1] [Language_to 2] [Language_to n..]`

### Flags:
- [ ] --output=files
---

### Cli Examples:


&rightarrow; input file `test.json`:
```json
{
    "Home": "Home",
    "Account": "Account"
}
```

<div style="text-align: center">
&downarrow; &downarrow; &downarrow; &downarrow;


</div>


```bash
go run main.go $(pwd)/cli/test.json 5 en de
```

&rightarrow; output file from &uparrow; command `Translated_en_de.json`:
```json
{
    "Account": "Konto",
    "Home": "Heim"
}
```

---

```bash
go run main.go $(pwd)/cli/test.json 5 en de it
```
&rightarrow; output file from &uparrow; command `Translated_en_de_it.json`:
```json
{
    "de": {
        "Account": "Konto",
        "Home": "Heim"
    },
    "it": {
        "Account": "Account",
        "Home": "Casa"
    }
}
```

---
### Testing

```shell
make test
```

```
2023/06/13 04:31:32 [en => de] Done:  1 / 2
2023/06/13 04:31:33 [en => de] Done:  2 / 2
2023/06/13 04:31:33 -----
2023/06/13 04:31:33 [en => de] Done:  1 / 2
2023/06/13 04:31:33 [en => de] Done:  2 / 2
2023/06/13 04:31:33 [en => it] Done:  1 / 2
2023/06/13 04:31:33 [en => it] Done:  2 / 2
2023/06/13 04:31:33 -----
2023/06/13 04:31:33 Error: when reading json file. open invalid/json/file/path/a.json: no such file or directory
2023/06/13 04:31:33 Error: when reading json file. [multi]
2023/06/13 04:31:33 -----
PASS
	github.com/khaledalam/json-translator	coverage: 70.8% of statements
ok  	github.com/khaledalam/json-translator	0.682s
```