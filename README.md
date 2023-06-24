# Blueprint CLI
Basic command line interface to manage project structure

## Architecture
- Processes and utils must not throw any error. They should be completly transparent about errors
- In order to do some queries in the database. The architecture is as follows:
 1. Processes creates the connection to the database
 2. Processes creates services and provides the connection
 3. Services creates their own repository and provides the connection
 4. This repositories creates an abstract interface repository providing the connection
 5. At the end, this abstract interface repository communicates with the database interface, being reponsible to openning and closing the database connection

## TODO
Should investigate about what kind of architecture about project structure is the best. Currently, the main approach is to have processes as modules and the rest of code is separated in utils, managers, handlers, etc. Another approach is to only have core and modules folders, and have everything as modules. Processes would be a module with its own folder structure.
