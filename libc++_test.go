package srcinfo

import (
	"path/filepath"
	"reflect"
	"testing"
)

var libcpp = &Srcinfo{
	PackageBase: PackageBase{
		Pkgbase: "libc++",
		Pkgver:  "6.0.0",
		Pkgrel:  "1",
		Epoch:   "",
		Source: []ArchDistroString{
			{Arch: "", Value: "https://releases.llvm.org/6.0.0/llvm-6.0.0.src.tar.xz"},
			{Arch: "", Value: "https://releases.llvm.org/6.0.0/llvm-6.0.0.src.tar.xz.sig"},
			{Arch: "", Value: "https://releases.llvm.org/6.0.0/libcxx-6.0.0.src.tar.xz"},
			{Arch: "", Value: "https://releases.llvm.org/6.0.0/libcxx-6.0.0.src.tar.xz.sig"},
			{Arch: "", Value: "https://releases.llvm.org/6.0.0/libcxxabi-6.0.0.src.tar.xz"},
			{Arch: "", Value: "https://releases.llvm.org/6.0.0/libcxxabi-6.0.0.src.tar.xz.sig"},
		},
		ValidPGPKeys: []string{
			"11E521D646982372EB577A1F8F0871F202119294",
			"B6C8F98282B944E3B0D5C2530FC3042E345AD05D",
		},
		NoExtract: []string{
			"llvm-6.0.0.src.tar.xz",
			"llvm-6.0.0.src.tar.xz.sig",
			"libcxx-6.0.0.src.tar.xz",
			"libcxx-6.0.0.src.tar.xz.sig",
			"libcxxabi-6.0.0.src.tar.xz",
			"libcxxabi-6.0.0.src.tar.xz.sig",
		},
		MD5Sums:    []ArchDistroString(nil),
		SHA1Sums:   []ArchDistroString(nil),
		SHA224Sums: []ArchDistroString(nil),
		SHA256Sums: []ArchDistroString(nil),
		SHA384Sums: []ArchDistroString(nil),
		SHA512Sums: []ArchDistroString{
			{Arch: "", Value: "a71fdd5ddc46f01327ad891cfcc198febdbe10769c57f14d8a4fb7d514621ee4080e1a641200d3353c16a16731d390270499ec6cd3dc98fadc570f3eb6b52b8c"},
			{Arch: "", Value: "SKIP"},
			{Arch: "", Value: "3d93910f85a778f36c5f7a4429639008acba5713a2c8ac79a9de09463af6f9a388af45d39af23423a7223660701697ba067f3391f25d5a970973691dd88635e3"},
			{Arch: "", Value: "SKIP"},
			{Arch: "", Value: "c5e4cc05105770b42b20595fdbda5e1483be4582bc94335da1a15531ba43a0ecf30e1e0a252f62d4d0e6c79cda9d44ff5fdbe69a0a295b2431fd6de158410e2e"},
			{Arch: "", Value: "SKIP"},
		},
		MakeDepends: []ArchDistroString{
			{Arch: "", Value: "clang"},
			{Arch: "", Value: "cmake"},
			{Arch: "", Value: "ninja"},
			{Arch: "", Value: "python"},
			{Arch: "", Value: "libunwind"},
		},
		CheckDepends: []ArchDistroString(nil),
	},

	Package: Package{
		Pkgname: "",
		Pkgdesc: "",
		Arch: []string{
			"i686",
			"x86_64",
		},
		URL: "https://libcxx.llvm.org/",
		License: []string{
			"MIT",
			"custom:University of Illinois/NCSA Open Source License",
		},
		Groups: []string(nil),
		Depends: []ArchDistroString{
			{Arch: "", Value: "gcc-libs"},
		},
		OptDepends: []ArchDistroString(nil),
		Provides:   []ArchDistroString(nil),
		Conflicts:  []ArchDistroString(nil),
		Replaces:   []ArchDistroString(nil),
		Backup:     []string(nil),
		Options:    []string(nil),
		Install:    "",
		Changelog:  "",
	},
	Packages: []Package{
		{
			Pkgname: "libc++",
			Pkgdesc: "LLVM C++ standard library.",
			Arch:    []string(nil),
			URL:     "",
			License: []string(nil),
			Groups:  []string(nil),
			Depends: []ArchDistroString{
				{Arch: "", Value: "libc++abi=6.0.0-1"},
			},
			OptDepends: []ArchDistroString(nil),
			Provides:   []ArchDistroString(nil),
			Conflicts:  []ArchDistroString(nil),
			Replaces:   []ArchDistroString(nil),
			Backup:     []string(nil),
			Options:    []string(nil),
			Install:    "",
			Changelog:  "",
		},
		Package{
			Pkgname:    "libc++abi",
			Pkgdesc:    "Low level support for the LLVM C++ standard library.",
			Arch:       []string(nil),
			URL:        "",
			License:    []string(nil),
			Groups:     []string(nil),
			Depends:    []ArchDistroString(nil),
			OptDepends: []ArchDistroString(nil),
			Provides:   []ArchDistroString(nil),
			Conflicts:  []ArchDistroString(nil),
			Replaces:   []ArchDistroString(nil),
			Backup:     []string(nil),
			Options:    []string(nil),
			Install:    "",
			Changelog:  "",
		},
		Package{
			Pkgname: "libc++experimental",
			Pkgdesc: "LLVM C++ experimental library.",
			Arch:    []string(nil),
			URL:     "",
			License: []string(nil),
			Groups:  []string(nil),
			Depends: []ArchDistroString{
				{Arch: "", Value: "libc++=6.0.0-1"},
			},
			OptDepends: []ArchDistroString(nil),
			Provides:   []ArchDistroString(nil),
			Conflicts:  []ArchDistroString(nil),
			Replaces:   []ArchDistroString(nil),
			Backup:     []string(nil),
			Options:    []string(nil),
			Install:    "",
			Changelog:  "",
		},
	},
}

