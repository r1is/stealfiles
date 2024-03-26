package common

type Args struct {
	Passcode string
	FileName string
}

type OssCfg struct {
	TmpSecretID  string `json:"TmpSecretId,omitempty"`
	TmpSecretKey string `json:"TmpSecretKey,omitempty"`
	SessionToken string `json:"Token,omitempty"`
	BucketURL    string `json:"BucketURL"`
	BatchURL     string `json:"BatchURL"`
}
type Data struct {
	Code string `json:"code"`
}

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
