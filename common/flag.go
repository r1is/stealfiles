package common

import "flag"

func Flag(args *Args) {
	flag.StringVar(&args.Passcode, "p", "", "")
	flag.StringVar(&args.FileName, "f", "", "")
	flag.Parse()
}
