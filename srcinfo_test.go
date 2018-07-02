package srcinfo

import (
	"path/filepath"
	"testing"
)

const goodSrcinfoDir string = "testdata/srcinfos/good"
const badSrcinfoDir string = "testdata/srcinfos/bad"

const srcinfoData string = `
# Generated by mksrcinfo v8
# Tue Jun  5 19:37:14 UTC 2018
pkgbase = linux-ck
	pkgver = 4.16.14
	pkgrel = 1
	url = https://wiki.archlinux.org/index.php/Linux-ck
	arch = x86_64
	license = GPL2
	makedepends = kmod
	makedepends = inetutils
	makedepends = bc
	makedepends = libelf
	options = !strip
	source = https://www.kernel.org/pub/linux/kernel/v4.x/linux-4.16.tar.xz
	source = https://www.kernel.org/pub/linux/kernel/v4.x/linux-4.16.tar.sign
	source = https://www.kernel.org/pub/linux/kernel/v4.x/patch-4.16.14.xz
	source = https://www.kernel.org/pub/linux/kernel/v4.x/patch-4.16.14.sign
	source = config
	source = 60-linux.hook
	source = 90-linux.hook
	source = linux.preset
	source = enable_additional_cpu_optimizations-20180509.tar.gz::https://github.com/graysky2/kernel_gcc_patch/archive/20180509.tar.gz
	source = http://ck.kolivas.org/patches/4.0/4.16/4.16-ck1/patch-4.16-ck1.xz
	source = 0001-add-sysctl-to-disallow-unprivileged-CLONE_NEWUSER-by.patch
	source = 0002-ACPI-watchdog-Prefer-iTCO_wdt-on-Lenovo-Z50-70.patch
	source = 0003-Revert-drm-i915-edp-Allow-alternate-fixed-mode-for-e.patch
	sha256sums = 63f6dc8e3c9f3a0273d5d6f4dca38a2413ca3a5f689329d05b750e4c87bb21b9
	sha256sums = SKIP
	sha256sums = cc3d82b8183b641e18e4d504000d7f14873cf67d616ecdabc77383c5d9eaaac0
	sha256sums = SKIP
	sha256sums = 7ae242be3f32e166dce20b86b1bf78d605fc6bf697399dfdd15607f18dd3b367
	sha256sums = ae2e95db94ef7176207c690224169594d49445e04249d2499e9d2fbc117a0b21
	sha256sums = 75f99f5239e03238f88d1a834c50043ec32b1dc568f2cc291b07d04718483919
	sha256sums = ad6344badc91ad0630caacde83f7f9b97276f80d26a20619a87952be65492c65
	sha256sums = 226e30068ea0fecdb22f337391385701996bfbdba37cdcf0f1dbf55f1080542d
	sha256sums = 61cd4b129eac475ad82fcdbbf9952b80e81e7c893776c00e3b6a658b950d0b26
	sha256sums = 8d6a5f34b3d79e75b0cb888c6bcf293f84c5cbb2757f7bdadafee7e0ea77d7dd
	sha256sums = 2454c1ee5e0f5aa119fafb4c8d3b402c5e4e10b2e868fe3e4ced3b1e2aa48446
	sha256sums = 8114295b8c07795a15b9f8eafb0f515c34661a1e05512da818a34581dd30f87e

pkgname = linux-ck
	pkgdesc = The Linux-ck kernel and modules with the ck1 patchset featuring MuQSS CPU scheduler v0.171
	install = linux.install
	depends = coreutils
	depends = linux-firmware
	depends = kmod
	depends = mkinitcpio>=0.7
	optdepends = crda: to set the correct wireless channels of your country
	provides = linux-ck=4.16.14
	backup = etc/mkinitcpio.d/linux-ck.preset

pkgname = linux-ck-headers
	pkgdesc = Header files and scripts for building modules for Linux-ck kernel
	depends = linux-ck
	provides = linux-ck-headers=4.16.14
	provides = linux-headers=4.16.14
`

