create table if not exists oauth_info (
    id int auto_increment primary key,
    deleted_at timestamp null,
	created_at timestamp not null,
	updated_at timestamp not null,
    user_id int  not null,
    access_token text not null,
    refresh_token text not null,
    meeting_provider_id int not null,
    constraint oauth_info_meeting_providers_id_fk
        foreign key (meeting_provider_id) references meeting_providers (id)
            on update cascade on delete cascade,
    constraint oauth_info_user_id_fk
        foreign key (user_id) references users (id)
            on update cascade on delete cascade
);
