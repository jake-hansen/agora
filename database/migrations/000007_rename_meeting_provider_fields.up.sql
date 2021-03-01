alter table oauth_info change meeting_provider_id meeting_platform_id int not null;

alter table oauth_info drop foreign key oauth_info_meeting_providers_id_fk;

alter table oauth_info
	add constraint oauth_info_meeting_platform_id_fk
		foreign key (meeting_platform_id) references meeting_platforms (id)
			on update cascade on delete cascade;

alter table meeting_platforms drop key meeting_providers_name_uindex;

alter table meeting_platforms
	add constraint meeting_platforms_name_uindex
		unique (name);

alter table meeting_platforms drop column redirect_url;