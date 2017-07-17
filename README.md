# put
Go tool to ensure a given line is inside a text file. If it is not, the tool will insert the line (configurable in which line number).


## how to use
`put -help` will list options:
`-file` describes the input text file
`-insert` is the line of code that is to be inserted
`-n` is the line number where the line should be inserted (default: append at end of file)

## why would someone need something like this?
mainly to script adding specific lines to config files only if that line does not already exist
