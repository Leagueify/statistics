# Statistics Service

The Leagueify statistics service is responsible for managing player
registration, team formation, and statistics updates. The Leagueify statistics
service uses  [Go][go-download] using version `1.23.0`.

## Key Functions

- 

## Development

### Prerequisites

- [Docker][docker-download] is installed and running.

### Getting Started

Local development of the Leagueify statistics service uses docker for a
consistent development environment. Running the Leagueify statistics service
locally can be accomplished using commands found in the
[Makefile][repo-makefile]. During local development changes will trigger a live
reload, eliminating the requirement to restart the docker image manually. This
is accomplished by using the wonderful tool [air][github-air]. The most common
commands have been outlined below:

```bash
# start leagueify docker image
make start

# stop leagueify docker image removing attached volumes
make clean
```

The Leagueify statistics service is ready for development once the banner output
is visible within the terminal. The banner blelow was created using the
[Text to ASCII Art Generator by Patorjk][patorjk-taag].

```
leagueify-statistics-1  |
leagueify-statistics-1  | '||'      '||''''|      |      ..|'''.|  '||'  '|' '||''''|  '||' '||''''| '||' '|'
leagueify-statistics-1  |  ||        ||  .       |||    .|'     '   ||    |   ||  .     ||   ||  .     || |
leagueify-statistics-1  |  ||        ||''|      |  ||   ||    ....  ||    |   ||''|     ||   ||''|      ||
leagueify-statistics-1  |  ||        ||        .''''|.  '|.    ||   ||    |   ||        ||   ||         ||
leagueify-statistics-1  | .||.....| .||.....| .|.  .||.  ''|...'|    '|..'   .||.....| .||. .||.       .||.
leagueify-statistics-1  |
leagueify-statistics-1  |  .|'''.|  |''||''|     |     |''||''| '||'  .|'''.|  |''||''| '||'   ..|'''.|  .|'''.|
leagueify-statistics-1  |  ||..  '     ||       |||       ||     ||   ||..  '     ||     ||  .|'     '   ||..  '
leagueify-statistics-1  |   ''|||.     ||      |  ||      ||     ||    ''|||.     ||     ||  ||           ''|||.
leagueify-statistics-1  | .     '||    ||     .''''|.     ||     ||  .     '||    ||     ||  '|.      . .     '||
leagueify-statistics-1  | |'....|'    .||.   .|.  .||.   .||.   .||. |'....|'    .||.   .||.  ''|....'  |'....|'
leagueify-statistics-1  |
```

[docker-download]: https://www.docker.com/get-started/
[github-air]: https://github.com/air-verse/air
[go-download]: https://go.dev/dl/
[patorjk-taag]: https://patorjk.com/software/taag/#p=display&f=Kban&t=LEAGUEIFY%0ASTATISTICS
[repo-makefile]: ./Makefile
