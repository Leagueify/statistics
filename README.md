# Statistics Service

The Leagueify statistics service is responsible for collecting statistics for
players, teams, and leagues. The Leagueify statistics service uses
[Go][go-download] using version `1.23.0`.

## Key Functions

- Gather match results and performance data from the League Service.
- Calculate and store player, team, and league stats.
- Provide advance analytics and reporting tools for users.

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
is visible within the terminal. By default the Leagueify statistics service api
docs are accessible at [http://localhost:6505/statistics/docs][service-url]. The
banner below was created using the
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
[service-url]: http://localhost:6505/statistics/docs
