alter table oauth_info add token_type text not null;
alter table oauth_info add expiry timestamp not null;