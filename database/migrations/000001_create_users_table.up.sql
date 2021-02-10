create table if not exists users
(
	id int not null auto_increment,
	username varchar(100) not null,
	password char(60) not null,
	deleted_at timestamp null,
	created_at timestamp not null,
	updated_at timestamp not null,
	firstname varchar(255) not null,
	lastname varchar(255) not null,
	constraint users_id_uindex
		unique (id),
	constraint users_username_uindex
		unique (username),
    primary key (id)
);