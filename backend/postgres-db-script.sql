CREATE TABLE musicians (
  id SERIAL PRIMARY KEY,
  musician_name VARCHAR(50) NOT NULL,
  email VARCHAR(50) NOT NULL,
  password VARCHAR(250) NOT NULL
);

CREATE TABLE bands (
  id SERIAL PRIMARY KEY,
  band_name VARCHAR(100) NOT NULL,
  owner_id SERIAL NOT NULL REFERENCES musicians
);

CREATE TABLE instruments (
  id SERIAL PRIMARY KEY,
  band_id SERIAL NOT NULL REFERENCES bands,
  instrument_name VARCHAR(50) NOT NULL
);

CREATE TABLE bandmembers (
  musician_id SERIAL NOT NULL REFERENCES musicians,
  band_id SERIAL NOT NULL REFERENCES bands,
  instrument_id int NULL REFERENCES instruments,
  PRIMARY KEY (musician_id, band_id)
);

CREATE TABLE songs (
  id SERIAL PRIMARY KEY,
  band_id SERIAL NOT NULL REFERENCES bands,
  song_name VARCHAR(50) NOT NULL,
  artist_name VARCHAR(50) NOT NULL
);

CREATE TABLE songInstruments (
  song_id SERIAL NOT NULL REFERENCES songs,
  instrument_id SERIAL NOT NULL REFERENCES instruments,
  pdf_location VARCHAR(150) NOT NULL,
  PRIMARY KEY (song_id, instrument_id)
);

CREATE TABLE joinRequests (
  musician_id SERIAL NOT NULL REFERENCES musicians,
  band_id SERIAL NOT NULL REFERENCES bands,
  PRIMARY KEY (musician_id, band_id)
);
