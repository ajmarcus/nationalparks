/*
Goal: learn how to connect information across tables using JOIN

Let's imagine you are planning a trip
You want to go hiking and camp at a national park
You will be traveling for a couple of weeks
So you want to make sure you'll be able to do laundry

You notice that there is a has_laundry_id column in the campground table
But you are confused by the values in the column
What does "1" mean?
*/
SELECT
    name, has_laundry_id
FROM campground
LIMIT 1;
-- result
-- +----------------------+----------------+
-- |         name         | has_laundry_id |
-- +----------------------+----------------+
-- | 277 North Campground | 1              |
-- +----------------------+----------------+
/*
First, you look at the number of campgrounds with each possible has_laundry_id value
*/
SELECT
    has_laundry_id,
    COUNT(1) AS num_campgrounds
FROM campground
GROUP BY 1
ORDER BY 1;
-- result
-- +----------------+-----------------+
-- | has_laundry_id | num_campgrounds |
-- +----------------+-----------------+
-- | 0              | 38              |
-- | 1              | 568             |
-- | 2              | 24              |
-- | 3              | 8               |
-- +----------------+-----------------+
/*
All of the rows in campground have 0, 1, 2, or 3 for has_laundry_id

Then look at the CREATE statement for the campground table using:

sqlite> .schema campground

You notice has_laundry_id has a foreign key reference to the id column in amenity_season table:

- FOREIGN KEY (has_laundry_id) REFERENCES amenity_season(id),

You look at the amenity_season table using:
*/
SELECT
    *
FROM amenity_season;
-- result
-- +----+------------------+
-- | id |       name       |
-- +----+------------------+
-- | 0  | None             |
-- | 1  | No               |
-- | 2  | Yes - seasonal   |
-- | 3  | Yes - year round |
-- +----+------------------+

/*
You see that the amenity_season has 4 unique rows
- With the values 0, 1, 2, and 3 for id column (same as has_laundry_id!)
- For example, the row with id=1 also has name="No"
- The foreign key reference from campground.has_laundry_id to amenity_season.id is a lookup
- Since has_laundry_id=1 for name="277 North Campground" in campground
- This corresponds to id=1 for name="No" in the amenity_season table
- Which indicates that "277 North Campground" does not have laundry

You can confirm your detective work by joining the campground and amenity_season tables:
*/
SELECT
    campground.name,
    campground.has_laundry_id,
    amenity_season.id,
    amenity_season.name
FROM campground
JOIN amenity_season
ON campground.has_laundry_id = amenity_season.id
WHERE campground.name = "277 North Campground";
-- result
-- +----------------------+----------------+----+------+
-- |         name         | has_laundry_id | id | name |
-- +----------------------+----------------+----+------+
-- | 277 North Campground | 1              | 1  | No   |
-- +----------------------+----------------+----+------+
