-- +migrate Up
CREATE TABLE "Channels" (
	"ChannelID" serial NOT NULL,
	"Name" VARCHAR(255) NOT NULL,
	CONSTRAINT "Channels_pk" PRIMARY KEY ("ChannelID")
);



CREATE TABLE "Users" (
	"UserID" serial NOT NULL,
	"Username" VARCHAR(255) NOT NULL,
	"Email" VARCHAR(255) NOT NULL,
	"Passwort" VARCHAR(255) NOT NULL,
	"Right" integer NOT NULL,
	"ChannelID" integer NOT NULL,
	CONSTRAINT "Users_pk" PRIMARY KEY ("UserID")
);



CREATE TABLE "Links" (
	"LinkID" serial NOT NULL,
	"ChannelID" integer NOT NULL,
	"URL" TEXT NOT NULL,
	"Image" TEXT NOT NULL,
	CONSTRAINT "Links_pk" PRIMARY KEY ("LinkID")
);



CREATE TABLE "Clicks" (
	"LinkID" integer NOT NULL,
	"Time" TIMESTAMP NOT NULL,
    CONSTRAINT "Clicks_pk" PRIMARY KEY ("Time")
);



CREATE TABLE "Tags" (
	"TagID" serial NOT NULL,
	"Name" integer NOT NULL,
	CONSTRAINT "Tags_pk" PRIMARY KEY ("TagID")
);



CREATE TABLE "LinkHasTag" (
	"LinkID" integer NOT NULL,
	"TagID" integer NOT NULL,
    CONSTRAINT "Link_Tags_pk" PRIMARY KEY ("LinkID", "TagID")
);




ALTER TABLE "Users" ADD CONSTRAINT "Users_fk0" FOREIGN KEY ("ChannelID") REFERENCES "Channels"("ChannelID");

ALTER TABLE "Links" ADD CONSTRAINT "Links_fk0" FOREIGN KEY ("ChannelID") REFERENCES "Channels"("ChannelID");

ALTER TABLE "Clicks" ADD CONSTRAINT "Clicks_fk0" FOREIGN KEY ("LinkID") REFERENCES "Links"("LinkID");


ALTER TABLE "LinkHasTag" ADD CONSTRAINT "LinkHasTag_fk0" FOREIGN KEY ("LinkID") REFERENCES "Links"("LinkID");
ALTER TABLE "LinkHasTag" ADD CONSTRAINT "LinkHasTag_fk1" FOREIGN KEY ("TagID") REFERENCES "Tags"("TagID");

-- +migrate Down

DROP TABLE "Users";
DROP TABLE "Clicks";
DROP TABLE "LinkHasTag";
DROP TABLE "Links";
DROP TABLE "Tags";
DROP TABLE "Channels";