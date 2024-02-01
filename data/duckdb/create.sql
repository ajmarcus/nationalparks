CREATE SEQUENCE amenity_season_id_seq START 1;
	CREATE TABLE amenity_season (
		id UINTEGER DEFAULT nextval('amenity_season_id_seq') PRIMARY KEY,
		name VARCHAR(63)
	);
	INSERT INTO amenity_season (id, name) VALUES (0, 'None');
	CREATE SEQUENCE duration_id_seq START 1;
	CREATE TABLE duration (
		id UINTEGER DEFAULT nextval('duration_id_seq') PRIMARY KEY,
		name CHAR(1)
	);
	INSERT INTO duration (id, name) VALUES (0, 'N');
	CREATE SEQUENCE state_code_id_seq START 1;
	CREATE TABLE state_code (
		id UINTEGER DEFAULT nextval('state_code_id_seq') PRIMARY KEY,
		name CHAR(2)
	);
	INSERT INTO state_code (id, name) VALUES (0, 'NA');CREATE SEQUENCE park_id_seq START 1;
		CREATE TABLE park (
		id UINTEGER DEFAULT nextval('park_id_seq') PRIMARY KEY,
		name VARCHAR(255),
		city VARCHAR(255),
		description TEXT,
		state_code_id UINTEGER,
		url VARCHAR(255),
		FOREIGN KEY (state_code_id) REFERENCES state_code(id)
	);
	CREATE SEQUENCE park_fee_id_seq START 1;
	CREATE TABLE park_fee (
		id UINTEGER DEFAULT nextval('park_fee_id_seq') PRIMARY KEY,
		park_id UINTEGER,
		name VARCHAR(255),
		cost_cents UINTEGER,
		FOREIGN KEY (park_id) REFERENCES park(id)
	);
CREATE SEQUENCE campground_id_seq START 1;
		CREATE TABLE campground (
		id UINTEGER DEFAULT nextval('campground_id_seq') PRIMARY KEY,
		name VARCHAR(255),
		campsites_electrical_hookups UINTEGER,
		campsites_first_come_first_serve UINTEGER,
		campsites_reservable UINTEGER,
		campsites_total UINTEGER,
		has_camp_store_id UINTEGER,
		has_cell_phone_reception_id UINTEGER,
		has_laundry_id UINTEGER,
		is_rv_allowed UINTEGER,
		park_id UINTEGER,
		reservation_url VARCHAR(255),
		FOREIGN KEY (has_camp_store_id) REFERENCES amenity_season(id),
		FOREIGN KEY (has_cell_phone_reception_id) REFERENCES amenity_season(id),
		FOREIGN KEY (has_laundry_id) REFERENCES amenity_season(id),
		FOREIGN KEY (park_id) REFERENCES park(id)
	);
	CREATE SEQUENCE campground_fee_id_seq START 1;
	CREATE TABLE campground_fee (
		id UINTEGER DEFAULT nextval('campground_fee_id_seq') PRIMARY KEY,
		campground_id UINTEGER,
		name VARCHAR(255),
		cost_cents UINTEGER,
		FOREIGN KEY (campground_id) REFERENCES campground(id)
	);
CREATE SEQUENCE tour_id_seq START 1;
		CREATE TABLE tour (
		id UINTEGER DEFAULT nextval('tour_id_seq') PRIMARY KEY,
		name VARCHAR(255),
		description TEXT,
		duration_max UINTEGER,
		duration_min UINTEGER,
		duration_id UINTEGER,
		park_id UINTEGER,
		FOREIGN KEY (park_id) REFERENCES park(id),
		FOREIGN KEY (duration_id) REFERENCES duration(id)
	);
