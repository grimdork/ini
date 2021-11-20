# ini
A simple INI file parser and builder.

## What
Read INI files, or build them section by section. Datatypes are kept very simple.

The main fuctions are:
- Load(): Load and parse into a structure.
- Save(): Store an INI file.
- New(): Create a new INI file.

## How
Load an INI file:
```go
	...
	f,err := ini.LoadFile("filename.ini")
	...
```

Add a section:
```go
	...
	sec := f.AddSection("server")
	...
```

Add entries to the section:
```go
	...
	sec.AddString("host", "localhost")
	sec.AddInt("port", 8080)
	sec.AddBool("ssl", false)
	sec.AddFloat("intensity", 10.0)
	...
```

Save the file:
```go
	...
	f.SaveTo("filename.ini")
	...
```
