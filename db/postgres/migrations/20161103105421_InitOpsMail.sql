
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table mail_domains (
  id serial primary key,
  name varchar(128) not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_mail_domains_name on mail_domains(name);

create table mail_users (
  id serial primary key,
  domain_id int not null,
  email varchar(255) not null,
  name varchar(128) not null,
  password varchar(255) not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_mail_users_email on mail_users(email);
create index idx_mail_users_name on mail_users(name);

create table mail_aliases (
  id serial primary key,
  domain_id int not null,
  source varchar(255) not null,
  destination varchar(255) not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_mail_aliases_source on mail_aliases(source);
create index idx_mail_aliases_destination on mail_aliases(destination);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table mail_domains;
drop table mail_users;
drop table mail_aliases;
