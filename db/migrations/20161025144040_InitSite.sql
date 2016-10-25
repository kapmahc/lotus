
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
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
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
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

create table pages(
  id serial primary key,
  loc varchar(32) not null,
  picture varchar(255) not null,
  title varchar(255) not null,
  summary varchar(255) not null,
  href varchar(255) not null,
  lang varchar(8) not null default 'en-US',
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_pages_lang_loc on pages(lang, loc);
create index idx_pages_lang on pages(lang);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table pages;
drop table attachments;
drop table leave_words;
drop table notices;
drop table locales;
drop table settings;
