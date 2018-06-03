package kong

import (
	"os"
	"strings"
)

type ResolverInfo struct {
	Flag func(flag *Flag) (string, bool)
}

//func JSONResolver(r io.Reader) *ResolverInfo {
//	return &ResolverInfo{
//		Grammar: func(grammar interface{}) error {
//			return json.NewDecoder(r).Decode(grammar)
//		},
//	}
//}

func EnvResolver(prefix string) *ResolverInfo {
	return &ResolverInfo{
		Flag: func(flag *Flag) (string, bool) {
			return os.LookupEnv(envString(prefix, flag))
		},
	}
}

func envString(prefix string, flag *Flag) string {
	//if env, ok := flag.Tag.Get("env"); ok {
	//	return env
	//}

	env := strings.ToUpper(flag.Name)
	env = strings.Replace(env, "-", "_", -1)
	env = prefix + env

	return env
}
