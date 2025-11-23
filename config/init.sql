DO
$do$
BEGIN
   IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'tenet-profile') THEN
      CREATE DATABASE "tenet-profile";
   END IF;
END
$do$;