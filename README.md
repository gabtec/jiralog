# jiralog

To log work hours into jira

## Install
### Option 1
- clone thos repo
- run included script to release it locally:
```sh
git clone <this>

./release-local.sh
# will ask sudo password to be able to write to /usr/local/bin path
```
### Option 2
- download the binary that match your system, from [releases, here](https://github.com/gabtec/jiralog/releases)
- in macOS
```sh
tar -xzf jiralog_Darwin_arm64.tar.gz

# move it to /usr/local/bin
sudo cp jiralog /usr/local/bin/jiralog


# if issues with macOS blacklist the binary
xattr -dr com.apple.quarantine ./your-binary
codesign -s - --deep --force ./your-binary


````

## Usage
- a jira apiKey is required as env var named: **JIRA_API_TOKEN**
- a jira server baseUrl is required as env var named: **JIRA_BASE_URL**
- we should provide a **worklog.yaml** file
- example:
```yaml
# data: is a mandatory key
data:
  # Day 1
  # "YYYY-MM-DD": the date to log work in
  "2025-08-29":
    # ABC-xxxx: the ticket ID
    ABC-1111:
      # start: the time task started, if defined must be HH:MM, default is 09:00
      start: "09:00" 
      # timeSpent: 30m, [1-8]h, or combination of "xh 30m", max 8h per day
      timeSpent: "30m" 
      # description: an optional comment to add to work log, default is ""
      description: Daily
    ABC-1122:
      timeSpent: "3h"
      description: Description is optional
  # Day 2
  "2025-08-30": {}   
```
## Usage
- we recommend using the --dry-run option first

```sh
# check version
jiralog version # or -v, --version, version

# it will read a worklog.yaml file
jiralog [-d]

# flags (optional):
# -d,--dry-run	- Will just print a summary table, with total hours worked per day
```

## Build

```sh
go build -o build/jiralog .
```

## ToDo
- [ ] Check if it duplicates same entry
- [ ] Allow a -c another-filename.yaml
- [ ] Validate date is from the cw
- [x] Validate ticket ABC-xxxx
- [x] Validate start, if existent, as HH:MM
- [x] Validate date yyy-mm-dd format
- [x] add total working hours per day
- [x] validate total working hours per day <=8h
- [x] Validate timeSpent, as oneOf 1h, etc,  30m