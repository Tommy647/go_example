CREATE TABLE IF NOT EXISTS public.routers (
    id serial NOT NULL PRIMARY KEY,
    vendor varchar (50),
    hostname varchar (50),
    mgmtIP varchar (50)
);