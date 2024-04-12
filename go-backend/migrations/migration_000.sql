create table m_user
(
    name  varchar(20)  not null,
    email varchar(100) not null
        constraint m_user_pk
            primary key
);

alter table m_user
    owner to root;

create table sessions
(
    id         varchar(128) not null
        constraint sessions_pk
            primary key,
    expires_at integer      not null,
    data       text         not null,
    "user"     varchar(100) not null
        constraint sessions_m_user_email_fk
            references m_user
);

alter table sessions
    owner to root;

create table event
(
    name        varchar(80)  not null,
    description text,
    adress_name varchar(200) not null,
    longitude   double precision,
    lattitude   double precision,
    creator     varchar(100) not null
        constraint event_m_user_email_fk
            references m_user,
    id          serial
        constraint event_pk
            primary key
);

alter table event
    owner to root;

create table participates
(
    participant varchar(100) not null,
    event_id    integer      not null,
    constraint participates_pk
        primary key (participant, event_id)
);

alter table participates
    owner to root;
