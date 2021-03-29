create table transactions
(
	transaction_id serial not null,
	unique_id varchar not null,
	state varchar null,
	amount decimal not null,
	status varchar not null,
	source_type varchar not null,
	account_id int
		constraint transactions_accounts_account_id_fk
			references accounts
);

create unique index transactions_unique_id_uindex
	on transactions (unique_id);

create unique index transactions_transaction_id_uindex
	on transactions (transaction_id);

alter table transactions
	add constraint transactions_pk
		primary key (transaction_id);

