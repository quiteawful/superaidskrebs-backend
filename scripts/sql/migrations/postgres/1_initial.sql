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
	CONSTRAINT "Users_pk" PRIMARY KEY ("UserID")
);



CREATE TABLE "Activations" (
	"Time" TIMESTAMP NOT NULL,
	"Key" VARCHAR(255) NOT NULL,
	"UserID" integer NOT NULL,
	CONSTRAINT "Activation_pk" PRIMARY KEY ("Time")
);



CREATE TABLE "ChannelhasUsers" (
	"ChannelID" integer NOT NULL,
	"UserID" integer NOT NULL,
	"Right" integer NOT NULL,
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
	"Name" integer NOT NULL,
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

ALTER TABLE "Activations" ADD CONSTRAINT "Activation_fk0" FOREIGN KEY ("UserID") REFERENCES "Users"("UserID");

-- +migrate Down

DROP TABLE "ChannelhasUsers";
DROP TABLE "Activations";
DROP TABLE "Users";
DROP TABLE "Clicks";
DROP TABLE "LinkHasTag";
DROP TABLE "Links";
DROP TABLE "Tags";
DROP TABLE "Channels";