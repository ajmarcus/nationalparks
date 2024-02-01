/*
Goal: learn how to LIMIT the size of your query results

Give me the name of one campground

SELECT allows us to choose which columns to return
This is called projection in Edgar Codd's original paper

FROM specifies the table for our query
In this case, we want rows from camp_example
*/
SELECT
    name
FROM
    camp_example
LIMIT 1;
-- result
-- +----------------------+
-- |         name         |
-- +----------------------+
-- | 277 North Campground |
-- +----------------------+
