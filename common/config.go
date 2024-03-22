package common

var version = "1.0.0"

type Args struct {
	Sm4key   string
	FileName string
}

type OssCfg struct {
	Ak        string `json:"ak"`
	Sk        string `json:"sk"`
	BucketURL string `json:"BucketURL"`
	BatchURL  string `json:"BatchURL"`
}
