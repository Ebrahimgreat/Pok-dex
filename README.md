



# Pokedex CLI


    ⬜⬜⬜⬜⬜⬛⬛⬛⬛⬜⬜⬜⬜⬜
⬜⬜⬜⬛⬛🟥🟥🟥🟥🟥⬛⬜⬜⬜
⬜⬜⬛🟥🟥🟥🟥🟥🟥🟥🟥⬛⬜⬜
⬜⬛🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬛⬜
⬜⬛🟥🟥🟥🟥🟥🟥🟥🟥🟥🟥⬛⬜
⬛🟥🟥🟥🟥🟥⬛⬛🟥🟥🟥🟥🟥⬛
⬛🟥🟥🟥🟥⬛⬜⬜⬛🟥🟥🟥🟥⬛
⬛⬛⬛⬛⬛⬛⬜⬜⬛⬛⬛⬛⬛⬛
⬛⬜⬜⬜⬜⬜⬛⬛⬜⬜⬜⬜⬜⬛
⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜
⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜
⬜⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬜
⬜⬜⬜⬛⬛⬜⬜⬜⬜⬛⬛⬜⬜⬜
⬜⬜⬜⬜⬜⬛⬛⬛⬛⬜⬜⬜⬜⬜


A command-line tool written in Go that lets you explore the Pokémon world using the [PokéAPI](https://pokeapi.co/).  
Navigate through regions, explore locations, and discover which Pokémon can be encountered — all from your terminal!

---

## Features

✅ Caching of API responses for faster repeated queries  
✅ Explore Pokémon by location area  
✅ Move through the map with `map`, `mapb` (back), and `mapf` (forward)  
✅ Simple CLI command system (`help`, `exit`, etc.)

---

## In Progress / Planned Features

🚧 Show detailed Pokémon stats (abilities, types, base experience)  
🚧 **Catch Pokémon** – users will be able to catch Pokémon they encounter, filling up their personal Pokédex 🎮  
🚧 More unit tests for stability and maintainability  
🚧 Improved error handling and user-friendly messages  

---

## Installation

Clone this repo:

```bash
git clone https://github.com/<your-username>/pokedex-cli.git
cd pokedex-cli
