# simple-clargs
Simple module to manage command line args in go

## Example usage
```
OptionInput := clargs.New("-i", "--input", "Input file path", true, true)
OptionOutput := clargs.New("-o", "--output", "Ouput file path", true, false)

options := []*clargs.Option{&OptionInput, &OptionOutput}
args := os.Args[1:]

usage := "./executable [options...]"

clargs.Init(usage, options, args)
clargs.CheckArgs()
```
