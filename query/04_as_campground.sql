/*
Goal: learn how to rename a column or table with AS

The AS keyword sets an alias for the column in the query results
The AS keyword also sets an alias for the table in the query statement
The AS keyword can be used implicitly or explicitly

We have seen this query before - notice the new column name
*/
SELECT
    COUNT(1) AS num_campgrounds
FROM
    campground;
-- result
-- +-----------------+
-- | num_campgrounds |
-- +-----------------+
-- | 638             |
-- +-----------------+
/*
The AS is not required - you can use a space and a string like this
*/
SELECT
    COUNT(1) num_campgrounds
FROM
    campground;
-- result
-- +-----------------+
-- | num_campgrounds |
-- +-----------------+
-- | 638             |
-- +-----------------+
/*
We can also use the AS keyword to rename a table in the query

By default you can refer to column using the full table name
*/
SELECT
    campground.name
FROM
    campground
LIMIT 1;
-- result
-- +----------------------+
-- |         name         |
-- +----------------------+
-- | 277 North Campground |
-- +----------------------+
/*
With one table this is not required
Once there are multiple tables in the query
We need to specify the table name for columns with the same name in both tables
Otherwise the database will not know which column to return

The following query is like a nested loop between the campground and park tables
Here is some python pseudo code to illustrate the query:

result = []
for (campground in campgrounds) {
    for (park in parks) {
        result.append((campground.name, park.name))
    }
}
return result

You usually would not want to do this
Since the performance of the nested loop is O(n^2)
For a better approach see the JOIN exercise
*/
SELECT
    campground.name,
    park.name
FROM
    campground,
    park
LIMIT 1;
-- result
-- +----------------------+----------------------------+
-- |         name         |            name            |
-- +----------------------+----------------------------+
-- | 277 North Campground | Abraham Lincoln Birthplace |
-- +----------------------+----------------------------+
/*
For now it is important to know that you can rename a table in the sql query with AS
*/
SELECT
    c.name,
    p.name
FROM
    campground AS c,
    park AS p
LIMIT 1;
-- result
-- +----------------------+----------------------------+
-- |         name         |            name            |
-- +----------------------+----------------------------+
-- | 277 North Campground | Abraham Lincoln Birthplace |
-- +----------------------+----------------------------+
/*
The AS keyword is not required
We can also rename tables with a space and a string
This is the same as the last query
*/
SELECT
    c.name,
    p.name
FROM
    campground c,
    park p
LIMIT 1;
-- result
-- +----------------------+----------------------------+
-- |         name         |            name            |
-- +----------------------+----------------------------+
-- | 277 North Campground | Abraham Lincoln Birthplace |
-- +----------------------+----------------------------+
/*
Everything else is identical about the query

For example we can count the number of campgrounds that do not allow RVs
*/
SELECT
    COUNT(1) no_rv
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
    COUNT(1) yes_rv
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
