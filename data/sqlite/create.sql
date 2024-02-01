CREATE TABLE amenity_season (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);
	INSERT INTO amenity_season (id, name) VALUES (0, 'None');
	CREATE TABLE duration (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);
	INSERT INTO duration (id, name) VALUES (0, 'N');
	CREATE TABLE state_code (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT
	);
	INSERT INTO state_code (id, name) VALUES (0, 'NA');CREATE TABLE park (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		city TEXT,
		description TEXT,
		state_code_id INTEGER,
		url TEXT,
		FOREIGN KEY (state_code_id) REFERENCES state_code(id)
	);
	CREATE TABLE park_fee (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		park_id INTEGER,
		name TEXT,
		cost_cents INTEGER,
		FOREIGN KEY (park_id) REFERENCES park(id)
	);
CREATE TABLE campground (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		campsites_electrical_hookups INTEGER,
		campsites_first_come_first_serve INTEGER,
		campsites_reservable INTEGER,
		campsites_total INTEGER,
		has_camp_store_id INTEGER,
		has_cell_phone_reception_id INTEGER,
		has_laundry_id INTEGER,
		is_rv_allowed INTEGER,
		park_id INTEGER,
		reservation_url TEXT,
		FOREIGN KEY (has_camp_store_id) REFERENCES amenity_season(id),
		FOREIGN KEY (has_cell_phone_reception_id) REFERENCES amenity_season(id),
		FOREIGN KEY (has_laundry_id) REFERENCES amenity_season(id),
		FOREIGN KEY (park_id) REFERENCES park(id)
	);
	CREATE TABLE campground_fee (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		campground_id INTEGER,
		name TEXT,
		cost_cents INTEGER,
		FOREIGN KEY (campground_id) REFERENCES campground(id)
	);
CREATE TABLE tour (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT,
		duration_max INTEGER,
		duration_min INTEGER,
		duration_id INTEGER,
		park_id INTEGER,
		FOREIGN KEY (park_id) REFERENCES park(id),
		FOREIGN KEY (duration_id) REFERENCES duration(id)
	);
