CREATE DATABASE IF NOT EXISTS source;

USE source;

CREATE TABLE IF NOT EXISTS sources(
    source_id varchar(100) not null primary key,
    card_number  varchar(16) not null,
    cvv int(4) unsigned not null 
);

-- source_id is encripted with SHA-256 following the pattern without any space: {"card":"xxxx","cvv":000}
INSERT INTO sources (source_id, card_number, cvv) values ("c674b5630d41cd9dbe08a28c3b11afc63ae4160be329404b42a2eaed0d527e05", "0000000000000001", 123);