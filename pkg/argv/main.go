package argv

import "os"

func GetArgs() []string {
	args := os.Args
	return args
}
