create table if not exists refresh_tokens
(
	id int not null auto_increment,
    deleted_at timestamp null,
	updated_at timestamp null,
	created_at timestamp not null,
    revoked int not null,
    expires_at timestamp not null,
	token_hash char(64) not null,
	token_nonce_hash char(64) not null,
    parent_token_hash char(64),
	user_id int not null,
	constraint refresh_tokens_pk
		primary key (id),
	constraint refresh_tokens_users_id_fk
		foreign key (user_id) references users (id)
			on update cascade on delete cascade
);