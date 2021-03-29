create table accounts
(
	account_id serial not null,
	account_name varchar not null,
	balance decimal not null
);

create unique index accounts_account_id_uindex
	on accounts (account_id);

alter table accounts
	add constraint accounts_pk
		primary key (account_id);

