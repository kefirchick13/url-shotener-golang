CREATE TABLE urls
(
    id       serial not null unique,
    url     varchar(255) not null,
    alias varchar(255) not null unique, 
);

