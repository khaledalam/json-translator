## JsonTranslator

#### Simple GoLang script that helps with translate json files using Google Translator.

## Usage:
in cli folder

```shell
# [file name] [Lang from] [Lang to]

$ go run main.go test.json en de
```

### Original test.json file

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


### Testing

```shell
$ go test
```

```
2023/06/11 14:09:37 Done:  1 / 2
2023/06/11 14:09:37 Done:  2 / 2
PASS
ok      github.com/khaledalam/json-translator   1.558s
```