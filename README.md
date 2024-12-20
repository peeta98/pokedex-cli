# Pokedex CLI 🎮

A simple Command-Line Interface (CLI) Pokedex project built as part of the Boot.dev backend curriculum. This interactive program allows users to explore the world of Pokémon through a REPL (Read-Eval-Print Loop). Players can visit different Pokémon areas, try to catch Pokémon, and inspect their stats. 

## Features

- **Explore Areas**: Visit various Pokémon habitats and discover which Pokémon can be found there.
- **Catch Pokémon**: Attempt to catch Pokémon you encounter. Success is not guaranteed!
- **View Caught Pokémon**: See all the Pokémon you've successfully caught.
- **Inspect Pokémon**: View detailed stats for any Pokémon in your collection.
- **Help Command**: Displays a list of available commands and how to use them.

## Commands

Use the following commands to interact with the CLI:

- `map` & `mapb`: Visit a Pokémon area to discover new Pokémon.
- `catch <pokemon_name>`: Try to catch the specified Pokémon.
- `inspect <pokemon_name>`: View detailed stats for a specific Pokémon in your collection.
- `pokedex`: List all Pokémon you've caught so far.
- `help`: Display a list of all available commands and their descriptions.
- `exit`: Exit the CLI.

## How to Use

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/pokedex-cli.git
   cd pokedex-cli
   ```

2. Build and run the CLI:
   ```bash
   go run main.go
   ```

3. Use the commands listed in the **Commands** section to explore, catch, and inspect Pokémon!

## Example

```
Welcome to the Pokedex CLI!
Type 'help' to see a list of available commands.

> catch pikachu
You caught a pikachu!
> pokedex
Your Pokémon:
  - pikachu
> inspect pikachu
Name: pikachu
Stats:
  - hp: 35
  - attack: 55
  - defense: 40
Type:
  - Electric
> exit
Goodbye, Trainer!
```

## Project Purpose

This project was built to practice CLI development, working with REPLs, and managing game-like logic in a Go program. It demonstrates fundamental skills in interactive programming, data handling, and user experience in a text-based environment.

## Acknowledgments

- Created as part of the [Boot.dev](https://boot.dev) curriculum.
- Inspired by the world of Pokémon, originally created by Nintendo/Game Freak.
