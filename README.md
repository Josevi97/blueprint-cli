# Blueprint CLI
Basic command line interface to manage project structure

## Architecture
- Processes and managers must not throw any error. They should completly transparent about errors

## TODO
Should investigate about what kind of architecture about project structure is the best. Currently, the main approach is to have processes as modules and the rest of code is separated in utils, managers, handlers, etc. Another approach is to only have core and modules folders, and have everything as modules. Processes would be a module with its own folder structure.
