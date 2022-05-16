# ASCII-ART

This programm takes user input and prints in as an ascii-art, only ascii-chars are supported. 

# INSTALLATION

To install the program, simply clone this repository into the folder and build using `go build -o ascii-art main.go`

# USAGE

## Color specification

### Paint whole string by color name
`./ascii-art [YOUR STRING] [--color=red]`


### Paint whole string by color hex value 
`./ascii-art [YOUR STRING] [--color=#ff0000]`

### Paint whole string by color rgb value
`./ascii-art [YOUR STRING] [--color=rgb(255,0,0)]`

### Paint whole string by color hsl value 
`./ascii-art [YOUR STRING] [--color=hsl(0,100,50)]` 

## Substring specification

### Paint only substring by chars
`./ascii-art [YOUR STRING] [--color=red] [CHARS FROM YOUR STRING]`

### Paint only substring by index
`./ascii-art [YOUR STRING] [--color=red] [start_index:end_index]`
                                                                              