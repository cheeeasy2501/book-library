CREATE TABLE users
(
    Id         serial primary key,
    FirstName  varchar(30) null,
    LastName   varchar(30) null,
    Email      varchar(30) unique null,
    UserName   varchar(50) unique,
    Password   varchar(255),
    Created_At timestamp with time zone,
    Updated_At timestamp with time zone
);

CREATE TABLE authors
(
    Id         serial primary key,
    FirstName  varchar(30),
    LastName   varchar(30),
    Created_At timestamp with time zone,
    Updated_At timestamp with time zone
);

CREATE TABLE books
(
    Id               serial primary key,
    House_Publish_Id integer,
    Title            varchar(30),
    Description      text,
    Link             varchar(100),
    In_Stock         smallint check (In_Stock >= 0),
    Created_At       timestamp with time zone,
    Updated_At       timestamp with time zone
);

CREATE TABLE author_books
(
    Id        serial primary key,
    Author_Id integer references authors (Id) on delete set null,
    Book_Id   integer references books (Id) on delete cascade
);

CREATE TABLE house_publishes
(
    Id         SERIAL PRIMARY KEY,
    Name       VARCHAR(100),
    Created_At timestamp with time zone,
    Updated_At timestamp with time zone
);

CREATE TABLE booking
(
    Id             uuid PRIMARY KEY,
    User_Id        integer REFERENCES users (id),
    Book_Id        integer REFERENCES books (id),
    Status         varchar(3),
    Start_DateTime timestamp with time zone,
    End_DateTime   timestamp with time zone,
    Created_At     timestamp with time zone,
    Updated_At     timestamp with time zone
);