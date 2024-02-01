/*
Goal: learn CREATE, INSERT, SELECT and FROM

A relational database defines a set of tables with typed columns
- A table is similar to a set of objects
- In Edgar Codd's original paper, he used the word relations for tables
- You can think of a table like a spreadsheet with named columns (fields) and rows (records)

If you don't want to install sqlite3 locally, or you have trouble you can run SELECT queries in your browser:

https://sqlime.org/#https://raw.githubusercontent.com/ajmarcus/nationalparks/main/data/nps.sqlite

Just note that the CREATE and INSERT queries covered below will not work in the browser.

If you want to install sqlite3, first check if you have it already:
- Open a terminal
- Type "sqlite3" and press enter
```bash
sqlite3
```
- If you see a prompt like "sqlite>", then you have sqlite3 installed
```bash
SQLite version 3.43.2 2023-10-10 13:08:14
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
sqlite>
```

If you don't already have it, this is how to install sqlite3:
- Go to https://www.sqlite.org/download.html
- Choose the "Precompiled Binaries For" Mac OS X or Windows section
- Download the sqlite-tools-*.zip file
- Unzip and you should have a sqlite3 executable
- On Mac you will need to right click on the sqlite3 file and choose "Open" to bypass the security warning
- On Windows, I'm not sure what you need to do
- On Mac, you can then move `sqlite3` to `/usr/local/bin` to make it available from the command line
```bash
mv ~/Downloads/sqlite-tools-osx-x64-3450100/sqlite3 /usr/local/bin
```
- Now you should be able to run `sqlite3` from the command line

Use these settings in your sqlite3 shell locally to get pretty output:
sqlite> .header on
sqlite> .mode table
*/
/*
Here's an example table about campgrounds:
- id is a unique identifier for each row
- name is the name of the campground
- campsites_reservable is the number of campsites that can be reserved
- is_rv_allowed is a boolean value indicating whether RVs are allowed (1 for true, 0 for false, NULL for unknown)
*/
CREATE TABLE IF NOT EXISTS camp_example (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL,
    campsites_reservable INTEGER NOT NULL,
    is_rv_allowed INTEGER
);
/*
Once we have CREATE-d a table, we can INSERT rows

SQL uses ternary logic:
- all columns can be NULL (unless you have set a NOT NULL constraint like above)
- a binary expression can evaluate to true, false, or NULL (aka unknown)

This query will fail because we required that
- campsites_reservable be set to a non-NULL value with:
- "campsites_reservable INTEGER NOT NULL"
*/
INSERT INTO camp_example (name, campsites_reservable, is_rv_allowed) VALUES ('ERROR!', NULL, 1);
-- result
-- Runtime error: NOT NULL constraint failed: camp_example.campsites_reservable (19)
/*
Here are some successful INSERTs

In this case we have set:
- name to '277 North Campground'
- campsites_reservable to 1
- is_rv_allowed to NULL (meaning unknown)
*/
INSERT INTO camp_example (name, campsites_reservable, is_rv_allowed) VALUES ('277 North Campground', 1, NULL);
/*
In this case we have set:
- name to 'Abrams Creek Campground'
- campsites_reservable to 16
- is_rv_allowed to 1 (meaning true)
*/
INSERT INTO camp_example (name, campsites_reservable, is_rv_allowed) VALUES ('Abrams Creek Campground', 16, 1);
/*
In this case we have set:
- name to 'Adirondack Shelters'
- campsites_reservable to 2
- is_rv_allowed to 0 (meaning false)
*/
INSERT INTO camp_example (name, campsites_reservable, is_rv_allowed) VALUES ('Adirondack Shelters', 2, 0);
/*
After INSERT-ing a row, we can SELECT columns from the table
*/
SELECT name, campsites_reservable, is_rv_allowed
FROM camp_example;
-- result
-- +-------------------------+----------------------+---------------+
-- |          name           | campsites_reservable | is_rv_allowed |
-- +-------------------------+----------------------+---------------+
-- | 277 North Campground    | 1                    |               |
-- | Abrams Creek Campground | 16                   | 1             |
-- | Adirondack Shelters     | 2                    | 0             |
-- +-------------------------+----------------------+---------------+
/*
The "*" is a wildcard character that means "all columns"
- The result includes the autoincrementing integer id column
- We did not explicitly INSERT ids
*/
SELECT *
FROM camp_example;
-- result
-- +----+-------------------------+----------------------+---------------+
-- | id |          name           | campsites_reservable | is_rv_allowed |
-- +----+-------------------------+----------------------+---------------+
-- | 1  | 277 North Campground    | 1                    |               |
-- | 2  | Abrams Creek Campground | 16                   | 1             |
-- | 3  | Adirondack Shelters     | 2                    | 0             |
-- +----+-------------------------+----------------------+---------------+
