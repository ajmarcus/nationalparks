/*
Goal: learn how aggregate over subsets with GROUP BY

How many campgrounds allow or do not allow RVs?
- Last two queries but expressed in a general way
*/
SELECT
    is_rv_allowed,
    count(1) as num_campgrounds
FROM
    campground
GROUP BY
    1;
-- result
-- +---------------+-----------------+
-- | is_rv_allowed | num_campgrounds |
-- +---------------+-----------------+
-- | 0             | 340             |
-- | 1             | 298             |
-- +---------------+-----------------+
