# ASCII-ART

This programm takes user input and prints in as an ascii-art, only ascii-chars are supported. 

# INSTALLATION

To install the program, simply clone this repository into the folder and build using `go build main.go`

# USAGE

##Color specification

### Paint whole string by color name
`./ascii-art [YOUR STRING] [--color=red]` 
EX: 

`./ascii-art "I love linux" "--color=red"`

OUTPUT:
 _____         _                                _   _
|_   _|       | |                              | | (_)
  | |         | |   ___   __   __   ___        | |  _   _ __    _   _  __  __
  | |         | |  / _ \  \ \ / /  / _ \       | | | | | '_ \  | | | | \ \/ /
 _| |_        | | | (_) |  \ V /  |  __/       | | | | | | | | | |_| |  >  <
|_____|       |_|  \___/    \_/    \___|       |_| |_| |_| |_|  \__,_| /_/\_\


### Paint whole string by color hex value 
`./ascii-art [YOUR STRING] [--color=#ff0000]`
####EX:

`./ascii-art "My favorite text editor is Vim" "--color=#ff0000"`

####OUTPUT:
 __  __                  __                                          _   _                  _                   _                      _   _   _                          _              __      __  _
|  \/  |                / _|                                        (_) | |                | |                 | |                    | | (_) | |                        (_)             \ \    / / (_)
| \  / |  _   _        | |_    __ _  __   __   ___    _   _   _ __   _  | |_    ___        | |_    ___  __  __ | |_          ___    __| |  _  | |_    ___    _ __         _   ___         \ \  / /   _   _ __ ___
| |\/| | | | | |       |  _|  / _` | \ \ / /  / _ \  | | | | | '__| | | | __|  / _ \       | __|  / _ \ \ \/ / | __|        / _ \  / _` | | | | __|  / _ \  | '__|       | | / __|         \ \/ /   | | | '_ ` _ \
| |  | | | |_| |       | |   | (_| |  \ V /  | (_) | | |_| | | |    | | \ |_  |  __/       \ |_  |  __/  >  <  \ |_        |  __/ | (_| | | | \ |_  | (_) | | |          | | \__ \          \  /    | | | | | | | |
|_|  |_|  \__, |       |_|    \__,_|   \_/    \___/   \__,_| |_|    |_|  \__|  \___|        \__|  \___| /_/\_\  \__|        \___|  \__,_| |_|  \__|  \___/  |_|          |_| |___/           \/     |_| |_| |_| |_|
          __/ /
         |___/


### Paint whole string by color rgb value
`./ascii-art [YOUR STRING] [--color=rgb(255,0,0)]`
####EX:

`./ascii-art "DEADBEEF" "--color=rgb(255,0,0)"`

####OUTPUT:
 _____    ______              _____    ____    ______   ______   ______  
|  __ \  |  ____|     /\     |  __ \  |  _ \  |  ____| |  ____| |  ____| 
| |  | | | |__       /  \    | |  | | | |_) | | |__    | |__    | |__    
| |  | | |  __|     / /\ \   | |  | | |  _ <  |  __|   |  __|   |  __|   
| |__| | | |____   / ____ \  | |__| | | |_) | | |____  | |____  | |      
|_____/  |______| /_/    \_\ |_____/  |____/  |______| |______| |_|      
                                                                         
                                                                         

### Paint whole string by color hsl value 
`./ascii-art [YOUR STRING] [--color=hsl(0,100,50)]` 
####EX:

`./ascii-art "rm -rf /" "--color=hsl(0,100,50)"`

####OUTPUT:
                                           __             __
                                          / _|           / /
 _ __   _ __ ___          ______   _ __  | |_           / /
| '__| | '_ ` _ \        |______| | '__| |  _|         / /
| |    | | | | | |                | |    | |          / /
|_|    |_| |_| |_|                |_|    |_|         /_/



##Substring specification
### Paint only substring by chars
`./ascii-art [YOUR STRING] [--color=red] [CHARS FROM YOUR STRING]
####EX:
`./ascii-art "This string has only i-s and s-es painted" "--color=red" is`
####OUTPUT:
 _______   _       _                     _            _                         _                                           _                 _                                            _                                                           _           _                _  
|__   __| | |     (_)                   | |          (_)                       | |                                         | |               (_)                                          | |                                                         (_)         | |              | | 
   | |    | |__    _   ___         ___  | |_   _ __   _   _ __     __ _        | |__     __ _   ___          ___    _ __   | |  _   _         _   ______   ___          __ _   _ __     __| |        ___   ______    ___   ___         _ __     __ _   _   _ __   | |_    ___    __| | 
   | |    |  _ \  | | / __|       / __| | __| | '__| | | | '_ \   / _` |       |  _ \   / _` | / __|        / _ \  | '_ \  | | | | | |       | | |______| / __|        / _` | | '_ \   / _` |       / __| |______|  / _ \ / __|       | '_ \   / _` | | | | '_ \  | __|  / _ \  / _` | 
   | |    | | | | | | \__ \       \__ \ \ |_  | |    | | | | | | | (_| |       | | | | | (_| | \__ \       | (_) | | | | | | | | |_| |       | |          \__ \       | (_| | | | | | | (_| |       \__ \          |  __/ \__ \       | |_) | | (_| | | | | | | | \ |_  |  __/ | (_| | 
   |_|    |_| |_| |_| |___/       |___/  \__| |_|    |_| |_| |_|  \__, |       |_| |_|  \__,_| |___/        \___/  |_| |_| |_|  \__, |       |_|          |___/        \__,_| |_| |_|  \__,_|       |___/           \___| |___/       | .__/   \__,_| |_| |_| |_|  \__|  \___|  \__,_| 
                                                                   __/ |                                                        __/ /                                                                                                 | |                                              
                                                                  |___/                                                        |___/                                                                                                  |_|                                              

### Paint only substring by index
`./ascii-art [YOUR STRING] [--color=red] [start_index:end_index]`
####EX:
`./ascii-art "Paint only second word" "--color=red" [6:10]"
####OUTPUT:
 _____            _           _                           _                                                         _                                       _  
|  __ \          (_)         | |                         | |                                                       | |                                     | | 
| |__) |   __ _   _   _ __   | |_          ___    _ __   | |  _   _         ___    ___    ___    ___    _ __     __| |       __      __   ___    _ __    __| | 
|  ___/   / _` | | | | '_ \  | __|        / _ \  | '_ \  | | | | | |       / __|  / _ \  / __|  / _ \  | '_ \   / _` |       \ \ /\ / /  / _ \  | '__|  / _` | 
| |      | (_| | | | | | | | \ |_        | (_) | | | | | | | | |_| |       \__ \ |  __/ | (__  | (_) | | | | | | (_| |        \ V  V /  | (_) | | |    | (_| | 
|_|       \__,_| |_| |_| |_|  \__|        \___/  |_| |_| |_|  \__, |       |___/  \___|  \___|  \___/  |_| |_|  \__,_|         \_/\_/    \___/  |_|     \__,_| 
                                                              __/ /                                                                                            
                                                             |___/                                                                                             
