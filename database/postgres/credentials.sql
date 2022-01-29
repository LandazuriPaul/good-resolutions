-- set your DB password here
CREATE USER good_resolutions_api WITH PASSWORD 'api-password';

-- preventive drops to avoid any conflict with precedent versions
DROP DATABASE if exists good_resolutions;

-- create the good_resolutions database
CREATE DATABASE good_resolutions
  WITH ENCODING = 'UTF8';
