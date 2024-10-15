package assert

import "log"

func NotError(err error, output string) {
	if err != nil {
		log.Panic(output)
	}
}

func IsEqual(a any, b any, output string) {
	if a != b {
		log.Panic(output)
	}
}

func NotEmptyString(s string) {
	if len(s)<= 0 {
		log.Panic("String is empty")
	}
}