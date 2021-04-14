create table if not exists invites (
    id int auto_increment primary key,
    deleted_at timestamp null,
    updated_at timestamp null,
    created_at timestamp not null,
    meeting_id varchar(30) not null,
    invitee_id int not null,
    inviter_id int not null,
    meeting_start_time timestamp not null,
    meeting_duration int not null,
    meeting_title text null,
    meeting_description text null,
    meeting_platform_id int not null,
    meeting_join_url text not null,
    constraint invites_meeting_platforms_id_fk foreign key (meeting_platform_id) references meeting_platforms (id) on update cascade on delete cascade,
    constraint invites_users_id_fk foreign key (inviter_id) references users (id) on update cascade on delete cascade,
    constraint invites_users_id_fk_2 foreign key (invitee_id) references users (id) on update cascade on delete cascade
);