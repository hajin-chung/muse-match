-- date must be in YYYYMMDD format
-- entity related to a single image: image id is entity id

CREATE TABLE user (
	id varchar(10) NOT NULL PRIMARY KEY,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	sub TEXT NOT NULL,
	picture TEXT NOT NULL,
	description TEXT DEFAULT "",
	note TEXT DEFAUTL "",
	is_owner INT NOT NULL,
	
	instagram_id TEXT,
	facebook_id TEXT,
	twitter_id TEXT
);

CREATE TABLE user_link (
	id varchar(10) NOT NULL PRIMARY KEY,
	user_id varchar(10) NOT NULL,
	content TEXT DEFAULT ""
);

CREATE TABLE user_history (
	id varchar(10) NOT NULL PRIMARY KEY,
	user_id varchar(10) NOT NULL,
	title TEXT NOT NULL,
	content TEXT NOT NULL
);

CREATE TABLE user_art_list (
	id varchar(10) NOT NULL PRIMARY KEY,
	user_id varchar(10) NOT NULL,
	title TEXT DEFAULT ""
);

CREATE TABLE user_art_list_item (
	list_id varchar(10) NOT NULL,
	art_id varchar(10) NOT NULL,
	idx INTEGER,
	PRIMARY KEY (list_id, art_id)
);

CREATE TABLE user_like_user (
	user_id varchar(10) NOT NULL,
	like_user_id varchar(10) NOT NULL,
	PRIMARY KEY(user_id, like_user_id)
);

CREATE TABLE user_like_art (
	user_id varchar(10) NOT NULL,
	art_id varchar(10) NOT NULL,
	PRIMARY KEY(user_id, art_id)
);

CREATE TABLE user_like_place (
	user_id varchar(10) NOT NULL,
	place_id varchar(10) NOT NULL,
	PRIMARY KEY(user_id, place_id)
);

CREATE TABLE art (
	id varchar(10) NOT NULL PRIMARY KEY,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	user_id varchar(10) NOT NULL,
	price INT,
	info TEXT NOT NULL
);

CREATE TABLE art_tag (
	art_id varchar(10) NOT NULL,
	tag TEXT NOT NULL,
	PRIMARY KEY(art_id, tag)
);

CREATE TABLE art_image (
	id varchar(10) NOT NULL PRIMARY KEY,
	art_id varchar(10) NOT NULL,
	idx INT NOT NULL
);

CREATE TABLE place (
	id varchar(10) NOT NULL PRIMARY KEY,
	title TEXT DEFAULT "",
	address TEXT DEFAULT "", 
	user_id varchar(10) NOT NULL,
	
	instagram_id TEXT,
	facebook_id TEXT,
	twitter_id TEXT
);

CREATE TABLE place_link (
	id varchar(10) NOT NULL PRIMARY KEY,
	place_id varchar(10) NOT NULL,
	content TEXT DEFAULT ""
);

CREATE TABLE place_image (
	id varchar(10) NOT NULL PRIMARY KEY,
	place_id varchar(10) NOT NULL,
	idx INT NOT NULL
);

CREATE TABLE place_location (
	id varchar(10) NOT NULL PRIMARY KEY,
	place_id varchar(10) NOT NULL,
	title TEXT NOT NULL,
	description TEXT NOT NULL
);

CREATE TABLE exhibit (
	id varchar(10) NOT NULL PRIMARY KEY,
	art_id varchar(10) NOT NULL,
	location_id varchar(10) NOT NULL,
	start_date TEXT NOT NULL,
	end_date TEXT NOT NULL,
	state TEXT NOT NULL -- "WAIT" | "INSTALL" | "FAIL" | "EXHIBIT" | "DONE"
);
