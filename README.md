# nationalparks

Snapshot of the united states national park service data from their api.

Data retrieved from the National Park Service.

Api Guide:

https://www.nps.gov/subjects/developer/guides.htm

Terms of Use:

https://www.nps.gov/aboutus/disclaimer.htm

Copyright law does not protect “any work of the U.S. Government” where “a work prepared by an officer or employee of the U.S. Government as part of that person's official duties” (See, 17 U.S.C. §§ 101, 105). Thus, material created by the NPS and presented on this website, unless otherwise indicated, is generally considered in the public domain. It may be distributed or copied as permitted by applicable law.

### Run

1. First create a .env file based on the example and replace the NPS_API_KEY

```bash
cp .env.example .env
```

2. Run the fetch and generate the sql

```bash
make run && cd data
```

3. Start sqlite

```bash
sqlite3 nps.sqlite
```

4. Execute the queries

```bash
.read ./sqlite/create.sql
.read ./insert/create.sql
.read ./insert/parks.sql
.read ./insert/campgrounds.sql
.read ./insert/tours.sql
```

5. Validate sqlite

```sql
select count(1) from amenity_season;
select count(1) from campground;
select count(1) from campground_fee;
select count(1) from duration;
select count(1) from park;
select count(1) from park_fee;
select count(1) from state_code;
select count(1) from tour;
```

Expected output:

```bash
sqlite> select count(1) from amenity_season;
4
sqlite> select count(1) from campground;
638
sqlite> select count(1) from campground_fee;
1148
sqlite> select count(1) from duration;
4
sqlite> select count(1) from park;
471
sqlite> select count(1) from park_fee;
522
sqlite> select count(1) from state_code;
54
sqlite> select count(1) from tour;
617
```

6. Start duckdb

```bash
duckdb nps.duckdb
```

7. Execute the queries

```bash
.read ./duckdb/create.sql
.read ./insert/create.sql
.read ./insert/parks.sql
.read ./insert/campgrounds.sql
.read ./insert/tours.sql
```

5. Validate duckdb

```sql
select count(1) from amenity_season;
select count(1) from campground;
select count(1) from campground_fee;
select count(1) from duration;
select count(1) from park;
select count(1) from park_fee;
select count(1) from state_code;
select count(1) from tour;
```

Expected output:

```bash
D select count(1) from amenity_season;
┌──────────┐
│ count(1) │
│  int64   │
├──────────┤
│        4 │
└──────────┘
D select count(1) from campground;
┌──────────┐
│ count(1) │
│  int64   │
├──────────┤
│      638 │
└──────────┘
D select count(1) from campground_fee;
┌──────────┐
│ count(1) │
│  int64   │
├──────────┤
│     1148 │
└──────────┘
D select count(1) from duration;
┌──────────┐
│ count(1) │
│  int64   │
├──────────┤
│        4 │
└──────────┘
D select count(1) from park;
┌──────────┐
│ count(1) │
│  int64   │
├──────────┤
│      471 │
└──────────┘
D select count(1) from park_fee;
┌──────────┐
│ count(1) │
│  int64   │
├──────────┤
│      522 │
└──────────┘
D select count(1) from state_code;
┌──────────┐
│ count(1) │
│  int64   │
├──────────┤
│       54 │
└──────────┘
D select count(1) from tour;
┌──────────┐
│ count(1) │
│  int64   │
├──────────┤
│      617 │
└──────────┘
```