var libcppPackage = &Package{
	Pkgname: "libc++",
	Pkgdesc: "LLVM C++ standard library.",
	Arch:    []string{"i686", "x86_64"},
	URL:     "https://libcxx.llvm.org/",
	License: []string{
		"MIT",
		"custom:University of Illinois/NCSA Open Source License",
	},
	Groups: []string(nil),
	Depends: []ArchDistroString{
		{Arch: "", Value: "libc++abi=6.0.0-1"},
	},
	OptDepends: []ArchDistroString(nil),
	Provides:   []ArchDistroString(nil),
	Conflicts:  []ArchDistroString(nil),
	Replaces:   []ArchDistroString(nil),
	Backup:     []string(nil),
	Options:    []string(nil),
	Install:    "",
	Changelog:  "",
}

var libcppABIPackage = &Package{
	Pkgname: "libc++abi",
	Pkgdesc: "Low level support for the LLVM C++ standard library.",
	Arch: []string{
		"i686",
		"x86_64",
	},
	URL: "https://libcxx.llvm.org/",
	License: []string{
		"MIT",
		"custom:University of Illinois/NCSA Open Source License",
	},
	Groups: []string(nil),
	Depends: []ArchDistroString{
		{Arch: "", Value: "gcc-libs"},
	},
	OptDepends: []ArchDistroString(nil),
	Provides:   []ArchDistroString(nil),
	Conflicts:  []ArchDistroString(nil),
	Replaces:   []ArchDistroString(nil),
	Backup:     []string(nil),
	Options:    []string(nil),
	Install:    "",
	Changelog:  "",
}

var libcppExperimentalPackage = &Package{
	Pkgname: "libc++experimental",
	Pkgdesc: "LLVM C++ experimental library.",
	Arch: []string{
		"i686",
		"x86_64",
	},
	URL: "https://libcxx.llvm.org/",
	License: []string{
		"MIT",
		"custom:University of Illinois/NCSA Open Source License",
	},
	Groups: []string(nil),
	Depends: []ArchDistroString{
		{Arch: "", Value: "libc++=6.0.0-1"},
	},
	OptDepends: []ArchDistroString(nil),
	Provides:   []ArchDistroString(nil),
	Conflicts:  []ArchDistroString(nil),
	Replaces:   []ArchDistroString(nil),
	Backup:     []string(nil),
	Options:    []string(nil),
	Install:    "",
	Changelog:  "",
}

func TestLibcpp(t *testing.T) {
	path := filepath.Join(goodSrcinfoDir, "libc++")
	srcinfo, err := ParseFile(path)
	if err != nil {
		t.Errorf("Error parsing data: %s", err)
	}

	if !reflect.DeepEqual(srcinfo, libcpp) {
		t.Errorf("srcinfos do not match for libc++:\n\n%#v\n\n%#v", libcpp, srcinfo)
	}

	srcinfoLibcpp, err := srcinfo.SplitPackage("libc++")
	if err != nil {
		t.Errorf("could not get package %s from %s: %s", "libc++", "libc++", err)
	}

	if !reflect.DeepEqual(srcinfoLibcpp, libcppPackage) {
		t.Errorf("srcinfos do not match for libc++:\n\n%#v\n\n%#v", libcppPackage, srcinfoLibcpp)
	}

	srcinfoLibcppABI, err := srcinfo.SplitPackage("libc++abi")
	if err != nil {
		t.Errorf("could not get package %s from %s: %s", "libc++", "libc++ABI", err)
	}

	if !reflect.DeepEqual(srcinfoLibcppABI, libcppABIPackage) {
		t.Errorf("srcinfos do not match for libc++:\n\n%#v\n\n%#v", libcppABIPackage, srcinfoLibcppABI)
	}

	srcinfoLibcppExperimental, err := srcinfo.SplitPackage("libc++experimental")
	if err != nil {
		t.Errorf("could not get package %s from %s: %s", "libc++", "libc++experimental", err)
	}

	if !reflect.DeepEqual(srcinfoLibcppExperimental, libcppExperimentalPackage) {
		t.Errorf("srcinfos do not match for libc++:\n\n%#v\n\n%#v", libcppExperimentalPackage, srcinfoLibcppExperimental)
	}
}