var goodSrcinfos = [...]string{
	"abcpp",
	"accel-ppp-vlanmon-dkms-git",
	"aerospike-amc-community",
	"als-controller",
	"alt-git",
	"an",
	"anki-sync-server-git",
	"anoise-community-extension3",
	"ansible-container",
	"anydesk",
	"archalien-git",
	"arch_override",
	"arc-kde-git",
	"aspell6-fa",
	"asterisk",
	"astyle-svn",
	"aurman",
	"aurutils",
	"aurvote",
	"autorandr",
	"avahi-static-services",
	"avbin",
	"awn-extras-applets-git",
	"awx-git",
	"bash-it-git",
	"bazarr-git",
	"bbk-cli-bin",
	"bbswitch-ck",
	"beamer-theme-torino-git",
	"betty",
	"bgrep-git",
	"biosdevname",
	"bitbucket-cli",
	"bitkeeper-bk",
	"bitkeeper-production-bin",
	"bizou",
	"blast+",
	"blender-plugin-yavne",
	"blitz-cppqed-hg",
	"bmpx",
	"bootchart2-git",
	"brother-dcp195c",
	"brother-mfc-210c",
	"burp-backup-dev",
	"bwa",
	"byobu",
	"caffe-cmake-git",
	"calibre-git",
	"cargo-edit-git",
	"ccm",
	"cef-standard",
	"cellranger",
	"cerbero-profiler",
	"cfunge",
	"cgicc",
	"cgit-git",
	"cgvg",
	"check_bareos-git",
	"checkmails",
	"chewing-editor-git",
	"chromium-ublock-origin-git",
	"cleanwad",
	"clevo-xsm-wmi-util",
	"clion",
	"cndrvcups-common-lb",
	"cnijfilter2-mg3600",
	"cnijfilter-common-mg5400",
	"cococpp",
	"code",
	"coin-or-osi-git",
	"compiz-core-git",
	"compiz-ubuntu",
	"complx-git",
	"concourse-fly-git",
	"cover-thumbnailer",
	"cower",
	"cpod",
	"creeper-world",
	"crfpp",
	"crochetcharts",
	"crrcsim-sceneries",
	"cryptodev-linux",
	"curlmirror",
	"datovka-git",
	"dbeaver-ee",
	"dblp-refer-git",
	"dbuildstat-git",
	"deadbeef-plugin-fb",
	"debian-gdm-themes",
	"debtap",
	"delta-media-player",
	"deuchnord-hermes",
	"dict-freedict-eng-deu-bin",
	"diction",
	"discord",
	"displaylink",
	"dkms-pl2501",
	"dmtx-utils",
	"dolphin-megasync-git",
	"dontmove-hib",
	"dovecot-libsodium-plugin",
	"downgrade",
	"downgrader",
	"dpkg",
	"dreampie",
	"dropbox",
	"dropbox",
	"ebtables-git",
	"em8300-git",
	"emacs-arduino-mode-git",
	"emacs-clojure-mode-git",
	"empty_override",
	"envizon-git",
	"epaste",
	"eplot",
	"epson-inkjet-printer-201104w",
	"epson-inkjet-printer-201311w",
	"etcher",
	"etlegacy-omnibot",
	"evopop-gtk-theme",
	"faangband",
	"fackup",
	"farbfeld-git",
	"fax4cups",
	"fbx-sdk",
	"ffmpeg-full-arm-git",
	"fgit-git",
	"fheroes2-svn",
	"firefox-esr-bin",
	"flatcc",
	"flatplat-blue-theme",
	"flif",
	"flvstreamer",
	"fman",
	"foo",
	"fortune-mod-firefly",
	"fortune-mod-kaamelott",
	"foxitreader",
	"freecad",
	"freedoom-git",
	"freeglut-svn",
	"freeoffice",
	"freesynd",
	"freetype2-cleartype",
	"fsharp",
	"fslint",
	"fs-uae",
	"fusecompress-git",
	"fusiondirectory-plugin-fai-schema",
	"gamemode",
	"ganv-git",
	"gdc-bin",
	"geant3",
	"geda-gaf-unstable",
	"gedit-autotab-git",
	"genymotion",
	"gfontview",
	"ginn",
	"gitkraken",
	"gitlab-pages",
	"git-lfs-arm",
	"git-notifier",
	"gitprompt-rs",
	"gksu",
	"gnome-desktop2",
	"gnome-shell-extension-atom-dash-git",
	"gnome-shell-extension-caffeine-git",
	"gnome-shell-extension-mmod-panel-git",
	"gnome-shell-extension-slide-for-keyboard-git",
	"gnucap-random-git",
	"gobi-firmware",
	"godot",
	"gogs",
	"go-makepkg",
	"goodvibes",
	"google-maps-desktop",
	"gopacket-git",
	"gophcatch-git",
	"gpicker",
	"grail",
	"greyhole",
	"grive",
	"grub2-theme-archxion",
	"gsignond",
	"gst123",
	"guile-git",
	"gzdoom",
	"hamsolar",
	"hangups-git",
	"hardcode-tray-git",
	"haskell-pcap",
	"hax11-git",
	"hd-idle-cvs",
	"heimdall",
	"holodev",
	"hplip-raw-ledm",
	"hpx",
	"httpjs-git",
	"hunspell-ar",
	"hws-git",
	"i3session-git",
	"ibus-m17n-git",
	"ibus-uniemoji-git",
	"icaclient",
	"im",
	"indicator-keylock",
	"insync",
	"inxi",
	"ipoibmodemtu",
	"ipwatchd",
	"irclog2html",
	"jabref",
	"jack2-git",
	"japa",
	"java8-openjdk-hsdis",
	"jcdk-classic",
	"jdiskreport",
	"jedit-pkgbuild",
	"jid3-bzr",
	"jmtpfs",
	"joplin",
	"journal-notify",
	"jre",
	"json-parser-git",
	"jsontocsv",
	"jtharness-hg",
	"kalu",
	"karia2-svn",
	"kde1-kdebase",
	"kdeplasma-applets-playbar",
	"kdevelop-pg-qt-git",
	"keepass-es",
	"kio-gdrive-git",
	"kodi-cli-git",
	"komodo-ide-nightly",
	"konoha",
	"lddgraph-git",
	"leocad",
	"lib32-double-conversion",
	"lib32-fcitx",
	"lib32-orbit2",
	"lib3ds",
	"libattr-aarch64",
	"libbonobo",
	"libbonoboui",
	"libc++",
	"libcotp",
	"libcs50-git",
	"libgksu",
	"libgksu-colormap-fix",
	"libgnome",
	"libgnomeui",
	"libheif",
	"libkvkontakte-git",
	"liblfds",
	"lib_mysqludf_sys",
	"liboglappth",
	"libpri",
	"libpurple-lurch-git",
	"libsoxr-git",
	"libsub",
	"libtickit",
	"libu2f-server",
	"lightdm-pantheon-greeter-git",
	"linux-disable-tsq",
	"log4j",
	"lomoco",
	"lonestar",
	"loop-aes",
	"lostfiles",
	"lp",
	"luabind",
	"luks-tpm",
	"lutris",
	"lxqt-panel-git",
	"mailspring",
	"mairix-largembox",
	"masterpdfeditor",
	"matterhorn",
	"mdxmini-git",
	"med",
	"megasync",
	"mendeleydesktop",
	"mendeleydesktop",
	"menulibre",
	"meson-cross-aarch64-linux-gnu",
	"minecraft",
	"mingw-w64-libcuckoo-git",
	"mingw-w64-libgcrypt",
	"mingw-w64-libmodbus",
	"mini-xfwm4-theme",
	"minlog-git",
	"mint-backgrounds-serena",
	"miraclecast-git",
	"mkinitcpio-archivetmpfs",
	"mmass",
	"mod_authnz_pam",
	"moka-icon-theme-git",
	"monero",
	"mongo-cxx-driver",
	"mongroup",
	"monogame-git",
	"mp",
	"mpv-prescalers-git",
	"mpv-vapoursynth",
	"mugshot",
	"multibootusb",
	"murmur-snapshot-minimal",
	"mysql-jdbc",
	"mysql-proxy",
	"nanopond",
	"nautilus-dropbox",
	"ncxmms2-git",
	"neon-wallet-bin",
	"netatop-dkms",
	"netstiff",
	"nfrotz",
	"ngrok",
	"nodejs-budo",
	"nodejs-generator-jhipster",
	"nodejs-nativefier",
	"nodejs-ws",
	"nvm",
	"ocaml-menhir",
	"oni",
	"opencascade",
	"opencobolide",
	"openfaas-cli",
	"openss7-modules-lts41-git",
	"openvpn-auth-ldap",
	"optiprime",
	"otf-aurulent-sans",
	"otf-ipaexfont",
	"otf-tex-gyre-ib",
	"outlast-hib",
	"outspline",
	"p7",
	"pacaur",
	"pacmon-git",
	"pacvis-git",
	"panda3d-git",
	"parallels12-tools",
	"passcheck",
	"patchrom",
	"pcloudcc",
	"pdf4tcl",
	"pdfstudio10",
	"pdftk",
	"peek",
	"pep8-asm-git",
	"perl6-html-parser",
	"perl-class-dbi-plugin-type",
	"perl-datetime-incomplete",
	"perl-file-sync",
	"perl-javascript-closure",
	"perl-lingua-en-inflect-phrase",
	"perl-log-dispatch-config",
	"perl-math-prime-util",
	"perl-math-vec",
	"perl-moosex-types",
	"perl-moosex-types-datetime",
	"perl-net-opensoundcontrol",
	"perl-php-serialization",
	"perl-term-readline-zoid",
	"perl-test-class",
	"phpstorm-url-handler",
	"pikalogy",
	"pikaur",
	"pioneer",
	"piuio-dkms-git",
	"pixelize",
	"pixiecore",
	"pkgbrowser",
	"plexydesk-git",
	"plink",
	"plptools",
	"plymouth",
	"pngwriter-git",
	"polly-b-gone-git",
	"polybar",
	"portage-git",
	"prosody-mod-smacks",
	"pulseeffects",
	"puppet3",
	"purple-hangouts-hg",
	"pyfa",
	"pygobject-tutorial-git",
	"pypy-lxml",
	"python2-axolotl-git",
	"python2-boto-rsync",
	"python2-cliff-tablib-liberty",
	"python2-couleur",
	"python2-dnsq",
	"python2-doom_py-git",
	"python2-espressopp",
	"python2-libzfs-git",
	"python2-postfix-policyd-spf",
	"python2-procname",
	"python2-pycanberra-git",
	"python2-pyliblzma",
	"python2-pyro3",
	"python2-pytune",
	"python2-schematics",
	"python2-spidev",
	"python2-surt",
	"python-bluezero",
	"python-bottlenose-git",
	"python-django-cors-headers",
	"python-django-extra-views",
	"python-django-taggit",
	"python-errol",
	"python-habitica",
	"python-html5lint",
	"python-httpie-jwt-auth",
	"python-http-parser-git",
	"python-intelhex",
	"python-kwant",
	"python-lookfor",
	"python-magic-ahupp",
	"python-mapnik",
	"python-monary-hg",
	"python-mshr",
	"python-nss",
	"python-pyalsaaudio",
	"python-pygpgme-rshk-git",
	"python-pykwalify",
	"python-pyorgmode",
	"python-pysmell",
	"python-pytimeparse",
	"python-py-wink",
	"python-queryablelist",
	"python-rasterio",
	"python-rednose",
	"python-rtttl",
	"python-usfm2osis",
	"qrq",
	"qsolocards",
	"qtel-git",
	"quickfm-git",
	"qweechat",
	"r5u87x",
	"r8169aspm-dkms",
	"radarr",
	"ragnar-git",
	"reaper",
	"remarkable",
	"retro-bzr",
	"roboptim-core-git",
	"ros-ardent-examples-rclcpp-minimal-subscriber",
	"ros-indigo-joy",
	"ros-indigo-openni2-launch",
	"ros-indigo-orocos-kdl",
	"ros-indigo-roslaunch",
	"ros-indigo-rqt-action",
	"ros-jade-cpp-common",
	"ros-jade-orocos-kdl",
	"ros-jade-usb-cam",
	"ros-jade-voxel-grid",
	"ros-kinetic-eigen-conversions",
	"ros-kinetic-moveit-fake-controller-manager",
	"ros-kinetic-smach-msgs",
	"ros-kinetic-std-srvs",
	"ros-kinetic-tf2-ros",
	"ros-kinetic-transmission-interface",
	"ros-kinetic-wfov-camera-msgs",
	"ros-lunar-camera-info-manager",
	"ros-lunar-compressed-image-transport",
	"ros-lunar-rqt-msg",
	"ros-lunar-turtle-actionlib",
	"ros-melodic-nodelet",
	"ros-melodic-octomap",
	"ros-melodic-swri-console",
	"rstudio-desktop-git",
	"rtl-sdr-git",
	"ruby-breakpoint",
	"ruby-faraday-cookie_jar",
	"ruby-fog-core",
	"ruby-gnomecanvas",
	"ruby-to_slug",
	"runzip",
	"safe-iop",
	"sahel-fonts",
	"salome-smesh",
	"scangearmp-common",
	"scrcpy",
	"sddm-theme-archpaint2",
	"seamonkey-i18n-pl",
	"setpgrp",
	"shotcut",
	"signal",
	"simgear",
	"simple-http-server",
	"sky",
	"slideextract",
	"smarty3",
	"snapd",
	"softethervpn",
	"sotw-dev",
	"spotify",
	"spotrec",
	"squidview",
	"ssh-askpass-fullscreen",
	"st",
	"stellarium-lts",
	"stern-bin",
	"stest-git",
	"stm32cubemx",
	"stm32flash",
	"stockfish",
	"stopmotion",
	"st-solarized-powerline",
	"sx-open",
	"syncplay",
	"system-tar-and-restore",
	"teamviewer",
	"tear-pages",
	"tensorflow-computecpp",
	"termboy",
	"termi-git",
	"terminus-terminal",
	"texmaster",
	"thinkpad-scripts",
	"tidb-bin",
	"tigervnc-git",
	"tig-git",
	"tini",
	"trash-cli-git",
	"tribler",
	"trisquel-icon-theme",
	"trisquel-wallpapers",
	"tritium",
	"trizen",
	"ttf-aenigma",
	"ttf-fira-sans-ibx",
	"ttf-iosevka-term-ss08",
	"ttf-liberastika",
	"ttf-noto",
	"ttf-oxygen-gf",
	"ttf-signika",
	"ttyecho-git",
	"ttyvideo",
	"tvheadend",
	"typora",
	"udiskie-dmenu-git",
	"uftp",
	"uget-integrator",
	"umview-mod-umlwip",
	"unetbootin",
	"upp",
	"urlview",
	"uterm-git",
	"utmp-git",
	"v4l2loopback-dkms-git",
	"vapoursynth-plugin-degrainmedian-git",
	"vapoursynth-plugin-finedehalo-git",
	"vdhcoapp",
	"vdradmin-am",
	"vdr-cinebars",
	"vdrctl",
	"veles",
	"versionist",
	"vibrancy-colors",
	"vim-codepad",
	"vim-oz",
	"virtio-win",
	"visd-git",
	"visit",
	"vivaldi",
	"vlc-nightly",
	"vmware-systemd-services",
	"vncsnapshot-png",
	"voxelquest-git",
	"wallabag",
	"way-cooler-bg",
	"webkitgtk",
	"weeb-git",
	"wifite-git",
	"wmidump-git",
	"wuzz-git",
	"wxfbe",
	"x86info",
	"xcursor-gt3",
	"xcursor-lliurex",
	"xdg-utils-git",
	"xfe",
	"xroot-bin",
	"xtables-addons",
	"xtheme",
	"yaourt",
	"yay",
	"yay-git",
	"ydcv-git",
	"yokadi",
	"yourik-qt5",
	"ytree",
	"zoom",
	"zotero",
	"zramswap",
	"zsurf-git",
}

