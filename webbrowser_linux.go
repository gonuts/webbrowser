package webbrowser

func init() {
	Register(GenericBrowser{"/bin/sh", []string{"-c", "xdg-open"}})
}
