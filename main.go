package cli_frmwk

var DefaultHandler Handler

func init() {
	DefaultHandler = NewHandler("~>")
	DefaultHandler.Init()
}
