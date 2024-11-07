package srcinfo

import (
	"bytes"
)

func appendHeader(buffer *bytes.Buffer, key string, value string) {
	if value == "" {
		return
	}

	buffer.WriteString(key + " = " + value + "\n")
}

func appendValue(buffer *bytes.Buffer, key string, value string) {
	if value == "" {
		return
	}

	if value == EmptyOverride {
		value = ""
	}

	buffer.WriteString("\t" + key + " = " + value + "\n")
}

func appendMultiValue(buffer *bytes.Buffer, key string, values []string) {
	for _, value := range values {
		if value == EmptyOverride {
			value = ""
		}

		buffer.WriteString("\t" + key + " = " + value + "\n")
	}
}

func appendMultiArchValue(buffer *bytes.Buffer, key string, values []ArchDistroString) {
	for _, value := range values {
		if value.Value == EmptyOverride {
			value.Value = ""
		}

		buffer.WriteString("\t" + key)
		if value.Distro != "" {
			buffer.WriteString("_" + value.Distro)
		}

		if value.Arch != "" {
			buffer.WriteString("_" + value.Arch)
		}

		buffer.WriteString(" = " + value.Value + "\n")
	}
}

// String generates a string that should be similar to the srcinfo data used to
// create this Srcinfo struct. Fields will be printed in order and with the same
// whitespace rules that `makepkg --printsrcinfo` uses.
//
// The order of each global field is as follows:
//
//	pkgdesc
//	pkgver
//	pkgrel
//	epoch
//	url
//	install
//	changelog
//	arch
//	groups
//	license
//	checkdepends
//	makedepends
//	depends
//	optdepends
//	provides
//	conflicts
//	replaces
//	noextract
//	options
//	backup
//	source
//	validpgpkeys
//	md5suns
//	sha1sums
//	sha224sums
//	sha256sums
//	sha384sums
//	sha512sums
//
// The order of each overwritten field is as follows:
//
//	pkgdesc
//	url
//	install
//	changelog
//	arch
//	groups
//	license
//	checkdepends
//	depends
//	optdepends
//	provides
//	conflicts
//	replaces
//	options
//	backup
func (si *Srcinfo) String() string {
	var buffer bytes.Buffer

	appendHeader(&buffer, "pkgbase", si.Pkgbase)

	appendValue(&buffer, "pkgdesc", si.Pkgdesc)
	appendValue(&buffer, "pkgver", si.Pkgver)
	appendValue(&buffer, "pkgrel", si.Pkgrel)
	appendValue(&buffer, "epoch", si.Epoch)
	appendValue(&buffer, "url", si.URL)
	appendValue(&buffer, "priority", si.Priority)
	appendMultiValue(&buffer, "arch", si.Arch)
	appendMultiValue(&buffer, "license", si.License)
	appendMultiArchValue(&buffer, "gives", si.Gives)
	appendMultiArchValue(&buffer, "depends", si.Depends)
	appendMultiArchValue(&buffer, "checkdepends", si.CheckDepends)
	appendMultiArchValue(&buffer, "makedepends", si.MakeDepends)
	appendMultiArchValue(&buffer, "optdepends", si.OptDepends)
	appendMultiArchValue(&buffer, "pacdeps", si.Pacdeps)
	appendMultiArchValue(&buffer, "checkconflicts", si.CheckConflicts)
	appendMultiArchValue(&buffer, "makeconflicts", si.MakeConflicts)
	appendMultiArchValue(&buffer, "conflicts", si.Conflicts)
	appendMultiArchValue(&buffer, "provides", si.Provides)
	appendMultiArchValue(&buffer, "breaks", si.Breaks)
	appendMultiArchValue(&buffer, "replaces", si.Replaces)
	appendMultiArchValue(&buffer, "enhances", si.Enhances)
	appendMultiArchValue(&buffer, "recommends", si.Recommends)
	appendMultiArchValue(&buffer, "suggests", si.Suggests)
	appendMultiValue(&buffer, "mask", si.Mask)
	appendMultiValue(&buffer, "compatible", si.Compatible)
	appendMultiValue(&buffer, "incompatible", si.Incompatible)
	appendMultiValue(&buffer, "maintainer", si.Maintainer)
	appendMultiArchValue(&buffer, "source", si.Source)
	appendMultiValue(&buffer, "noextract", si.NoExtract)
	appendMultiValue(&buffer, "nosubmodules", si.NoSubmodules)
	appendMultiArchValue(&buffer, "md5sums", si.MD5Sums)
	appendMultiArchValue(&buffer, "sha1sums", si.SHA1Sums)
	appendMultiArchValue(&buffer, "sha224sums", si.SHA224Sums)
	appendMultiArchValue(&buffer, "sha256sums", si.SHA256Sums)
	appendMultiArchValue(&buffer, "sha384sums", si.SHA384Sums)
	appendMultiArchValue(&buffer, "sha512sums", si.SHA512Sums)
	appendMultiValue(&buffer, "backup", si.Backup)
	appendMultiValue(&buffer, "repology", si.Repology)

	for n, pkg := range si.Packages {
		appendHeader(&buffer, "\npkgname", si.Packages[n].Pkgname)

		appendValue(&buffer, "pkgdesc", pkg.Pkgdesc)
		appendValue(&buffer, "url", pkg.URL)
		appendValue(&buffer, "priority", pkg.Priority)
		appendMultiValue(&buffer, "arch", pkg.Arch)
		appendMultiValue(&buffer, "license", pkg.License)
		appendMultiArchValue(&buffer, "gives", pkg.Gives)
		appendMultiArchValue(&buffer, "depends", pkg.Depends)
		appendMultiArchValue(&buffer, "checkdepends", si.CheckDepends)
		appendMultiArchValue(&buffer, "optdepends", pkg.OptDepends)
		appendMultiArchValue(&buffer, "pacdeps", pkg.Pacdeps)
		appendMultiArchValue(&buffer, "checkconflicts", pkg.CheckConflicts)
		appendMultiArchValue(&buffer, "conflicts", pkg.Conflicts)
		appendMultiArchValue(&buffer, "provides", pkg.Provides)
		appendMultiArchValue(&buffer, "breaks", pkg.Breaks)
		appendMultiArchValue(&buffer, "replaces", pkg.Replaces)
		appendMultiArchValue(&buffer, "enhances", pkg.Enhances)
		appendMultiArchValue(&buffer, "recommends", pkg.Recommends)
		appendMultiArchValue(&buffer, "suggests", pkg.Suggests)
		appendMultiValue(&buffer, "backup", pkg.Backup)
		appendMultiValue(&buffer, "repology", pkg.Repology)
	}

	return buffer.String()
}
