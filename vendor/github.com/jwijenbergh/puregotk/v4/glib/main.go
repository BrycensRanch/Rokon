// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xBookmarkFileErrorQuark func() Quark

func BookmarkFileErrorQuark() Quark {

	cret := xBookmarkFileErrorQuark()
	return cret
}

var xConvertErrorQuark func() Quark

func ConvertErrorQuark() Quark {

	cret := xConvertErrorQuark()
	return cret
}

var xFileErrorQuark func() Quark

func FileErrorQuark() Quark {

	cret := xFileErrorQuark()
	return cret
}

var xIoChannelErrorQuark func() Quark

func IoChannelErrorQuark() Quark {

	cret := xIoChannelErrorQuark()
	return cret
}

var xKeyFileErrorQuark func() Quark

func KeyFileErrorQuark() Quark {

	cret := xKeyFileErrorQuark()
	return cret
}

var xMarkupErrorQuark func() Quark

func MarkupErrorQuark() Quark {

	cret := xMarkupErrorQuark()
	return cret
}

var xNumberParserErrorQuark func() Quark

func NumberParserErrorQuark() Quark {

	cret := xNumberParserErrorQuark()
	return cret
}

var xOptionErrorQuark func() Quark

func OptionErrorQuark() Quark {

	cret := xOptionErrorQuark()
	return cret
}

var xRegexErrorQuark func() Quark

func RegexErrorQuark() Quark {

	cret := xRegexErrorQuark()
	return cret
}

var xShellErrorQuark func() Quark

func ShellErrorQuark() Quark {

	cret := xShellErrorQuark()
	return cret
}

var xSpawnErrorQuark func() Quark

func SpawnErrorQuark() Quark {

	cret := xSpawnErrorQuark()
	return cret
}

var xSpawnExitErrorQuark func() Quark

func SpawnExitErrorQuark() Quark {

	cret := xSpawnExitErrorQuark()
	return cret
}

var xThreadErrorQuark func() Quark

func ThreadErrorQuark() Quark {

	cret := xThreadErrorQuark()
	return cret
}

var xUnixErrorQuark func() Quark

func UnixErrorQuark() Quark {

	cret := xUnixErrorQuark()
	return cret
}

var xUriErrorQuark func() Quark

func UriErrorQuark() Quark {

	cret := xUriErrorQuark()
	return cret
}

var xVariantParseErrorQuark func() Quark

func VariantParseErrorQuark() Quark {

	cret := xVariantParseErrorQuark()
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xBookmarkFileErrorQuark, lib, "g_bookmark_file_error_quark")
	core.PuregoSafeRegister(&xConvertErrorQuark, lib, "g_convert_error_quark")
	core.PuregoSafeRegister(&xFileErrorQuark, lib, "g_file_error_quark")
	core.PuregoSafeRegister(&xIoChannelErrorQuark, lib, "g_io_channel_error_quark")
	core.PuregoSafeRegister(&xKeyFileErrorQuark, lib, "g_key_file_error_quark")
	core.PuregoSafeRegister(&xMarkupErrorQuark, lib, "g_markup_error_quark")
	core.PuregoSafeRegister(&xNumberParserErrorQuark, lib, "g_number_parser_error_quark")
	core.PuregoSafeRegister(&xOptionErrorQuark, lib, "g_option_error_quark")
	core.PuregoSafeRegister(&xRegexErrorQuark, lib, "g_regex_error_quark")
	core.PuregoSafeRegister(&xShellErrorQuark, lib, "g_shell_error_quark")
	core.PuregoSafeRegister(&xSpawnErrorQuark, lib, "g_spawn_error_quark")
	core.PuregoSafeRegister(&xSpawnExitErrorQuark, lib, "g_spawn_exit_error_quark")
	core.PuregoSafeRegister(&xThreadErrorQuark, lib, "g_thread_error_quark")
	core.PuregoSafeRegister(&xUnixErrorQuark, lib, "g_unix_error_quark")
	core.PuregoSafeRegister(&xUriErrorQuark, lib, "g_uri_error_quark")
	core.PuregoSafeRegister(&xVariantParseErrorQuark, lib, "g_variant_parse_error_quark")

}
