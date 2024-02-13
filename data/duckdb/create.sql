CREATE SEQUENCE amenity_season_id_seq START 1;
	CREATE TABLE amenity_season (
		id UINTEGER DEFAULT nextval('amenity_season_id_seq') PRIMARY KEY NOT NULL,
		name VARCHAR(63) NOT NULL
	);
	INSERT INTO amenity_season (id, name) VALUES (0, 'None');
	CREATE SEQUENCE duration_id_seq START 1;
	CREATE TABLE duration (
		id UINTEGER DEFAULT nextval('duration_id_seq') PRIMARY KEY NOT NULL,
		name CHAR(1) NOT NULL
	);
	INSERT INTO duration (id, name) VALUES (0, 'N');
	CREATE SEQUENCE fee_id_seq START 1;
	CREATE TABLE fee (
		id UINTEGER DEFAULT nextval('fee_id_seq') PRIMARY KEY NOT NULL,
		name VARCHAR(63) NOT NULL
	);
	INSERT INTO fee (id, name) VALUES (0, 'None');
	CREATE SEQUENCE state_code_id_seq START 1;
	CREATE TABLE state_code (
		id UINTEGER DEFAULT nextval('state_code_id_seq') PRIMARY KEY NOT NULL,
		name CHAR(2) NOT NULL
	);
	INSERT INTO state_code (id, name) VALUES (0, 'NA');CREATE SEQUENCE park_id_seq START 1;
		CREATE TABLE park (
		id UINTEGER DEFAULT nextval('park_id_seq') PRIMARY KEY NOT NULL,
		name VARCHAR(255) NOT NULL,
		city VARCHAR(255) NOT NULL,
		state_code_id UINTEGER NOT NULL,
		url VARCHAR(255) NOT NULL,
		FOREIGN KEY (state_code_id) REFERENCES state_code(id)
	);
	CREATE SEQUENCE park_fee_id_seq START 1;
	CREATE TABLE park_fee (
		id UINTEGER DEFAULT nextval('park_fee_id_seq') PRIMARY KEY NOT NULL,
		park_id UINTEGER NOT NULL,
		fee_id UINTEGER NOT NULL,
		cost_cents UINTEGER NOT NULL,
		FOREIGN KEY (park_id) REFERENCES park(id),
		FOREIGN KEY (fee_id) REFERENCES fee(id)
	);
CREATE SEQUENCE campground_id_seq START 1;
		CREATE TABLE campground (
		id UINTEGER DEFAULT nextval('campground_id_seq') PRIMARY KEY NOT NULL,
		name VARCHAR(255) NOT NULL,
		campsites_electrical_hookups UINTEGER NOT NULL,
		campsites_first_come_first_serve UINTEGER NOT NULL,
		campsites_reservable UINTEGER NOT NULL,
		campsites_total UINTEGER NOT NULL,
		has_camp_store_id UINTEGER NOT NULL,
		has_cell_phone_reception_id UINTEGER NOT NULL,
		has_laundry_id UINTEGER NOT NULL,
		is_rv_allowed UINTEGER NOT NULL,
		park_id UINTEGER NOT NULL,
		reservation_url VARCHAR(255) NOT NULL,
		FOREIGN KEY (has_camp_store_id) REFERENCES amenity_season(id),
		FOREIGN KEY (has_cell_phone_reception_id) REFERENCES amenity_season(id),
		FOREIGN KEY (has_laundry_id) REFERENCES amenity_season(id),
		FOREIGN KEY (park_id) REFERENCES park(id)
	);
CREATE SEQUENCE tour_id_seq START 1;
		CREATE TABLE tour (
		id UINTEGER DEFAULT nextval('tour_id_seq') PRIMARY KEY NOT NULL,
		name VARCHAR(255) NOT NULL,
		duration_max UINTEGER NOT NULL,
		duration_min UINTEGER NOT NULL,
		duration_id UINTEGER NOT NULL,
		park_id UINTEGER NOT NULL,
		FOREIGN KEY (park_id) REFERENCES park(id),
		FOREIGN KEY (duration_id) REFERENCES duration(id)
	);
