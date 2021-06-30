-- +migrate Up
CREATE TABLE "Channels" (
	"ChannelID" serial NOT NULL,
	"Name" TEXT NOT NULL,
	CONSTRAINT "Channels_pk" PRIMARY KEY ("ChannelID")
);



CREATE TABLE "Users" (
	"UserID" serial NOT NULL,
	"Username" TEXT NOT NULL,
	"Email" TEXT NOT NULL,
	"Passwort" TEXT NOT NULL,
	"Active" Boolean NOT NULL DEFAULT 'false',
	CONSTRAINT "Users_pk" PRIMARY KEY ("UserID")
);



CREATE TABLE "OnetimePads" (
	"Time" TIMESTAMP NOT NULL,
	"Key" TEXT NOT NULL,
	"UserID" integer NOT NULL,
	"Type" TEXT NOT NULL,
	CONSTRAINT "OnetimePads_pk" PRIMARY KEY ("Time")
);



CREATE TABLE "ChannelhasUsers" (
	"ChannelID" integer NOT NULL,
	"UserID" integer NOT NULL,
	"Right" TEXT NOT NULL,
	CONSTRAINT "Channel_Users_pk" PRIMARY KEY ("ChannelID", "UserID")
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
	"Name" TEXT NOT NULL,
	CONSTRAINT "Tags_pk" PRIMARY KEY ("TagID")
);



CREATE TABLE "LinkHasTag" (
	"LinkID" integer NOT NULL,
	"TagID" integer NOT NULL,
    CONSTRAINT "Link_Tags_pk" PRIMARY KEY ("LinkID", "TagID")
);


ALTER TABLE "Links" ADD CONSTRAINT "Links_fk0" FOREIGN KEY ("ChannelID") REFERENCES "Channels"("ChannelID");

ALTER TABLE "Clicks" ADD CONSTRAINT "Clicks_fk0" FOREIGN KEY ("LinkID") REFERENCES "Links"("LinkID");


ALTER TABLE "LinkHasTag" ADD CONSTRAINT "LinkHasTag_fk0" FOREIGN KEY ("LinkID") REFERENCES "Links"("LinkID");
ALTER TABLE "LinkHasTag" ADD CONSTRAINT "LinkHasTag_fk1" FOREIGN KEY ("TagID") REFERENCES "Tags"("TagID");

ALTER TABLE "ChannelhasUsers" ADD CONSTRAINT "ChannelhasUsers_fk0" FOREIGN KEY ("UserID") REFERENCES "Users"("UserID");
ALTER TABLE "ChannelhasUsers" ADD CONSTRAINT "ChannelhasUsers_fk1" FOREIGN KEY ("ChannelID") REFERENCES "Channels"("ChannelID");

ALTER TABLE "OnetimePads" ADD CONSTRAINT "OnetimePads_fk0" FOREIGN KEY ("UserID") REFERENCES "Users"("UserID");

-- +migrate Down

DROP TABLE "ChannelhasUsers";
DROP TABLE "OnetimePads";
DROP TABLE "Users";
DROP TABLE "Clicks";
DROP TABLE "LinkHasTag";
DROP TABLE "Links";
DROP TABLE "Tags";
DROP TABLE "Channels";