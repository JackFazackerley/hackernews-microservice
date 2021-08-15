CREATE TABLE hackernews (
    id int8 NOT NULL,
    "type" varchar NOT NULL,
    "by" varchar NOT NULL,
    url varchar NULL,
    score int4 NOT NULL,
    title varchar NOT NULL,
    "time" int4 NULL,
    CONSTRAINT hackernews_pk PRIMARY KEY (id)
);

CREATE INDEX hackernews_type_idx ON hackernews USING btree (type);