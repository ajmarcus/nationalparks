/*
Goal: learn CREATE, INSERT, SELECT and FROM

A relational database defines a set of tables with typed columns
- A table is similar to a set of objects
- In Edgar Codd's original paper, he used the word relations for tables
- You can think of a table like a spreadsheet with named columns (fields) and rows (records)
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
