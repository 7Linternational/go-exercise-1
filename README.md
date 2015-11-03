# Setup GO development environment
- https://golang.org/doc/install

# Setup Oracle EXPRESS
- http://www.oracle.com/technetwork/database/database-technologies/express-edition/overview/index.html

# Setup Oracle environment and access from GO
1. Install the listed packages for your target platform from the following links (Needs and
Oracle.com account, registration is free)
    1. Linux - http://www.oracle.com/technetwork/topics/linuxx86­64soft­092277.html​
    2. Windows - http://www.oracle.com/technetwork/topics/winx64soft­089540.html
2. Make sure all zip files are extracted to the same folder. For example
‘/home/thanasis/instantclient_12_1’ .
3. Setup the following environmental variables.
    1. `ORACLE_HOME=/path/to/instantclient_12_1`
    2. `LD_LIBRARY_FLAGS=$ORACLE_HOME`
    3. `LD_LIBRARY_PATH=$ORACLE_HOME:$LD_LIBRARY_PATH`
    4. `CGO_CFLAGS=­I$ORACLE_HOME`
    5. `CGO_LDFLAGS="­L$ORACLE_HOME ­lclntsh"`
    6. `C_INCLUDE_PATH=$ORACLE_HOME/sdk/include/`
4. Install the `libaio` library if needed (ex. `sudo apt­get install libaio1 libaio­dev`).
5. Install the Oracle GO library, by executing: `go get gopkg.in/rana/ora.v3`

# Quick GO tutorial
- https://gobyexample.com/

# Exercise flow
- Connect with Oracle Express DB
- Create Table `users` with the following fields
    - `id`
    - `username`
    - `password`
    - `description`
    - `timestamp`
        - The `timestamp` field should be automatically inserted and updated after every operation in the `users` table.
- Execute the following operations step by step using GO code and the ora library
    - Single execution mode
        1. Create 4 `users` records
        2. Select and display all `users` records
        3. Update 2 `users` records' description field
        4. Delete the ones that where not updated
        5. Select and display all `users` records
    - Concurrent mode
        1. Create 4 `users` records by calling a different go routine for 2 of them each time.
        2. When both of the previous routines finishes select and display all `users` records
        3. Update 2 `users` records' description field, create a go routine for each record update and a third go routine that will delete the ones that will not be updated. Once all 3 routines finish select and display all `users` records.

