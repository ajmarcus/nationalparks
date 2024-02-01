/*
Goal: learn how to COUNT

We will move from the camp_example -> campground table
Since it has a complete list of US national park campgrounds

COUNT(1) means count the number of rows returned by the query
- COUNT is an aggregate function that operates on a set of rows
- Without a WHERE clause, it counts all rows in the table
*/
SELECT
    COUNT(1)
FROM
    camp_example;
-- result
-- +----------+
-- | COUNT(1) |
-- +----------+
-- | 3        |
-- +----------+
/*
With a WHERE clause, COUNT returns number the rows that satisfy the condition

For example there is 1 row where is_rv_allowed is NULL
*/
SELECT
    COUNT(1)
FROM
    camp_example
WHERE
    is_rv_allowed IS NULL;
-- result
-- +----------+
-- | COUNT(1) |
-- +----------+
-- | 1        |
-- +----------+
/*
There are 2 rows where is_rv_allowed is not NULL
*/
SELECT
    COUNT(1)
FROM
    camp_example
WHERE
    is_rv_allowed IS NOT NULL;
-- result
-- +----------+
-- | COUNT(1) |
-- +----------+
-- | 2        |
-- +----------+
/*
We will now swith to use the campground table for the rest of the queries
There are 638 rows in the campground table
*/
SELECT
    COUNT(1)
FROM
    campground;
-- result
-- +----------+
-- | COUNT(1) |
-- +----------+
-- | 638      |
-- +----------+
/*
COUNT(*) usually is a synonym for COUNT(1)
- Technically for some databases it means count the number of rows with any non-null value
*/
SELECT
    COUNT(*)
FROM
    campground;
-- result
-- +----------+
-- | COUNT(*) |
-- +----------+
-- | 638      |
-- +----------+
