CREATE TABLE amenity_season (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL
	);
	INSERT INTO amenity_season (id, name) VALUES (0, 'None');
	CREATE TABLE duration (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL
	);
	INSERT INTO duration (id, name) VALUES (0, 'N');
	CREATE TABLE state_code (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL
	);
	INSERT INTO state_code (id, name) VALUES (0, 'NA');CREATE TABLE park (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL,
		city TEXT NOT NULL,
		state_code_id INTEGER NOT NULL,
		url TEXT NOT NULL,
		FOREIGN KEY (state_code_id) REFERENCES state_code(id)
	);
	CREATE TABLE park_fee (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		park_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		cost_cents INTEGER NOT NULL,
		FOREIGN KEY (park_id) REFERENCES park(id)
	);
CREATE TABLE campground (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL,
		campsites_electrical_hookups INTEGER NOT NULL,
		campsites_first_come_first_serve INTEGER NOT NULL,
		campsites_reservable INTEGER NOT NULL,
		campsites_total INTEGER NOT NULL,
		has_camp_store_id INTEGER NOT NULL,
		has_cell_phone_reception_id INTEGER NOT NULL,
		has_laundry_id INTEGER NOT NULL,
		is_rv_allowed INTEGER NOT NULL,
		park_id INTEGER NOT NULL,
		reservation_url TEXT NOT NULL,
		FOREIGN KEY (has_camp_store_id) REFERENCES amenity_season(id),
		FOREIGN KEY (has_cell_phone_reception_id) REFERENCES amenity_season(id),
		FOREIGN KEY (has_laundry_id) REFERENCES amenity_season(id),
		FOREIGN KEY (park_id) REFERENCES park(id)
	);
	CREATE TABLE campground_fee (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		campground_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		cost_cents INTEGER NOT NULL,
		FOREIGN KEY (campground_id) REFERENCES campground(id)
	);
CREATE TABLE tour (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL,
		duration_max INTEGER NOT NULL,
		duration_min INTEGER NOT NULL,
		duration_id INTEGER NOT NULL,
		park_id INTEGER NOT NULL,
		FOREIGN KEY (park_id) REFERENCES park(id),
		FOREIGN KEY (duration_id) REFERENCES duration(id)
	);
