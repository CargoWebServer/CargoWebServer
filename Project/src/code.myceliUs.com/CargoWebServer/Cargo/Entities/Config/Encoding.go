// +build Config

package Config

type Encoding int
const(
	Encoding_UTF8 Encoding = 1+iota
	Encoding_WINDOWS_1250
	Encoding_WINDOWS_1251
	Encoding_WINDOWS_1252
	Encoding_WINDOWS_1253
	Encoding_WINDOWS_1254
	Encoding_WINDOWS_1255
	Encoding_WINDOWS_1256
	Encoding_WINDOWS_1257
	Encoding_WINDOWS_1258
	Encoding_ISO8859_1
	Encoding_KOI8R
	Encoding_KOI8U
)
