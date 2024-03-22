package common

import "flag"

func Flag(args *Args) {
	flag.StringVar(&args.Sm4key, "k", "", "")
	flag.StringVar(&args.FileName, "f", "", "")
	flag.Parse()
}
