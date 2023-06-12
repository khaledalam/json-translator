## JsonTranslator

#### Simple GoLang script that helps with translate json files using Google Translator.

## Usage:
in cli folder

```shell
# [file name] [interval in millisecond] [Language from] [Languages to]

$ go run cli/main.go $(pwd)/cli/test.json en de
```

### Single translate language

```bash
go run main.go test.json en de
```

Input:

```json
{
  "Home": "Home",
  "Account": "Account"
}
```

Output:

```json
{
  "Account": "Konto",
  "Home": "Heim"
}
```


### Multiple translate languages

```bash
go run main.go test.json en de it
```

Input:
```json
{
  "Home": "Home",
  "Account": "Account"
}
```

Output:

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


### Testing

```shell
$ go test
```

```
[en => de] Done:  1 / 2
[en => de] Done:  2 / 2
-----
[en => de] Done:  1 / 2
[en => de] Done:  2 / 2
[en => it] Done:  1 / 2
[en => it] Done:  2 / 2
-----
PASS
ok      github.com/khaledalam/json-translator   2.144s
```