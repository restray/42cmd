
## Commands

Syntax:
```bash
ftcli [global-options] subcommand [subcommand-options]
```

### Agenda

```bash
ftcli agenda
ftcli agenda available
ftcli agenda incoming
ftcli agenda reserved
```

### Correction

```bash
ftcli corrections add {begin} {end}
ftcli corrections remove {begin} {end}
ftcli corrections list
ftcli corrections view {correction_id}
```

### Me

```bash
ftcli me [-wallet=false] [-eval-points=true] [-grade=false] [-place=false] [-black-hole=false] [-prefix=false] [-tags=false] [-coalition=false]
ftcli me level
ftcli me projects
ftcli me logtime
```

### Projects

```bash
ftcli projects [-cursus=current_cursus]
ftcli projects available # Get available projects
ftcli projects list # List all the projects for the current cursus
```

#### Project
```bash
ftcli project {project_name}            # Get detail informations about the project
ftcli project {project_name} register   # Register to the specified project
ftcli project {project_name} clone      # Clone the project repo in current folder
ftcli project {project_name} correct    # Display an interactive list of available correction slots
# not possible because 42 API doesn't allow it
# ftcli project {project_name} subject    # Open the subject.pdf in the default browser
```

### User