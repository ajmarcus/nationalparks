/*
Goal: learn about NULL & how to return rows that match a boolean expression with WHERE

Which campgrounds allow RVs?

Remember the full camp_example table from 00_create_camp_example:
+-------------------------+----------------------+---------------+
|          name           | campsites_reservable | is_rv_allowed |
+-------------------------+----------------------+---------------+
| 277 North Campground    | 1                    |               |
| Abrams Creek Campground | 16                   | 1             |
| Adirondack Shelters     | 2                    | 0             |
+-------------------------+----------------------+---------------+

Sqlite does not have a boolean type
We use:
- 1 for true
- 0 for false
- NULL (blank) for unknown
*/

/*
We don't know if 227 North Campground allows RVs
*/
SELECT
    name
FROM
    camp_example
WHERE
    is_rv_allowed IS NULL;
-- result
-- +----------------------+
-- |         name         |
-- +----------------------+
-- | 277 North Campground |
-- +----------------------+

/*
We have NOT NULL (aka known) values for the other two campgrounds
*/
SELECT
    name
FROM
    camp_example
WHERE
    is_rv_allowed IS NOT NULL;
-- result
-- +-------------------------+
-- |          name           |
-- +-------------------------+
-- | Abrams Creek Campground |
-- | Adirondack Shelters     |
-- +-------------------------+

/*
Abrams Creek Campground allows RVs
*/
SELECT
    name
FROM
    camp_example
WHERE
    is_rv_allowed  = 1;
-- result
-- +-------------------------+
-- |          name           |
-- +-------------------------+
-- | Abrams Creek Campground |
-- +-------------------------+

/*
Adirondack Shelters does not allow RVs
*/
SELECT
    name
FROM
    camp_example
WHERE
    is_rv_allowed  = 0;
-- result
-- +---------------------+
-- |        name         |
-- +---------------------+
-- | Adirondack Shelters |
-- +---------------------+
