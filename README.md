# qc
Change query string values with a user-supplied value

## Installation
### From Source
1. Install Go on your system (min v17.0.0 and GO111MODULE env should be "on")
2. `go install github.com/burak0x01/qc@latest`

### From Binary
You can download the pre-built binaries from the releases page and run. For example: </br> </br>
`unzip qc_1-0-0_linux_amd64.zip` </br>
`mv qc /usr/bin` </br>
`qc -h`

## Usage
Findsecret requires 3 parameters to run: -i (input), -o (output), -p(parameter).

### For example 
`qc -i domains.txt -p FUZZ -o result.txt` 

