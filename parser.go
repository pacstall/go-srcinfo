package srcinfo

import (
	"fmt"
	"os"
	"strings"
)

// parser is used to track our current state as we parse the srcinfo.
type parser struct {
	// srcinfo is a Pointer to the Srcinfo we are currently building.
	srcinfo *Srcinfo

	// seenPkgnames is a set of pkgnames we have seen
	seenPkgnames map[string]struct{}
}

func (psr *parser) currentPackage() (*Package, error) {
	if psr.srcinfo.Pkgbase == "" {
		return nil, fmt.Errorf("Not in pkgbase or pkgname")
	} else if len(psr.srcinfo.Packages) == 0 {
		return &psr.srcinfo.Package, nil
	} else {
		return &psr.srcinfo.Packages[len(psr.srcinfo.Packages)-1], nil
	}
}

func (psr *parser) setHeaderOrField(key, value string) error {
	pkgbase := &psr.srcinfo.PackageBase

	switch key {
	case "pkgbase":
		if psr.srcinfo.Pkgbase != "" {
			return fmt.Errorf("key \"%s\" can not occur after pkgbase or pkgname", key)
		}

		pkgbase.Pkgbase = value
		return nil
	case "pkgname":
		if psr.srcinfo.Pkgbase == "" {
			return fmt.Errorf("key \"%s\" can not occur before pkgbase", key)
		}
		if _, ok := psr.seenPkgnames[value]; ok {
			return fmt.Errorf("pkgname \"%s\" can not occur more than once", value)
		}
		psr.seenPkgnames[value] = struct{}{}

		psr.srcinfo.Packages = append(psr.srcinfo.Packages, Package{Pkgname: value})
		return nil
	}

	if psr.srcinfo.Pkgbase == "" {
		return fmt.Errorf("key \"%s\" can not occur before pkgbase or pkgname", key)
	}

	return psr.setField(key, value)
}

func (psr *parser) setField(archKey, value string) error {
	pkg, err := psr.currentPackage()
	if err != nil {
		return err
	}

	pkgbase := &psr.srcinfo.PackageBase
	key, distro, arch := splitDistroArchFromKey(psr.srcinfo.Arch, archKey)
	err = checkArch(psr.srcinfo.Arch, archKey, arch)
	if err != nil {
		return err
	}

	if value == "" {
		value = EmptyOverride
	}

	// pkgbase only + not arch dependent
	found := true
	switch archKey {
	case "pkgver":
		pkgbase.Pkgver = value
	case "pkgrel":
		pkgbase.Pkgrel = value
	case "epoch":
		pkgbase.Epoch = value
	case "mask":
		pkgbase.Mask = append(pkgbase.Mask, value)
	case "compatible":
		pkgbase.Compatible = append(pkgbase.Compatible, value)
	case "incompatible":
		pkgbase.Incompatible = append(pkgbase.Incompatible, value)
	case "maintainer":
		pkgbase.Maintainer = append(pkgbase.Maintainer, value)
	case "noextract":
		pkgbase.NoExtract = append(pkgbase.NoExtract, value)
	case "nosubmodules":
		pkgbase.NoSubmodules = append(pkgbase.NoSubmodules, value)
	default:
		found = false
	}

	if found {
		if len(psr.srcinfo.Packages) > 0 {
			return fmt.Errorf("key \"%s\" can not occur after pkgname", archKey)
		}

		return nil
	}

	// pkgbase only + arch dependent
	found = true
	switch key {
	case "source":
		pkgbase.Source = append(pkgbase.Source, ArchDistroString{arch, distro, value})
	case "md5sums":
		pkgbase.MD5Sums = append(pkgbase.MD5Sums, ArchDistroString{arch, distro, value})
	case "sha1sums":
		pkgbase.SHA1Sums = append(pkgbase.SHA1Sums, ArchDistroString{arch, distro, value})
	case "sha224sums":
		pkgbase.SHA224Sums = append(pkgbase.SHA224Sums, ArchDistroString{arch, distro, value})
	case "sha256sums":
		pkgbase.SHA256Sums = append(pkgbase.SHA256Sums, ArchDistroString{arch, distro, value})
	case "sha384sums":
		pkgbase.SHA384Sums = append(pkgbase.SHA384Sums, ArchDistroString{arch, distro, value})
	case "sha512sums":
		pkgbase.SHA512Sums = append(pkgbase.SHA512Sums, ArchDistroString{arch, distro, value})
	case "b2sums":
		pkgbase.B2Sums = append(pkgbase.B2Sums, ArchDistroString{arch, distro, value})
	case "makedepends":
		pkgbase.MakeDepends = append(pkgbase.MakeDepends, ArchDistroString{arch, distro, value})
	case "makeconflicts":
		pkgbase.MakeConflicts = append(pkgbase.MakeConflicts, ArchDistroString{arch, distro, value})
	default:
		found = false
	}

	if found {
		if len(psr.srcinfo.Packages) > 0 {
			return fmt.Errorf("key \"%s\" can not occur after pkgname", archKey)
		}

		return nil
	}

	// pkgbase or pkgname + not arch dependent
	found = true
	switch archKey {
	case "pkgdesc":
		pkg.Pkgdesc = value
	case "url":
		pkg.URL = value
	case "priority":
		pkg.Priority = value
	case "arch":
		pkg.Arch = append(pkg.Arch, value)
	case "license":
		pkg.License = append(pkg.License, value)
	case "backup":
		pkg.Backup = append(pkg.Backup, value)
	case "repology":
		pkg.Repology = append(pkg.Repology, value)
	default:
		found = false
	}

	if found {
		return nil
	}

	// pkgbase or pkgname + arch dependent
	switch key {
	case "gives":
		pkg.Gives = append(pkg.Gives, ArchDistroString{arch, distro, value})
	case "depends":
		pkg.Depends = append(pkg.Depends, ArchDistroString{arch, distro, value})
	case "checkdepends":
		pkg.CheckDepends = append(pkg.CheckDepends, ArchDistroString{arch, distro, value})
	case "optdepends":
		pkg.OptDepends = append(pkg.OptDepends, ArchDistroString{arch, distro, value})
	case "pacdeps":
		pkg.Pacdeps = append(pkg.Pacdeps, ArchDistroString{arch, distro, value})
	case "checkconflicts":
		pkg.CheckConflicts = append(pkg.CheckConflicts, ArchDistroString{arch, distro, value})
	case "conflicts":
		pkg.Conflicts = append(pkg.Conflicts, ArchDistroString{arch, distro, value})
	case "provides":
		pkg.Provides = append(pkg.Provides, ArchDistroString{arch, distro, value})
	case "breaks":
		pkg.Breaks = append(pkg.Breaks, ArchDistroString{arch, distro, value})
	case "replaces":
		pkg.Replaces = append(pkg.Replaces, ArchDistroString{arch, distro, value})
	case "enhances":
		pkg.Enhances = append(pkg.Enhances, ArchDistroString{arch, distro, value})
	case "recommends":
		pkg.Recommends = append(pkg.Recommends, ArchDistroString{arch, distro, value})
	case "suggests":
		pkg.Suggests = append(pkg.Suggests, ArchDistroString{arch, distro, value})
	}

	return nil
}

