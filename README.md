## pathmem

#### Index
- [About](#about)
- [Instalation](#installation)
- [Commands](#commands)
- [Tech Stack](#stack)

### About
Path memory is a simple cli app that helps you navigate folder path faster. You can save any system path with a name, and can later access the absolute path with it's name.

### Installation
- If you only want to use: For mac users, build is alredy present in build folder. You can simply set alias in bash.rc (or .zsh) as `alias path="/ansolute/path/pathmem/build/main"`. This lets you use path as a regular terminal command.
- If you want to develop: Clone the repository. Build the project using `go build` and you are good to go.

### Commands
| Command | Description |
| --- | --- |
| save any_name | Saves currend folder path under alias `any_name` |
| open any_name | Opens the absolute path corresponding to `any_name` in new terminal window. (currently supports only macos) |

### Tech Stack
go 1.13