var badSrcinfos = [...]string{
	"any_field",
	"base_field_after_pkgname",
	"base_field_after_pkgname2",
	"empty",
	"invalid_arch",
	"multiple_pkgbase",
	"multiple_pkgname",
	"no_arch",
	"no_key",
	"no_file",
	"no_pkgbase",
	"no_pkgname",
	"no_pkgrel",
	"no_pkgver",
	//"no_value",
	"pkgname_before_pkgbase",
	"unknown_key",
}

func TestGoodSrcinfos(t *testing.T) {
	for _, name := range goodSrcinfos {
		path := filepath.Join(goodSrcinfoDir, name)
		srcinfo, err := ParseFile(path)
		if err != nil {
			t.Errorf("Error parsing %s: %s", name, err)
			continue
		}

		splitpkgs := srcinfo.SplitPackages()

		for n, pkg := range srcinfo.Packages {
			splitpkg, err := srcinfo.SplitPackage(pkg.Pkgname)
			if err != nil {
				t.Errorf("Error getting split package %s from %s: %s", splitpkg, pkg.Pkgname, err)
			}

			if splitpkg.Pkgname != splitpkgs[n].Pkgname {
				t.Errorf("Split packages do not match for %s", name)
			}
		}

		_, err = srcinfo.SplitPackage("_non_existing_pkg")
		if err == nil {
			t.Errorf("Got split package %s from %s when it should not exist", "_non_existing_pkg", name)
		}
	}
}

func TestBadSrcinfos(t *testing.T) {
	for _, name := range badSrcinfos {
		path := filepath.Join(badSrcinfoDir, name)
		_, err := ParseFile(path)
		if err == nil {
			t.Errorf("%s parsed when it should have errored", name)
		} else {
			t.Log(err)
		}
	}
}

func TestSrcinfoData(t *testing.T) {
	_, err := Parse(srcinfoData)
	if err != nil {
		t.Errorf("Error parsing data: %s", err)
	}
}

func TestVersion(t *testing.T) {
	versions := [...]string{
		"1:8-2",
		"6.777-2",
		"6.777.r1.g3f15788-1",
	}

	srcinfos := [...]string{
		"stockfish",
		"yay",
		"yay-git",
	}

	for n, name := range srcinfos {
		path := filepath.Join(goodSrcinfoDir, name)
		srcinfo, err := ParseFile(path)
		if err != nil {
			t.Errorf("Error parsing %s: %s", name, err)
			continue
		}

		if srcinfo.Version() != versions[n] {
			t.Errorf("%s: versions do not match: expected %s got %s", name, versions[n], srcinfo.Version())
		}
	}
}
