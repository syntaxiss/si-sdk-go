package autenticacion

type Request struct {
	Username  string `json:"username"`
	AccessKey string `json:"access_key"`
}
