create table if not exists meeting_providers
(
    id int auto_increment primary key,
    name varchar(100) not null,
    constraint meeting_providers_name_uindex unique (name)
)