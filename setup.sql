/*CREATE TABLE IF NOT EXISTS intent.intents (
		id int primary key AUTO_INCREMENT,
		name varchar(255),
		label varchar(255),
		day_of_the_week varchar(255),
		start_tiime varchar(255),
		end_time varchar(255),
		minimum_cell_offset int,
		maximum_cell_offset int
);*/


CREATE TABLE IF NOT EXISTS intent.intents (
		id int primary key AUTO_INCREMENT,
		name varchar(255),
		description varchar(255),
		ric_id varchar(255),
		policy_id int,
		service_id varchar(255),
		policy_type_id int
);