func parse(data string) (*Srcinfo, error) {
	psr := &parser{
		&Srcinfo{},
		make(map[string]struct{}),
	}

	lines := strings.Split(data, "\n")

	for n, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, err := splitPair(line)
		if err != nil {
			return nil, Error(n+1, line, err.Error())
		}

		err = psr.setHeaderOrField(key, value)
		if err != nil {
			return nil, Error(n+1, line, err.Error())
		}
	}

	if psr.srcinfo.Pkgbase == "" {
		return nil, fmt.Errorf("No pkgbase field")
	}

	if len(psr.srcinfo.Packages) == 0 {
		return nil, fmt.Errorf("No pkgname field")
	}

	if psr.srcinfo.Pkgver == "" {
		return nil, fmt.Errorf("No pkgver field")
	}

	if psr.srcinfo.Pkgrel == "" {
		// Needed because Pacstall implicitely has `1` for empty pkgrel.
		psr.srcinfo.Pkgrel = "1"
	}

	if len(psr.srcinfo.Arch) == 0 {
		psr.srcinfo.Arch = []string{"any"}
		// return nil, fmt.Errorf("No arch field")
	}

	return psr.srcinfo, nil
}

// splitPair splits a key value string in the form of "key = value",
// whitespace being ignored. The key and the value is returned.
func splitPair(line string) (string, string, error) {
	split := strings.SplitN(line, "=", 2)

	if len(split) != 2 {
		return "", "", fmt.Errorf("Line does not contain =")
	}

	key := strings.TrimSpace(split[0])
	value := strings.TrimSpace(split[1])

	if key == "" {
		return "", "", fmt.Errorf("Key is empty")
	}

	return key, value, nil
}

// splitArchFromKey splits up architecture dependent field names, separating
// the field name from the architecture they depend on.
func splitDistroArchFromKey(arches []string, key string) ( /* name */ string /* distro */, string /* arch */, string) {
	split := strings.SplitN(key, "_", 3)

	// possible cases: name_distro_arch, name_distro, name_arch
	if len(split) == 3 {
		arch := split[2]
		// treat cases like name_x86_64
		if checkArch(arches, key, arch) != nil {
			// possibly in a case like name_x86_64 or invalid arch
			arch = split[1] + "_" + split[2]
			if checkArch(arches, key, arch) != nil {
				// invalid arch so it's a distro
				return split[0], arch, ""
			}

			// valid arch but no distro name
			return split[0], "", arch

		}
		// name_distro_arch
		return split[0], split[1], split[2]
	}

	// name_arch, name_distro
	if len(split) == 2 {
		if checkArch(arches, key, split[1]) == nil {
			return split[0], "", split[1]
		}

		return split[0], split[1], ""
	}

	return split[0], "", ""
}

// checkArg checks that the arch from an arch dependent string is actually
// defined inside of the srcinfo and specifically disallows the arch "any" as it
// is not a real arch
func checkArch(arches []string, key string, arch string) error {
	if arch == "" {
		return nil
	}

	if arch == "any" {
		return fmt.Errorf("Invalid key \"%s\" arch \"%s\" is not allowed", key, arch)
	}

	for _, a := range arches {
		if a == arch {
			return nil
		}
	}

	return fmt.Errorf("Invalid key \"%s\" unsupported arch \"%s\"", key, arch)
}

// ParseFile parses a srcinfo file as specified by path.
func ParseFile(path string) (*Srcinfo, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to read file: %s: %s", path, err.Error())
	}

	return Parse(string(file))
}

// Parse parses a srcinfo in string form. Parsing will fail if:
//
//	A srcinfo does not contain all required fields
//	The same pkgname is specified more then once
//	arch is missing
//	pkgver is mising
//	pkgrel is missing
//	An architecture specific field is defined for an architecture that does not exist
//	An unknown key is specified
//	An empty value is specified
//
// Required fields are:
//
//	pkgbase
//	pkname
//	arch
//	pkgrel
//	pkgver
func Parse(data string) (*Srcinfo, error) {
	return parse(data)
}
