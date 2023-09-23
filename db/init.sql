CREATE TABLE user (
	id varchar(10) NOT NULL PRIMARY KEY,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	sub TEXT NOT NULL,
	picture TEXT,
  description TEXT DEFAULT "",
	history TEXT DEFAULT ""
);

CREATE TABLE art (
	id varchar(10) NOT NULL PRIMARY KEY,
	name varchar(255) NOT NULL,
	description text,
	userId varchar(10),
	price INT,
	status text
);

CREATE TABLE image (
	id varchar(10) NOT NULL PRIMARY KEY,
	url varchar(2048) NOT NULL,
	artId varchar(10)
);