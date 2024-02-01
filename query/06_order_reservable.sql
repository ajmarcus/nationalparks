/*
Goal: learn how ORDER BY your results

Which campgroud has the most reservable campsites?
- Notice we use LIMIT 1 to only return the first result
*/
SELECT
    name, campsites_reservable
FROM
    campground
ORDER BY
    campsites_reservable DESC
LIMIT 1;
-- result
-- +-----------------------+----------------------+
-- |         name          | campsites_reservable |
-- +-----------------------+----------------------+
-- | Bridge Bay Campground | 432                  |
-- +-----------------------+----------------------+

/*
From campgrounds that allow reservations, which one has the fewest reservable campsites?
- Notice that we use WHERE to filter campgrounds with 0 reservable campsites
*/
SELECT
    name, campsites_reservable
FROM
    campground
WHERE campsites_reservable > 0
ORDER BY
    campsites_reservable ASC
LIMIT 1;
-- result
-- +----------------------+----------------------+
-- |         name         | campsites_reservable |
-- +----------------------+----------------------+
-- | 277 North Campground | 1                    |
-- +----------------------+----------------------+
