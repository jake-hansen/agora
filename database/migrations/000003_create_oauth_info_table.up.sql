create table if not exists oauth_info (
    id int auto_increment primary key,
    user_id int  not null,
    access_token text not null,
    refresh_token text not null,
    meeting_service_id int not null,
    constraint oauth_info_meeting_providers_id_fk
        foreign key (meeting_service_id) references meeting_providers (id)
            on update cascade on delete cascade,
    constraint oauth_info_user_id_fk
        foreign key (user_id) references users (id)
            on update cascade on delete cascade
);
