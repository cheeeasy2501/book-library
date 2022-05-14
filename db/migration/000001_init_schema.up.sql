CREATE TABLE users
(
    Id         SERIAL PRIMARY KEY,
    FirstName  VARCHAR(30)        NULL,
    LastName   VARCHAR(30)        NULL,
    Email      VARCHAR(30) UNIQUE NULL,
    UserName   VARCHAR(50) UNIQUE,
    Password   varchar(255),
    Created_At timestamp with time zone,
    Updated_At timestamp with time zone
);

CREATE TABLE author
(
    Id         SERIAL PRIMARY KEY,
    FirstName  VARCHAR(30),
    LastName   VARCHAR(30),
    Created_At timestamp with time zone,
    Updated_At timestamp with time zone
);

CREATE TABLE books
(
    Id               SERIAL PRIMARY KEY,
    Author_Id        INTEGER REFERENCES author (Id) ON DELETE SET NULL,
    Book_Name        VARCHAR(30),
    Book_Description TEXT,
    Link             VARCHAR(100),
    Created_At       timestamp with time zone,
    Updated_At       timestamp with time zone
);

CREATE TABLE publish_house
(
    Id         SERIAL PRIMARY KEY,
    Name       VARCHAR(100),
    Created_At timestamp with time zone,
    Updated_At timestamp with time zone
);

CREATE TABLE booking
(
    Id         uuid PRIMARY KEY,
    User_Id    integer REFERENCES users (id),
    Book_Id    integer REFERENCES books (id),
    Start_Date timestamp with time zone,
    End_Time   timestamp with time zone,
    Created_At timestamp with time zone,
    Updated_At timestamp with time zone
);