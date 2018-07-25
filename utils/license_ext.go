package utils

const (
	// licenses data
	LicenseTxt      = "license.txt"
	LicenseDot      = "license."
	LicenseMd       = "license.md"
	License         = "license"
	Bsdl            = "bsdl"
	Copying         = "copying"
	COPYINGv3       = "COPYINGv3"
	CopyingDot      = "copying."
	CopyingDash     = "copying-"
	Legal           = "legal"
	Readme          = "readme"
	Copyright       = "copyright"
	Ftl             = "ftl.txt"
	GPLv2           = "gplv2.TXT"
	gpl20           = "gpl-2.0"
	MITtxt          = "mit.txt"
	LisenseRst      = "lisense.rst"
	LisenceHtml     = "lisence.html"
	LicenseHtml     = "license.html"
	Licenses2       = "LICENSE-2.0.txt"
	LICENSE_UNKNOWN = "Unknown"
	Licenses        = "/licenses/"
	EMPTY_STRING    = ""
)

func GetLicensesFiles() []string {
	return []string{LicenseTxt, LicenseMd, License, Copying, Legal, COPYINGv3,
		Readme, Ftl, GPLv2, gpl20, Bsdl, Copyright, MITtxt, LisenseRst,
		LisenceHtml, LicenseHtml, Licenses2}
}

func GetLicensesFilesPrefix() []string {
	return []string{LicenseDot, CopyingDot, Copyright, CopyingDash}
}
