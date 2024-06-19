-- CREATE DATABASE mydb;

-- CREATE USER user1 WITH PASSWORD 'zxc';

GRANT ALL PRIVILEGES ON DATABASE mydb TO user1;

\c mydb;

create table beans (
    bin integer primary key,
    issuer varchar
);

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO user1;