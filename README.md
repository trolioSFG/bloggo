# bloggo
Blog aggregator in Go

Tutorial from boot.dev

## TODO
1. Create a README.md file in the root of your repo if you don't have one already. (You should be tracking your changes with Git.)
2. Explain to the user that they'll need Postgres and Go installed to run the program.
3. Explain to the user how to install the gator CLI using go install.
4. Explain to the user how to set up the config file and run the program. Tell them about a few of the commands they can run.
5. Push gator up to GitHub, then submit the link to your remote repo. Your link should look something like this: https://github.com/github-username/repo-name.
6. Fix bug with PublishedAt on browse command

## EXTENDING THE PROJECT

- Add sorting and filtering options to the browse command
- Add pagination to the browse command
- Add concurrency to the agg command so that it can fetch more frequently
- Add a search command that allows for fuzzy searching of posts
- Add bookmarking or liking posts
- Add a TUI that allows you to select a post in the terminal and view it in a more readable format (either in the terminal or open in a browser)
- Add an HTTP API (and authentication/authorization) that allows other users to interact with the service remotely
Write a service manager that keeps the agg command running in the background and restarts it if it crashes

