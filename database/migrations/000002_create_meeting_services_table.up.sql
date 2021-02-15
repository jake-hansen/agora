create table if not exists meeting_services 
(
    id int auto_increment primary key,
    name varchar(100) not null,
    constraint meeting_services_name_uindex unique (name)
)