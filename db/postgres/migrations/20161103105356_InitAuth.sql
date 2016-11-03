
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table users(
  id serial primary key,
  name varchar(32) not null,
  email varchar(255) not null,
  uid varchar(36) not null,
  password bytea,
  provider_id varchar(255) not null,
  provider_type varchar(32) not null,
  home varchar(255),
  logo varchar(255),
  sign_in_count int not null default 0,
  current_sign_in_at timestamp without time zone,
  current_sign_in_ip inet,
  last_sign_in_at timestamp without time zone,
  last_sign_in_ip inet,
  confirmed_at timestamp without time zone,
  locked_at timestamp without time zone,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_users_uid on users(uid);
create unique index idx_users_email on users(email);
create unique index idx_users_provider_id_type on users(provider_id, provider_type);
create index idx_users_name on users(name);
create index idx_users_provider_id on users(provider_id);
create index idx_users_provider_type on users(provider_type);

create table logs(
  id serial primary key,
  user_id int not null,
  type varchar(8) not null default 'info',
  message varchar(255) not null,
  created_at timestamp without time zone not null default now()
);
create index idx_logs_type on logs(type);

create table roles(
  id serial primary key,
  name varchar(32) not null,
  resource_id int not null default 0,
  resource_type varchar(255) not null default '-',
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_roles_name_resource_type_id on roles(name, resource_type, resource_id);
create index idx_roles_name on roles(name);
create index idx_roles_resource_type on roles(resource_type);

create table permissions(
  id serial primary key,
  user_id int not null,
  role_id int not null,
  start_up date not null default current_date,
  shut_down date not null default '2016-10-07',
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_permissions_user_role on permissions(user_id, role_id);

create table settings(
  id serial primary key,
  key varchar(255) not null,
  val bytea not null,
  flag boolean not null default false,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_settings_key on settings(key);

create table locales(
  id serial primary key,
  code varchar(255) not null,
  lang varchar(8) not null default 'en-US',
  message text not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_locales_code_lang on locales(code, lang);
create index idx_locales_code on locales(code);
create index idx_locales_lang on locales(lang);

create table notices(
  id serial primary key,
  lang varchar(8) not null default 'en-US',
  content text not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create index idx_notices_lang on notices(lang);

create table leave_words(
  id serial primary key,
  content text not null,
  created_at timestamp without time zone not null default now()
);

create table attachments(
  id serial primary key,
  title varchar(255) not null,
  name varchar(255) not null,
  media_type varchar(32) not null,
  summary varchar(255),
  user_id int not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_attachments_name on attachments(name);
create index idx_attachments_title on attachments(title);
create index idx_attachments_media_type on attachments(media_type);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table attachments;
drop table leave_words;
drop table notices;
drop table locales;
drop table settings;

drop table permissions;
drop table roles;
drop table logs;
drop table users;
