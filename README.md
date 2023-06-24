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
