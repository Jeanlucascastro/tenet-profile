DO
$do$
BEGIN
   IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'tenet_profile') THEN
      CREATE DATABASE "tenet_profile";
   END IF;
END
$do$;