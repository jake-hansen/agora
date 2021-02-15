create table if not exists oauth_tokens (
    id int auto_increment primary key,
    user_id int  not null,
    access_token text not null,
    refresh_token text not null,
    meeting_service_id int not null,
    constraint oauth_tokens_meeting_services_id_fk
        foreign key (meeting_service_id) references meeting_services (id)
            on update cascade on delete cascade,
    constraint oauth_tokens_user_id_fk
        foreign key (user_id) references users (id)
            on update cascade on delete cascade
);
