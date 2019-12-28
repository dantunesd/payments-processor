CREATE DATABASE IF NOT EXISTS source;

USE source;

CREATE TABLE IF NOT EXISTS sources(
    source_id varchar(100) not null primary key,
    card_number  varchar(16) not null,
    cvv int(4) unsigned not null 
);

-- source_id is encripted with SHA-256 following the pattern without any space: {"card":"xxxx","cvv":000}

--  CIELO valid card / Card for MockServer
INSERT INTO sources (source_id, card_number, cvv) values ("c674b5630d41cd9dbe08a28c3b11afc63ae4160be329404b42a2eaed0d527e05", "0000000000000001", 123);

--  CIELO unauthorized
INSERT INTO sources (source_id, card_number, cvv) values ("d3b2e8e32f5d281044166acc1d5d9b54f9a49934c28a7223e58aea90ebf8703d", "0000000000000002", 123);

--  CIELO expired
INSERT INTO sources (source_id, card_number, cvv) values ("e6bb70e5b16e05c7182762de0ef78275810dcf1391f47fb739378762424a0488", "0000000000000003", 123);

-- REDE valid mastercard
INSERT INTO sources (source_id, card_number, cvv) values ("56d473c92fab87e9d0ea8ef53e4eea485e687b04a52ffcf991b04d1c58097c7f", "5448280000000007", 123);

-- REDE invalid
INSERT INTO sources (source_id, card_number, cvv) values ("6de80c6378acf79971822b51462366b803b158688a262cc9c69afe4fbde2c348", "5000000000000008", 123);