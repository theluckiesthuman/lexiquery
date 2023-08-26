# LexiQuery
 A CLI tool to query merriam-webster dictionary. Currently, the tool only supports querying the dictionary. The tool is written in `go` ans uses inbuilt command line tools for CLI. The tool is tested on `macOS 13.4.1 (22F82) (Apple M1 Pro)`.

## Run Locally
The pre-requisite is to have `go` installed.
We can install `go` by following the instructions [here](https://golang.org/doc/install).
Once `go` is installed, we need to replace the `API_KEY` in the `.env` file with the actual API key. The API key can be obtained by registering on [Merriam-Webster](https://dictionaryapi.com/).
Then we can run the tool by running the following command:
```bash
make run
```

## Installnation
We can install the tool by running the following command. The pr
```bash
make install
```


## Note
The tool is still under development. The following features are yet to be implemented:
- [ ] Add support for multiple dictionaries.
- [ ] Add support for multiple languages.
- [ ] Add support for multiple thesaurus.
- [ ] Add support for multiple thesaurus languages.
- [ ] Add support for multiple thesaurus dictionaries.
- [ ] Robust error handling.
- [ ] Improve the code quality and add more tests.
- [ ] Improve the user experience.