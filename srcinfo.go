// Package srcinfo is a parser for srcinfo files. Typically generated by
// makepkg, part of the pacman package manager.
//
// Split packages and architecture dependent fields are fully supported.
//
// This Package aims to parse srcinfos but not interpret them in any way.
// All values are fundamentally strings, other tools should be used for
// things such as dependency parsing, validity checking etc.
package srcinfo

import (
	"fmt"
)

// ArchDistroString describes string values that may be architecture dependent.
// For Example depends_x86_64.
// If Arch is an empty string then the field is not architecture dependent.
type ArchDistroString struct {
	Arch   string // Architecture name
	Distro string // Distribution
	Value  string // Value
}

// Package describes the fields of a pkgbuild that may be overwritten by
// in build_<pkgname> function.
type Package struct {
	Pkgname        string
	Pkgdesc        string
	URL            string
	Priority       string
	Arch           []string
	License        []string
	Gives          []ArchDistroString
	Depends        []ArchDistroString
	CheckDepends   []ArchDistroString
	OptDepends     []ArchDistroString
	Pacdeps        []ArchDistroString
	CheckConflicts []ArchDistroString
	Conflicts      []ArchDistroString
	Provides       []ArchDistroString
	Breaks         []ArchDistroString
	Replaces       []ArchDistroString
	Enhances       []ArchDistroString
	Recommends     []ArchDistroString
	Suggests       []ArchDistroString
	Backup         []string
	Repology       []string
}

// PackageBase describes the fields of a pkgbuild that may not be overwritten
// in package_<pkgname> function.
type PackageBase struct {
	Pkgbase         string
	Pkgver          string
	Pkgrel          string
	Epoch           string
	Mask            []string
	Compatible      []string
	Incompatible    []string
	Maintainer      []string
	Source          []ArchDistroString
	NoExtract       []string
	NoSubmodules    []string
	MD5Sums         []ArchDistroString
	SHA1Sums        []ArchDistroString
	SHA224Sums      []ArchDistroString
	SHA256Sums      []ArchDistroString
	SHA384Sums      []ArchDistroString
	SHA512Sums      []ArchDistroString
	B2Sums          []ArchDistroString
	MakeDepends     []ArchDistroString
	MakeConflicts   []ArchDistroString
}

// Srcinfo represents a full srcinfo. All global fields are defined here while
// fields overwritten in the package_<pkgname> function are defined in the
// Packages field.
//
// Note: The Packages field only contains the values that each package
// overrides, global fields will be missing. A Package containing both global
// and overwritten fields can be generated using the SplitPackage function.
type Srcinfo struct {
	PackageBase           // Fields that only apply to the package base
	Package               // Fields that apply to the package globally
	Packages    []Package // Fields for each package this package base contains
}

// EmptyOverride is used to signal when a value has been overridden with an
// empty value. An empty ovrride is when a value is defined in the pkgbuild but
// then overridden inside the package function to be empty.
//
// For example "pkgdesc=”" is an empty override on the pkgdesc which would
// lead to the line "pkgdesc=" in the srcinfo.
//
// This value is used internally to store empty overrides, mainly to avoid
// using string pointers. It is possible to check for empty overrides using
// the Packages slice in Packagebase.
//
// During normal use with the SplitPackage function this value will be
// converted back to an empty string, or removed entirely for slice values.
// This means the this value can be completley ignored unless you are
// explicitly looking for empty overrides.
const EmptyOverride = "\x00"

// Version formats a version string from the epoch, pkgver and pkgrel of the
// srcinfo. In the format [epoch:]pkgver-pkgrel.
func (si *Srcinfo) Version() string {
	if si.Epoch == "" {
		return si.Pkgver + "-" + si.Pkgrel
	}

	return si.Epoch + ":" + si.Pkgver + "-" + si.Pkgrel
}

