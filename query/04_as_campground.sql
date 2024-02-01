/*
Goal: learn how to rename a column with AS

The AS keyword sets an alias for the column in the query results

We have seen this query before - notice the new column name
*/
SELECT
    COUNT(1) as num_campgrounds
FROM
    campground;
-- result
-- +-----------------+
-- | num_campgrounds |
-- +-----------------+
-- | 638             |
-- +-----------------+
/*
Everything else is identical about the query

For example we can count the number of campgrounds that do not allow RVs
*/
SELECT
    COUNT(1) as no_rv
FROM
    campground
WHERE
    is_rv_allowed = 0;
-- result
-- +-------+
-- | no_rv |
-- +-------+
-- | 340   |
-- +-------+
/*
And we can count the number of campgrounds that allow RVs
*/
SELECT
    COUNT(1) as yes_rv
FROM
    campground
WHERE
    is_rv_allowed = 1;
-- result
-- +--------+
-- | yes_rv |
-- +--------+
-- | 298    |
-- +--------+
