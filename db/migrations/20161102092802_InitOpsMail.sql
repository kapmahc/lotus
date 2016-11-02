
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table mail_transport (
  id serial primary key,
  domain varchar(128) not null,
  transport varchar(128) not null default 'virtual:',
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_mail_transport_domain on mail_transport(domain);

create table mail_users (
  id serial primary key,
  mail varchar(255) not null,
  password varchar(255) not null,
  name varchar(128) not null,
  uid INTEGER not null,
  gid INTEGER not null,
  home varchar(255) not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_mail_users_mail on mail_users(mail);
create index idx_mail_users_name on mail_users(name);

create table mail_virtual (
  id serial primary key,
  source varchar(255) not null,
  target varchar(255) not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_mail_virtual_source on mail_virtual(source);
create index idx_mail_virtual_target on mail_virtual(target);

create view postfix_mailboxes as
  select mail, home||'/' as mailbox from mail_users
  union all
  select domain as mail, 'dummy' as mailbox from mail_transport;
  
create view postfix_virtual as
  select mail, mail as target from mail_users
  union all
  select target, source from mail_virtual;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop view postfix_virtual;
drop view postfix_mailboxes;
drop table mail_virtual;
drop table mail_users;
drop table mail_transport;
