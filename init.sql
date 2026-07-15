CREATE TABLE "contacts" (
    "id" BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,

    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    "name" VARCHAR NOT NULL,
    "dob" DATE,
    "address" VARCHAR,
    "phone" VARCHAR,
    "email" VARCHAR
);

CREATE FUNCTION "update_updated_at"()
RETURNS TRIGGER AS $$
BEGIN
    IF row(NEW.*) IS DISTINCT FROM row(OLD.*) THEN
      NEW."updated_at" = CURRENT_TIMESTAMP;
      RETURN NEW;
   ELSE
      RETURN OLD;
   END IF;
END;
$$ language 'plpgsql';

CREATE TRIGGER "contacts_updated_at"
BEFORE UPDATE ON "contacts"
FOR EACH ROW EXECUTE PROCEDURE "update_updated_at"();

--

INSERT INTO "contacts" ("name", "dob", "address", "phone", "email") VALUES
('Adinda Prawira',   '1996-03-14', Null,                                           '+6281234567890', 'adinda.prawira@gmail.com'),
('Bagus Setiawan',   NULL,         'Jl. Kenanga No. 45, Bandung, Jawa Barat',      '+6285678901234', 'bagus.setiawan@yahoo.com'),
('Citra Handayani',  '2001-07-25', 'Jl. Anggrek Raya No. 7, Jakarta Selatan',      '+6287712345678', 'citra.h@outlook.com'),
('Dimas Nugroho',    '1993-01-19', 'Jl. Cendana No. 88, Yogyakarta',               '+6281398765432', 'dimas.nugroho@gmail.com'),
('Eka Wulandari',    '1999-09-30', 'Jl. Mawar No. 3, Surabaya, Jawa Timur',        NULL,             NULL),
('Fajar Ramadhan',   '1985-05-08', NULL,                                           '+6289956781234', 'fajar.ramadhan@gmail.com'),
('Gita Permatasari', '2003-12-11', 'Jl. Flamboyan No. 16, Malang, Jawa Timur',     '+6281567890123', 'gita.permata@icloud.com'),
('Hendra Kusuma',    '1991-08-04', 'Jl. Teratai No. 9, Medan, Sumatera Utara',     '+6281890123456', NULL),
('Indah Lestari',    NULL,         'Jl. Dahlia No. 33, Denpasar, Bali',            '+6283812345670', 'indah.lestari@yahoo.com'),
('Joko Santoso',     '1982-10-23', 'Jl. Seroja No. 5, Makassar, Sulawesi Selatan', '+6285234567891', 'joko.santoso@gmail.com');