// SplitPackages generates a splice of all packages that are part of this
// srcinfo. This is equivalent to calling SplitPackage on every pkgname.
func (si *Srcinfo) SplitPackages() []*Package {
	pkgs := make([]*Package, 0, len(si.Packages))

	for _, pkg := range si.Packages {
		pkgs = append(pkgs, mergeSplitPackage(&si.Package, &pkg))
	}

	return pkgs
}

// SplitPackage generates a Package that contains all fields that the specified
// pkgname has. But will fall back on global fields if they are not defined in
// the Package.
//
// Note slice values will be passed by reference, it is not recommended you
// modify this struct after it is returned.
func (si *Srcinfo) SplitPackage(pkgname string) (*Package, error) {
	for n := range si.Packages {
		if si.Packages[n].Pkgname == pkgname {
			return mergeSplitPackage(&si.Package, &si.Packages[n]), nil
		}
	}

	return nil, fmt.Errorf("Package \"%s\" is not part of the package base \"%s\"", pkgname, si.Pkgbase)
}

func mergeArchSlice(global, override []ArchDistroString) []ArchDistroString {
	overridden := make(map[string]struct{})
	merged := make([]ArchDistroString, 0, len(override))

	for _, v := range override {
		overridden[v.Arch] = struct{}{}
		if v.Value == EmptyOverride {
			continue
		}
		merged = append(merged, v)
	}

	for _, v := range global {
		if _, ok := overridden[v.Arch]; !ok {
			merged = append(merged, v)
		}
	}

	return merged
}

func mergeSplitPackage(base, split *Package) *Package {
	pkg := &Package{}
	*pkg = *base

	pkg.Pkgname = split.Pkgname

	if split.Pkgdesc != "" {
		pkg.Pkgdesc = split.Pkgdesc
	}

	if split.URL != "" {
		pkg.URL = split.URL
	}
	
	if split.Priority != "" {
		pkg.Priority = split.Priority
	}
	
	if len(split.Arch) != 0 {
		pkg.Arch = split.Arch
	}

	if len(split.License) != 0 {
		pkg.License = split.License
	}

	if len(split.Gives) != 0 {
		pkg.Gives = mergeArchSlice(pkg.Gives, split.Gives)
	}

	if len(split.Depends) != 0 {
		pkg.Depends = mergeArchSlice(pkg.Depends, split.Depends)
	}

	if len(split.CheckDepends) != 0 {
		pkg.CheckDepends = mergeArchSlice(pkg.CheckDepends, split.CheckDepends)
	}
	
	if len(split.OptDepends) != 0 {
		pkg.OptDepends = mergeArchSlice(pkg.OptDepends, split.OptDepends)
	}

	if len(split.Pacdeps) != 0 {
		pkg.Pacdeps = mergeArchSlice(pkg.Pacdeps, split.Pacdeps)
	}

	if len(split.CheckConflicts) != 0 {
		pkg.CheckConflicts = mergeArchSlice(pkg.CheckConflicts, split.CheckConflicts)
	}
	
	if len(split.Conflicts) != 0 {
		pkg.Conflicts = mergeArchSlice(pkg.Conflicts, split.Conflicts)
	}
	
	if len(split.Provides) != 0 {
		pkg.Provides = mergeArchSlice(pkg.Provides, split.Provides)
	}

	if len(split.Breaks) != 0 {
		pkg.Breaks = mergeArchSlice(pkg.Breaks, split.Breaks)
	}
	
	if len(split.Replaces) != 0 {
		pkg.Replaces = mergeArchSlice(pkg.Replaces, split.Replaces)
	}

	if len(split.Enhances) != 0 {
		pkg.Enhances = mergeArchSlice(pkg.Enhances, split.Enhances)
	}

	if len(split.Recommends) != 0 {
		pkg.Recommends = mergeArchSlice(pkg.Recommends, split.Recommends)
	}

	if len(split.Suggests) != 0 {
		pkg.Suggests = mergeArchSlice(pkg.Suggests, split.Suggests)
	}

	if len(split.Backup) != 0 {
		pkg.Backup = split.Backup
	}

	if len(split.Repology) != 0 {
		pkg.Repology = split.Repology
	}

	return pkg
}
