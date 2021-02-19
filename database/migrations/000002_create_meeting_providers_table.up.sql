create table if not exists meeting_providers
(
    id int auto_increment primary key,
    deleted_at timestamp null,
	created_at timestamp not null,
	updated_at timestamp not null,
    name varchar(100) not null,
    constraint meeting_providers_name_uindex unique (name)
)