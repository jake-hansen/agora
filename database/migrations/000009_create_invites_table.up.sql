create table if not exists invites (
    id int auto_increment primary key,
    deleted_at timestamp null,
	updated_at timestamp null,
	created_at timestamp not null,
    meeting_id varchar(30) not null,
    invitee_id int not null,
    inviter_id int not null,
    meeting_end_time timestamp not null,
    constraint invites_users_id_fk foreign key (inviter_id) references users (id) on update cascade on delete cascade,
    constraint invites_users_id_fk_2 foreign key (invitee_id) references users (id) on update cascade on delete cascade
);