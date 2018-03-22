package args

import "os"

func GetCommandArg(arg string) string {
	for idx, s := range os.Args {
		if s == "-"+arg || s == "--"+arg {
			return os.Args[idx+1]
		}
	}
	return ""
}
