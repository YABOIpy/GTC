package utils

type Util interface {
	WriteFileArray(files string, item []string)
	WriteFile(files, item string) ([]string, error)
	HalfToken(T string, v int) string
	ReadFile(path string) ([]string, error)
	Input(text string) string
}

const (
	LoadingChace      = "Loading Cache: \033[36m[\u001B[39m%s\u001B[36m]\u001B[39m"
	CheckerFormat     = "[\u001B[32m✓\033[39m] (TIME): %sMs (\033[33mLOCKED\033[39m): %v (\033[31mINVALID\033[39m): %v (\033[32mVALID\033[39m): %v (TOTAL): %v \n"
	WriteValidMention = "\u001B[30;1m[!] Write Unlocked Tokens To Tokens.txt? y/n:\u001B[0m"
)
