
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table vpn_users (
    id serial primary key,
    name varchar(32) not null,
    password varchar(255) not null,
    details text,
    online bool not null default false,
    enable bool not null default false,
    start_up date not null default '2016-11-02',
    shut_down date not null default current_date,
    created_at timestamp without time zone not null default now(),
    updated_at timestamp without time zone not null
);
create unique index idx_vpn_users_name on vpn_users(name);

create table vpn_logs (
    id serial primary key,
    user_id int not null,
    trusted_ip inet,
    trusted_port smallint,
    remote_ip inet,
    remote_port smallint,
    start_time timestamp without time zone not null default CURRENT_TIMESTAMP,
    end_time timestamp without time zone,
    received float not null default '0',
    send float not null default '0'
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table vpn_logs;
drop table vpn_users;
