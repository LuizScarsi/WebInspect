# WebInspect
Tool that provides informations about websites

## How to use:
### Clone this repo
- `git clone https://github.com/LuizScarsi/WebInspect.git`
### Change to WebInspect directory
- `cd WebInspect`
### Create binary file
- `go build`
### Search IPs
- Flags: 
    - `--host`
- Example:
    - `./webinspect ip --host example.com`
### Search server names
- Flags: 
    - `--domain`, `--d`
- Example:
    - `./webinspect servers --d example.com`
### Check server health
- Flags:
    - `--domain`, `--d`
    - `--port`, `--p` (optional)
- Example:
    - `./webinspect health --d example.com --p 80`